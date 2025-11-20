package datasource

import (
    "crypto/sha256"
    "database/sql"
    "encoding/hex"
    "errors"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "time"

    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
    _ "github.com/microsoft/go-mssqldb"

    "ifaceconf/internal/config"
)

// 管理器 / Manager
type Manager struct {
    SQL map[string]*sql.DB
    Cache map[string]*TTLCache
    Blob map[string]BlobStore
    Vector map[string]VectorStore
}

// 初始化 / Initialize
func NewManager(cfg *config.DatasourcesConfig) (*Manager, error) {
    m := &Manager{SQL: map[string]*sql.DB{}, Cache: map[string]*TTLCache{}, Blob: map[string]BlobStore{}, Vector: map[string]VectorStore{}}
    // SQL
    for name, sc := range cfg.SQL {
        db, err := sql.Open(sc.Driver, sc.DSN)
        if err != nil { return nil, err }
        if sc.MaxOpenConns > 0 { db.SetMaxOpenConns(sc.MaxOpenConns) }
        if sc.MaxIdleConns > 0 { db.SetMaxIdleConns(sc.MaxIdleConns) }
        m.SQL[name] = db
    }
    // Cache
    for name, cc := range cfg.Cache {
        ttl := parseDuration(cc.DefaultTTL)
        m.Cache[name] = NewTTLCache(cc.Capacity, ttl)
    }
    // Blob
    for name, bc := range cfg.Blob {
        switch bc.Type {
        case "local":
            if bc.BasePath == "" { bc.BasePath = "./storage" }
            _ = os.MkdirAll(bc.BasePath, 0o755)
            m.Blob[name] = &LocalBlob{BasePath: bc.BasePath}
        default:
            return nil, fmt.Errorf("不支持的 Blob 类型: %s", bc.Type)
        }
    }
    // Vector: 默认内存实现，可后续扩展为 qdrant/milvus/pinecone
    for name, vc := range cfg.Vector {
        switch vc.Type {
        case "memory", "":
            m.Vector[name] = NewInMemoryVector()
        case "qdrant":
            m.Vector[name] = NewQdrantVector(vc.Endpoint, vc.APIKey)
        default:
            // 未来适配外部向量库
            m.Vector[name] = NewInMemoryVector()
        }
    }
    return m, nil
}

// SQL 查询 / SQL query
func (m *Manager) SQLQuery(ds string, sqlStr string, params map[string]any) ([]map[string]any, error) {
    db := m.SQL[ds]
    if db == nil { return nil, errors.New("SQL 数据源不存在") }
    // 简易命名参数支持：使用 ? 绑定顺序参数（示例实现）
    rows, err := db.Query(sqlStr, ValuesFrom(params)...)
    if err != nil { return nil, err }
    defer rows.Close()
    cols, _ := rows.Columns()
    res := []map[string]any{}
    for rows.Next() {
        vals := make([]any, len(cols))
        ptrs := make([]any, len(cols))
        for i := range vals { ptrs[i] = &vals[i] }
        if err := rows.Scan(ptrs...); err != nil { return nil, err }
        mrow := map[string]any{}
        for i, c := range cols {
            if b, ok := vals[i].([]byte); ok { mrow[c] = string(b) } else { mrow[c] = vals[i] }
        }
        res = append(res, mrow)
    }
    return res, nil
}

// SQL 执行 / SQL exec
func (m *Manager) SQLExec(ds string, sqlStr string, params map[string]any) error {
    db := m.SQL[ds]
    if db == nil { return errors.New("SQL 数据源不存在") }
    _, err := db.Exec(sqlStr, ValuesFrom(params)...)
    return err
}

// KV 示例：用 Cache 代替（占位），实际应对接 Redis/Etcd/Badger
func (m *Manager) KVSet(ds string, key string, val any) error {
    c := m.Cache[ds]
    if c == nil { return errors.New("KV 数据源不存在") }
    c.Set(key, val, 0)
    return nil
}

func (m *Manager) KVGet(ds string, key string) (any, error) {
    c := m.Cache[ds]
    if c == nil { return nil, errors.New("KV 数据源不存在") }
    v, ok := c.Get(key)
    if !ok { return nil, errors.New("KV 未命中") }
    return v, nil
}

// 上传保存到本地 / Save upload to local blob
func (m *Manager) UploadSave(ds string, input any, naming string) ([]string, error) {
    lb := m.Blob[ds]
    if lb == nil { return nil, errors.New("Blob 数据源不存在") }
    var files []UploadedFile
    switch v := input.(type) {
    case []UploadedFile:
        files = v
    case []any:
        for _, it := range v {
            if uf, ok := it.(UploadedFile); ok { files = append(files, uf) }
        }
        if len(files) == 0 { return nil, errors.New("输入类型错误，应为文件列表") }
    default:
        return nil, errors.New("输入类型错误，应为文件列表")
    }
    ids := make([]string, 0, len(files))
    for _, f := range files {
        var name string
        switch naming {
        case "sha256":
            h := sha256.New()
            if _, err := io.Copy(h, f.Content); err != nil { return nil, err }
            sum := h.Sum(nil)
            name = hex.EncodeToString(sum)
            if _, err := f.Content.Seek(0, 0); err != nil { return nil, err }
        case "uuid":
            name = fmt.Sprintf("%d-%s", time.Now().UnixNano(), f.Filename)
        default:
            name = f.Filename
        }
        p := filepath.Join(lb.Base(), name)
        out, err := os.Create(p)
        if err != nil { return nil, err }
        if _, err := io.Copy(out, f.Content); err != nil { out.Close(); return nil, err }
        out.Close()
        ids = append(ids, name)
    }
    return ids, nil
}

func ValuesFrom(m map[string]any) []any {
    out := make([]any, 0, len(m))
    for _, v := range m { out = append(out, v) }
    return out
}

func ValuesFromOrder(m map[string]any, order []string) []any {
    if len(order) == 0 { return ValuesFrom(m) }
    out := make([]any, 0, len(order))
    for _, k := range order { out = append(out, m[k]) }
    return out
}

func parseDuration(s string) time.Duration {
    if s == "" { return 0 }
    d, _ := time.ParseDuration(s)
    return d
}

// TTL 内存缓存 / TTL in-memory cache
type item struct{ v any; exp int64 }

type TTLCache struct {
    data map[string]item
    cap  int
    ttl  time.Duration
}

func NewTTLCache(capacity int, ttl time.Duration) *TTLCache {
    return &TTLCache{data: make(map[string]item), cap: capacity, ttl: ttl}
}

func (c *TTLCache) Set(key string, v any, ttlOverride time.Duration) {
    ttl := c.ttl
    if ttlOverride > 0 { ttl = ttlOverride }
    c.data[key] = item{v: v, exp: time.Now().Add(ttl).UnixNano()}
}

func (c *TTLCache) Get(key string) (any, bool) {
    it, ok := c.data[key]
    if !ok { return nil, false }
    if it.exp > 0 && time.Now().UnixNano() > it.exp { delete(c.data, key); return nil, false }
    return it.v, true
}

// Blob 存储接口 / Blob storage interface
type BlobStore interface { Base() string }

type LocalBlob struct { BasePath string }

func (l *LocalBlob) Base() string { return l.BasePath }

// 上传文件模型 / Uploaded file model
type UploadedFile struct {
    Filename string
    Content  ReadSeekCloser
}

// 读写接口 / ReadSeekCloser interface
type ReadSeekCloser interface {
    io.Reader
    io.Seeker
    io.Closer
}