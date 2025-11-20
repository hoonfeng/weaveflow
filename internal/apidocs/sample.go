package apidocs

func SampleFromSchema(nodes []DocSchema) any {
    if len(nodes) == 0 { return map[string]any{} }
    out := map[string]any{}
    for _, n := range nodes {
        out[n.Name] = sampleValue(n)
    }
    return out
}

func sampleValue(n DocSchema) any {
    switch lower(n.Type) {
    case "string":
        return "string"
    case "integer", "int":
        return 123
    case "number", "float", "double":
        return 1.23
    case "boolean", "bool":
        return true
    case "array":
        if len(n.Children) > 0 {
            return []any{sampleValue(n.Children[0])}
        }
        return []any{"string"}
    case "object":
        m := map[string]any{}
        for _, c := range n.Children { m[c.Name] = sampleValue(c) }
        return m
    default:
        if len(n.Children) > 0 {
            m := map[string]any{}
            for _, c := range n.Children { m[c.Name] = sampleValue(c) }
            return m
        }
        return "string"
    }
}

func lower(s string) string {
    b := make([]byte, 0, len(s))
    for i := 0; i < len(s); i++ { c := s[i]; if c >= 'A' && c <= 'Z' { b = append(b, c+('a'-'A')) } else { b = append(b, c) } }
    return string(b)
}