package router

import (
    "net/http"
    "strings"
    jwt "github.com/golang-jwt/jwt/v5"
    "ifaceconf/internal/config"
)

func requireJWTWithRoles(secret string, roles []string, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        auth := r.Header.Get("Authorization")
        if !strings.HasPrefix(auth, "Bearer ") { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        tokenStr := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
        t, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) { return []byte(secret), nil })
        if err != nil || !t.Valid { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        if len(roles) > 0 {
            if claims, ok := t.Claims.(jwt.MapClaims); ok {
                if !hasAnyRole(claims, roles) { http.Error(w, "forbidden", http.StatusForbidden); return }
            }
        }
        next(w, r)
    }
}

func requireJWTWithRolesConfig(sec config.SecurityConfig, roles []string, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        auth := r.Header.Get("Authorization")
        if !strings.HasPrefix(auth, "Bearer ") { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        tokenStr := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
        t, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) { return []byte(sec.JWTSecret), nil })
        if err != nil || !t.Valid { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
        if claims, ok := t.Claims.(jwt.MapClaims); ok {
            if sec.JWTIssuer != "" {
                if iss, _ := claims["iss"].(string); iss != sec.JWTIssuer { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
            }
            if sec.JWTAudience != "" {
                switch aud := claims["aud"].(type) {
                case string:
                    if aud != sec.JWTAudience { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
                case []any:
                    okAud := false
                    for _, x := range aud { if s, ok := x.(string); ok && s == sec.JWTAudience { okAud = true; break } }
                    if !okAud { http.Error(w, "unauthorized", http.StatusUnauthorized); return }
                }
            }
        }
        if len(roles) > 0 {
            if claims, ok := t.Claims.(jwt.MapClaims); ok {
                if !hasAnyRole(claims, roles) { http.Error(w, "forbidden", http.StatusForbidden); return }
            }
        }
        next(w, r)
    }
}

func hasAnyRole(claims jwt.MapClaims, roles []string) bool {
    m := map[string]struct{}{}
    for _, r := range roles { m[r] = struct{}{} }
    if v, ok := claims["roles"]; ok {
        switch arr := v.(type) {
        case []any:
            for _, x := range arr { if s, ok := x.(string); ok { if _, ex := m[s]; ex { return true } } }
        case []string:
            for _, s := range arr { if _, ex := m[s]; ex { return true } }
        }
    }
    if v, ok := claims["role"].(string); ok { _, ex := m[v]; return ex }
    return false
}
