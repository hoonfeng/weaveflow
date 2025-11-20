package router

import (
    "sync"
    "time"
    "net/http"
)

type Limiter struct {
    rate  float64
    burst float64
    tokens float64
    last  time.Time
    mu    sync.Mutex
}

func NewLimiter(rps float64, burst int) *Limiter {
    return &Limiter{rate: rps, burst: float64(burst), tokens: float64(burst), last: time.Now()}
}

func (l *Limiter) Allow() bool {
    l.mu.Lock()
    defer l.mu.Unlock()
    now := time.Now()
    elapsed := now.Sub(l.last).Seconds()
    l.tokens += elapsed * l.rate
    if l.tokens > l.burst { l.tokens = l.burst }
    l.last = now
    if l.tokens >= 1 {
        l.tokens -= 1
        return true
    }
    return false
}

func NewGlobalMiddleware(rps float64, burst int, perIp bool) func(next http.Handler) http.Handler {
    base := NewLimiter(rps, burst)
    lm := map[string]*Limiter{}
    mu := sync.Mutex{}
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if rps <= 0 { next.ServeHTTP(w, r); return }
            if !perIp {
                if !base.Allow() { http.Error(w, "too many requests", http.StatusTooManyRequests); return }
                next.ServeHTTP(w, r)
                return
            }
            key := r.RemoteAddr
            mu.Lock()
            lim := lm[key]
            if lim == nil { lim = NewLimiter(base.rate, int(base.burst)); lm[key] = lim }
            mu.Unlock()
            if !lim.Allow() { http.Error(w, "too many requests", http.StatusTooManyRequests); return }
            next.ServeHTTP(w, r)
        })
    }
}