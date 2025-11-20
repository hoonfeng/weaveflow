package main

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "ifaceconf/internal/apidocs"
    "ifaceconf/internal/config"
)

func main() {
    outDir := "docs/manual"
    for i := 1; i < len(os.Args); i++ { if os.Args[i] == "--out" && i+1 < len(os.Args) { outDir = os.Args[i+1]; i++ } }
    list, err := config.LoadInterfaceConfigs("configs/interfaces")
    if err != nil { panic(err) }
    eps := apidocs.BuildDocs(list)
    byMod := map[string][]apidocs.DocEndpoint{}
    for _, e := range eps { byMod[e.Module] = append(byMod[e.Module], e) }
    _ = os.MkdirAll(outDir, 0o755)
    idx := &strings.Builder{}
    fmt.Fprintf(idx, "# 接口文档（按模块）\n\n")
    for mod, arr := range byMod {
        // 模块汇总文档
        mdPath := filepath.Join(outDir, fmt.Sprintf("%s.md", safe(mod)))
        md := renderManual(mod, arr)
        _ = os.WriteFile(mdPath, []byte(md), 0o644)
        fmt.Fprintf(idx, "- %s → %s\n", mod, filepath.Base(mdPath))
        // 按接口拆分文档
        modDir := filepath.Join(outDir, safe(mod))
        _ = os.MkdirAll(modDir, 0o755)
        for _, e := range arr {
            fname := endpointFileName(e)
            one := renderManual(mod, []apidocs.DocEndpoint{e})
            _ = os.WriteFile(filepath.Join(modDir, fname), []byte(one), 0o644)
        }
    }
    _ = os.WriteFile(filepath.Join(outDir, "index.md"), []byte(idx.String()), 0o644)
}

func renderManual(mod string, arr []apidocs.DocEndpoint) string {
    b := &strings.Builder{}
    fmt.Fprintf(b, "# 模块：%s\n\n", mod)
    fmt.Fprintf(b, "说明：本模块包含以下接口，基于配置目录编写，覆盖接口名、路径、说明、请求参数、返回参数与示例数据。\n\n")
    for _, e := range arr {
        title := nonEmpty(e.Title, e.Endpoint)
        fmt.Fprintf(b, "## %s\n", title)
        fmt.Fprintf(b, "- 接口方法：%s\n", strings.ToUpper(e.Method))
        fmt.Fprintf(b, "- 接口路径：%s\n", e.Path)
        if e.Description != "" { fmt.Fprintf(b, "- 接口说明：%s\n", e.Description) }
        if e.Auth != "" { fmt.Fprintf(b, "- 认证方式：%s\n", e.Auth) }
        if len(e.Roles) > 0 { fmt.Fprintf(b, "- 访问角色：%s\n", strings.Join(e.Roles, ", ")) }
        // 请求参数
        fmt.Fprintf(b, "\n### 请求参数\n")
        if len(e.Params) == 0 && len(e.Files) == 0 {
            fmt.Fprintf(b, "（无显式参数，或仅 Body/Headers 由业务生成）\n")
        } else {
            fmt.Fprintf(b, "| 来源 | 名称 | 类型 | 必填 | 说明 | 约束 |\n|---|---|---|---|---|---|\n")
            for _, p := range e.Params { fmt.Fprintf(b, "| %s | %s | %s | %t | %s | %s |\n", p.Source, p.Name, p.Type, p.Required, p.Description, constraintsStr(p.Constraints)) }
            for _, p := range e.Files { fmt.Fprintf(b, "| file | %s | %s | %t | %s | %s |\n", p.Name, p.Type, p.Required, p.Description, constraintsStr(p.Constraints)) }
        }
        // 请求示例
        ce, be := buildCurlExample(e)
        if ce != "" {
            fmt.Fprintf(b, "\n### 请求示例\n")
            fmt.Fprintf(b, "```bash\n%s\n```\n", ce)
            if be != "" { fmt.Fprintf(b, "```json\n%s\n```\n", be) }
        }
        // 失败请求示例
        fe := buildFailureExample(e)
        if fe != "" {
            fmt.Fprintf(b, "\n### 失败请求示例\n")
            fmt.Fprintf(b, "```bash\n%s\n```\n", fe)
        }
        // 认证说明（HMAC）
        if strings.ToLower(e.Auth) == "apikey" {
            fmt.Fprintf(b, "\n### 认证说明（HMAC-SHA256）\n")
            fmt.Fprintf(b, "必需请求头：`X-Api-Key`、`X-Timestamp(RFC3339)`、`X-Nonce`、`X-Signature`。\n")
            fmt.Fprintf(b, "计算步骤：\n")
            fmt.Fprintf(b, "1) 计算 BodyHash：`base64url(sha256(body))`（空体则为空字符串）\n")
            fmt.Fprintf(b, "2) 构造消息：\n")
            fmt.Fprintf(b, "```")
            fmt.Fprintf(b, "\n%s\n%s\n%s\n%s\n%s\n", strings.ToUpper(e.Method), e.Path, "<timestamp>", "<nonce>", "<bodyHash>")
            fmt.Fprintf(b, "```\n")
            fmt.Fprintf(b, "\n3) 计算签名：`base64url( HMAC-SHA256(secret, message) )`\n")
            sn := hmacSnippets()
            fmt.Fprintf(b, "\n#### 代码示例（Go）\n")
            fmt.Fprintf(b, "```go\n%s\n```\n", sn["go"])
            fmt.Fprintf(b, "\n#### 代码示例（JavaScript/Node.js）\n")
            fmt.Fprintf(b, "```js\n%s\n```\n", sn["js"])
            fmt.Fprintf(b, "\n#### 代码示例（Python）\n")
            fmt.Fprintf(b, "```python\n%s\n```\n", sn["py"])
        }
        // 响应参数
        fmt.Fprintf(b, "\n### 返回参数\n")
        if len(e.Response.Schema) > 0 {
            renderSchemaTable(b, e.Response.Schema, 0)
        } else {
            fmt.Fprintf(b, "（未声明结构，参考示例数据）\n")
        }
        // 错误码
        if len(e.Errors) > 0 {
            fmt.Fprintf(b, "\n### 错误码\n")
            fmt.Fprintf(b, "| Code | Message | Description |\n|---|---|---|\n")
            for _, er := range e.Errors { fmt.Fprintf(b, "| %s | %s | %s |\n", er.Code, er.Message, er.Description) }
            // 错误响应示例（取首条）
            er := e.Errors[0]
            em := map[string]any{"code": er.Code, "msg": er.Message}
            jb, _ := json.MarshalIndent(em, "", "  ")
            fmt.Fprintf(b, "\n#### 错误响应示例\n")
            fmt.Fprintf(b, "```json\n%s\n```\n", string(jb))
        }
        // 返回说明与头
        if e.Response.BodyDescription != "" { fmt.Fprintf(b, "\n- 返回说明：%s\n", e.Response.BodyDescription) }
        if len(e.Response.Headers) > 0 {
            fmt.Fprintf(b, "- 响应头：\n")
            for _, h := range e.Response.Headers { fmt.Fprintf(b, "  - %s：%s\n", h.Name, nonEmpty(h.Description, h.Value)) }
        }
        // 示例数据
        fmt.Fprintf(b, "\n### 示例数据\n")
        if len(e.Response.Schema) > 0 {
            sample := apidocs.SampleFromSchema(e.Response.Schema)
            jb, _ := json.MarshalIndent(sample, "", "  ")
            fmt.Fprintf(b, "```json\n%s\n```\n", string(jb))
        } else if e.Response.Body != "" {
            fmt.Fprintf(b, "```\n%s\n```\n", e.Response.Body)
        } else {
            fmt.Fprintf(b, "（无示例）\n")
        }
        fmt.Fprintf(b, "\n")
    }
    return b.String()
}

