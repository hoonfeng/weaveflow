package core

import (
    "errors"
    "fmt"
    "regexp"
    "strings"
    "io"
    "net/http"
    "net/url"
    "path/filepath"
    "ifaceconf/internal/datasource"
    "ifaceconf/internal/config"
    "time"
    "encoding/json"
    "crypto/sha256"
    "crypto/hmac"
    "encoding/hex"
    "encoding/base64"
    "crypto/rand"
    "strconv"
    "math"
    "sort"
)

// 步骤执行结果 / Execution result
type ExecResult struct {
    Status  int
    Headers map[string]string
    Body    any
}

// 执行步骤列表 / Execute steps list
func ExecuteSteps(ctx *RequestContext, steps []map[string]any, req map[string]any, helpers *Helpers) (*ExecResult, error) {
    var result *ExecResult
    for _, step := range steps {
        for k, v := range step {
            switch k {
            case "response":
                m := v.(map[string]any)
                status := intFrom(m["status"], 200)
                headers := mapFromString(m["headers"]) 
                body := m["body"]
                if wm, ok := m["wrap"].(map[string]any); ok {
                    code := evalTemplate(req, ctx, wm["code"])
                    msg := evalTemplate(req, ctx, wm["msg"]) 
                    data := evalTemplate(req, ctx, wm["data"])
                    if dm, ok := data.(map[string]any); ok { data = evalMap(req, ctx, dm) }
                    if data == nil { data = evalTemplate(req, ctx, body) }
                    body = map[string]any{"code": code, "msg": msg, "data": data}
                } else {
                    bb := evalTemplate(req, ctx, body)
                    if bm, ok := bb.(map[string]any); ok { body = evalMap(req, ctx, bm) } else { body = bb }
                }
                result = &ExecResult{Status: status, Headers: headers, Body: body}
            case "transform":
                m := v.(map[string]any)
                if mapping, ok := m["mapping"].(map[string]any); ok {
                    for mk, mv := range mapping {
                        ctx.Vars[mk] = evalTemplate(req, ctx, mv)
                    }
                }
            case "validate":
                m := v.(map[string]any)
                target := stringFrom(m["target"], "")
                // 尝试从 req 与 ctx 中读取目标值 / read target value
                val, ok := getByPath(req, target)
                if !ok { val, ok = getByPath(ctx.Vars, target) }
                // required
                if boolFrom(m["required"]) && (val == nil || (isEmptyString(val))) {
                    return nil, fmt.Errorf("参数缺失或为空: %s", target)
                }
                // 文件列表校验
                if strings.HasPrefix(target, "file") {
                    files := []datasource.UploadedFile{}
                    if list, ok := val.([]datasource.UploadedFile); ok { files = list } else if arr, ok := val.([]any); ok { files = toUploadedFiles(arr) }
                    if len(files) == 0 { return nil, fmt.Errorf("参数缺失或为空: %s", target) }
                    if s := stringFrom(m["maxSize"], ""); s != "" {
                        lim := parseSize(s)
                        for _, f := range files { if sizeOf(f.Content) > lim { return nil, fmt.Errorf("文件过大: %s", f.Filename) } }
                    }
                    if ts, ok := m["types"].([]any); ok && len(ts) > 0 {
                        allow := toStringSet(ts)
                        for _, f := range files {
                            ct := sniffContentType(f.Content)
                            ext := strings.ToLower(filepath.Ext(f.Filename))
                            if _, ok := allow[ct]; ok { continue }
                            if _, ok := allow[ext]; ok { continue }
                            return nil, fmt.Errorf("文件类型不允许: %s", f.Filename)
                        }
                    }
                }
                // type
                if t := stringFrom(m["type"], ""); t != "" {
                    if !typeMatches(t, val) { return nil, fmt.Errorf("类型不匹配: %s 期望 %s", target, t) }
                }
                // string constraints
                if s, ok := val.(string); ok {
                    if mi := intFrom(m["minLen"], -1); mi >= 0 && len(s) < mi { return nil, fmt.Errorf("长度过短: %s", target) }
                    if ma := intFrom(m["maxLen"], -1); ma >= 0 && len(s) > ma { return nil, fmt.Errorf("长度过长: %s", target) }
                    if rg := stringFrom(m["regex"], ""); rg != "" { r, _ := regexp.Compile(rg); if !r.MatchString(s) { return nil, fmt.Errorf("正则不匹配: %s", target) } }
                }
                // numeric
                if isNumber(val) {
                    fv := toFloat(val)
                    if mi := floatFrom(m["min"], -1e308); fv < mi { return nil, fmt.Errorf("数值过小: %s", target) }
                    if ma := floatFrom(m["max"], 1e308); fv > ma { return nil, fmt.Errorf("数值过大: %s", target) }
                }
                // enum
                if en, ok := m["enum"].([]any); ok && len(en) > 0 {
                    matched := false
                    for _, ev := range en { if fmt.Sprint(ev) == fmt.Sprint(val) { matched = true; break } }
                    if !matched { return nil, fmt.Errorf("不在枚举集合: %s", target) }
                }
            case "cache.set":
                m := v.(map[string]any)
                key := fmt.Sprint(evalTemplate(req, ctx, m["key"]))
                val := evalTemplate(req, ctx, m["value"]) 
                ttl := stringFrom(m["ttl"], "")
                if helpers != nil && helpers.CacheSet != nil {
                    if err := helpers.CacheSet(key, val, ttl); err != nil { return nil, err }
                }
            case "cache.get":
                m := v.(map[string]any)
                key := fmt.Sprint(evalTemplate(req, ctx, m["key"]))
                out := stringFrom(m["out"], "")
                if helpers != nil && helpers.CacheGet != nil {
                    val, _ := helpers.CacheGet(key)
                    if out != "" { ctx.Vars[out] = val }
                }
            case "sql.query":
                m := v.(map[string]any)
                ds := stringFrom(m["ds"], "")
                // 支持模板渲染，允许在 SQL 文本中拼接分页等数值
                sqlAny := evalTemplate(req, ctx, m["sql"]) 
                sql := fmt.Sprint(sqlAny)
                params := mapToParams(req, ctx, m["params"]) 
                out := stringFrom(m["out"], "")
                order := toStringSlice(evalTemplate(req, ctx, m["order"]))
                {
                    entry := map[string]any{"type": "query", "sql": sql, "params": params}
                    if len(order) > 0 { entry["order"] = order }
                    if arr, ok := ctx.Vars["sql_log"].([]map[string]any); ok { ctx.Vars["sql_log"] = append(arr, entry) } else { ctx.Vars["sql_log"] = []map[string]any{entry} }
                }
                if helpers != nil {
                    var rows []map[string]any
                    var err error
                    if len(order) > 0 && helpers.SQLQueryOrder != nil {
                        rows, err = helpers.SQLQueryOrder(ds, sql, params, order)
                    } else if helpers.SQLQuery != nil {
                        rows, err = helpers.SQLQuery(ds, sql, params)
                    }
                    if err != nil { return nil, err }
                    if out != "" { ctx.Vars[out] = rows }
                }
            case "sql.exec":
                m := v.(map[string]any)
                ds := stringFrom(m["ds"], "")
                sqlAny := evalTemplate(req, ctx, m["sql"]) 
                sql := fmt.Sprint(sqlAny)
                params := mapToParams(req, ctx, m["params"]) 
                order := toStringSlice(evalTemplate(req, ctx, m["order"]))
                {
                    entry := map[string]any{"type": "exec", "sql": sql, "params": params}
                    if len(order) > 0 { entry["order"] = order }
                    if arr, ok := ctx.Vars["sql_log"].([]map[string]any); ok { ctx.Vars["sql_log"] = append(arr, entry) } else { ctx.Vars["sql_log"] = []map[string]any{entry} }
                }
                if helpers != nil {
                    var err error
                    if len(order) > 0 && helpers.SQLExecOrder != nil {
                        err = helpers.SQLExecOrder(ds, sql, params, order)
                    } else if helpers.SQLExec != nil {
                        err = helpers.SQLExec(ds, sql, params)
                    }
                    if err != nil { return nil, err }
                }
            case "auth.jwt":
                m := v.(map[string]any)
                out := stringFrom(m["out"], "token")
                claims := map[string]any{}
                if cm, ok := m["claims"].(map[string]any); ok { for k, vv := range cm { claims[k] = evalTemplate(req, ctx, vv) } }
                if rs, ok := m["roles"].([]any); ok { claims["roles"] = rs }
                if helpers != nil && helpers.JWTSign != nil {
                    tok, err := helpers.JWTSign(claims)
                    if err != nil { return nil, err }
                    if out != "" { ctx.Vars[out] = tok }
                }
            case "kv.set":
                m := v.(map[string]any)
                ds := stringFrom(m["ds"], "")
                key := fmt.Sprint(evalTemplate(req, ctx, m["key"]))
                val := evalTemplate(req, ctx, m["value"]) 
                if helpers != nil && helpers.KVSet != nil {
                    if err := helpers.KVSet(ds, key, val); err != nil { return nil, err }
                }
            case "kv.get":
                m := v.(map[string]any)
                ds := stringFrom(m["ds"], "")
                key := fmt.Sprint(evalTemplate(req, ctx, m["key"]))
                out := stringFrom(m["out"], "")
                if helpers != nil && helpers.KVGet != nil {
                    val, err := helpers.KVGet(ds, key)
                    if err != nil { return nil, err }
                    if out != "" { ctx.Vars[out] = val }
                }
            case "http.request":
                m := v.(map[string]any)
                method := strings.ToUpper(stringFrom(m["method"], "GET"))
                url := stringFrom(m["url"], "")
                headers := mapFromString(m["headers"]) 
                bodyVal := evalTemplate(req, ctx, m["body"]) 
                out := stringFrom(m["out"], "resp")
                retry := intFrom(m["retry"], 0)
                backoff := intFrom(m["backoff"], 0)
                timeoutMs := intFrom(m["timeoutMs"], 5000)
                var hasCircuit bool
                var circuitKey string
                var threshold int
                var openMs int
                var fallback any
                if cm, ok := m["circuit"].(map[string]any); ok {
                    hasCircuit = true
                    circuitKey = stringFrom(cm["key"], url)
                    threshold = intFrom(cm["threshold"], 3)
                    openMs = intFrom(cm["openMs"], 30000)
                    fallback = cm["fallback"]
                }
                var bodyStr string
                if s, ok := bodyVal.(string); ok { bodyStr = s } else if bodyVal != nil { b, _ := json.Marshal(bodyVal); bodyStr = string(b) }
                if hasCircuit && isCircuitOpen(circuitKey) {
                    fb := evalTemplate(req, ctx, fallback)
                    if mm, ok := fb.(map[string]any); ok { fb = evalMap(req, ctx, mm) }
                    if out != "" { ctx.Vars[out] = fb }
                    break
                }
                var lastErr error
                for attempt := 0; attempt <= retry; attempt++ {
                    var hreq *http.Request
                    if method == "GET" { hreq, _ = http.NewRequest(method, url, nil) } else { hreq, _ = http.NewRequest(method, url, strings.NewReader(bodyStr)) }
                    for k, v := range headers { hreq.Header.Set(k, v) }
                    cli := &http.Client{ Timeout: time.Duration(timeoutMs) * time.Millisecond }
                    resp, err := cli.Do(hreq)
                    if err == nil && resp != nil {
                        defer resp.Body.Close()
                        buf := new(strings.Builder)
                        _, _ = io.Copy(buf, resp.Body)
                        if resp.StatusCode >= 500 {
                            lastErr = fmt.Errorf("http %d", resp.StatusCode)
                        } else {
                            ctx.Vars[out] = map[string]any{"status": resp.StatusCode, "body": buf.String()}
                            resetCircuit(circuitKey)
                            lastErr = nil
                            break
                        }
                    } else {
                        lastErr = err
                    }
                    if attempt < retry && backoff > 0 { time.Sleep(time.Duration(backoff) * time.Millisecond) }
                }
                if lastErr != nil {
                    if hasCircuit { recordCircuitFail(circuitKey, threshold, openMs) }
                    if fallback != nil {
                        fb := evalTemplate(req, ctx, fallback)
                        if mm, ok := fb.(map[string]any); ok { fb = evalMap(req, ctx, mm) }
                        if out != "" { ctx.Vars[out] = fb }
                        lastErr = nil
                    }
                    if lastErr != nil { return nil, lastErr }
                }
            case "upload.save":
                m := v.(map[string]any)
                ds := stringFrom(m["ds"], "")
                input := evalTemplate(req, ctx, m["input"]) 
                if input == nil {
                    if v, ok := req["file.files"]; ok { input = v }
                }
                naming := stringFrom(m["naming"], "sha256")
                out := stringFrom(m["out"], "")
                if helpers != nil && helpers.UploadSave != nil {
                    ids, err := helpers.UploadSave(ds, input, naming)
                    if err != nil { return nil, err }
                    if out != "" { ctx.Vars[out] = ids }
                }
            case "admin.reload":
                m := v.(map[string]any)
                out := stringFrom(m["out"], "reload")
                if helpers != nil && helpers.AdminReload != nil {
                    diff, err := helpers.AdminReload()
                    if err != nil { return nil, err }
                    if out != "" { ctx.Vars[out] = diff }
                }
            case "plugins.status":
                m := v.(map[string]any)
                out := stringFrom(m["out"], "plugins")
                if helpers != nil && helpers.PluginsStatus != nil {
                    res := helpers.PluginsStatus()
                    if out != "" { ctx.Vars[out] = res }
                }
            case "plugins.control":
                m := v.(map[string]any)
                action := stringFrom(m["action"], "")
                namesAny := evalTemplate(req, ctx, m["names"])
                enabled := boolFrom(m["enabled"]) 
                names := []string{}
                if arr, ok := namesAny.([]any); ok { for _, it := range arr { names = append(names, fmt.Sprint(it)) } }
                if helpers != nil && helpers.PluginsControl != nil {
                    if err := helpers.PluginsControl(action, names, enabled); err != nil { return nil, err }
                }
            case "plugins.add":
                m := v.(map[string]any)
                var item config.PluginRegistryItem
                item.Name = stringFrom(m["name"], "")
                if em, ok := m["executable"].(map[string]any); ok {
                    item.Executable.Windows = fmt.Sprint(evalTemplate(req, ctx, em["windows"]))
                    item.Executable.Unix = fmt.Sprint(evalTemplate(req, ctx, em["unix"]))
                }
                item.Instances = intFrom(m["instances"], 1)
                item.Timeout = stringFrom(m["timeout"], "300ms")
                item.QueueSize = intFrom(m["queueSize"], 1024)
                if fs, ok := evalTemplate(req, ctx, m["functions"]).([]any); ok { for _, f := range fs { item.Functions = append(item.Functions, fmt.Sprint(f)) } }
                if helpers != nil && helpers.PluginsAdd != nil { if err := helpers.PluginsAdd(item); err != nil { return nil, err } }
            case "plugins.remove":
                m := v.(map[string]any)
                namesAny := evalTemplate(req, ctx, m["names"]) 
                names := []string{}
                if arr, ok := namesAny.([]any); ok { for _, it := range arr { names = append(names, fmt.Sprint(it)) } }
                if helpers != nil && helpers.PluginsRemove != nil {
                    for _, n := range names { if err := helpers.PluginsRemove(n); err != nil { return nil, err } }
                }
            case "plugins.restart":
                m := v.(map[string]any)
                namesAny := evalTemplate(req, ctx, m["names"]) 
                names := []string{}
                if arr, ok := namesAny.([]any); ok { for _, it := range arr { names = append(names, fmt.Sprint(it)) } }
                if helpers != nil && helpers.PluginsRestart != nil {
                    for _, n := range names { if err := helpers.PluginsRestart(n); err != nil { return nil, err } }
                }
            case "plugins.stop":
                m := v.(map[string]any)
                namesAny := evalTemplate(req, ctx, m["names"]) 
                names := []string{}
                if arr, ok := namesAny.([]any); ok { for _, it := range arr { names = append(names, fmt.Sprint(it)) } }
                if helpers != nil && helpers.PluginsStop != nil {
                    for _, n := range names { if err := helpers.PluginsStop(n); err != nil { return nil, err } }
                }
            case "plugins.start":
                m := v.(map[string]any)
                namesAny := evalTemplate(req, ctx, m["names"]) 
                names := []string{}
                if arr, ok := namesAny.([]any); ok { for _, it := range arr { names = append(names, fmt.Sprint(it)) } }
                if helpers != nil && helpers.PluginsStart != nil {
                    for _, n := range names { if err := helpers.PluginsStart(n); err != nil { return nil, err } }
                }
            case "hooks.inspect":
                m := v.(map[string]any)
                mv := evalTemplate(req, ctx, m["method"]) 
                pv := evalTemplate(req, ctx, m["path"]) 
                method := fmt.Sprint(mv)
                path := fmt.Sprint(pv)
                out := stringFrom(m["out"], "inspect")
                if helpers != nil && helpers.HooksInspect != nil {
                    res, err := helpers.HooksInspect(method, path)
                    if err != nil { return nil, err }
                    if out != "" { ctx.Vars[out] = res }
                }
            case "admin.lint":
                m := v.(map[string]any)
                out := stringFrom(m["out"], "lint")
                if helpers != nil && helpers.LintInterfaces != nil {
                    res := helpers.LintInterfaces()
                    if out != "" { ctx.Vars[out] = res }
                }
            case "admin.docs":
                m := v.(map[string]any)
                out := stringFrom(m["out"], "docs")
                if helpers != nil && helpers.DocsEndpoints != nil {
                    res := helpers.DocsEndpoints()
                    if out != "" { ctx.Vars[out] = res }
                }
            case "admin.openapi":
                m := v.(map[string]any)
                out := stringFrom(m["out"], "openapi")
                if helpers != nil && helpers.OpenAPI != nil {
                    spec := helpers.OpenAPI()
                    if out != "" { ctx.Vars[out] = spec }
                }
            case "admin.builtin":
                m := v.(map[string]any)
                out := stringFrom(m["out"], "builtin")
                if helpers != nil && helpers.BuiltinsList != nil {
                    list := helpers.BuiltinsList()
                    if out != "" { ctx.Vars[out] = list }
                }
            case "admin.plugins_usage":
                m := v.(map[string]any)
                out := stringFrom(m["out"], "usage")
                if helpers != nil && helpers.PluginsUsage != nil {
                    usage := helpers.PluginsUsage()
                    if out != "" { ctx.Vars[out] = usage }
                }
            case "admin.permissions_scan":
                m := v.(map[string]any)
                out := stringFrom(m["out"], "perms")
                if helpers != nil && helpers.PermissionsScan != nil {
                    list := helpers.PermissionsScan()
                    if out != "" { ctx.Vars[out] = list }
                }
            case "model.apply":
                m := v.(map[string]any)
                ds := stringFrom(m["ds"], "main")
                dir := stringFrom(m["dir"], "configs/models")
                out := stringFrom(m["out"], "models")
                if helpers != nil && helpers.ModelApply != nil {
                    names, err := helpers.ModelApply(ds, dir)
                    if err != nil { return nil, err }
                    if out != "" { ctx.Vars[out] = names }
                }
            case "obs.metrics":
                m := v.(map[string]any)
                out := stringFrom(m["out"], "metrics")
                if helpers != nil && helpers.MetricsText != nil {
                    txt := helpers.MetricsText()
                    if out != "" { ctx.Vars[out] = txt }
                }
            case "plugin.call":
                m := v.(map[string]any)
                plugin := stringFrom(m["plugin"], "")
                function := stringFrom(m["function"], "")
                params := mapToParams(req, ctx, m["params"]) 
                out := stringFrom(m["out"], "")
                retry := intFrom(m["retry"], 0)
                backoffMs := intFrom(m["backoff"], 0)
                var hasCircuit bool
                var circuitKey string
                var threshold int
                var openMs int
                var fallback any
                if cm, ok := m["circuit"].(map[string]any); ok {
                    hasCircuit = true
                    circuitKey = stringFrom(cm["key"], plugin+"."+function)
                    threshold = intFrom(cm["threshold"], 3)
                    openMs = intFrom(cm["openMs"], 30000)
                    fallback = cm["fallback"]
                }
                if helpers != nil && helpers.PluginCall != nil {
                    if hasCircuit && isCircuitOpen(circuitKey) {
                        if out != "" { ctx.Vars[out] = evalTemplate(req, ctx, fallback) }
                        break
                    }
                    var res any
                    var err error
                    for attempt := 0; attempt <= retry; attempt++ {
                        res, err = helpers.PluginCall(plugin, function, params)
                        if err == nil { break }
                        if attempt < retry && backoffMs > 0 { time.Sleep(time.Duration(backoffMs) * time.Millisecond) }
                    }
                    if err != nil {
                        if hasCircuit { recordCircuitFail(circuitKey, threshold, openMs) }
                        if fallback != nil {
                            if out != "" { ctx.Vars[out] = evalTemplate(req, ctx, fallback) }
                            err = nil
                        }
                        if err != nil { return nil, err }
                    } else {
                        if hasCircuit { resetCircuit(circuitKey) }
                    }
                    if out != "" { ctx.Vars[out] = res }
                }
            case "branch":
                m := v.(map[string]any)
                cond := evalTemplate(req, ctx, m["if"]) 
                var path []map[string]any
                if truthy(cond) {
                    if then, ok := m["then"].([]any); ok { path = toStepList(then) }
                } else {
                    if els, ok := m["else"].([]any); ok { path = toStepList(els) }
                }
                if len(path) > 0 {
                    r, err := ExecuteSteps(ctx, path, req, helpers)
                    if err != nil {
                        msg := fmt.Sprint(err)
                        if !(strings.Contains(msg, "未生成响应") || strings.Contains(msg, "no response")) { return nil, err }
                    }
                    result = r
                }
            case "loop":
                m := v.(map[string]any)
                var items []any
                x := evalTemplate(req, ctx, m["items"])
                switch t := x.(type) {
                case []any:
                    items = t
                case []map[string]any:
                    items = make([]any, 0, len(t))
                    for _, mm := range t { items = append(items, mm) }
                default:
                    items = []any{}
                }
                doSteps := toStepList(m["do"].([]any))
                varName := stringFrom(m["var"], "item")
                for _, it := range items {
                    ctx.Vars[varName] = it
                    r, err := ExecuteSteps(ctx, doSteps, req, helpers)
                    if err != nil {
                        msg := fmt.Sprint(err)
                        if strings.Contains(msg, "未生成响应") || strings.Contains(strings.ToLower(msg), "no response") {
                            continue
                        }
                        return nil, err
                    }
                    result = r
                }
            case "vector.search":
                m := v.(map[string]any)
                ds := stringFrom(m["ds"], "")
                collection := stringFrom(m["collection"], "default")
                vec := toFloatSlice(evalTemplate(req, ctx, m["vec"]))
                topK := intFrom(m["topK"], 10)
                out := stringFrom(m["out"], "")
                var options map[string]any
                if ov := evalTemplate(req, ctx, m["options"]); ov != nil {
                    if om, ok := ov.(map[string]any); ok { options = om }
                }
                if helpers != nil {
                    var res []map[string]any
                    var err error
                    if options != nil && helpers.VectorSearchOpt != nil {
                        res, err = helpers.VectorSearchOpt(ds, collection, vec, topK, options)
                    } else if helpers.VectorSearch != nil {
                        res, err = helpers.VectorSearch(ds, collection, vec, topK)
                    }
                    if err != nil { return nil, err }
                    if out != "" { ctx.Vars[out] = res }
                }
            case "vector.upsert":
                m := v.(map[string]any)
                ds := stringFrom(m["ds"], "")
                collection := stringFrom(m["collection"], "default")
                id := fmt.Sprint(evalTemplate(req, ctx, m["id"]))
                vec := toFloatSlice(evalTemplate(req, ctx, m["vec"]))
                meta, _ := evalTemplate(req, ctx, m["meta"]).(map[string]any)
                if helpers != nil && helpers.VectorUpsert != nil {
                    if err := helpers.VectorUpsert(ds, collection, id, vec, meta); err != nil { return nil, err }
                }
            case "vector.delete":
                m := v.(map[string]any)
                ds := stringFrom(m["ds"], "")
                collection := stringFrom(m["collection"], "default")
                id := fmt.Sprint(evalTemplate(req, ctx, m["id"]))
                if helpers != nil && helpers.VectorDelete != nil {
                    if err := helpers.VectorDelete(ds, collection, id); err != nil { return nil, err }
                }
            case "vector.ensure":
                m := v.(map[string]any)
                ds := stringFrom(m["ds"], "")
                collection := stringFrom(m["collection"], "default")
                size := intFrom(m["size"], 0)
                metric := stringFrom(m["metric"], "Cosine")
                if helpers != nil && helpers.VectorEnsure != nil {
                    if err := helpers.VectorEnsure(ds, collection, size, metric); err != nil { return nil, err }
                }
            case "vector.upsert_batch":
                m := v.(map[string]any)
                ds := stringFrom(m["ds"], "")
                collection := stringFrom(m["collection"], "default")
                items, _ := evalTemplate(req, ctx, m["items"]).([]any)
                if helpers != nil && helpers.VectorUpsert != nil {
                    for _, it := range items {
                        if mp, ok := it.(map[string]any); ok {
                            id := fmt.Sprint(evalTemplate(req, ctx, mp["id"]))
                            vec := toFloatSlice(evalTemplate(req, ctx, mp["vec"]))
                            meta, _ := evalTemplate(req, ctx, mp["meta"]).(map[string]any)
                            if err := helpers.VectorUpsert(ds, collection, id, vec, meta); err != nil { return nil, err }
                        }
                    }
                }
            case "transaction":
                m := v.(map[string]any)
                action := stringFrom(m["action"], "")
                ds := stringFrom(m["ds"], "")
                switch action {
                case "begin":
                    if helpers != nil && helpers.BeginTx != nil { if err := helpers.BeginTx(ds); err != nil { return nil, err } }
                case "commit":
                    if helpers != nil && helpers.CommitTx != nil { if err := helpers.CommitTx(ds); err != nil { return nil, err } }
                case "rollback":
                    if helpers != nil && helpers.RollbackTx != nil { if err := helpers.RollbackTx(ds); err != nil { return nil, err } }
                default:
                    return nil, fmt.Errorf("未知事务动作: %s", action)
                }
            default:
                return nil, fmt.Errorf("未知步骤 / unknown step: %s", k)
            }
        }
    }
    if result == nil {
        return nil, errors.New("未生成响应 / no response step executed")
    }
    return result, nil
}

