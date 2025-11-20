package config

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// 接口配置结构 / Interface configuration root
type InterfaceConfig struct {
    Module      string           `yaml:"module" json:"module"`
    Endpoint    string           `yaml:"endpoint" json:"endpoint"`
    Method      string           `yaml:"method" json:"method"`
    Path        string           `yaml:"path" json:"path"`
    Auth        string           `yaml:"auth" json:"auth"`
    Roles       []string         `yaml:"roles" json:"roles"`
    Permissions []string         `yaml:"permissions" json:"permissions"`
    Labels      []string         `yaml:"labels" json:"labels"`
    Docs        map[string]any   `yaml:"docs" json:"docs"`
    Request     map[string]any   `yaml:"request" json:"request"`
    Steps       []map[string]any `yaml:"steps" json:"steps"`
    Errors      map[string]any   `yaml:"errors" json:"errors"`
    Transaction map[string]any   `yaml:"transaction" json:"transaction"`
    RateLimit   map[string]any   `yaml:"rateLimit" json:"rateLimit"`
    Hooks       HooksOverride    `yaml:"hooks" json:"hooks"`
}

type HooksOverride struct {
	Disable []string         `yaml:"disable" json:"disable"`
	Auth    []map[string]any `yaml:"auth" json:"auth"`
	Before  []map[string]any `yaml:"before" json:"before"`
	After   []map[string]any `yaml:"after" json:"after"`
}

// 加载所有接口配置 / Load all interface configs
func LoadInterfaceConfigs(root string) ([]*InterfaceConfig, error) {
	var out []*InterfaceConfig
	errs := []string{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			errs = append(errs, fmt.Sprintf("walk error: %s: %v", path, err))
			return nil
		}
		if d.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".yaml" && ext != ".yml" && ext != ".json" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			errs = append(errs, fmt.Sprintf("read error: %s: %v", path, err))
			return nil
		}
		var ic InterfaceConfig
		if err := yaml.Unmarshal(data, &ic); err != nil {
			errs = append(errs, fmt.Sprintf("parse error: %s: %v", path, err))
			return nil
		}
		if ic.Request != nil {
			expandMap(ic.Request)
		}
		if ic.Docs != nil {
			expandMap(ic.Docs)
		}
		if ic.Transaction != nil {
			expandMap(ic.Transaction)
		}
		if ic.Module == "" || ic.Path == "" || ic.Method == "" {
			errs = append(errs, fmt.Sprintf("missing required fields: %s", path))
			return nil
		}
		out = append(out, &ic)
		return nil
	})
	if err != nil {
		return nil, err
	}
    if len(errs) > 0 {
        return nil, fmt.Errorf("接口配置加载失败: %s", strings.Join(errs, "; "))
    }
    return out, nil
}
