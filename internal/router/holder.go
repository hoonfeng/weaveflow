package router

import (
    "net/http"
    "sync/atomic"
)

type HandlerHolder struct {
    h atomic.Value // http.Handler
}

func NewHandlerHolder() *HandlerHolder { return &HandlerHolder{} }

func (hh *HandlerHolder) Set(h http.Handler) { hh.h.Store(h) }

func (hh *HandlerHolder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if v := hh.h.Load(); v != nil {
        v.(http.Handler).ServeHTTP(w, r)
        return
    }
    http.NotFound(w, r)
}