// 助手接口聚合 / Helpers interfaces aggregation
    type Helpers struct {
        CacheSet   func(key string, val any, ttl string) error
        CacheGet   func(key string) (any, bool)
        SQLQuery   func(ds string, sql string, params map[string]any) ([]map[string]any, error)
        SQLQueryOrder func(ds string, sql string, params map[string]any, order []string) ([]map[string]any, error)
        SQLExec    func(ds string, sql string, params map[string]any) error
        SQLExecOrder func(ds string, sql string, params map[string]any, order []string) error
        KVSet      func(ds string, key string, val any) error
        KVGet      func(ds string, key string) (any, error)
        UploadSave func(ds string, input any, naming string) ([]string, error)
        PluginCall func(plugin string, function string, params map[string]any) (any, error)
        JWTSign    func(claims map[string]any) (string, error)
        BeginTx    func(ds string) error
        CommitTx   func(ds string) error
        RollbackTx func(ds string) error
        VectorSearch func(ds string, collection string, vec []float64, topK int) ([]map[string]any, error)
        VectorSearchOpt func(ds string, collection string, vec []float64, topK int, options map[string]any) ([]map[string]any, error)
        VectorUpsert func(ds string, collection string, id string, vec []float64, meta map[string]any) error
        VectorDelete func(ds string, collection string, id string) error
        VectorEnsure func(ds string, collection string, size int, metric string) error
        ModelApply func(ds string, dir string) ([]string, error)
        AdminReload func() (map[string]any, error)
        HooksInspect func(method string, path string) (any, error)
        LintInterfaces func() any
        PluginsStatus func() any
        PluginsControl func(action string, names []string, enabled bool) error
        PluginsAdd func(item config.PluginRegistryItem) error
        PluginsRemove func(name string) error
        PluginsRestart func(name string) error
        PluginsStop func(name string) error
        PluginsStart func(name string) error
        DocsEndpoints func() any
        OpenAPI func() any
        MetricsText func() string
        BuiltinsList func() any
        PluginsUsage func() any
        PermissionsScan func() any
    }

