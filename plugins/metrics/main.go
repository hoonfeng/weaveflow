package main

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
    sdk "github.com/hoonfeng/goproc/sdk"
)

func record(params map[string]interface{}) (interface{}, error) {
    tenantID := fmt.Sprint(params["tenant_id"]) // from header X-Tenant-Id
    endpoint := fmt.Sprint(params["endpoint"])  // method+path
    provider := fmt.Sprint(params["provider"])  // optional
    status := fmt.Sprint(params["status"])      // response status
    duration := fmt.Sprint(params["duration_ms"]) // int ms
    dsn := fmt.Sprint(params["dsn"]) ; if dsn == "" { dsn = os.Getenv("DB_DSN") } ; if dsn == "" { return nil, fmt.Errorf("missing dsn") }
    db, err := sql.Open("mysql", dsn)
    if err != nil { return nil, err }
    defer db.Close()
    _, err = db.Exec("INSERT INTO usage_logs(tenant_id, endpoint, provider, status, duration_ms, ts) VALUES(?,?,?,?,?, NOW())", tenantID, endpoint, provider, status, duration)
    if err != nil { return nil, err }
    return map[string]interface{}{"ok": true}, nil
}

func main() {
    sdk.RegisterFunction("record", record)
    _ = sdk.Start()
    sdk.Wait()
}