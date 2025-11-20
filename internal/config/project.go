package config

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v3"
)

// 服务器配置 / Server configuration
type ServerConfig struct {
    Port          int    `yaml:"port"`
    Timeout       string `yaml:"timeout"`
    MaxUploadSize string `yaml:"maxUploadSize"`
    AutoReload    bool   `yaml:"autoReload"`
}

// 文档配置 / Docs configuration
type DocsConfig struct {
    Enabled  bool   `yaml:"enabled"`
    Title    string `yaml:"title"`
    BasePath string `yaml:"basePath"`
}

type CorsConfig struct {
    Enabled        bool     `yaml:"enabled"`
    AllowedOrigins []string `yaml:"allowedOrigins"`
    AllowedMethods []string `yaml:"allowedMethods"`
    AllowedHeaders []string `yaml:"allowedHeaders"`
}

type RateLimitConfig struct {
    Enabled bool    `yaml:"enabled"`
    RPS     float64 `yaml:"rps"`
    Burst   int     `yaml:"burst"`
    PerIp   bool    `yaml:"perIp"`
}


// 数据源集合 / Datasources group
type DatasourcesConfig struct {
    SQL    map[string]SQLConfig    `yaml:"sql"`
    KV     map[string]KVConfig     `yaml:"kv"`
    Vector map[string]VectorConfig `yaml:"vector"`
    Cache  map[string]CacheConfig  `yaml:"cache"`
    Blob   map[string]BlobConfig   `yaml:"blob"`
}

// SQL 数据源 / SQL datasource
type SQLConfig struct {
    Driver          string `yaml:"driver"`
    DSN             string `yaml:"dsn"`
    MaxOpenConns    int    `yaml:"maxOpenConns"`
    MaxIdleConns    int    `yaml:"maxIdleConns"`
    ConnMaxLifetime string `yaml:"connMaxLifetime"`
}

// KV 数据源 / KV datasource
type KVConfig struct {
    Type string `yaml:"type"`
    Addr string `yaml:"addr"`
    DB   int    `yaml:"db"`
}

// 向量库数据源 / Vector datasource
type VectorConfig struct {
    Type     string `yaml:"type"`
    Endpoint string `yaml:"endpoint"`
    APIKey   string `yaml:"apiKey"`
    Collection string `yaml:"collection"`
}

// 内存缓存配置 / Cache configuration
type CacheConfig struct {
    Policy     string `yaml:"policy"`
    Capacity   int    `yaml:"capacity"`
    DefaultTTL string `yaml:"defaultTTL"`
}

// Blob 存储配置 / Blob storage
type BlobConfig struct {
    Type     string `yaml:"type"`
    BasePath string `yaml:"basePath"`
    Bucket   string `yaml:"bucket"`
}

// 插件配置（goproc） / Plugin configuration (goproc)
type PluginRegistryItem struct {
    Name       string         `yaml:"name"`
    Executable ExecutablePath `yaml:"executable"`
    Instances  int            `yaml:"instances"`
    Timeout    string         `yaml:"timeout"`
    QueueSize  int            `yaml:"queueSize"`
    Functions  []string       `yaml:"functions"`
    Env        map[string]string `yaml:"env"`
    EnvFrom    []string           `yaml:"envFrom"`
}

// 插件可执行路径 / Plugin executable paths
type ExecutablePath struct {
    Windows string `yaml:"windows"`
    Unix    string `yaml:"unix"`
}

type PluginsConfig struct {
    Enabled          bool                 `yaml:"enabled"`
    Runtime          string               `yaml:"runtime"`
    InstancesDefault int                  `yaml:"instancesDefault"`
    Registry         []PluginRegistryItem `yaml:"registry"`
}