func hmacSnippets() map[string]string {
    goCode := `package main
import (
  "crypto/hmac"
  "crypto/sha256"
  "encoding/base64"
)
func sign(secret, method, path, ts, nonce, bodyHash string) string {
  msg := method + "\n" + path + "\n" + ts + "\n" + nonce + "\n" + bodyHash
  mac := hmac.New(sha256.New, []byte(secret))
  mac.Write([]byte(msg))
  return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}`
    jsCode := "const crypto = require('crypto')\n" +
        "function sign(secret, method, path, ts, nonce, bodyHash){\n" +
        "  const msg = method+'\\n'+path+'\\n'+ts+'\\n'+nonce+'\\n'+bodyHash\n" +
        "  const h = crypto.createHmac('sha256', secret).update(msg).digest()\n" +
        "  return h.toString('base64url')\n" +
        "}"
    pyCode := `import hmac, hashlib, base64
def sign(secret: str, method: str, path: str, ts: str, nonce: str, body_hash: str) -> str:
    msg = "\n".join([method, path, ts, nonce, body_hash]).encode()
    dig = hmac.new(secret.encode(), msg, hashlib.sha256).digest()
    return base64.urlsafe_b64encode(dig).decode().rstrip('=')`
    javaCode := `import javax.crypto.Mac; import javax.crypto.spec.SecretKeySpec; import java.util.Base64;
public class Signer {
  public static String sign(String secret, String method, String path, String ts, String nonce, String bodyHash) throws Exception {
    String msg = String.join("\n", method, path, ts, nonce, bodyHash);
    Mac mac = Mac.getInstance("HmacSHA256");
    mac.init(new SecretKeySpec(secret.getBytes(), "HmacSHA256"));
    byte[] dig = mac.doFinal(msg.getBytes());
    return Base64.getUrlEncoder().withoutPadding().encodeToString(dig);
  }
}`
    csCode := `using System; using System.Security.Cryptography; using System.Text;
public class Signer {
  public static string Sign(string secret, string method, string path, string ts, string nonce, string bodyHash){
    var msg = string.Join("\n", method, path, ts, nonce, bodyHash);
    using var hmac = new HMACSHA256(Encoding.UTF8.GetBytes(secret));
    var dig = hmac.ComputeHash(Encoding.UTF8.GetBytes(msg));
    return Convert.ToBase64String(dig).TrimEnd('=').Replace('+','-').Replace('/','_');
  }
}`
    return map[string]string{"go": goCode, "js": jsCode, "py": pyCode, "java": javaCode, "cs": csCode}
}