// 简易模板求值 / Simple template evaluation (placeholder)
// parseTemplateExpression 解析模板表达式，支持嵌套函数调用
func parseTemplateExpression(expr string) (string, []string) {
	expr = strings.TrimSpace(expr)
	
	// 查找第一个空格，分离函数名和参数部分
	firstSpace := strings.Index(expr, " ")
	if firstSpace == -1 {
		// 没有参数，可能是变量名
		return "", []string{expr}
	}
	
	fn := expr[:firstSpace]
	argsStr := strings.TrimSpace(expr[firstSpace+1:])
	
	if argsStr == "" {
		return fn, []string{}
	}
	
	// 解析参数，支持嵌套括号
	var args []string
	var currentArg strings.Builder
	parenDepth := 0
	inString := false
	stringChar := rune(0)
	
	for i, ch := range argsStr {
		if !inString {
			switch ch {
			case '"', '\'':
				inString = true
				stringChar = ch
				currentArg.WriteRune(ch)
			case '(':
				parenDepth++
				currentArg.WriteRune(ch)
			case ')':
				parenDepth--
				currentArg.WriteRune(ch)
			case ' ':
				if parenDepth == 0 {
					// 参数结束
					if currentArg.Len() > 0 {
						args = append(args, strings.TrimSpace(currentArg.String()))
						currentArg.Reset()
					}
				} else {
					currentArg.WriteRune(ch)
				}
			default:
				currentArg.WriteRune(ch)
			}
		} else {
			// 在字符串中
			currentArg.WriteRune(ch)
			if ch == stringChar && (i == 0 || argsStr[i-1] != '\\') {
				inString = false
			}
		}
	}
	
	// 添加最后一个参数
	if currentArg.Len() > 0 {
		args = append(args, strings.TrimSpace(currentArg.String()))
	}
	
	return fn, args
}

