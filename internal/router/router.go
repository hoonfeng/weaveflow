package router

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
    "context"
    "strings"
    "database/sql"
    "fmt"
    "os"
    "time"
    "bytes"
    "strconv"
    

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    jwt "github.com/golang-jwt/jwt/v5"

    "ifaceconf/internal/config"
    "ifaceconf/internal/core"
    "ifaceconf/internal/datasource"
    "ifaceconf/internal/plugin"
    "ifaceconf/internal/ext"
    "ifaceconf/internal/model"
    "ifaceconf/modules/apphooks"
    "ifaceconf/internal/apidocs"
    "ifaceconf/internal/obs"
    idocs "ifaceconf/internal/docs"
)

// 构建路由 / Build router from configs
func BuildRouter(baseCtx context.Context, project *config.ProjectConfig, interfaces []*config.InterfaceConfig) (http.Handler, error) {
    currentInterfaces = interfaces
    interfaceIndex = make(map[string]*config.InterfaceConfig, len(interfaces))
    for _, x := range interfaces { interfaceIndex[strings.ToUpper(x.Method)+" "+strings.TrimSuffix(x.Path, "/")] = x }
    r := chi.NewRouter()
    r.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"*"}, AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, AllowedHeaders: []string{"*"}}))

    // 数据源管理器 / Datasource manager
    dsMgr, err := datasource.NewManager(&project.Datasources)
    if err != nil { return nil, err }

    // 插件运行时 / Plugin runtime
    prt, _ := plugin.BuildRuntime(&project.Plugins)

    // removed built-in debug route to keep framework exposing no endpoints by default
    

    // 注册接口 / Register endpoints
    limiterMap := map[string]*Limiter{}
    labels := map[string]string{}
    for _, ic := range interfaces {
        method := strings.ToUpper(ic.Method)
        baseHandler := makeHandler(baseCtx, project, ic, dsMgr, prt)
        handler := baseHandler
        if h := ext.GetHooks().Auth; h != nil {
            handler = h(project, ic, dsMgr, prt, handler)
        } else {
            authType := strings.ToLower(ic.Auth)
            if authType == "" { authType = strings.ToLower(project.Security.AuthType) }
            switch authType {
            case "jwt":
                if project.Security.JWTSecret != "" {
                    handler = requireJWTWithRolesConfig(project.Security, ic.Roles, handler)
                    if len(ic.Permissions) > 0 {
                        handler = requirePermissionsClaims(project.Security, ic.Permissions, handler)
                    }
                }
            case "apikey":
                handler = RequireApiKey(dsMgr, RequireHMACSignature(dsMgr, handler))
            }
        }
        perIp := false
        perTenant := false
        perKey := false
        if ic.RateLimit != nil {
            rpsAny, brAny := ic.RateLimit["rps"], ic.RateLimit["burst"]
            var rps float64; var br int
            switch x := rpsAny.(type) { case int: rps=float64(x); case float64: rps=x }
            switch y := brAny.(type) { case int: br=y; case float64: br=int(y) }
            if p, ok := ic.RateLimit["perIp"].(bool); ok { perIp = p }
            if p, ok := ic.RateLimit["perTenant"].(bool); ok { perTenant = p }
            if p, ok := ic.RateLimit["perKey"].(bool); ok { perKey = p }
            if rps > 0 { limiterMap[ic.Path] = NewLimiter(rps, br) }
        }
        if lim := limiterMap[ic.Path]; lim != nil {
            orig := handler
            handler = func(w http.ResponseWriter, r *http.Request) {
                key := ic.Path
                if perTenant { key = ic.Path+"|tenant:"+r.Header.Get("X-Tenant-Id") }
                if perKey { if k := r.Header.Get("X-Api-Key"); k != "" { key = ic.Path+"|key:"+k } }
                if perIp && !perTenant && !perKey { key = ic.Path+"|"+r.RemoteAddr }
                if perTenant || perKey || perIp {
                if limiterMap[key] == nil {
                    rate := lim.rate; burst := int(lim.burst)
                    limiterMap[key] = NewLimiter(rate, burst)
                }
                    if !limiterMap[key].Allow() { http.Error(w, "too many requests", http.StatusTooManyRequests); return }
                } else {
                    if !lim.Allow() { http.Error(w, "too many requests", http.StatusTooManyRequests); return }
                }
                orig(w, r)
            }
        }
        switch method {
        case http.MethodGet:
            r.Get(ic.Path, handler)
        case http.MethodPost:
            r.Post(ic.Path, handler)
        case http.MethodPut:
            r.Put(ic.Path, handler)
        case http.MethodDelete:
            r.Delete(ic.Path, handler)
        default:
            log.Printf("未知方法: %s", ic.Method)
        }
        labels[method+" "+ic.Path] = ic.Module
    }
    if len(labels) > 0 { obs.SetRouteLabels(labels) }
    // 收集权限声明列表 / collect permission declarations
    perms := make([]map[string]any, 0, 64)
    for _, ic := range interfaces {
        if len(ic.Permissions) == 0 { continue }
        nm := ic.Endpoint
        if ic.Docs != nil { if t, ok := ic.Docs["title"].(string); ok && t != "" { nm = t } }
        for _, code := range ic.Permissions { perms = append(perms, map[string]any{"code": code, "name": nm}) }
    }
    currentPermissions = perms
    return r, nil
}

