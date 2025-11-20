package core

import "context"

// 请求上下文 / Request context
type RequestContext struct {
    Base context.Context
    Vars map[string]any
}

// 创建上下文 / Create new context
func NewRequestContext(base context.Context) *RequestContext {
    return &RequestContext{Base: base, Vars: make(map[string]any)}
}