package config

func LintInterfaces(list []*InterfaceConfig) []string {
    issues := []string{}
    seen := map[string]struct{}{}
    for _, ic := range list {
        if ic.Method == "" || ic.Path == "" { issues = append(issues, "missing method or path") ; continue }
        key := ic.Method+" "+ic.Path
        if _, ok := seen[key]; ok { issues = append(issues, "duplicate "+key) } else { seen[key] = struct{}{} }
    }
    return issues
}