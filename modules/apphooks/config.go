package apphooks

import (
    "net/http"
    "strings"
    "fmt"
    "time"
    "sort"
    "bytes"
    "io"
    "crypto/sha256"
    "encoding/hex"
    "regexp"

    "ifaceconf/internal/config"
    "ifaceconf/internal/ext"
    "ifaceconf/internal/datasource"
    "ifaceconf/internal/plugin"
    "ifaceconf/internal/core"
)

type hookDef struct{ Kind string `yaml:"kind" json:"kind"`; Name string `yaml:"name" json:"name"`; Match map[string]any `yaml:"match" json:"match"`; Params map[string]any `yaml:"params" json:"params"`; Order int `yaml:"order" json:"order"` }

func RegisterFromConfigs(project *config.ProjectConfig, bundle *config.HooksBundle) {
    var authDefs []hookDef
    var beforeDefs []hookDef
    var afterDefs []hookDef
    for _, m := range bundle.Auth { authDefs = append(authDefs, toDef(m)) }
    for _, m := range bundle.Before { beforeDefs = append(beforeDefs, toDef(m)) }
    for _, m := range bundle.After { afterDefs = append(afterDefs, toDef(m)) }
    loadedAuthDefs = authDefs
    loadedBeforeDefs = beforeDefs
    loadedAfterDefs = afterDefs
    ext.SetHooks(ext.Hooks{
        Auth: func(project *config.ProjectConfig, ic *config.InterfaceConfig, ds *datasource.Manager, prt plugin.Runtime, base http.HandlerFunc) http.HandlerFunc {
            defs := effectiveDefs(authDefs, ic.Hooks.Auth, ic.Hooks.Disable)
            h := base
            if len(defs) == 0 { return h }
            wrap := h
            h = func(w http.ResponseWriter, r *http.Request) {
                for _, d := range defs {
                    if !matches(d.Match, ic) { continue }
                    if d.Kind != "plugin" { continue }
                    pluginName := fmt.Sprint(d.Params["plugin"]) ; fn := fmt.Sprint(d.Params["function"]) ; params := cloneMap(d.Params)
                    if params["dsn"] == nil || fmt.Sprint(params["dsn"]) == "" { params["dsn"] = project.Datasources.SQL["main"].DSN }
                    applyReqParams(&params, r)
                    if _, err := prt.Call(pluginName, fn, params); err != nil { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
                }
                wrap(w, r)
            }
            return h
        },
        Before: func(project *config.ProjectConfig, ic *config.InterfaceConfig, ds *datasource.Manager, prt plugin.Runtime, r *http.Request) error {
            defs := effectiveDefs(beforeDefs, ic.Hooks.Before, ic.Hooks.Disable)
            for _, d := range defs {
                if !matches(d.Match, ic) { continue }
                if d.Kind != "plugin" { continue }
                pluginName := fmt.Sprint(d.Params["plugin"]) ; fn := fmt.Sprint(d.Params["function"]) ; params := cloneMap(d.Params)
                if params["dsn"] == nil || fmt.Sprint(params["dsn"]) == "" { params["dsn"] = project.Datasources.SQL["main"].DSN }
                applyReqParams(&params, r)
                if _, err := prt.Call(pluginName, fn, params); err != nil { return err }
            }
            return nil
        },
        After: func(project *config.ProjectConfig, ic *config.InterfaceConfig, ds *datasource.Manager, prt plugin.Runtime, r *http.Request, res *core.ExecResult, dur time.Duration) {
            defs := effectiveDefs(afterDefs, ic.Hooks.After, ic.Hooks.Disable)
            for _, d := range defs {
                if !matches(d.Match, ic) { continue }
                if d.Kind != "plugin" { continue }
                pluginName := fmt.Sprint(d.Params["plugin"]) ; fn := fmt.Sprint(d.Params["function"]) ; params := cloneMap(d.Params)
                if params["dsn"] == nil || fmt.Sprint(params["dsn"]) == "" { params["dsn"] = project.Datasources.SQL["main"].DSN }
                applyReqParams(&params, r)
                // inject response info
                params["status"] = res.Status
                params["endpoint"] = ic.Method + " " + ic.Path
                params["duration_ms"] = int(dur.Milliseconds())
                _, _ = prt.Call(pluginName, fn, params)
            }
        },
    })
}

var loadedAuthDefs []hookDef
var loadedBeforeDefs []hookDef
var loadedAfterDefs []hookDef

type InspectResult struct{
    Auth []hookDef `json:"auth"`
    Before []hookDef `json:"before"`
    After []hookDef `json:"after"`
}

func Inspect(ic *config.InterfaceConfig) InspectResult {
    res := InspectResult{}
    res.Auth = filterDefs(loadedAuthDefs, ic.Hooks.Auth, ic.Hooks.Disable, ic)
    res.Before = filterDefs(loadedBeforeDefs, ic.Hooks.Before, ic.Hooks.Disable, ic)
    res.After = filterDefs(loadedAfterDefs, ic.Hooks.After, ic.Hooks.Disable, ic)
    return res
}

func filterDefs(base []hookDef, add []map[string]any, disable []string, ic *config.InterfaceConfig) []hookDef {
    defs := effectiveDefs(base, add, disable)
    out := make([]hookDef, 0, len(defs))
    for _, d := range defs { if matches(d.Match, ic) { out = append(out, d) } }
    return out
}

func toDef(m map[string]any) hookDef {
    d := hookDef{}
    if v, ok := m["kind"].(string); ok { d.Kind = v }
    if v, ok := m["name"].(string); ok { d.Name = v }
    if v, ok := m["match"].(map[string]any); ok { d.Match = v }
    if v, ok := m["params"].(map[string]any); ok { d.Params = v }
    if v, ok := m["order"].(int); ok { d.Order = v } else if vv, ok := m["order"].(float64); ok { d.Order = int(vv) }
    return d
}

func matches(m map[string]any, ic *config.InterfaceConfig) bool {
    if m == nil { return true }
    if v, ok := m["auth"].(string); ok { if strings.ToLower(v) != strings.ToLower(ic.Auth) { return false } }
    if v, ok := m["pathPrefix"].(string); ok { if !strings.HasPrefix(ic.Path, v) { return false } }
    if v, ok := m["path"].(string); ok { if ic.Path != v { return false } }
    if v, ok := m["pathRegex"].(string); ok && v != "" {
        re, err := regexp.Compile(v)
        if err == nil { if !re.MatchString(ic.Path) { return false } }
    }
    if v, ok := m["method"].(string); ok { if strings.ToUpper(v) != strings.ToUpper(ic.Method) { return false } }
    if arr, ok := m["method"].([]any); ok {
        found := false
        for _, x := range arr { if strings.ToUpper(fmt.Sprint(x)) == strings.ToUpper(ic.Method) { found = true; break } }
        if !found { return false }
    }
    if v, ok := m["module"].(string); ok { if v != ic.Module { return false } }
    if v, ok := m["endpoint"].(string); ok { if v != ic.Endpoint { return false } }
    if arr, ok := m["interfaces"].([]any); ok {
        // interfaces: ["POST /api/service/compute", ...]
        found := false
        key := strings.ToUpper(ic.Method) + " " + ic.Path
        for _, x := range arr { if strings.ToUpper(fmt.Sprint(x)) == key { found = true; break } }
        if !found { return false }
    }
    // labels matching
    if v, ok := m["label"].(string); ok && v != "" {
        if !hasLabel(ic.Labels, v) { return false }
    }
    if anyArr, ok := m["labelsAny"].([]any); ok {
        if !anyLabel(ic.Labels, anyArr) { return false }
    }
    if allArr, ok := m["labelsAll"].([]any); ok {
        if !allLabels(ic.Labels, allArr) { return false }
    }
    return true
}

func hasLabel(labels []string, target string) bool {
    for _, l := range labels { if l == target { return true } }
    return false
}

func anyLabel(labels []string, arr []any) bool {
    set := map[string]struct{}{}
    for _, l := range labels { set[l] = struct{}{} }
    for _, x := range arr { if _, ok := set[fmt.Sprint(x)]; ok { return true } }
    return len(arr) == 0
}

func allLabels(labels []string, arr []any) bool {
    set := map[string]struct{}{}
    for _, l := range labels { set[l] = struct{}{} }
    for _, x := range arr { if _, ok := set[fmt.Sprint(x)]; !ok { return false } }
    return true
}

func effectiveDefs(base []hookDef, add []map[string]any, disable []string) []hookDef {
    ds := map[string]struct{}{}
    for _, n := range disable { ds[n] = struct{}{} }
    out := make([]hookDef, 0, len(base)+len(add))
    for _, d := range base { if _, skip := ds[d.Name]; !skip { out = append(out, d) } }
    for _, m := range add { d := toDef(m); if _, skip := ds[d.Name]; !skip { out = append(out, d) } }
    sort.Slice(out, func(i, j int) bool { return out[i].Order < out[j].Order })
    return out
}

func cloneMap(m map[string]any) map[string]any { if m == nil { return map[string]any{} } ; out := make(map[string]any, len(m)); for k, v := range m { out[k] = v } ; return out }

func applyReqParams(params *map[string]any, r *http.Request) {
    if params == nil || r == nil { return }
    p := *params
    if hm, ok := p["headers"].(map[string]any); ok {
        for key, hv := range hm {
            hs := fmt.Sprint(hv)
            p[key] = r.Header.Get(hs)
        }
    }
    p["method"] = r.Method
    p["path"] = r.URL.Path
    if pv, ok := p["provider"]; !ok || fmt.Sprint(pv) == "" {
        prov := r.URL.Query().Get("provider")
        if prov == "" {
            u := strings.ToLower(r.URL.Path)
            if strings.Contains(u, "alipay") { prov = "alipay" } else if strings.Contains(u, "wechat") { prov = "wechat" } else if strings.Contains(u, "stripe") { prov = "stripe" }
        }
        if prov != "" { p["provider"] = prov }
    }
    if v, ok := p["bodyHash"].(string); ok && strings.ToLower(v) == "sha256" {
        b, _ := io.ReadAll(r.Body)
        r.Body = io.NopCloser(bytes.NewReader(b))
        h := sha256.Sum256(b)
        p["body_hash"] = hex.EncodeToString(h[:])
    }
    *params = p
}

// no idempotency cache handled here; keep simple mapping for config-driven hooks