// 项目配置根结构 / Project configuration root
type ProjectConfig struct {
    Server      ServerConfig      `yaml:"server"`
    Docs        DocsConfig        `yaml:"docs"`
    Cors        CorsConfig        `yaml:"cors"`
    RateLimit   RateLimitConfig   `yaml:"rateLimit"`
    Datasources DatasourcesConfig `yaml:"datasources"`
    Plugins     PluginsConfig     `yaml:"plugins"`
    Security    SecurityConfig    `yaml:"security"`
    Static      StaticConfig      `yaml:"static"`
    Init        InitDataConfig    `yaml:"init"`
}

type SecurityConfig struct {
    AuthType  string `yaml:"authType"`
    JWTSecret string `yaml:"jwtSecret"`
    StripeWebhookSecret string `yaml:"stripeWebhookSecret"`
    JWTIssuer string `yaml:"jwtIssuer"`
    JWTAudience string `yaml:"jwtAudience"`
}

// 加载项目配置 / Load project config
func LoadProjectConfig(path string) (*ProjectConfig, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("读取配置文件失败: %w", err)
    }
    var cfg ProjectConfig
    if err := yaml.Unmarshal(data, &cfg); err != nil {
        return nil, fmt.Errorf("解析配置失败: %w", err)
    }
    // ENV 展开
    if cfg.Datasources.SQL != nil {
        for name, sc := range cfg.Datasources.SQL { sc.DSN = expandString(sc.DSN); cfg.Datasources.SQL[name] = sc }
    }
    if cfg.Datasources.Vector != nil {
        for name, vc := range cfg.Datasources.Vector { vc.Endpoint = expandString(vc.Endpoint); vc.APIKey = expandString(vc.APIKey); cfg.Datasources.Vector[name] = vc }
    }
    for i := range cfg.Plugins.Registry {
        cfg.Plugins.Registry[i].Executable.Windows = expandString(cfg.Plugins.Registry[i].Executable.Windows)
        cfg.Plugins.Registry[i].Executable.Unix = expandString(cfg.Plugins.Registry[i].Executable.Unix)
        if cfg.Plugins.Registry[i].Env != nil { 
            envMap := make(map[string]any)
            for k, v := range cfg.Plugins.Registry[i].Env {
                envMap[k] = v
            }
            expandMap(envMap)
            // 将修改后的值重新赋值给Env
            for k, v := range envMap {
                cfg.Plugins.Registry[i].Env[k] = v.(string)
            }
        }
    }
    for i := range cfg.Static.Items {
        cfg.Static.Items[i].Route = expandString(cfg.Static.Items[i].Route)
        cfg.Static.Items[i].Dir = expandString(cfg.Static.Items[i].Dir)
        cfg.Static.Items[i].File = expandString(cfg.Static.Items[i].File)
        cfg.Static.Items[i].Index = expandString(cfg.Static.Items[i].Index)
    }
    if cfg.Server.Port == 0 {
        cfg.Server.Port = 8080
    }
    if cfg.Docs.BasePath == "" {
        cfg.Docs.BasePath = "/docs"
    }
    if len(cfg.Cors.AllowedOrigins) == 0 { cfg.Cors.AllowedOrigins = []string{"*"} }
    if len(cfg.Cors.AllowedMethods) == 0 { cfg.Cors.AllowedMethods = []string{"GET","POST","PUT","DELETE"} }
    if len(cfg.Cors.AllowedHeaders) == 0 { cfg.Cors.AllowedHeaders = []string{"*"} }
    return &cfg, nil
}

type StaticConfig struct {
    Enabled bool         `yaml:"enabled"`
    Items   []StaticItem `yaml:"items"`
}

type StaticItem struct {
    Route string `yaml:"route"`
    Dir   string `yaml:"dir"`
    File  string `yaml:"file"`
    Index string `yaml:"index"`
    SPA   bool   `yaml:"spa"`
    Strip bool   `yaml:"strip"`
}

type InitDataConfig struct {
    Enabled                 bool     `yaml:"enabled"`
    EnsureAdminUser         bool     `yaml:"ensureAdminUser"`
    EnsureAdminRoleBinding  bool     `yaml:"ensureAdminRoleBinding"`
    Scripts                 []string `yaml:"scripts"`
}