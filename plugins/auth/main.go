package main

import (
    "crypto/sha256"
    "database/sql"
    "encoding/hex"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
    sdk "github.com/hoonfeng/goproc/sdk"
)

func sha256Hex(s string) string { h := sha256.Sum256([]byte(s)); return hex.EncodeToString(h[:]) }

func apikey(params map[string]interface{}) (interface{}, error) {
    key := fmt.Sprint(params["key"]) // plain api key from header
    if key == "" { return nil, fmt.Errorf("missing api key") }
    dsn := fmt.Sprint(params["dsn"])
    if dsn == "" { dsn = os.Getenv("DB_DSN") }
    if dsn == "" { return nil, fmt.Errorf("missing dsn") }
    db, err := sql.Open("mysql", dsn)
    if err != nil { return nil, err }
    defer db.Close()
    var status string
    row := db.QueryRow("SELECT status FROM api_keys WHERE key_hash = SHA2(?,256) LIMIT 1", key)
    if err := row.Scan(&status); err != nil { return nil, fmt.Errorf("unauthorized") }
    if status == "revoked" { return nil, fmt.Errorf("forbidden") }
    return map[string]interface{}{"ok": true}, nil
}

func main() {
    sdk.RegisterFunction("apikey", apikey)
    _ = sdk.Start()
    sdk.Wait()
}