// evalTemplateArg 评估模板参数（可能是嵌套函数调用或变量）
func evalTemplateArg(req map[string]any, ctx *RequestContext, arg string) any {
	arg = strings.TrimSpace(arg)
	
	// 如果是字符串字面量
	if (strings.HasPrefix(arg, "\"") && strings.HasSuffix(arg, "\"")) ||
		(strings.HasPrefix(arg, "'") && strings.HasSuffix(arg, "'")) {
		return arg[1 : len(arg)-1]
	}
	
	// 如果是数字字面量
	if num, err := strconv.ParseFloat(arg, 64); err == nil {
		return num
	}
	
	// 如果是函数调用（包含括号）
	if strings.HasPrefix(arg, "(") && strings.HasSuffix(arg, ")") {
		// 递归评估嵌套函数调用 - 去掉括号并作为模板表达式评估
		innerExpr := strings.TrimSpace(arg[1 : len(arg)-1])
		return evalTemplate(req, ctx, "{{"+innerExpr+"}}")
	}
	
	// 否则作为变量名处理
	if val, ok := getByPath(ctx.Vars, arg); ok {
		return val
	}
	if val, ok := getByPath(req, arg); ok {
		return val
	}
	
	return nil
}

// processTemplateString 处理字符串中的模板变量
func processTemplateString(req map[string]any, ctx *RequestContext, str string) string {
	result := str
	// 查找所有 {{ variable }} 模式
	start := 0
	for {
		idx := strings.Index(result[start:], "{{")
		if idx == -1 {
			break
		}
		idx += start
		endIdx := strings.Index(result[idx:], "}}")
		if endIdx == -1 {
			break
		}
		endIdx += idx
		
		// 提取模板表达式
		templateExpr := strings.TrimSpace(result[idx+2 : endIdx])
		
		// 处理模板表达式
		processed := evalTemplate(req, ctx, "{{"+templateExpr+"}}")
		
		// 替换原始字符串
		result = result[:idx] + fmt.Sprint(processed) + result[endIdx+2:]
		start = idx + len(fmt.Sprint(processed))
	}
	return result
}

