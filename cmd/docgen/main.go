package main

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "ifaceconf/internal/apidocs"
    "ifaceconf/internal/config"
    idocs "ifaceconf/internal/docs"
)

func main() {
    outDir := "docs/generated"
    for i := 1; i < len(os.Args); i++ {
        if os.Args[i] == "--out" && i+1 < len(os.Args) { outDir = os.Args[i+1]; i++ }
    }
    proj, err := config.LoadProjectConfig("configs/project.yaml")
    if err != nil { panic(err) }
    list, err := config.LoadInterfaceConfigs("configs/interfaces")
    if err != nil { panic(err) }
    eps := apidocs.BuildDocs(list)
    byMod := map[string][]apidocs.DocEndpoint{}
    for _, e := range eps { byMod[e.Module] = append(byMod[e.Module], e) }
    _ = os.MkdirAll(outDir, 0o755)
    idx := &strings.Builder{}
    fmt.Fprintf(idx, "# 接口文档索引\n")
    for mod, arr := range byMod {
        mdPath := filepath.Join(outDir, fmt.Sprintf("%s.md", safe(mod)))
        jsonPath := filepath.Join(outDir, fmt.Sprintf("openapi.%s.json", safe(mod)))
        md := renderModuleMarkdown(mod, arr)
        _ = os.WriteFile(mdPath, []byte(md), 0o644)
        oa := idocs.BuildOpenAPI(proj.Docs.Title, arr)
        jb, _ := json.MarshalIndent(oa, "", "  ")
        _ = os.WriteFile(jsonPath, jb, 0o644)
        fmt.Fprintf(idx, "- %s: [%s](%s) | OpenAPI: %s\n", mod, filepath.Base(mdPath), mdPath, filepath.Base(jsonPath))
    }
    _ = os.WriteFile(filepath.Join(outDir, "index.md"), []byte(idx.String()), 0o644)
}

func renderModuleMarkdown(mod string, arr []apidocs.DocEndpoint) string {
    b := &strings.Builder{}
    fmt.Fprintf(b, "# 模块: %s\n\n", mod)
    for _, e := range arr {
        fmt.Fprintf(b, "## 接口: %s\n", nonEmpty(e.Title, e.Endpoint))
        fmt.Fprintf(b, "- 方法: %s\n", e.Method)
        fmt.Fprintf(b, "- 路径: %s\n", e.Path)
        if e.Description != "" { fmt.Fprintf(b, "- 说明: %s\n", e.Description) }
        if e.Auth != "" { fmt.Fprintf(b, "- 认证: %s\n", e.Auth) }
        if len(e.Roles) > 0 { fmt.Fprintf(b, "- 角色: %s\n", strings.Join(e.Roles, ", ")) }
        if len(e.Params) > 0 || len(e.Files) > 0 {
            fmt.Fprintf(b, "\n### 请求参数\n")
            for _, p := range e.Params { fmt.Fprintf(b, "- [%s] %s (%s) 必填:%v 说明:%s\n", p.Source, p.Name, p.Type, p.Required, p.Description) }
            for _, p := range e.Files { fmt.Fprintf(b, "- [file] %s (%s) 必填:%v 说明:%s\n", p.Name, p.Type, p.Required, p.Description) }
        }
        fmt.Fprintf(b, "\n### 响应\n")
        fmt.Fprintf(b, "- 状态码: %d\n", e.Response.Status)
        if e.Response.BodyDescription != "" { fmt.Fprintf(b, "- 返回说明: %s\n", e.Response.BodyDescription) }
        if len(e.Response.Headers) > 0 { fmt.Fprintf(b, "- 响应头:\n") ; for _, h := range e.Response.Headers { fmt.Fprintf(b, "  - %s: %s\n", h.Name, nonEmpty(h.Description, h.Value)) } }
        if len(e.Response.Schema) > 0 {
            fmt.Fprintf(b, "\n#### 返回参数\n")
            for _, s := range e.Response.Schema { renderSchema(b, s, 0) }
            sample := apidocs.SampleFromSchema(e.Response.Schema)
            sb, _ := json.MarshalIndent(sample, "", "  ")
            fmt.Fprintf(b, "\n#### 示例数据\n")
            fmt.Fprintf(b, "```json\n%s\n```\n", string(sb))
        } else if e.Response.Body != "" {
            fmt.Fprintf(b, "\n#### 示例数据\n")
            fmt.Fprintf(b, "```\n%s\n```\n", e.Response.Body)
        }
        fmt.Fprintf(b, "\n")
    }
    return b.String()
}

func renderSchema(b *strings.Builder, s apidocs.DocSchema, indent int) {
    pad := strings.Repeat("  ", indent)
    name := s.Name
    if name == "" { name = "(root)" }
    if s.Description != "" {
        fmt.Fprintf(b, "%s- %s (%s) — %s\n", pad, name, s.Type, s.Description)
    } else {
        fmt.Fprintf(b, "%s- %s (%s)\n", pad, name, s.Type)
    }
    for _, c := range s.Children { renderSchema(b, c, indent+1) }
}

func nonEmpty(a, b string) string { if a != "" { return a } ; return b }
func safe(s string) string {
    if s == "" { return "default" }
    r := make([]byte, 0, len(s))
    for i := 0; i < len(s); i++ {
        c := s[i]
        if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == '_' { r = append(r, c) } else { r = append(r, '_') }
    }
    return string(r)
}