package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "ifaceconf/internal/config"
)

func main() {
    cfg, err := config.LoadProjectConfig("configs/project.yaml")
    if err != nil { fmt.Println("load config error:", err); return }
    dsn := cfg.Datasources.SQL["main"].DSN
    db, err := sql.Open("mysql", dsn)
    if err != nil { fmt.Println("open db error:", err); return }
    defer db.Close()
    var id int64
    var username, email, passwordHash, salt string
    err = db.QueryRow("SELECT id, username, email, password_hash, salt FROM users WHERE username='admin' LIMIT 1").Scan(&id, &username, &email, &passwordHash, &salt)
    if err != nil { fmt.Println("query admin error:", err); return }
    fmt.Println("admin id:", id)
    fmt.Println("username:", username)
    fmt.Println("email:", email)
    fmt.Println("salt length:", len(salt))
    fmt.Println("salt:", salt)
    fmt.Println("password_hash length:", len(passwordHash))
    fmt.Println("password_hash:", passwordHash)
    var calc string
    err = db.QueryRow("SELECT LOWER(SHA2(CONCAT(?, ?), 256))", "admin123", salt).Scan(&calc)
    if err != nil { fmt.Println("calc error:", err); return }
    fmt.Println("calc:", calc)
    fmt.Println("match:", calc == passwordHash)
    var uid int64
    err = db.QueryRow("SELECT id FROM users WHERE username = ? AND password_hash = LOWER(SHA2(CONCAT(?, salt), 256)) LIMIT 1", "admin", "admin123").Scan(&uid)
    if err != nil { fmt.Println("login query error:", err) } else { fmt.Println("login query uid:", uid) }
}