func evalTemplate(req map[string]any, ctx *RequestContext, v any) any {
    switch vv := v.(type) {
    case string:
        // 首先检查是否是完全的模板字符串
        if len(vv) >= 4 && vv[:2] == "{{" && vv[len(vv)-2:] == "}}" {
            expr := strings.TrimSpace(vv[2:len(vv)-2])
            
            // 解析函数调用，支持嵌套括号
            fn, args := parseTemplateExpression(expr)
            
            if fn != "" && len(args) > 0 {
                // 处理第一个参数
                firstArg := evalTemplateArg(req, ctx, args[0])
                
                switch fn {
                case "len":
                    switch t := firstArg.(type) { case string: return len(t); case []any: return len(t); case []map[string]any: return len(t); case []datasource.UploadedFile: return len(t); case map[string]any: return len(t) }
                    return 0
                case "count":
                    switch t := firstArg.(type) { case string: return len(t); case []any: return len(t); case []map[string]any: return len(t); case []datasource.UploadedFile: return len(t); case map[string]any: return len(t) }
                    return 0
                case "upper":
                    return strings.ToUpper(fmt.Sprint(firstArg))
                case "lower":
                    return strings.ToLower(fmt.Sprint(firstArg))
                case "json_encode":
                    b, _ := json.Marshal(firstArg); return string(b)
                case "sha256":
                    h := sha256.Sum256([]byte(fmt.Sprint(firstArg)))
                    return hex.EncodeToString(h[:])
                case "sha256_concat":
                    a := fmt.Sprint(firstArg)
                    b := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    h := sha256.Sum256([]byte(a + b))
                    return hex.EncodeToString(h[:])
                case "uuid":
                    buf := make([]byte, 16); _, _ = rand.Read(buf); return hex.EncodeToString(buf)
                case "add":
                    a := toFloat(firstArg)
                    b := toFloat(evalTemplateArg(req, ctx, args[1]))
                    return a + b
                case "eq":
                    a := fmt.Sprint(firstArg)
                    b := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    return a == b
                case "ne":
                    a := fmt.Sprint(firstArg)
                    b := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    return a != b
                case "gt":
                    a := toFloat(firstArg)
                    b := toFloat(evalTemplateArg(req, ctx, args[1]))
                    return a > b
                case "lt":
                    a := toFloat(firstArg)
                    b := toFloat(evalTemplateArg(req, ctx, args[1]))
                    return a < b
                case "ge":
                    a := toFloat(firstArg)
                    b := toFloat(evalTemplateArg(req, ctx, args[1]))
                    return a >= b
                case "le":
                    a := toFloat(firstArg)
                    b := toFloat(evalTemplateArg(req, ctx, args[1]))
                    return a <= b
                case "and":
                    a := truthy(firstArg)
                    b := truthy(evalTemplateArg(req, ctx, args[1]))
                    return a && b
                case "or":
                    a := firstArg
                    b := evalTemplateArg(req, ctx, args[1])
                    if truthy(a) {
                        return a
                    }
                    return b
                case "if":
                    cond := truthy(firstArg)
                    if len(args) > 1 {
                        if cond { return evalTemplateArg(req, ctx, args[1]) }
                        if len(args) > 2 { return evalTemplateArg(req, ctx, args[2]) }
                    }
                    return nil
                case "coalesce":
                    for i := 0; i < len(args); i++ {
                        val := evalTemplateArg(req, ctx, args[i])
                        if truthy(val) { return val }
                    }
                    return nil
                case "not":
                    return !truthy(firstArg)
                case "toint":
                    return toIntAny(firstArg)
                case "tofloat":
                    return toFloat(firstArg)
                case "sub":
                    a := toFloat(firstArg)
                    b := toFloat(evalTemplateArg(req, ctx, args[1]))
                    return a - b
                case "mul":
                    a := toFloat(firstArg)
                    b := toFloat(evalTemplateArg(req, ctx, args[1]))
                    return a * b
                case "div":
                    a := toFloat(firstArg)
                    b := toFloat(evalTemplateArg(req, ctx, args[1]))
                    if b == 0 { return 0 }
                    return a / b
                case "mod":
                    a := toFloat(firstArg)
                    b := toFloat(evalTemplateArg(req, ctx, args[1]))
                    if b == 0 { return 0 }
                    return math.Mod(a, b)
                case "sum":
                    switch t := firstArg.(type) {
                    case []any:
                        var s float64
                        for _, x := range t { s += toFloat(x) }
                        return s
                    case []float64:
                        var s float64
                        for _, x := range t { s += x }
                        return s
                    default:
                        return 0
                    }
                case "avg":
                    switch t := firstArg.(type) {
                    case []any:
                        if len(t) == 0 { return 0 }
                        var s float64
                        for _, x := range t { s += toFloat(x) }
                        return s / float64(len(t))
                    case []float64:
                        if len(t) == 0 { return 0 }
                        var s float64
                        for _, x := range t { s += x }
                        return s / float64(len(t))
                    default:
                        return 0
                    }
                case "min":
                    switch t := firstArg.(type) {
                    case []any:
                        if len(t) == 0 { return 0 }
                        m := toFloat(t[0])
                        for _, x := range t[1:] { v := toFloat(x); if v < m { m = v } }
                        return m
                    case []float64:
                        if len(t) == 0 { return 0 }
                        m := t[0]
                        for _, x := range t[1:] { if x < m { m = x } }
                        return m
                    default:
                        return 0
                    }
                case "max":
                    switch t := firstArg.(type) {
                    case []any:
                        if len(t) == 0 { return 0 }
                        m := toFloat(t[0])
                        for _, x := range t[1:] { v := toFloat(x); if v > m { m = v } }
                        return m
                    case []float64:
                        if len(t) == 0 { return 0 }
                        m := t[0]
                        for _, x := range t[1:] { if x > m { m = x } }
                        return m
                    default:
                        return 0
                    }
                case "join":
                    sep := func() string { if len(args) > 1 { return fmt.Sprint(evalTemplateArg(req, ctx, args[1])) } ; return "," }()
                    ss := toStringSlice(firstArg)
                    if ss == nil { return "" }
                    return strings.Join(ss, sep)
                case "split":
                    sep := func() string { if len(args) > 1 { return fmt.Sprint(evalTemplateArg(req, ctx, args[1])) } ; return "," }()
                    s := fmt.Sprint(firstArg)
                    return strings.Split(s, sep)
                case "format":
                    fs := fmt.Sprint(firstArg)
                    var argv []any
                    if len(args) > 1 { argv = make([]any, 0, len(args)-1); for i := 1; i < len(args); i++ { argv = append(argv, evalTemplateArg(req, ctx, args[i])) } }
                    return fmt.Sprintf(fs, argv...)
                case "concat":
                    var b strings.Builder
                    b.WriteString(fmt.Sprint(firstArg))
                    for i := 1; i < len(args); i++ { b.WriteString(fmt.Sprint(evalTemplateArg(req, ctx, args[i]))) }
                    return b.String()
                case "tostring":
                    return fmt.Sprint(firstArg)
                case "tobool":
                    return truthy(firstArg)
                case "now":
                    mode := strings.ToLower(fmt.Sprint(firstArg))
                    t := time.Now()
                    if mode == "unix" { return float64(t.Unix()) }
                    if mode == "ms" { return float64(t.UnixMilli()) }
                    return t.Format(time.RFC3339)
                case "format_time":
                    layout := time.RFC3339
                    if len(args) > 1 { layout = fmt.Sprint(evalTemplateArg(req, ctx, args[1])) }
                    switch v := firstArg.(type) {
                    case float64:
                        return time.Unix(int64(v), 0).Format(layout)
                    case int:
                        return time.Unix(int64(v), 0).Format(layout)
                    default:
                        ts := fmt.Sprint(v)
                        tt, err := time.Parse(time.RFC3339, ts)
                        if err != nil { return ts }
                        return tt.Format(layout)
                    }
                case "parse_time":
                    layout := time.RFC3339
                    if len(args) > 1 { layout = fmt.Sprint(evalTemplateArg(req, ctx, args[1])) }
                    ts := fmt.Sprint(firstArg)
                    tt, err := time.Parse(layout, ts)
                    if err != nil { return 0 }
                    return float64(tt.Unix())
                case "add_duration":
                    durStr := ""
                    if len(args) > 1 { durStr = fmt.Sprint(evalTemplateArg(req, ctx, args[1])) }
                    d, err := time.ParseDuration(durStr)
                    if err != nil { return firstArg }
                    var base time.Time
                    switch v := firstArg.(type) {
                    case float64:
                        base = time.Unix(int64(v), 0)
                    case int:
                        base = time.Unix(int64(v), 0)
                    default:
                        ts := fmt.Sprint(v)
                        tt, err := time.Parse(time.RFC3339, ts)
                        if err != nil { return ts }
                        base = tt
                    }
                    out := base.Add(d)
                    if len(args) > 2 { layout := fmt.Sprint(evalTemplateArg(req, ctx, args[2])); return out.Format(layout) }
                    return float64(out.Unix())
                case "base64_encode":
                    return base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(firstArg)))
                case "base64_decode":
                    s := fmt.Sprint(firstArg)
                    b, err := base64.StdEncoding.DecodeString(s)
                    if err != nil { return "" }
                    return string(b)
                case "url_encode":
                    return url.QueryEscape(fmt.Sprint(firstArg))
                case "url_decode":
                    s := fmt.Sprint(firstArg)
                    u, err := url.QueryUnescape(s)
                    if err != nil { return s }
                    return u
                case "hmac_sha256":
                    key := ""
                    if len(args) > 1 { key = fmt.Sprint(evalTemplateArg(req, ctx, args[1])) }
                    mac := hmac.New(sha256.New, []byte(key))
                    mac.Write([]byte(fmt.Sprint(firstArg)))
                    return hex.EncodeToString(mac.Sum(nil))
                case "contains":
                    b := evalTemplateArg(req, ctx, args[1])
                    switch t := firstArg.(type) {
                    case string:
                        return strings.Contains(t, fmt.Sprint(b))
                    case []string:
                        for _, x := range t { if x == fmt.Sprint(b) { return true } }
                        return false
                    case []any:
                        bs := fmt.Sprint(b)
                        for _, x := range t { if fmt.Sprint(x) == bs { return true } }
                        return false
                    case map[string]any:
                        _, ok := t[fmt.Sprint(b)]
                        return ok
                    default:
                        return false
                    }
                case "startswith":
                    b := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    return strings.HasPrefix(fmt.Sprint(firstArg), b)
                case "endswith":
                    b := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    return strings.HasSuffix(fmt.Sprint(firstArg), b)
                case "replace":
                    old := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    newv := ""
                    if len(args) > 2 { newv = fmt.Sprint(evalTemplateArg(req, ctx, args[2])) }
                    return strings.ReplaceAll(fmt.Sprint(firstArg), old, newv)
                case "json_decode":
                    s := fmt.Sprint(firstArg)
                    var out any
                    if err := json.Unmarshal([]byte(s), &out); err != nil { return nil }
                    return out
                case "round":
                    return math.Round(toFloat(firstArg))
                case "floor":
                    return math.Floor(toFloat(firstArg))
                case "ceil":
                    return math.Ceil(toFloat(firstArg))
                case "abs":
                    return math.Abs(toFloat(firstArg))
                case "pow":
                    a := toFloat(firstArg)
                    b := toFloat(evalTemplateArg(req, ctx, args[1]))
                    return math.Pow(a, b)
                case "unique":
                    switch t := firstArg.(type) {
                    case []string:
                        seen := map[string]struct{}{}
                        out := make([]string, 0, len(t))
                        for _, x := range t { if _, ok := seen[x]; !ok { seen[x] = struct{}{}; out = append(out, x) } }
                        return out
                    case []any:
                        seen := map[string]struct{}{}
                        out := make([]any, 0, len(t))
                        for _, x := range t { k := fmt.Sprint(x); if _, ok := seen[k]; !ok { seen[k] = struct{}{}; out = append(out, x) } }
                        return out
                    default:
                        return firstArg
                    }
                case "slice":
                    start := 0
                    end := -1
                    if len(args) > 1 { start = toIntAny(evalTemplateArg(req, ctx, args[1])) }
                    if len(args) > 2 { end = toIntAny(evalTemplateArg(req, ctx, args[2])) }
                    switch t := firstArg.(type) {
                    case []any:
                        if start < 0 { start = 0 }
                        if end < 0 || end > len(t) { end = len(t) }
                        if start > end { return []any{} }
                        out := make([]any, end-start)
                        copy(out, t[start:end])
                        return out
                    case []string:
                        if start < 0 { start = 0 }
                        if end < 0 || end > len(t) { end = len(t) }
                        if start > end { return []string{} }
                        out := make([]string, end-start)
                        copy(out, t[start:end])
                        return out
                    default:
                        return nil
                    }
                case "range":
                    start := toIntAny(firstArg)
                    end := start
                    step := 1
                    if len(args) > 1 { end = toIntAny(evalTemplateArg(req, ctx, args[1])) }
                    if len(args) > 2 { step = toIntAny(evalTemplateArg(req, ctx, args[2])) }
                    if step <= 0 { step = 1 }
                    if end < start { return []any{} }
                    n := (end - start) / step + 1
                    out := make([]any, 0, n)
                    for i := start; i <= end; i += step { out = append(out, i) }
                    return out
                case "union":
                    a := []any{}
                    if arr, ok := firstArg.([]any); ok { a = arr }
                    b := []any{}
                    if len(args) > 1 { if arr, ok := evalTemplateArg(req, ctx, args[1]).([]any); ok { b = arr } }
                    seen := map[string]struct{}{}
                    out := make([]any, 0, len(a)+len(b))
                    for _, x := range a { k := fmt.Sprint(x); if _, ok := seen[k]; !ok { seen[k] = struct{}{}; out = append(out, x) } }
                    for _, x := range b { k := fmt.Sprint(x); if _, ok := seen[k]; !ok { seen[k] = struct{}{}; out = append(out, x) } }
                    return out
                case "intersect":
                    a := []any{}
                    if arr, ok := firstArg.([]any); ok { a = arr }
                    b := []any{}
                    if len(args) > 1 { if arr, ok := evalTemplateArg(req, ctx, args[1]).([]any); ok { b = arr } }
                    setB := map[string]struct{}{}
                    for _, x := range b { setB[fmt.Sprint(x)] = struct{}{} }
                    out := make([]any, 0, len(a))
                    for _, x := range a { if _, ok := setB[fmt.Sprint(x)]; ok { out = append(out, x) } }
                    return out
                case "diff":
                    a := []any{}
                    if arr, ok := firstArg.([]any); ok { a = arr }
                    b := []any{}
                    if len(args) > 1 { if arr, ok := evalTemplateArg(req, ctx, args[1]).([]any); ok { b = arr } }
                    setB := map[string]struct{}{}
                    for _, x := range b { setB[fmt.Sprint(x)] = struct{}{} }
                    out := make([]any, 0, len(a))
                    for _, x := range a { if _, ok := setB[fmt.Sprint(x)]; !ok { out = append(out, x) } }
                    return out
                case "keys":
                    if m, ok := firstArg.(map[string]any); ok {
                        ks := make([]string, 0, len(m))
                        for k := range m { ks = append(ks, k) }
                        return ks
                    }
                    return nil
                case "values":
                    if m, ok := firstArg.(map[string]any); ok {
                        vs := make([]any, 0, len(m))
                        for _, v := range m { vs = append(vs, v) }
                        return vs
                    }
                    return nil
                case "get":
                    key := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    v, ok := getFrom(firstArg, key)
                    if !ok { return nil }
                    return v
                case "indexby":
                    path := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    switch t := firstArg.(type) {
                    case []map[string]any:
                        out := make(map[string]any, len(t))
                        for _, it := range t { v, _ := getFrom(it, path); out[fmt.Sprint(v)] = it }
                        return out
                    case []any:
                        out := map[string]any{}
                        for _, it := range t { if m, ok := it.(map[string]any); ok { v, _ := getFrom(m, path); out[fmt.Sprint(v)] = m } }
                        return out
                    default:
                        return nil
                    }
                case "map":
                    path := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    switch t := firstArg.(type) {
                    case []map[string]any:
                        out := make([]any, 0, len(t))
                        for _, it := range t { v, _ := getFrom(it, path); out = append(out, v) }
                        return out
                    case []any:
                        out := make([]any, 0, len(t))
                        for _, it := range t { if m, ok := it.(map[string]any); ok { v, _ := getFrom(m, path); out = append(out, v) } }
                        return out
                    default:
                        return nil
                    }
                case "filter":
                    path := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    var val any
                    if len(args) > 2 { val = evalTemplateArg(req, ctx, args[2]) }
                    switch t := firstArg.(type) {
                    case []map[string]any:
                        out := make([]map[string]any, 0, len(t))
                        for _, it := range t { v, _ := getFrom(it, path); if len(args) > 2 { if fmt.Sprint(v) == fmt.Sprint(val) { out = append(out, it) } } else { if truthy(v) { out = append(out, it) } } }
                        return out
                    case []any:
                        out := make([]any, 0, len(t))
                        for _, it := range t { if m, ok := it.(map[string]any); ok { v, _ := getFrom(m, path); if len(args) > 2 { if fmt.Sprint(v) == fmt.Sprint(val) { out = append(out, it) } } else { if truthy(v) { out = append(out, it) } } } }
                        return out
                    default:
                        return nil
                    }
                case "sort":
                    path := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    desc := false
                    if len(args) > 2 { desc = strings.ToLower(fmt.Sprint(evalTemplateArg(req, ctx, args[2]))) == "desc" }
                    switch t := firstArg.(type) {
                    case []map[string]any:
                        out := make([]map[string]any, len(t))
                        copy(out, t)
                        sort.Slice(out, func(i, j int) bool {
                            vi, _ := getFrom(out[i], path)
                            vj, _ := getFrom(out[j], path)
                            si := fmt.Sprint(vi)
                            sj := fmt.Sprint(vj)
                            if desc { return si > sj } ; return si < sj
                        })
                        return out
                    case []any:
                        out := make([]any, len(t))
                        copy(out, t)
                        sort.Slice(out, func(i, j int) bool {
                            mi, _ := out[i].(map[string]any)
                            mj, _ := out[j].(map[string]any)
                            vi, _ := getFrom(mi, path)
                            vj, _ := getFrom(mj, path)
                            si := fmt.Sprint(vi)
                            sj := fmt.Sprint(vj)
                            if desc { return si > sj } ; return si < sj
                        })
                        return out
                    default:
                        return nil
                    }
                case "sortn":
                    path := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    desc := false
                    if len(args) > 2 { desc = strings.ToLower(fmt.Sprint(evalTemplateArg(req, ctx, args[2]))) == "desc" }
                    switch t := firstArg.(type) {
                    case []map[string]any:
                        out := make([]map[string]any, len(t))
                        copy(out, t)
                        sort.Slice(out, func(i, j int) bool {
                            vi, _ := getFrom(out[i], path)
                            vj, _ := getFrom(out[j], path)
                            fi := toFloat(vi)
                            fj := toFloat(vj)
                            if desc { return fi > fj } ; return fi < fj
                        })
                        return out
                    case []any:
                        out := make([]any, len(t))
                        copy(out, t)
                        sort.Slice(out, func(i, j int) bool {
                            mi, _ := out[i].(map[string]any)
                            mj, _ := out[j].(map[string]any)
                            vi, _ := getFrom(mi, path)
                            vj, _ := getFrom(mj, path)
                            fi := toFloat(vi)
                            fj := toFloat(vj)
                            if desc { return fi > fj } ; return fi < fj
                        })
                        return out
                    default:
                        return nil
                    }
                case "groupby":
                    path := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    out := map[string][]map[string]any{}
                    switch t := firstArg.(type) {
                    case []map[string]any:
                        for _, it := range t { v, _ := getFrom(it, path); k := fmt.Sprint(v); out[k] = append(out[k], it) }
                        return out
                    case []any:
                        for _, it := range t { if m, ok := it.(map[string]any); ok { v, _ := getFrom(m, path); k := fmt.Sprint(v); out[k] = append(out[k], m) } }
                        return out
                    default:
                        return nil
                    }
                case "sumby":
                    path := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    var s float64
                    switch t := firstArg.(type) {
                    case []map[string]any:
                        for _, it := range t { v, _ := getFrom(it, path); s += toFloat(v) }
                        return s
                    case []any:
                        for _, it := range t { if m, ok := it.(map[string]any); ok { v, _ := getFrom(m, path); s += toFloat(v) } }
                        return s
                    default:
                        return 0
                    }
                case "avgby":
                    path := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    var s float64
                    var n int
                    switch t := firstArg.(type) {
                    case []map[string]any:
                        for _, it := range t { v, _ := getFrom(it, path); s += toFloat(v); n++ }
                    case []any:
                        for _, it := range t { if m, ok := it.(map[string]any); ok { v, _ := getFrom(m, path); s += toFloat(v); n++ } }
                    }
                    if n == 0 { return 0 }
                    return s / float64(n)
                case "minby":
                    path := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    var set bool
                    var mval float64
                    switch t := firstArg.(type) {
                    case []map[string]any:
                        for _, it := range t { v, _ := getFrom(it, path); fv := toFloat(v); if !set || fv < mval { mval = fv; set = true } }
                    case []any:
                        for _, it := range t { if mm, ok := it.(map[string]any); ok { v, _ := getFrom(mm, path); fv := toFloat(v); if !set || fv < mval { mval = fv; set = true } } }
                    }
                    if !set { return 0 }
                    return mval
                case "maxby":
                    path := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    var set bool
                    var mval float64
                    switch t := firstArg.(type) {
                    case []map[string]any:
                        for _, it := range t { v, _ := getFrom(it, path); fv := toFloat(v); if !set || fv > mval { mval = fv; set = true } }
                    case []any:
                        for _, it := range t { if mm, ok := it.(map[string]any); ok { v, _ := getFrom(mm, path); fv := toFloat(v); if !set || fv > mval { mval = fv; set = true } } }
                    }
                    if !set { return 0 }
                    return mval
                case "uniqby":
                    path := fmt.Sprint(evalTemplateArg(req, ctx, args[1]))
                    seen := map[string]struct{}{}
                    switch t := firstArg.(type) {
                    case []map[string]any:
                        out := make([]map[string]any, 0, len(t))
                        for _, it := range t { v, _ := getFrom(it, path); k := fmt.Sprint(v); if _, ok := seen[k]; !ok { seen[k] = struct{}{}; out = append(out, it) } }
                        return out
                    case []any:
                        out := make([]any, 0, len(t))
                        for _, it := range t { if m, ok := it.(map[string]any); ok { v, _ := getFrom(m, path); k := fmt.Sprint(v); if _, ok := seen[k]; !ok { seen[k] = struct{}{}; out = append(out, it) } } }
                        return out
                    default:
                        return nil
                    }
                case "pick":
                    if m, ok := firstArg.(map[string]any); ok {
                        out := map[string]any{}
                        if len(args) == 1 { return out }
                        var keys []string
                        if len(args) == 2 {
                            if arr, ok := evalTemplateArg(req, ctx, args[1]).([]any); ok { for _, x := range arr { keys = append(keys, fmt.Sprint(x)) } }
                            if arrs, ok := evalTemplateArg(req, ctx, args[1]).([]string); ok { keys = append(keys, arrs...) }
                        }
                        if len(keys) == 0 { for i := 1; i < len(args); i++ { keys = append(keys, fmt.Sprint(evalTemplateArg(req, ctx, args[i]))) } }
                        for _, k := range keys { if v, ok := m[k]; ok { out[k] = v } }
                        return out
                    }
                    return nil
                case "omit":
                    if m, ok := firstArg.(map[string]any); ok {
                        out := map[string]any{}
                        for k, v := range m { out[k] = v }
                        var keys []string
                        if len(args) == 2 {
                            if arr, ok := evalTemplateArg(req, ctx, args[1]).([]any); ok { for _, x := range arr { keys = append(keys, fmt.Sprint(x)) } }
                            if arrs, ok := evalTemplateArg(req, ctx, args[1]).([]string); ok { keys = append(keys, arrs...) }
                        }
                        if len(keys) == 0 { for i := 1; i < len(args); i++ { keys = append(keys, fmt.Sprint(evalTemplateArg(req, ctx, args[i]))) } }
                        for _, k := range keys { delete(out, k) }
                        return out
                    }
                    return nil
                case "merge":
                    a, aok := firstArg.(map[string]any)
                    if !aok { return nil }
                    b, bok := evalTemplateArg(req, ctx, args[1]).(map[string]any)
                    if !bok { return a }
                    out := map[string]any{}
                    for k, v := range a { out[k] = v }
                    for k, v := range b { out[k] = v }
                    return out
                case "flatten":
                    switch t := firstArg.(type) {
                    case []any:
                        out := make([]any, 0, len(t))
                        for _, it := range t { if arr, ok := it.([]any); ok { out = append(out, arr...) } else { out = append(out, it) } }
                        return out
                    case []string:
                        out := make([]string, 0, len(t))
                        for _, it := range t { out = append(out, it) }
                        return out
                    default:
                        return nil
                    }
                case "compact":
                    switch t := firstArg.(type) {
                    case []any:
                        out := make([]any, 0, len(t))
                        for _, it := range t { if truthy(it) { out = append(out, it) } }
                        return out
                    case []string:
                        out := make([]string, 0, len(t))
                        for _, it := range t { if truthy(it) { out = append(out, it) } }
                        return out
                    default:
                        return nil
                    }
                default:
                    // 如果是变量名，直接返回变量值
                    if val, ok := getByPath(ctx.Vars, expr); ok { return val }
                    if val, ok := getByPath(req, expr); ok { return val }
                    return nil
                }
            } else {
                key := expr
                if val, ok := getByPath(ctx.Vars, key); ok { return val }
                if val, ok := getByPath(req, key); ok { return val }
                return nil
            }
        } else if strings.Contains(vv, "{{") && strings.Contains(vv, "}}") {
            // 处理包含模板变量的字符串
            return processTemplateString(req, ctx, vv)
        }
        return vv
    default:
        return vv
    }
}

