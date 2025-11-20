package apidocs

import (
    "fmt"
    "ifaceconf/internal/config"
)

func BuildDocs(list []*config.InterfaceConfig) []DocEndpoint {
    out := make([]DocEndpoint, 0, len(list))
    for _, ic := range list {
        de := DocEndpoint{
            Module:      ic.Module,
            Endpoint:    ic.Endpoint,
            Method:      ic.Method,
            Path:        ic.Path,
            Title:       strFromMap(ic.Docs, "title"),
            Description: strFromMap(ic.Docs, "description"),
            Auth:        ic.Auth,
            Roles:       ic.Roles,
        }
        params, files := collectParams(ic.Request)
        de.Params = params
        de.Files = files
        de.Response = collectResponse(ic.Steps, ic.Docs)
        de.Response.Schema = BuildResponseSchema(ic.Docs)
        de.Errors = collectErrors(ic.Docs)
        out = append(out, de)
    }
    return out
}

func collectParams(req map[string]any) ([]DocParam, []DocParam) {
    if req == nil { return nil, nil }
    ps := []DocParam{}
    fs := []DocParam{}
    for source, block := range req {
        bm, ok := block.(map[string]any)
        if !ok { continue }
        for name, spec := range bm {
            sm, ok := spec.(map[string]any)
            if !ok { continue }
            dp := DocParam{
                Source:      source,
                Name:        name,
                Type:        stringFrom(sm["type"]),
                Required:    boolFrom(sm["required"]),
                Description: firstString(sm, "desc", "description"),
                Constraints: toConstraints(sm),
            }
            if source == "file" {
                fs = append(fs, dp)
            } else {
                ps = append(ps, dp)
            }
        }
    }
    return ps, fs
}

func collectResponse(steps []map[string]any, docsMeta map[string]any) DocResponse {
    dr := DocResponse{}
    for _, st := range steps {
        if v, ok := st["response"]; ok {
            m, mok := v.(map[string]any)
            if !mok || m == nil { continue }
            dr.Status = intFrom(m["status"], 200)
            dr.Body = stringify(m["body"]) 
            if hm, ok := m["headers"].(map[string]any); ok {
                dr.Headers = mergeHeaderDocs(mapToHeaders(hm), docsMeta)
            } else {
                dr.Headers = mergeHeaderDocs(nil, docsMeta)
            }
        }
    }
    if docsMeta != nil {
        if rm, ok := docsMeta["response"].(map[string]any); ok {
            dr.BodyDescription = stringFrom(rm["description"]) 
            if ex, ok := rm["example"]; ok { dr.Body = stringify(ex) }
            if dr.Headers == nil {
                dr.Headers = mergeHeaderDocs(nil, docsMeta)
            }
        }
    }
    return dr
}

func strFromMap(m map[string]any, k string) string {
    if m == nil { return "" }
    if v, ok := m[k]; ok { return stringFrom(v) }
    return ""
}

func stringFrom(v any) string {
    switch t := v.(type) {
    case string:
        return t
    default:
        return fmt.Sprint(v)
    }
}

func boolFrom(v any) bool {
    if b, ok := v.(bool); ok { return b }
    return false
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

func toConstraints(m map[string]any) map[string]string {
    out := map[string]string{}
    for _, k := range []string{"min","max","minLen","maxLen","enum","regex","default","maxSize","types"} {
        if v, ok := m[k]; ok { out[k] = stringFrom(v) }
    }
    return out
}

func stringify(v any) string {
    switch t := v.(type) {
    case string:
        return t
    default:
        return fmt.Sprintf("%v", v)
    }
}

func firstString(m map[string]any, keys ...string) string {
    for _, k := range keys {
        if v, ok := m[k]; ok { return stringFrom(v) }
    }
    return ""
}

func mapToHeaders(hm map[string]any) []DocHeader {
    if hm == nil { return nil }
    out := []DocHeader{}
    for k, v := range hm {
        out = append(out, DocHeader{Name: k, Value: stringFrom(v)})
    }
    return out
}

func mergeHeaderDocs(headers []DocHeader, docsMeta map[string]any) []DocHeader {
    var descs map[string]any
    if docsMeta != nil {
        if rm, ok := docsMeta["response"].(map[string]any); ok {
            if hm, ok := rm["headers"].(map[string]any); ok { descs = hm }
        }
    }
    if headers == nil && descs != nil {
        out := []DocHeader{}
        for k, v := range descs { out = append(out, DocHeader{Name: k, Description: stringFrom(v)}) }
        return out
    }
    if headers == nil { return nil }
    for i := range headers {
        if descs != nil {
            if v, ok := descs[headers[i].Name]; ok { headers[i].Description = stringFrom(v) }
        }
    }
    return headers
}

func BuildResponseSchema(docsMeta map[string]any) []DocSchema {
    if docsMeta == nil { return nil }
    rm, ok := docsMeta["response"].(map[string]any)
    if !ok { return nil }
    sm, ok := rm["schema"].(map[string]any)
    if !ok { return nil }
    return parseSchemaProps("", sm)
}

func parseSchemaProps(parent string, sm map[string]any) []DocSchema {
    var roots []DocSchema
    if pm, ok := sm["props"].(map[string]any); ok {
        for name, spec := range pm {
            node := DocSchema{Name: name}
            if s, ok := spec.(map[string]any); ok {
                node.Type = stringFrom(s["type"]) 
                node.Description = firstString(s, "desc", "description")
                if child, ok := s["props"].(map[string]any); ok {
                    node.Children = parseSchemaProps(name, map[string]any{"props": child})
                }
            }
            roots = append(roots, node)
        }
        return roots
    }
    return roots
}

func collectErrors(docsMeta map[string]any) []DocError {
    if docsMeta == nil { return nil }
    em := docsMeta["errors"]
    if em == nil { return nil }
    out := []DocError{}
    if arr, ok := em.([]any); ok {
        for _, item := range arr {
            if m, ok := item.(map[string]any); ok {
                out = append(out, DocError{Code: stringFrom(m["code"]), Message: stringFrom(m["message"]), Description: firstString(m, "desc", "description")})
            }
        }
        return out
    }
    if mp, ok := em.(map[string]any); ok {
        for k, v := range mp {
            if m, ok := v.(map[string]any); ok {
                out = append(out, DocError{Code: k, Message: stringFrom(m["message"]), Description: firstString(m, "desc", "description")})
            } else {
                out = append(out, DocError{Code: k, Message: stringFrom(v)})
            }
        }
        return out
    }
    return out
}