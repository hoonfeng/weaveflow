package docs

import (
    "strings"
    "strconv"
    "ifaceconf/internal/apidocs"
    "unicode"
)

func BuildOpenAPI(title string, endpoints []apidocs.DocEndpoint) map[string]any {
    paths := map[string]any{}
    securitySchemes := map[string]any{"bearerAuth": map[string]any{"type": "http", "scheme": "bearer", "bearerFormat": "JWT"}}
    for _, e := range endpoints {
        m := strings.ToLower(e.Method)
        if paths[e.Path] == nil { paths[e.Path] = map[string]any{} }
        op := map[string]any{
            "summary": e.Title,
            "description": e.Description,
            "tags": []string{e.Module},
            "parameters": buildParams(e),
            "responses": buildResponses(e),
        }
        if len(e.Roles) > 0 { op["x-roles"] = e.Roles }
        if rb := buildRequestBody(e); rb != nil { op["requestBody"] = rb }
        if strings.ToLower(e.Auth) == "jwt" {
            op["security"] = []any{map[string]any{"bearerAuth": []any{}}}
        }
        paths[e.Path].(map[string]any)[m] = op
    }
    return map[string]any{
        "openapi": "3.0.3",
        "info": map[string]any{"title": title, "version": "v1"},
        "components": map[string]any{"securitySchemes": securitySchemes},
        "paths": paths,
    }
}

func buildParams(e apidocs.DocEndpoint) []any {
    ps := []any{}
    add := func(p apidocs.DocParam) {
        // Only path/query/header belong to parameters; body/form/file go to requestBody
        in := normalizeIn(p.Source)
        if in != "path" && in != "query" && in != "header" { return }
        sch := map[string]any{"type": normalizeType(p.Type)}
        applyConstraints(sch, p)
        ps = append(ps, map[string]any{
            "name": p.Name,
            "in": in,
            "required": p.Required,
            "description": p.Description,
            "schema": sch,
        })
    }
    for _, p := range e.Params { add(p) }
    for _, f := range e.Files { add(f) }
    return ps
}

func buildResponses(e apidocs.DocEndpoint) map[string]any {
    resp := map[string]any{}
    code := e.Response.Status
    if code == 0 { code = 200 }
    bodySchema := schemaFromDoc(e.Response.Schema)
    headers := map[string]any{}
    for _, h := range e.Response.Headers {
        headers[h.Name] = map[string]any{
            "description": h.Description,
            "schema": map[string]any{"type": "string"},
        }
    }
    resp[itoa(code)] = map[string]any{
        "description": e.Response.BodyDescription,
        "headers": headers,
        "content": map[string]any{
            "application/json": map[string]any{"schema": bodySchema},
        },
    }
    return resp
}

func schemaFromDoc(nodes []apidocs.DocSchema) map[string]any {
    if len(nodes) == 0 { return map[string]any{"type": "object"} }
    obj := map[string]any{"type": "object", "properties": map[string]any{}}
    props := obj["properties"].(map[string]any)
    for _, n := range nodes {
        if len(n.Children) == 0 {
            props[n.Name] = map[string]any{"type": normalizeType(n.Type), "description": n.Description}
        } else {
            props[n.Name] = schemaFromDoc(n.Children)
        }
    }
    return obj
}

func normalizeIn(s string) string {
    switch s {
    case "path":
        return "path"
    case "query":
        return "query"
    case "header":
        return "header"
    default:
        return "query"
    }
}

func normalizeType(t string) string {
    switch t {
    case "int":
        return "integer"
    default:
        return strings.ToLower(t)
    }
}

func itoa(i int) string { return strconv.Itoa(i) }

// Build requestBody for JSON and multipart form
func buildRequestBody(e apidocs.DocEndpoint) map[string]any {
    jsonProps := map[string]any{}
    jsonRequired := []string{}
    formProps := map[string]any{}
    formRequired := []string{}
    // Body params → application/json
    for _, p := range e.Params {
        if p.Source == "body" {
            sch := map[string]any{"type": normalizeType(p.Type), "description": p.Description}
            applyConstraints(sch, p)
            jsonProps[p.Name] = sch
            if p.Required { jsonRequired = append(jsonRequired, p.Name) }
        }
        if p.Source == "form" {
            sch := map[string]any{"type": normalizeType(p.Type), "description": p.Description}
            applyConstraints(sch, p)
            formProps[p.Name] = sch
            if p.Required { formRequired = append(formRequired, p.Name) }
        }
    }
    // Files → multipart/form-data with binary fields
    for _, f := range e.Files {
        // default to string/binary; if field name is 'files' treat as array of binary
        if strings.ToLower(f.Name) == "files" {
            formProps[f.Name] = map[string]any{"type": "array", "items": map[string]any{"type": "string", "format": "binary"}, "description": f.Description}
        } else {
            formProps[f.Name] = map[string]any{"type": "string", "format": "binary", "description": f.Description}
        }
        if f.Required { formRequired = append(formRequired, f.Name) }
    }
    content := map[string]any{}
    if len(jsonProps) > 0 {
        sch := map[string]any{"type": "object", "properties": jsonProps}
        if len(jsonRequired) > 0 { sch["required"] = jsonRequired }
        content["application/json"] = map[string]any{"schema": sch}
    }
    if len(formProps) > 0 {
        sch := map[string]any{"type": "object", "properties": formProps}
        if len(formRequired) > 0 { sch["required"] = formRequired }
        content["multipart/form-data"] = map[string]any{"schema": sch}
    }
    if len(content) == 0 { return nil }
    return map[string]any{"required": true, "content": content}
}

// Apply common constraints from DocParam to OpenAPI schema
func applyConstraints(sch map[string]any, p apidocs.DocParam) {
    cs := p.Constraints
    if cs == nil { return }
    if v, ok := cs["minLen"]; ok { if n := atoiSafe(v); n >= 0 { sch["minLength"] = n } }
    if v, ok := cs["maxLen"]; ok { if n := atoiSafe(v); n >= 0 { sch["maxLength"] = n } }
    if v, ok := cs["min"]; ok { if f := atofSafe(v); f == f { sch["minimum"] = f } }
    if v, ok := cs["max"]; ok { if f := atofSafe(v); f == f { sch["maximum"] = f } }
    if v, ok := cs["regex"]; ok { if v != "" { sch["pattern"] = v } }
    if v, ok := cs["enum"]; ok {
        if v != "" { sch["enum"] = splitEnum(v) }
    }
}

func atoiSafe(s string) int {
    i := 0; sign := 1; started := false
    for _, r := range s {
        if !started && (r == '+' || r == '-') { if r == '-' { sign = -1 }; started = true; continue }
        if unicode.IsDigit(r) { i = i*10 + int(r-'0'); started = true } else if started { break }
    }
    return i * sign
}

func atofSafe(s string) float64 {
    // very simple parse; rely on strconv when possible
    f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
    if err != nil { return 0 }
    return f
}

func splitEnum(s string) []string {
    // support comma or space separated lists, trim brackets
    t := strings.TrimSpace(s)
    t = strings.TrimPrefix(t, "[")
    t = strings.TrimSuffix(t, "]")
    var parts []string
    if strings.Contains(t, ",") { parts = strings.Split(t, ",") } else { parts = strings.Fields(t) }
    out := make([]string, 0, len(parts))
    for _, p := range parts { p = strings.TrimSpace(p); if p != "" { out = append(out, p) } }
    return out
}