package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "database/sql"
    "encoding/base64"
    "fmt"
    "os"
    "time"
    _ "github.com/go-sql-driver/mysql"
    sdk "github.com/hoonfeng/goproc/sdk"
)

func hmacVerify(params map[string]interface{}) (interface{}, error) {
    key := fmt.Sprint(params["key"])      // api key header
    sig := fmt.Sprint(params["sig"])      // signature header
    ts := fmt.Sprint(params["ts"])        // timestamp header
    nonce := fmt.Sprint(params["nonce"])  // nonce header
    method := fmt.Sprint(params["method"])
    path := fmt.Sprint(params["path"])
    bodyHash := fmt.Sprint(params["body_hash"])
    if key == "" || sig == "" || ts == "" || method == "" || path == "" { return nil, fmt.Errorf("missing fields") }
    dsn := fmt.Sprint(params["dsn"]) ; if dsn == "" { dsn = os.Getenv("DB_DSN") } ; if dsn == "" { return nil, fmt.Errorf("missing dsn") }
    db, err := sql.Open("mysql", dsn)
    if err != nil { return nil, err }
    defer db.Close()
    var secret string
    row := db.QueryRow("SELECT secret FROM api_keys WHERE key_hash = SHA2(?,256) AND status='active' LIMIT 1", key)
    if err := row.Scan(&secret); err != nil { return nil, fmt.Errorf("unauthorized") }
    skew := 300
    if v := fmt.Sprint(params["skewSeconds"]); v != "" { var vv int; _, _ = fmt.Sscanf(v, "%d", &vv); if vv > 0 { skew = vv } }
    tParsed, err := time.Parse(time.RFC3339, ts)
    if err != nil { return nil, fmt.Errorf("unauthorized") }
    if d := time.Since(tParsed.UTC()); d < -time.Duration(skew)*time.Second || d > time.Duration(skew)*time.Second { return nil, fmt.Errorf("unauthorized") }
    var exists int
    err = db.QueryRow("SELECT COUNT(1) FROM hmac_nonces WHERE key_hash = SHA2(?,256) AND nonce = ? LIMIT 1", key, nonce).Scan(&exists)
    if err == nil && exists > 0 { return nil, fmt.Errorf("unauthorized") }
    // compute message
    msg := method + "\n" + path + "\n" + ts + "\n" + nonce + "\n" + bodyHash
    mac := hmac.New(sha256.New, []byte(secret))
    mac.Write([]byte(msg))
    expect := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
    if expect != sig { return nil, fmt.Errorf("unauthorized") }
    _, _ = db.Exec("INSERT INTO hmac_nonces(key_hash, nonce, ts_text) VALUES (SHA2(?,256), ?, ?)", key, nonce, ts)
    return map[string]interface{}{"ok": true}, nil
}

func main() {
    sdk.RegisterFunction("hmac", hmacVerify)
    _ = sdk.Start()
    sdk.Wait()
}