// removed internal permission ingestion to keep framework decoupled

func requirePermissionsClaims(sec config.SecurityConfig, perms []string, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if len(perms) == 0 { next(w,r); return }
        auth := r.Header.Get("Authorization")
        if !strings.HasPrefix(auth, "Bearer ") { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        tokenStr := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
        t, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) { return []byte(sec.JWTSecret), nil })
        if err != nil || !t.Valid { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        have := map[string]struct{}{}
        if claims, ok := t.Claims.(jwt.MapClaims); ok {
            if v, ok := claims["perms"]; ok {
                switch arr := v.(type) {
                case []any:
                    for _, x := range arr { if s, ok := x.(string); ok { have[s] = struct{}{} } }
                case []string:
                    for _, s := range arr { have[s] = struct{}{} }
                }
            }
        }
        for _, p := range perms { if _, ok := have[p]; !ok { http.Error(w, "forbidden", http.StatusForbidden); return } }
        next(w, r)
    }
}

// 处理函数 / Handler
func makeHandler(baseCtx context.Context, project *config.ProjectConfig, ic *config.InterfaceConfig, ds *datasource.Manager, prt plugin.Runtime) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := core.NewRequestContext(baseCtx)
        start := time.Now()
        tx := map[string]*sql.Tx{}
        // 参数收集 / Collect params
        req := map[string]any{}
        // 路径参数（chi 路由器会解析）
        // Path params
        rc := chi.RouteContext(r.Context())
        for i, k := range rc.URLParams.Keys {
            if i < len(rc.URLParams.Values) { req[k] = rc.URLParams.Values[i] }
        }
        // Query params
        for k, v := range r.URL.Query() { if len(v) > 0 { req[k] = v[0] } }
        // Headers
        for k, v := range r.Header { if len(v) > 0 { req["header."+k] = v[0] } }
        // Body JSON
        if r.Body != nil && strings.Contains(r.Header.Get("Content-Type"), "application/json") {
            var m map[string]any
            _ = json.NewDecoder(r.Body).Decode(&m)
            for k, v := range m { req[k] = v }
            req["body"] = m
        }
        // Form & Files
        if strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
            lim := int64(32 << 20)
            if s := project.Server.MaxUploadSize; s != "" { lim = parseSize(s) }
            _ = r.ParseMultipartForm(lim)
            if r.MultipartForm != nil {
                for k, v := range r.MultipartForm.Value { if len(v) > 0 { req[k] = v[0] } }
                files := []datasource.UploadedFile{}
                for _, fhList := range r.MultipartForm.File {
                    for _, fh := range fhList {
                        mf, err := fh.Open()
                        if err != nil { continue }
                        if rsc, ok := mf.(datasource.ReadSeekCloser); ok {
                            files = append(files, datasource.UploadedFile{Filename: fh.Filename, Content: rsc})
                            continue
                        }
                        tmp, err := os.CreateTemp("", "upload-*")
                        if err != nil { continue }
                        _, _ = io.Copy(tmp, mf)
                        _, _ = tmp.Seek(0, 0)
                        files = append(files, datasource.UploadedFile{Filename: fh.Filename, Content: tmp})
                    }
                }
                if len(files) > 0 {
                    req["file.files"] = files
                    req["file"] = map[string]any{"files": files}
                }
            }
            // Fallback: if no parsed files and body present, treat body as single file
            if req["file.files"] == nil && r.Body != nil {
                // 兼容无边界/无表单字段的 InFile 原始上传
                tmp, err := os.CreateTemp("", "upload-*")
                if err == nil {
                    _, _ = io.Copy(tmp, r.Body)
                    _, _ = tmp.Seek(0, 0)
                    files := []datasource.UploadedFile{{Filename: "infile", Content: tmp}}
                    req["file.files"] = files
                    req["file"] = map[string]any{"files": files}
                }
            }
        }
        // Global fallback: if still no files and body present with non-JSON content, treat as file
        if req["file.files"] == nil && r.Body != nil && !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
            tmp, err := os.CreateTemp("", "upload-*")
            if err == nil {
                _, _ = io.Copy(tmp, r.Body)
                _, _ = tmp.Seek(0, 0)
                files := []datasource.UploadedFile{{Filename: "infile", Content: tmp}}
                req["file.files"] = files
                req["file"] = map[string]any{"files": files}
            }
        }
        // Helpers bridge
        helpers := &core.Helpers{
            CacheSet: func(key string, val any, ttl string) error { ds.Cache["default"].Set(key, val, 0); return nil },
            CacheGet: func(key string) (any, bool) { return ds.Cache["default"].Get(key) },
            SQLQuery: func(dsName string, sqlStr string, params map[string]any) ([]map[string]any, error) {
                if t := tx[dsName]; t != nil { rows, err := t.Query(sqlStr, datasource.ValuesFrom(params)...); if err != nil { return nil, err } ; defer rows.Close(); cols, _ := rows.Columns(); res := []map[string]any{}; for rows.Next(){ vals := make([]any, len(cols)); ptrs := make([]any, len(cols)); for i := range vals { ptrs[i] = &vals[i] } ; if err := rows.Scan(ptrs...); err != nil { return nil, err } ; m := map[string]any{}; for i, c := range cols { if b, ok := vals[i].([]byte); ok { m[c] = string(b) } else { m[c] = vals[i] } } ; res = append(res, m) } ; return res, nil }
                return ds.SQLQuery(dsName, sqlStr, params)
            },
            SQLQueryOrder: func(dsName string, sqlStr string, params map[string]any, order []string) ([]map[string]any, error) {
                if t := tx[dsName]; t != nil { rows, err := t.Query(sqlStr, datasource.ValuesFromOrder(params, order)...); if err != nil { return nil, err } ; defer rows.Close(); cols, _ := rows.Columns(); res := []map[string]any{}; for rows.Next(){ vals := make([]any, len(cols)); ptrs := make([]any, len(cols)); for i := range vals { ptrs[i] = &vals[i] } ; if err := rows.Scan(ptrs...); err != nil { return nil, err } ; m := map[string]any{}; for i, c := range cols { if b, ok := vals[i].([]byte); ok { m[c] = string(b) } else { m[c] = vals[i] } } ; res = append(res, m) } ; return res, nil }
                db := ds.SQL[dsName]
                if db == nil { return nil, fmt.Errorf("SQL 数据源不存在") }
                rows, err := db.Query(sqlStr, datasource.ValuesFromOrder(params, order)...)
                if err != nil { return nil, err }
                defer rows.Close()
                cols, _ := rows.Columns()
                res := []map[string]any{}
                for rows.Next(){ vals := make([]any, len(cols)); ptrs := make([]any, len(cols)); for i := range vals { ptrs[i] = &vals[i] } ; if err := rows.Scan(ptrs...); err != nil { return nil, err } ; m := map[string]any{}; for i, c := range cols { if b, ok := vals[i].([]byte); ok { m[c] = string(b) } else { m[c] = vals[i] } } ; res = append(res, m) }
                return res, nil
            },
            SQLExec:  func(dsName string, sqlStr string, params map[string]any) error {
                if t := tx[dsName]; t != nil { _, err := t.Exec(sqlStr, datasource.ValuesFrom(params)...); return err }
                return ds.SQLExec(dsName, sqlStr, params)
            },
            SQLExecOrder: func(dsName string, sqlStr string, params map[string]any, order []string) error {
                if t := tx[dsName]; t != nil { _, err := t.Exec(sqlStr, datasource.ValuesFromOrder(params, order)...); return err }
                db := ds.SQL[dsName]
                if db == nil { return fmt.Errorf("SQL 数据源不存在") }
                _, err := db.Exec(sqlStr, datasource.ValuesFromOrder(params, order)...)
                return err
            },
            KVSet:    ds.KVSet,
            KVGet:    ds.KVGet,
            UploadSave: func(dsName string, input any, naming string) ([]string, error) { return ds.UploadSave(dsName, input, naming) },
            PluginCall: func(p, f string, params map[string]any) (any, error) { return prt.Call(p, f, params) },
            JWTSign: func(claims map[string]any) (string, error) {
                if project.Security.JWTSecret == "" { return "", fmt.Errorf("JWT 未配置") }
                header := map[string]any{"alg":"HS256","typ":"JWT"}
                hb, _ := json.Marshal(header)
                cb, _ := json.Marshal(claims)
                h := base64.RawURLEncoding.EncodeToString(hb)
                p := base64.RawURLEncoding.EncodeToString(cb)
                unsigned := h + "." + p
                mac := hmac.New(sha256.New, []byte(project.Security.JWTSecret))
                _, _ = mac.Write([]byte(unsigned))
                sig := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
                return unsigned + "." + sig, nil
            },
            BeginTx: func(dsName string) error { db := ds.SQL[dsName]; if db == nil { return fmt.Errorf("SQL 数据源不存在") } ; t, err := db.Begin(); if err != nil { return err } ; tx[dsName] = t; return nil },
            CommitTx: func(dsName string) error { t := tx[dsName]; if t == nil { return nil } ; err := t.Commit(); delete(tx, dsName); return err },
            RollbackTx: func(dsName string) error { t := tx[dsName]; if t == nil { return nil } ; err := t.Rollback(); delete(tx, dsName); return err },
            VectorSearch: func(dsName, collection string, vec []float64, topK int) ([]map[string]any, error) {
                vs := ds.Vector[dsName]
                if vs == nil { return nil, fmt.Errorf("Vector 数据源不存在") }
                res, err := vs.Search(collection, vec, topK)
                if err != nil { return nil, err }
                out := make([]map[string]any, 0, len(res))
                for _, r := range res { out = append(out, map[string]any{"id": r.ID, "score": r.Score, "meta": r.Meta}) }
                return out, nil
            },
            VectorSearchOpt: func(dsName, collection string, vec []float64, topK int, options map[string]any) ([]map[string]any, error) {
                vs := ds.Vector[dsName]
                if vs == nil { return nil, fmt.Errorf("Vector 数据源不存在") }
                type qdr interface{ SearchWithOptions(string, []float64, int, map[string]any) ([]datasource.VectorResult, error) }
                var res []datasource.VectorResult
                var err error
                if q, ok := vs.(qdr); ok { res, err = q.SearchWithOptions(collection, vec, topK, options) } else { res, err = vs.Search(collection, vec, topK) }
                if err != nil { return nil, err }
                out := make([]map[string]any, 0, len(res))
                for _, r := range res { out = append(out, map[string]any{"id": r.ID, "score": r.Score, "meta": r.Meta}) }
                return out, nil
            },
            VectorUpsert: func(dsName, collection, id string, vec []float64, meta map[string]any) error {
                vs := ds.Vector[dsName]
                if vs == nil { return fmt.Errorf("Vector 数据源不存在") }
                return vs.Upsert(collection, id, vec, meta)
            },
            VectorDelete: func(dsName, collection, id string) error {
                vs := ds.Vector[dsName]
                if vs == nil { return fmt.Errorf("Vector 数据源不存在") }
                return vs.Delete(collection, id)
            },
            VectorEnsure: func(dsName, collection string, size int, metric string) error {
                vs := ds.Vector[dsName]
                if vs == nil { return fmt.Errorf("Vector 数据源不存在") }
                return vs.EnsureCollection(collection, size, metric)
            },
            ModelApply: func(dsName, dir string) ([]string, error) {
                db := ds.SQL[dsName]
                if db == nil { return nil, fmt.Errorf("SQL 数据源不存在") }
                tabs, err := model.Load(dir)
                if err != nil { return nil, err }
                if err := model.ApplyModels(db, tabs); err != nil { return nil, err }
                names := make([]string, 0, len(tabs))
                for _, t := range tabs { names = append(names, t.Name) }
                return names, nil
            },
            DocsEndpoints: func() any { return apidocs.BuildDocs(currentInterfaces) },
            OpenAPI: func() any { return idocs.BuildOpenAPI(project.Docs.Title, apidocs.BuildDocs(currentInterfaces)) },
            MetricsText: func() string { return obs.Render() },
            BuiltinsList: func() any {
                return []map[string]any{
                    {"name":"auth.jwt","desc":"JWT token signer"},
                    {"name":"obs.metrics","desc":"Prometheus metrics text"},
                    {"name":"admin.docs","desc":"Docs endpoints list"},
                    {"name":"admin.openapi","desc":"OpenAPI spec build"},
                    {"name":"hooks.inspect","desc":"Inspect hooks chain"},
                    {"name":"upload.save","desc":"Blob upload save"},
                    {"name":"http.request","desc":"HTTP client"},
                    {"name":"sql.query","desc":"SQL query"},
                    {"name":"sql.exec","desc":"SQL exec"},
                    {"name":"cache.set","desc":"Cache set"},
                    {"name":"cache.get","desc":"Cache get"},
                    {"name":"vector.search","desc":"Vector search"},
                    {"name":"vector.upsert","desc":"Vector upsert"},
                    {"name":"vector.delete","desc":"Vector delete"},
                    {"name":"vector.ensure","desc":"Vector ensure"},
                    {"name":"model.apply","desc":"Apply SQL models"},
                    {"name":"admin.reload","desc":"Reload router"},
                    {"name":"admin.lint","desc":"Lint interfaces"},
                }
            },
            PermissionsScan: func() any { return currentPermissions },
            PluginsUsage: func() any {
                type info struct{ endpoints []string; external bool }
                usage := map[string]*info{}
                for _, ic := range currentInterfaces {
                    isAdmin := strings.ToLower(ic.Module) == "admin"
                    authType := strings.ToLower(ic.Auth)
                    if authType == "" { authType = strings.ToLower(project.Security.AuthType) }
                    hasPaid := false
                    for _, l := range ic.Labels { if strings.ToLower(l) == "paid" { hasPaid = true; break } }
                    ext := (!isAdmin) && (authType == "apikey" || hasPaid)
                    for _, st := range ic.Steps {
                        if _, ok := st["plugin.call"]; ok {
                            pm := st["plugin.call"].(map[string]any)
                            pn := fmt.Sprint(pm["plugin"])
                            key := strings.ToUpper(ic.Method)+" "+ic.Path
                            if usage[pn] == nil { usage[pn] = &info{} }
                            usage[pn].endpoints = append(usage[pn].endpoints, key)
                            if ext { usage[pn].external = true }
                        }
                    }
                }
                st := map[string]any{"external": []map[string]any{}, "internal": []map[string]any{}}
                var pm map[string]any
                if prt != nil { if m, ok := prt.Status()["plugins"].(map[string]any); ok { pm = m } }
                for name, in := range usage {
                    item := map[string]any{"name": name, "endpoints": in.endpoints}
                    if pm != nil { if mv, ok := pm[name].(map[string]any); ok { item["status"] = mv } }
                    if in.external { st["external"] = append(st["external"].([]map[string]any), item) } else { st["internal"] = append(st["internal"].([]map[string]any), item) }
                }
                return st
            },
            AdminReload: func() (map[string]any, error) {
                newProject, err := config.LoadProjectConfig("configs/project.yaml")
                if err != nil { return nil, err }
                newInterfaces, err := config.LoadInterfaceConfigs("configs/interfaces")
                if err != nil { return nil, err }
                if hb, err := config.LoadHooksConfig("configs/hooks"); err == nil && (len(hb.Auth)+len(hb.Before)+len(hb.After) > 0) {
                    apphooks.RegisterFromConfigs(newProject, hb)
                }
                newRouter, err := BuildRouter(baseCtx, newProject, newInterfaces)
                if err != nil { return nil, err }
                if globalHolder != nil { globalHolder.Set(newRouter) }
                oldMap := map[string]struct{}{}
                newMap := map[string]struct{}{}
                for _, o := range currentInterfaces { oldMap[o.Method+" "+o.Path] = struct{}{} }
                for _, n := range newInterfaces { newMap[n.Method+" "+n.Path] = struct{}{} }
                added := []string{}
                removed := []string{}
                unchanged := []string{}
                for k := range newMap { if _, ok := oldMap[k]; !ok { added = append(added, k) } else { unchanged = append(unchanged, k) } }
                for k := range oldMap { if _, ok := newMap[k]; !ok { removed = append(removed, k) } }
                currentInterfaces = newInterfaces
                interfaceIndex = make(map[string]*config.InterfaceConfig, len(newInterfaces))
                for _, x := range newInterfaces { interfaceIndex[strings.ToUpper(x.Method)+" "+strings.TrimSuffix(x.Path, "/")] = x }
                return map[string]any{"Added": added, "Removed": removed, "Unchanged": unchanged}, nil
            },
            HooksInspect: func(method string, path string) (any, error) {
                m := strings.ToUpper(method)
                p := strings.TrimSuffix(path, "/")
                var target *config.InterfaceConfig
                if m != "" { target = interfaceIndex[m+" "+p] }
                if target == nil { target = interfaceIndex["GET "+p] }
                if target == nil { target = interfaceIndex["POST "+p] }
                if target == nil { return nil, fmt.Errorf("interface not found") }
                return apphooks.Inspect(target), nil
            },
            LintInterfaces: func() any {
                return config.LintInterfaces(currentInterfaces)
            },
            PluginsStatus: func() any {
                if prt == nil { return map[string]any{"enabled": false} }
                return prt.Status()
            },
            PluginsControl: func(action string, names []string, enabled bool) error {
                if prt == nil { return fmt.Errorf("插件未启用") }
                for _, n := range names {
                    switch action {
                    case "enable":
                        prt.SetEnabled(n, true)
                    case "disable":
                        prt.SetEnabled(n, false)
                    default:
                        prt.SetEnabled(n, enabled)
                    }
                }
                return nil
            },
            PluginsAdd: func(item config.PluginRegistryItem) error {
                if prt == nil { return fmt.Errorf("插件未启用") }
                return prt.AddPlugin(item)
            },
            PluginsRemove: func(name string) error {
                if prt == nil { return fmt.Errorf("插件未启用") }
                return prt.RemovePlugin(name)
            },
            PluginsRestart: func(name string) error {
                if prt == nil { return fmt.Errorf("插件未启用") }
                return prt.RestartPlugin(name)
            },
            PluginsStop: func(name string) error {
                if prt == nil { return fmt.Errorf("插件未启用") }
                return prt.StopPlugin(name)
            },
            PluginsStart: func(name string) error {
                if prt == nil { return fmt.Errorf("插件未启用") }
                return prt.StartPlugin(name)
            },
        }
        // 事务包装 / Transaction wrap
        stepsToRun := ic.Steps
        var txDS string
        if ic.Transaction != nil {
            if dsName, ok := ic.Transaction["ds"].(string); ok && dsName != "" {
                txDS = dsName
                begin := map[string]any{"transaction": map[string]any{"action": "begin", "ds": txDS}}
                commit := map[string]any{"transaction": map[string]any{"action": "commit", "ds": txDS}}
                stepsToRun = append([]map[string]any{begin}, stepsToRun...)
                stepsToRun = append(stepsToRun, commit)
            }
        }
        if ext.GetHooks().Before != nil {
            if err := ext.GetHooks().Before(project, ic, ds, prt, r); err != nil {
                w.WriteHeader(http.StatusForbidden)
                _ = json.NewEncoder(w).Encode(map[string]any{"error": err.Error()})
                return
            }
        }
        if cfg := config.GetCustom(); cfg != nil { ctx.Vars["custom"] = cfg }
        // 执行 / Execute
        res, err := core.ExecuteSteps(ctx, stepsToRun, req, helpers)
        if err != nil {
            if txDS != "" { _ = helpers.RollbackTx(txDS) }
            if ic.Errors != nil {
                if def, ok := ic.Errors["default"].(map[string]any); ok {
                    st := core.IntFromAny(def["status"], 500)
                    w.WriteHeader(st)
                    _ = json.NewEncoder(w).Encode(def["body"])
                    return
                }
            }
            // 特例：未生成响应视为成功但返回调试信息 / treat no-response as success
            msg := err.Error()
            if strings.Contains(msg, "未生成响应") || strings.Contains(strings.ToLower(msg), "no response") {
                w.WriteHeader(http.StatusOK)
                body := map[string]any{"code": 0, "msg": "ok"}
                if logs, ok := ctx.Vars["sql_log"]; ok { body["sql"] = logs }
                _ = json.NewEncoder(w).Encode(body)
                return
            }
            w.WriteHeader(http.StatusInternalServerError)
            _ = json.NewEncoder(w).Encode(map[string]any{"error": msg})
            return
        }
        for k, v := range res.Headers { w.Header().Set(k, v) }
        w.WriteHeader(res.Status)
        switch b := res.Body.(type) {
        case string:
            _, _ = io.WriteString(w, b)
        default:
            _ = json.NewEncoder(w).Encode(b)
        }
        dur := time.Since(start)
        if ext.GetHooks().After != nil {
            ext.GetHooks().After(project, ic, ds, prt, r, res, dur)
        }
    }
}