func intFrom(v any, def int) int {
    switch t := v.(type) {
    case int:
        return t
    case int64:
        return int(t)
    case float64:
        return int(t)
    default:
        return def
    }
}

func stringFrom(v any, def string) string {
    if s, ok := v.(string); ok { return s }
    return def
}

func mapFromString(v any) map[string]string {
    out := map[string]string{}
    if m, ok := v.(map[string]any); ok {
        for k, vv := range m { out[k] = fmt.Sprint(vv) }
    }
    return out
}

func mapToParams(req map[string]any, ctx *RequestContext, v any) map[string]any {
    out := map[string]any{}
    if m, ok := v.(map[string]any); ok {
        for k, vv := range m { out[k] = evalTemplate(req, ctx, vv) }
    }
    return out
}

func getByPath(m map[string]any, path string) (any, bool) {
    if v, ok := m[path]; ok { return v, true }
    parts := splitPath(path)
    cur := any(m)
    for _, p := range parts {
        switch c := cur.(type) {
        case map[string]any:
            v, ok := c[p]
            if !ok { return nil, false }
            cur = v
        case []any:
            i, err := strconv.Atoi(p)
            if err != nil || i < 0 || i >= len(c) { return nil, false }
            cur = c[i]
        case []map[string]any:
            i, err := strconv.Atoi(p)
            if err != nil || i < 0 || i >= len(c) { return nil, false }
            cur = c[i]
        default:
            return nil, false
        }
    }
    return cur, true
}

