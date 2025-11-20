package router

import (
    "crypto/sha256"
    "encoding/hex"
    "net/http"
    "strings"
    "fmt"
    "ifaceconf/internal/datasource"
)

func RequireApiKey(ds *datasource.Manager, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        key := r.Header.Get("X-Api-Key")
        if key == "" {
            auth := r.Header.Get("Authorization")
            if strings.HasPrefix(auth, "ApiKey ") { key = strings.TrimSpace(strings.TrimPrefix(auth, "ApiKey ")) }
        }
        if key == "" { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        h := sha256.Sum256([]byte(key))
        hh := hex.EncodeToString(h[:])
        rows, err := ds.SQLQuery("main", "SELECT tenant_id, status FROM api_keys WHERE key_hash = ? LIMIT 1", map[string]any{"key_hash": hh})
        if err != nil || len(rows) == 0 { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        if s, _ := rows[0]["status"].(string); s == "revoked" { http.Error(w, "forbidden", http.StatusForbidden); return }
        tid := fmt.Sprint(rows[0]["tenant_id"])
        if tid != "" { r.Header.Set("X-Tenant-Id", tid) }
        next(w, r)
    }
}