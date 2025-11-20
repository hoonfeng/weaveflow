package config

import (
    "os"
    "strings"
)

func expandString(s string) string {
    out := s
    for {
        i := strings.Index(out, "${")
        if i < 0 { break }
        j := strings.Index(out[i:], "}")
        if j < 0 { break }
        token := out[i+2 : i+j]
        key := token
        def := ""
        if k := strings.Index(token, ":"); k >= 0 { key = token[:k]; def = token[k+1:] }
        val := os.Getenv(key)
        if val == "" { val = def }
        out = out[:i] + val + out[i+j+1:]
    }
    return out
}

func expandMap(m map[string]any) {
    for k, v := range m {
        switch t := v.(type) {
        case string:
            m[k] = expandString(t)
        case map[string]any:
            expandMap(t)
        }
    }
}