func renderSchemaTable(b *strings.Builder, nodes []apidocs.DocSchema, indent int) {
    for _, s := range nodes {
        pad := strings.Repeat("  ", indent)
        if s.Description != "" {
            fmt.Fprintf(b, "%s- %s (%s)：%s\n", pad, s.Name, s.Type, s.Description)
        } else {
            fmt.Fprintf(b, "%s- %s (%s)\n", pad, s.Name, s.Type)
        }
        if len(s.Children) > 0 { renderSchemaTable(b, s.Children, indent+1) }
    }
}

func constraintsStr(m map[string]string) string {
    if len(m) == 0 { return "" }
    kv := make([]string, 0, len(m))
    for k, v := range m { kv = append(kv, fmt.Sprintf("%s=%s", k, v)) }
    return strings.Join(kv, ", ")
}

func nonEmpty(a, b string) string { if a != "" { return a } ; return b }
func safe(s string) string {
    if s == "" { return "default" }
    r := make([]byte, 0, len(s))
    for i := 0; i < len(s); i++ { c := s[i]; if (c>='a'&&c<='z')||(c>='A'&&c<='Z')||(c>='0'&&c<='9')||c=='-'||c=='_' { r = append(r, c) } else { r = append(r, '_') } }
    return string(r)
}

func endpointFileName(e apidocs.DocEndpoint) string {
    // 文件名：<METHOD>_<PATH>.md，路径非字母数字替换为下划线
    p := e.Path
    r := make([]byte, 0, len(p))
    for i := 0; i < len(p); i++ {
        c := p[i]
        if (c>='a'&&c<='z')||(c>='A'&&c<='Z')||(c>='0'&&c<='9') { r = append(r, c) } else { r = append(r, '_') }
    }
    return fmt.Sprintf("%s_%s.md", strings.ToUpper(e.Method), string(r))
}

func buildCurlExample(e apidocs.DocEndpoint) (string, string) {
    method := strings.ToUpper(e.Method)
    urlPath := e.Path
    qp := []string{}
    headers := []string{}
    body := map[string]any{}
    for _, p := range e.Params {
        switch strings.ToLower(p.Source) {
        case "query":
            qp = append(qp, fmt.Sprintf("%s=%v", p.Name, sampleByType(p.Type)))
        case "path":
            ph := "{" + p.Name + "}"
            if strings.Contains(urlPath, ph) {
                urlPath = strings.ReplaceAll(urlPath, ph, fmt.Sprintf("%v", sampleByType(p.Type)))
            }
        case "header":
            headers = append(headers, fmt.Sprintf("%s: %v", p.Name, sampleByType(p.Type)))
        case "body":
            body[p.Name] = sampleByType(p.Type)
        }
    }
    if strings.ToLower(e.Auth) == "jwt" {
        headers = append(headers, "Authorization: Bearer <token>")
    } else if strings.ToLower(e.Auth) == "apikey" {
        headers = append(headers, "X-API-Key: <key>")
        headers = append(headers, "X-Signature: <hmac>")
    }
    url := "http://localhost:8080" + urlPath
    if len(qp) > 0 {
        url = url + "?" + strings.Join(qp, "&")
    }
    curl := &strings.Builder{}
    fmt.Fprintf(curl, "curl -X %s \"%s\"", method, url)
    for _, h := range headers { fmt.Fprintf(curl, " -H \"%s\"", h) }
    bodyStr := ""
    if len(body) > 0 {
        jb, _ := json.MarshalIndent(body, "", "  ")
        bodyStr = string(jb)
        fmt.Fprintf(curl, " -H \"Content-Type: application/json\" -d @- << 'JSON'\n%s\nJSON", bodyStr)
        return curl.String(), bodyStr
    }
    return curl.String(), ""
}

func buildFailureExample(e apidocs.DocEndpoint) string {
    method := strings.ToUpper(e.Method)
    url := "http://localhost:8080" + e.Path
    curl := &strings.Builder{}
    fmt.Fprintf(curl, "curl -X %s \"%s\"", method, url)
    if strings.ToLower(e.Auth) == "jwt" {
        fmt.Fprintf(curl, " -H \"Authorization: Bearer invalid\"")
    } else if strings.ToLower(e.Auth) == "apikey" {
        fmt.Fprintf(curl, " -H \"X-API-Key: <key>\" -H \"X-Signature: invalid\"")
    }
    return curl.String()
}

func sampleByType(t string) any {
    lt := strings.ToLower(strings.TrimSpace(t))
    switch lt {
    case "int", "integer":
        return 1
    case "number", "float", "double":
        return 1.23
    case "bool", "boolean":
        return true
    case "array":
        return []any{"string"}
    case "object":
        return map[string]any{"key": "value"}
    default:
        return "string"
    }
}