func splitPath(s string) []string {
    out := []string{}
    start := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '.' { out = append(out, s[start:i]); start = i+1 }
    }
    out = append(out, s[start:])
    return out
}

func boolFrom(v any) bool {
    if b, ok := v.(bool); ok { return b }
    return false
}

func isEmptyString(v any) bool {
    s, ok := v.(string)
    return ok && len(strings.TrimSpace(s)) == 0
}

func typeMatches(t string, v any) bool {
    switch t {
    case "string":
        _, ok := v.(string); return ok
    case "int":
        switch v.(type) { case int, int64, float64: return true; default: return false }
    case "float":
        switch v.(type) { case float32, float64, int, int64: return true; default: return false }
    case "bool":
        _, ok := v.(bool); return ok
    case "array":
        _, ok := v.([]any); return ok
    case "object":
        _, ok := v.(map[string]any); return ok
    default:
        return true
    }
}

func isNumber(v any) bool {
    switch v.(type) { case int, int64, float32, float64: return true }
    return false
}

func toFloat(v any) float64 {
    switch t := v.(type) {
    case int:
        return float64(t)
    case int64:
        return float64(t)
    case float32:
        return float64(t)
    case float64:
        return t
    default:
        return 0
    }
}

func floatFrom(v any, def float64) float64 {
    switch t := v.(type) {
    case float64:
        return t
    case float32:
        return float64(t)
    case int:
        return float64(t)
    case int64:
        return float64(t)
    default:
        return def
    }
}

