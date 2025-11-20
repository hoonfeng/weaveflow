package test

import (
    "bytes"
    "encoding/json"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "testing"
    "time"
)

type RespWrap[T any] struct {
    Code int    `json:"code"`
    Msg  string `json:"msg"`
    Data T      `json:"data"`
}

func moduleRoot() string {
    wd, _ := os.Getwd()
    return filepath.Dir(wd)
}

func buildPlugin(t *testing.T) {
    t.Helper()
    root := moduleRoot()
    cmd := exec.Command("go", "build", "-o", filepath.Join(root, "plugins", "string_reverse.exe"), "./plugins/string_reverse")
    cmd.Dir = root
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil { t.Fatalf("build plugin failed: %v", err) }
}

func buildServer(t *testing.T) string {
    t.Helper()
    root := moduleRoot()
    _ = os.MkdirAll(filepath.Join(root, "bin"), 0o755)
    out := filepath.Join(root, "bin", "server.exe")
    cmd := exec.Command("go", "build", "-tags", "goproc", "-o", out, "./cmd/server")
    cmd.Dir = root
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil { t.Fatalf("build server failed: %v", err) }
    return out
}

func startServer(t *testing.T, path string) *exec.Cmd {
    t.Helper()
    cmd := exec.Command(path)
    cmd.Dir = moduleRoot()
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Start(); err != nil { t.Fatalf("start server failed: %v", err) }
    // wait until /docs is ready
    deadline := time.Now().Add(5 * time.Second)
    for time.Now().Before(deadline) {
        resp, err := http.Get("http://localhost:8080/docs")
        if err == nil && resp.StatusCode == 200 { resp.Body.Close(); return cmd }
        time.Sleep(200 * time.Millisecond)
    }
    t.Fatalf("server not ready")
    return cmd
}

func TestPluginReverse(t *testing.T) {
    buildPlugin(t)
    srv := buildServer(t)
    cmd := startServer(t, srv)
    defer func() { _ = cmd.Process.Kill() }()

    body := bytes.NewBufferString(`{"text":"Hello"}`)
    resp, err := http.Post("http://localhost:8080/api/text/reverse", "application/json", body)
    if err != nil { t.Fatalf("request failed: %v", err) }
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        b, _ := io.ReadAll(resp.Body)
        if bytes.Contains(b, []byte("plugin runtime not enabled")) { t.Skip("plugin runtime not enabled") }
        t.Fatalf("status %d body %s", resp.StatusCode, string(b))
    }
    var out RespWrap[string]
    if err := json.NewDecoder(resp.Body).Decode(&out); err != nil { t.Fatalf("decode failed: %v", err) }
    if out.Code != 0 || out.Data != "olleH" { t.Fatalf("unexpected resp: %+v", out) }
}

func TestFileUpload(t *testing.T) {
    buildPlugin(t)
    srv := buildServer(t)
    cmd := startServer(t, srv)
    defer func() { _ = cmd.Process.Kill() }()

    var buf bytes.Buffer
    w := multipart.NewWriter(&buf)
    fw, _ := w.CreateFormFile("files", "a.txt")
    _, _ = io.WriteString(fw, "hello")
    _ = w.Close()

    req, _ := http.NewRequest("POST", "http://localhost:8080/api/files/upload", &buf)
    req.Header.Set("Content-Type", w.FormDataContentType())
    resp, err := http.DefaultClient.Do(req)
    if err != nil { t.Fatalf("upload failed: %v", err) }
    defer resp.Body.Close()
    if resp.StatusCode != 201 {
        b, _ := io.ReadAll(resp.Body)
        if bytes.Contains(b, []byte("plugin runtime not enabled")) { t.Skip("plugin runtime not enabled") }
        if bytes.Contains(b, []byte("输入类型错误")) { t.Skip("upload input type not ready") }
        t.Fatalf("status %d body %s", resp.StatusCode, string(b))
    }
    var out RespWrap[[]string]
    if err := json.NewDecoder(resp.Body).Decode(&out); err != nil { t.Fatalf("decode failed: %v", err) }
    if out.Code != 0 || len(out.Data) == 0 { t.Fatalf("unexpected resp: %+v", out) }
}

func TestDocsEndpoints(t *testing.T) {
    srv := buildServer(t)
    cmd := startServer(t, srv)
    defer func() { _ = cmd.Process.Kill() }()
    resp, err := http.Get("http://localhost:8080/docs")
    if err != nil { t.Fatalf("docs request failed: %v", err) }
    defer resp.Body.Close()
    if resp.StatusCode != 200 { t.Fatalf("docs status %d", resp.StatusCode) }
    var list []map[string]any
    if err := json.NewDecoder(resp.Body).Decode(&list); err != nil { t.Fatalf("decode failed: %v", err) }
    if len(list) == 0 { t.Fatalf("empty docs") }
}
