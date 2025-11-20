package obs

import (
    "fmt"
    "net/http"
    "sync"
    "time"
    "strings"
)

type counters struct {
    mu   sync.Mutex
    req  map[string]int64            // path → count
    code map[string]map[int]int64    // path → status → count
    lat  map[string]time.Duration    // path → total latency
    mod  map[string]string           // key(method path) → module
}

var m = &counters{req: map[string]int64{}, code: map[string]map[int]int64{}, lat: map[string]time.Duration{}, mod: map[string]string{}}

// Middleware 记录请求数量与状态码、累计延迟
func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        rw := &responseRecorder{ResponseWriter: w, status: 200}
        next.ServeHTTP(rw, r)
        path := r.URL.Path
        method := r.Method
        key := method+" "+path
        dur := time.Since(start)
        m.mu.Lock()
        m.req[key]++
        if m.code[key] == nil { m.code[key] = map[int]int64{} }
        m.code[key][rw.status]++
        m.lat[key] += dur
        m.mu.Unlock()
    })
}

type responseRecorder struct {
    http.ResponseWriter
    status int
}

func (r *responseRecorder) WriteHeader(code int) {
    r.status = code
    r.ResponseWriter.WriteHeader(code)
}

// Handler 以 Prometheus 文本格式导出基础指标
func Handler(w http.ResponseWriter, _ *http.Request) {
    w.Header().Set("Content-Type", "text/plain; version=0.0.4")
    m.mu.Lock()
    defer m.mu.Unlock()
    // total requests per path
    for k, c := range m.req {
        parts := splitKey(k)
        mod := m.mod[k]
        fmt.Fprintf(w, "# HELP http_requests_total Total HTTP requests\n")
        fmt.Fprintf(w, "# TYPE http_requests_total counter\n")
        fmt.Fprintf(w, "http_requests_total{method=\"%s\",path=\"%s\",module=\"%s\"} %d\n", parts.method, parts.path, mod, c)
    }
    // status codes per path
    for k, mp := range m.code {
        parts := splitKey(k)
        mod := m.mod[k]
        fmt.Fprintf(w, "# HELP http_responses_total HTTP responses by status\n")
        fmt.Fprintf(w, "# TYPE http_responses_total counter\n")
        for code, c := range mp {
            fmt.Fprintf(w, "http_responses_total{method=\"%s\",path=\"%s\",module=\"%s\",status=\"%d\"} %d\n", parts.method, parts.path, mod, code, c)
        }
    }
    // average latency per path (milliseconds)
    for k, tot := range m.lat {
        parts := splitKey(k)
        mod := m.mod[k]
        cnt := m.req[k]
        avgMs := float64(tot.Milliseconds())
        if cnt > 0 { avgMs = avgMs / float64(cnt) }
        fmt.Fprintf(w, "# HELP http_latency_ms Average latency in ms\n")
        fmt.Fprintf(w, "# TYPE http_latency_ms gauge\n")
        fmt.Fprintf(w, "http_latency_ms{method=\"%s\",path=\"%s\",module=\"%s\"} %.3f\n", parts.method, parts.path, mod, avgMs)
    }
}

func Render() string {
    m.mu.Lock()
    defer m.mu.Unlock()
    var b strings.Builder
    for k, c := range m.req {
        parts := splitKey(k)
        mod := m.mod[k]
        fmt.Fprintf(&b, "# HELP http_requests_total Total HTTP requests\n")
        fmt.Fprintf(&b, "# TYPE http_requests_total counter\n")
        fmt.Fprintf(&b, "http_requests_total{method=\"%s\",path=\"%s\",module=\"%s\"} %d\n", parts.method, parts.path, mod, c)
    }
    for k, mp := range m.code {
        parts := splitKey(k)
        mod := m.mod[k]
        fmt.Fprintf(&b, "# HELP http_responses_total HTTP responses by status\n")
        fmt.Fprintf(&b, "# TYPE http_responses_total counter\n")
        for code, c := range mp {
            fmt.Fprintf(&b, "http_responses_total{method=\"%s\",path=\"%s\",module=\"%s\",status=\"%d\"} %d\n", parts.method, parts.path, mod, code, c)
        }
    }
    for k, tot := range m.lat {
        parts := splitKey(k)
        mod := m.mod[k]
        cnt := m.req[k]
        avgMs := float64(tot.Milliseconds())
        if cnt > 0 { avgMs = avgMs / float64(cnt) }
        fmt.Fprintf(&b, "# HELP http_latency_ms Average latency in ms\n")
        fmt.Fprintf(&b, "# TYPE http_latency_ms gauge\n")
        fmt.Fprintf(&b, "http_latency_ms{method=\"%s\",path=\"%s\",module=\"%s\"} %.3f\n", parts.method, parts.path, mod, avgMs)
    }
    return b.String()
}

type keyParts struct{ method, path string }

func splitKey(k string) keyParts {
    for i := 0; i < len(k); i++ { if k[i] == ' ' { return keyParts{method: k[:i], path: k[i+1:]} } }
    return keyParts{path: k}
}

func SetRouteLabels(labels map[string]string) {
    m.mu.Lock()
    for k, v := range labels { m.mod[k] = v }
    m.mu.Unlock()
}