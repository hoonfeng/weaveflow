package main

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
    sdk "github.com/hoonfeng/goproc/sdk"
)

func monthly(params map[string]interface{}) (interface{}, error) {
    tenantID := fmt.Sprint(params["tenant_id"]) // from header X-Tenant-Id
    dsn := fmt.Sprint(params["dsn"]) ; if dsn == "" { dsn = os.Getenv("DB_DSN") } ; if dsn == "" { return nil, fmt.Errorf("missing dsn") }
    db, err := sql.Open("mysql", dsn)
    if err != nil { return nil, err }
    defer db.Close()
    if tenantID == "" {
        key := fmt.Sprint(params["key"]) // from header X-Api-Key
        if key == "" { return nil, fmt.Errorf("missing tenant_id and key") }
        var tid string
        row := db.QueryRow("SELECT tenant_id FROM api_keys WHERE key_hash = SHA2(?,256) AND status='active' LIMIT 1", key)
        if err := row.Scan(&tid); err != nil { return nil, fmt.Errorf("unauthorized") }
        tenantID = tid
    }
    var quota int
    row := db.QueryRow("SELECT p.monthly_quota FROM subscriptions s JOIN plans p ON s.plan_id=p.id WHERE s.tenant_id=? AND s.status='active' LIMIT 1", tenantID)
    if err := row.Scan(&quota); err != nil { return nil, fmt.Errorf("quota not found") }
    var cnt int
    row2 := db.QueryRow("SELECT COUNT(1) FROM usage_logs WHERE tenant_id=? AND DATE_FORMAT(ts,'%Y-%m')=DATE_FORMAT(NOW(),'%Y-%m')", tenantID)
    _ = row2.Scan(&cnt)
    if quota > 0 && cnt >= quota { return nil, fmt.Errorf("quota exceeded") }
    return map[string]interface{}{"ok": true, "used": cnt, "quota": quota}, nil
}

func main() {
    sdk.RegisterFunction("monthly", monthly)
    _ = sdk.Start()
    sdk.Wait()
}