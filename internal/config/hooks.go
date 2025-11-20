package config

import (
    "io/fs"
    "os"
    "path/filepath"
    "gopkg.in/yaml.v3"
)

type HooksBundle struct{ Auth []map[string]any; Before []map[string]any; After []map[string]any }

func LoadHooksConfig(root string) (*HooksBundle, error) {
    b := &HooksBundle{}
    _ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
        if err != nil { return nil }
        if d.IsDir() { return nil }
        ext := filepath.Ext(path)
        if ext != ".yaml" && ext != ".yml" && ext != ".json" { return nil }
        data, err := os.ReadFile(path)
        if err != nil { return nil }
        var hf struct{ Auth []map[string]any `yaml:"auth"`; Before []map[string]any `yaml:"before"`; After []map[string]any `yaml:"after"` }
        if err := yaml.Unmarshal(data, &hf); err != nil { return nil }
        for _, m := range hf.Auth { expandMap(m); b.Auth = append(b.Auth, m) }
        for _, m := range hf.Before { expandMap(m); b.Before = append(b.Before, m) }
        for _, m := range hf.After { expandMap(m); b.After = append(b.After, m) }
        return nil
    })
    return b, nil
}