func parseSize(s string) int64 {
    t := strings.ToLower(strings.TrimSpace(s))
    if strings.HasSuffix(t, "kb") { return int64(float64(toInt(strings.TrimSuffix(t, "kb"))) * 1024) }
    if strings.HasSuffix(t, "mb") { return int64(float64(toInt(strings.TrimSuffix(t, "mb"))) * 1024 * 1024) }
    if strings.HasSuffix(t, "b") { return int64(toInt(strings.TrimSuffix(t, "b"))) }
    return int64(toInt(t))
}

type recordWriter struct{ h http.Header; b bytes.Buffer; s int; w http.ResponseWriter }
func (rw *recordWriter) Header() http.Header { if rw.h == nil { rw.h = http.Header{} } ; return rw.h }
func (rw *recordWriter) Write(b []byte) (int, error) { return rw.b.Write(b) }
func (rw *recordWriter) WriteHeader(statusCode int) { rw.s = statusCode }
func toInt(s string) int { n := 0; for i := 0; i < len(s); i++ { c := s[i]; if c >= '0' && c <= '9' { n = n*10 + int(c-'0') } } ; return n }
func toFloatSafe(v any) float64 { switch t := v.(type) { case float64: return t; case float32: return float64(t); case int: return float64(t); case int64: return float64(t); case string: f, _ := strconv.ParseFloat(t, 64); return f; default: return 0 } }
var globalHolder *HandlerHolder
func SetGlobalHolder(h *HandlerHolder) { globalHolder = h }
var currentInterfaces []*config.InterfaceConfig
var interfaceIndex map[string]*config.InterfaceConfig
var currentPermissions []map[string]any