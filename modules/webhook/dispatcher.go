package webhook

import (
    "bytes"
    "net/http"
    "time"
    "ifaceconf/internal/datasource"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
)

type Dispatcher struct{ ds *datasource.Manager; client *http.Client }

func New(ds *datasource.Manager) *Dispatcher { return &Dispatcher{ds: ds, client: &http.Client{Timeout: 5 * time.Second}} }

func (d *Dispatcher) Run(stop <-chan struct{}) {
    t := time.NewTicker(2 * time.Second)
    defer t.Stop()
    for {
        select {
        case <-t.C:
            d.tick()
        case <-stop:
            return
        }
    }
}

func (d *Dispatcher) tick() {
    rows, err := d.ds.SQLQuery("main", "SELECT id, endpoint_id, event, payload, retries FROM webhook_tasks WHERE status IN ('pending','retry') AND (next_try_at IS NULL OR next_try_at<=NOW()) LIMIT 10", map[string]any{})
    if err != nil { return }
    for _, r := range rows {
        id := r["id"]
        eid := r["endpoint_id"]
        epRows, err := d.ds.SQLQuery("main", "SELECT url, secret FROM webhook_endpoints WHERE id = ?", map[string]any{"id": eid})
        if err != nil || len(epRows) == 0 { continue }
        url := epRows[0]["url"].(string)
        secret, _ := epRows[0]["secret"].(string)
        payload := []byte(r["payload"].(string))
        req, _ := http.NewRequest("POST", url, bytes.NewReader(payload))
        req.Header.Set("Content-Type", "application/json")
        if secret != "" {
            ts := time.Now().UTC().Format(time.RFC3339)
            sig := sign(secret, ts, payload)
            req.Header.Set("X-Timestamp", ts)
            req.Header.Set("X-Signature", sig)
        }
        resp, err := d.client.Do(req)
        if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
            _ = d.ds.SQLExec("main", "UPDATE webhook_tasks SET status='done' WHERE id=?", map[string]any{"id": id})
            if resp.Body != nil { resp.Body.Close() }
            continue
        }
        if resp != nil && resp.Body != nil { resp.Body.Close() }
        retries := toIntSafe(r["retries"])
        retries++
        next := time.Now().Add(time.Duration(1<<minInt(retries, 5)) * time.Second)
        _ = d.ds.SQLExec("main", "UPDATE webhook_tasks SET status='retry', retries=?, next_try_at=? WHERE id=?", map[string]any{"retries": retries, "next": next, "id": id})
    }
}

func sign(secret string, ts string, body []byte) string {
    mac := hmac.New(sha256.New, []byte(secret))
    mac.Write([]byte(ts))
    mac.Write([]byte("\n"))
    mac.Write(body)
    return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func toIntSafe(v any) int { switch t := v.(type) { case int: return t; case int64: return int(t); case float64: return int(t); default: return 0 } }
func minInt(a, b int) int { if a < b { return a }; return b }