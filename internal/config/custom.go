package config

import (
    "io/fs"
    "os"
    "path/filepath"
    "gopkg.in/yaml.v3"
)

var customData map[string]any

func LoadCustomConfig(root string) (map[string]any, error) {
    out := map[string]any{}
    _ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
        if err != nil { return nil }
        if d.IsDir() { return nil }
        ext := filepath.Ext(path)
        if ext != ".yaml" && ext != ".yml" && ext != ".json" { return nil }
        data, err := os.ReadFile(path)
        if err != nil { return nil }
        m := map[string]any{}
        if err := yaml.Unmarshal(data, &m); err != nil { return nil }
        for k, v := range m { out[k] = v }
        return nil
    })
    customData = out
    return out, nil
}

func GetCustom() map[string]any { return customData }
func SetCustom(m map[string]any) { customData = m }