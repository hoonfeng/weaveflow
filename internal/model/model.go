package model

import (
    "database/sql"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "gopkg.in/yaml.v3"
)

type Table struct{ Name string `yaml:"table"`; Columns map[string]Column `yaml:"columns"`; Indexes []Index `yaml:"indexes"`; Comment string `yaml:"comment"` }
type Column struct{ Type string `yaml:"type"`; PK bool `yaml:"pk"`; Auto bool `yaml:"auto"`; NotNull bool `yaml:"notNull"`; Default string `yaml:"default"`; Unique bool `yaml:"unique"`; Comment string `yaml:"comment"` }
type Index struct{ Name string `yaml:"name"`; Columns []string `yaml:"columns"`; Unique bool `yaml:"unique"` }

func Load(dir string) ([]Table, error) {
    var out []Table
    ents, err := os.ReadDir(dir)
    if err != nil { return nil, err }
    for _, e := range ents {
        if e.IsDir() { continue }
        ext := strings.ToLower(filepath.Ext(e.Name()))
        if ext != ".yaml" && ext != ".yml" { continue }
        b, err := os.ReadFile(filepath.Join(dir, e.Name()))
        if err != nil { return nil, err }
        var t Table
        if err := yaml.Unmarshal(b, &t); err != nil { return nil, err }
        if t.Name == "" { continue }
        out = append(out, t)
    }
    return out, nil
}

func GenerateDDL(t Table) []string {
    cols := []string{}
    pks := []string{}
    for name, c := range t.Columns {
        d := fmt.Sprintf("`%s` %s", name, c.Type)
        if c.Auto { d += " AUTO_INCREMENT" }
        if c.NotNull { d += " NOT NULL" }
        if c.Default != "" { d += " DEFAULT " + c.Default }
        if c.Comment != "" { d += " COMMENT '" + escapeSQLString(c.Comment) + "'" }
        cols = append(cols, d)
        if c.PK { pks = append(pks, fmt.Sprintf("`%s`", name)) }
    }
    if len(pks) > 0 { cols = append(cols, "PRIMARY KEY ("+strings.Join(pks, ",")+")") }
    create := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (%s) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4", t.Name, strings.Join(cols, ","))
    if t.Comment != "" { create += " COMMENT='" + escapeSQLString(t.Comment) + "'" }
    ddls := []string{create}
    for _, idx := range t.Indexes {
        uniq := ""
        if idx.Unique { uniq = "UNIQUE " }
        ddls = append(ddls, fmt.Sprintf("CREATE %sINDEX IF NOT EXISTS `%s` ON `%s` (%s)", uniq, idx.Name, t.Name, strings.Join(wrapCols(idx.Columns), ",")))
    }
    return ddls
}

func wrapCols(cols []string) []string { out := make([]string, len(cols)); for i, c := range cols { out[i] = "`"+c+"`" } ; return out }

func ApplyModels(db *sql.DB, tables []Table) error {
    for _, t := range tables {
        if err := ensureTable(db, t); err != nil { return err }
        if err := ensureColumns(db, t); err != nil { return err }
        if err := ensureIndexes(db, t); err != nil { return err }
    }
    return nil
}

func ensureTable(db *sql.DB, t Table) error {
    ddls := GenerateDDL(t)
    if len(ddls) > 0 {
        if _, err := db.Exec(ddls[0]); err != nil { return fmt.Errorf("ddl failed: %s ; %w", ddls[0], err) }
    }
    return nil
}

func ensureColumns(db *sql.DB, t Table) error {
    rows, err := db.Query("SELECT COLUMN_NAME FROM information_schema.COLUMNS WHERE TABLE_SCHEMA=DATABASE() AND TABLE_NAME=?", t.Name)
    if err != nil { return err }
    defer rows.Close()
    existing := map[string]struct{}{}
    for rows.Next() { var c string; _ = rows.Scan(&c); existing[c] = struct{}{} }
    for name, c := range t.Columns {
        if _, ok := existing[name]; ok { continue }
        d := fmt.Sprintf("ALTER TABLE `%s` ADD COLUMN `%s` %s", t.Name, name, c.Type)
        if c.NotNull { d += " NOT NULL" }
        if c.Default != "" { d += " DEFAULT " + c.Default }
        if c.Comment != "" { d += " COMMENT '" + escapeSQLString(c.Comment) + "'" }
        if _, err := db.Exec(d); err != nil { return err }
    }
    return nil
}

func ensureIndexes(db *sql.DB, t Table) error {
    rows, err := db.Query("SELECT INDEX_NAME FROM information_schema.STATISTICS WHERE TABLE_SCHEMA=DATABASE() AND TABLE_NAME=?", t.Name)
    if err != nil { return err }
    defer rows.Close()
    existing := map[string]struct{}{}
    for rows.Next() { var n string; _ = rows.Scan(&n); existing[n] = struct{}{} }
    for _, idx := range t.Indexes {
        if _, ok := existing[idx.Name]; ok { continue }
        uniq := ""
        if idx.Unique { uniq = "UNIQUE " }
        ddl := fmt.Sprintf("CREATE %sINDEX `%s` ON `%s` (%s)", uniq, idx.Name, t.Name, strings.Join(wrapCols(idx.Columns), ","))
        if _, err := db.Exec(ddl); err != nil { return err }
    }
    for colName, c := range t.Columns {
        if !c.Unique { continue }
        idxName := fmt.Sprintf("uniq_%s_%s", t.Name, colName)
        if _, ok := existing[idxName]; ok { continue }
        ddl := fmt.Sprintf("CREATE UNIQUE INDEX `%s` ON `%s` (`%s`)", idxName, t.Name, colName)
        if _, err := db.Exec(ddl); err != nil { return err }
    }
    return nil
}
func escapeSQLString(s string) string { return strings.ReplaceAll(s, "'", "''") }