func toFloatSlice(v any) []float64 {
    switch t := v.(type) {
    case []float64:
        return t
    case []any:
        out := make([]float64, 0, len(t))
        for _, x := range t { out = append(out, floatFrom(x, 0)) }
        return out
    default:
        return nil
    }
}

func toStepList(arr []any) []map[string]any {
    out := []map[string]any{}
    for _, x := range arr {
        if m, ok := x.(map[string]any); ok { out = append(out, m) }
    }
    return out
}

func toStringSlice(v any) []string {
    switch t := v.(type) {
    case []string:
        return t
    case []any:
        out := make([]string, 0, len(t))
        for _, x := range t { out = append(out, fmt.Sprint(x)) }
        return out
    default:
        return nil
    }
}

func truthy(v any) bool {
    switch t := v.(type) {
    case bool:
        return t
    case string:
        s := strings.TrimSpace(strings.ToLower(t))
        return s != "" && s != "false" && s != "0"
    case int:
        return t != 0
    case float64:
        return t != 0
    default:
        return v != nil
    }
}

func toUploadedFiles(arr []any) []datasource.UploadedFile {
    out := []datasource.UploadedFile{}
    for _, it := range arr {
        if uf, ok := it.(datasource.UploadedFile); ok { out = append(out, uf) }
    }
    return out
}

func sizeOf(r datasource.ReadSeekCloser) int64 {
    cur, _ := r.Seek(0, io.SeekCurrent)
    end, _ := r.Seek(0, io.SeekEnd)
    _, _ = r.Seek(cur, io.SeekStart)
    return end
}

func sniffContentType(r datasource.ReadSeekCloser) string {
    cur, _ := r.Seek(0, io.SeekCurrent)
    buf := make([]byte, 512)
    n, _ := r.Read(buf)
    _, _ = r.Seek(cur, io.SeekStart)
    return http.DetectContentType(buf[:n])
}

func parseSize(s string) int64 {
    t := strings.ToLower(strings.TrimSpace(s))
    if strings.HasSuffix(t, "kb") { return int64(float64(toInt(strings.TrimSuffix(t, "kb"))) * 1024) }
    if strings.HasSuffix(t, "mb") { return int64(float64(toInt(strings.TrimSuffix(t, "mb"))) * 1024 * 1024) }
    if strings.HasSuffix(t, "b") { return int64(toInt(strings.TrimSuffix(t, "b"))) }
    return int64(toInt(t))
}

func toInt(s string) int {
    var x int
    for i := 0; i < len(s); i++ { c := s[i]; if c >= '0' && c <= '9' { x = x*10 + int(c-'0') } }
    return x
}
func toIntAny(v any) int {
    switch t := v.(type) {
    case int:
        return t
    case int64:
        return int(t)
    case float64:
        return int(t)
    case float32:
        return int(t)
    case string:
        s := strings.TrimSpace(t)
        if i, err := strconv.Atoi(s); err == nil { return i }
        return toInt(s)
    default:
        return 0
    }
}

func toStringSet(ts []any) map[string]struct{} {
    m := map[string]struct{}{}
    for _, v := range ts { m[strings.ToLower(fmt.Sprint(v))] = struct{}{} }
    return m
}

func IntFromAny(v any, def int) int { return intFrom(v, def) }

// 简单插件熔断 / Simple circuit breaker
type circuitState struct{ fails int; openUntil int64 }
var pluginCircuits = map[string]*circuitState{}

func isCircuitOpen(key string) bool {
    if st, ok := pluginCircuits[key]; ok { return st.openUntil > time.Now().UnixMilli() }
    return false
}

func recordCircuitFail(key string, threshold, openMs int) {
    st := pluginCircuits[key]
    if st == nil { st = &circuitState{}; pluginCircuits[key] = st }
    st.fails++
    if st.fails >= threshold && openMs > 0 { st.openUntil = time.Now().Add(time.Duration(openMs) * time.Millisecond).UnixMilli(); st.fails = 0 }
}

func resetCircuit(key string) { if st, ok := pluginCircuits[key]; ok { st.fails = 0; st.openUntil = 0 } }
func argVal(req map[string]any, ctx *RequestContext, parts []string, idx int) any {
    if len(parts) <= idx { return "" }
    p := parts[idx]
    if v, ok := getByPath(ctx.Vars, p); ok { return v }
    if v, ok := getByPath(req, p); ok { return v }
    return p
}

func evalMap(req map[string]any, ctx *RequestContext, m map[string]any) map[string]any {
    out := make(map[string]any, len(m))
    for k, vv := range m {
        val := evalTemplate(req, ctx, vv)
        if mm, ok := val.(map[string]any); ok {
            out[k] = evalMap(req, ctx, mm)
        } else {
            out[k] = val
        }
    }
    return out
}
func getFrom(root any, path string) (any, bool) {
    parts := splitPath(path)
    cur := root
    for _, p := range parts {
        switch c := cur.(type) {
        case map[string]any:
            v, ok := c[p]
            if !ok { return nil, false }
            cur = v
        case []any:
            i, err := strconv.Atoi(p)
            if err != nil || i < 0 || i >= len(c) { return nil, false }
            cur = c[i]
        case []map[string]any:
            i, err := strconv.Atoi(p)
            if err != nil || i < 0 || i >= len(c) { return nil, false }
            cur = c[i]
        default:
            return nil, false
        }
    }
    return cur, true
}