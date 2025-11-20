package router

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "encoding/hex"
    "bytes"
    "io"
    "net/http"
    "time"
    "ifaceconf/internal/datasource"
)

func RequireHMACSignature(ds *datasource.Manager, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        sig := r.Header.Get("X-Signature")
        ts := r.Header.Get("X-Timestamp")
        nonce := r.Header.Get("X-Nonce")
        key := r.Header.Get("X-Api-Key")
        if sig == "" || ts == "" || nonce == "" || key == "" { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        rows, err := ds.SQLQuery("main", "SELECT secret FROM api_keys WHERE key_hash = ? AND status='active' LIMIT 1", map[string]any{"key_hash": sha256Hex(key)})
        if err != nil || len(rows) == 0 { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        secret := rows[0]["secret"].(string)
        if !withinSkew(ts, 300) { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        if used(ds, nonce) { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        bodyHash := sha256Body(r)
        msg := r.Method + "\n" + r.URL.Path + "\n" + ts + "\n" + nonce + "\n" + bodyHash
        mac := hmac.New(sha256.New, []byte(secret))
        _, _ = mac.Write([]byte(msg))
        expect := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
        if expect != sig { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        next(w, r)
    }
}

func sha256Hex(s string) string { h := sha256.Sum256([]byte(s)); return hex.EncodeToString(h[:]) }

func withinSkew(ts string, skewSec int) bool {
    t, err := time.Parse(time.RFC3339, ts)
    if err != nil { return false }
    d := time.Since(t)
    if d < 0 { d = -d }
    return d <= time.Duration(skewSec)*time.Second
}

func sha256Body(r *http.Request) string {
    if r.Body == nil { return "" }
    b, _ := io.ReadAll(r.Body)
    r.Body = io.NopCloser(bytes.NewReader(b))
    h := sha256.Sum256(b)
    return base64.RawURLEncoding.EncodeToString(h[:])
}

func used(ds *datasource.Manager, nonce string) bool {
    if ds.Cache["default"] == nil { return false }
    if _, ok := ds.Cache["default"].Get("nonce:"+nonce); ok { return true }
    ds.Cache["default"].Set("nonce:"+nonce, 1, 5*time.Minute)
    return false
}