package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	gcfg "github.com/hoonfeng/goproc/config"
	gpm "github.com/hoonfeng/goproc/plugin"

	"ifaceconf/internal/config"
)

type GoprocRuntime struct {
	mu       sync.RWMutex
	manager  *gpm.PluginManager
	disabled map[string]bool
}

func BuildRuntime(cfg *config.PluginsConfig) (Runtime, error) {
	if cfg == nil || !cfg.Enabled {
		return &NoopRuntime{}, nil
	}
	if strings.ToLower(cfg.Runtime) != "goproc" {
		return &NoopRuntime{}, nil
	}

	sys := &gcfg.SystemConfig{
		Version: "v1",
		Plugins: map[string]gcfg.PluginConfig{},
		System: gcfg.SystemSettings{
			MaxConcurrentCalls: 0,
			CallTimeout:        firstNonEmpty(cfgTimeouts(cfg), "300ms"),
		},
		Platform:    gcfg.PlatformConfig{},
		Development: gcfg.DevelopmentConfig{DebugMode: false},
	}

	for _, item := range cfg.Registry {
		path := item.Executable.Windows
		if runtime.GOOS != "windows" {
			path = item.Executable.Unix
		}
		if abs, err := filepath.Abs(path); err == nil {
			path = abs
		}
		// build environment for plugin process
		env := map[string]string{}
		// EnvFrom: explicit keys or "*" for all
		if len(item.EnvFrom) > 0 {
			hasAll := false
			for _, k := range item.EnvFrom {
				if k == "*" {
					hasAll = true
					break
				}
			}
			if hasAll {
				for _, kv := range envPairs() {
					env[kv[0]] = kv[1]
				}
			} else {
				for _, k := range item.EnvFrom {
					env[k] = getenv(k)
				}
			}
		}
		for k, v := range item.Env {
			env[k] = v
		}

		pc := gcfg.PluginConfig{
			Type:         gcfg.PluginTypeBinary,
			Path:         path,
			PoolSize:     max(1, item.Instances),
			MaxInstances: max(1, item.Instances),
			Functions:    item.Functions,
			Args:         []string{},
			Environment:  env,
		}
		sys.Plugins[item.Name] = pc
		if item.QueueSize > 0 && sys.System.MaxConcurrentCalls == 0 {
			sys.System.MaxConcurrentCalls = item.QueueSize
		}
		if sys.System.CallTimeout == "" {
			sys.System.CallTimeout = item.Timeout
		}
	}

	if err := gcfg.ValidateConfig(sys); err != nil {
		return &NoopRuntime{}, fmt.Errorf("goproc config invalid: %w", err)
	}
	pm := gpm.NewPluginManager(sys)
	if err := pm.Start(); err != nil {
		return &NoopRuntime{}, fmt.Errorf("goproc manager start failed: %w", err)
	}
	return &GoprocRuntime{manager: pm, disabled: map[string]bool{}}, nil
}

func (g *GoprocRuntime) Call(plugin string, function string, params map[string]any) (any, error) {
	g.mu.RLock()
	pm := g.manager
	g.mu.RUnlock()
	if pm == nil {
		return nil, fmt.Errorf("goproc manager not initialized")
	}
	g.mu.RLock()
	dis := g.disabled[plugin]
	g.mu.RUnlock()
	if dis {
		return nil, fmt.Errorf("plugin disabled")
	}
	return pm.CallFunction(plugin, function, params)
}

func (g *GoprocRuntime) Status() map[string]any {
	g.mu.RLock()
	pm := g.manager
	g.mu.RUnlock()
	if pm == nil {
		return map[string]any{"enabled": false}
	}
	st := pm.GetAllStatus()
	g.mu.RLock()
	dis := g.disabled
	g.mu.RUnlock()
	if dis != nil {
		st["disabled"] = dis
	}
	return st
}

func (g *GoprocRuntime) SetEnabled(name string, enabled bool) {
	g.mu.Lock()
	if g.disabled == nil {
		g.disabled = map[string]bool{}
	}
	g.disabled[name] = !enabled
	g.mu.Unlock()
}

func (g *GoprocRuntime) AddPlugin(item config.PluginRegistryItem) error {
	g.mu.RLock()
	pm := g.manager
	g.mu.RUnlock()
	if pm == nil {
		return fmt.Errorf("goproc manager not initialized")
	}
	path := item.Executable.Windows
	if runtime.GOOS != "windows" {
		path = item.Executable.Unix
	}
	if abs, err := filepath.Abs(path); err == nil {
		path = abs
	}
	env := map[string]string{}
	if len(item.EnvFrom) > 0 {
		hasAll := false
		for _, k := range item.EnvFrom {
			if k == "*" {
				hasAll = true
				break
			}
		}
		if hasAll {
			for _, kv := range envPairs() {
				env[kv[0]] = kv[1]
			}
		} else {
			for _, k := range item.EnvFrom {
				env[k] = getenv(k)
			}
		}
	}
	for k, v := range item.Env {
		env[k] = v
	}
	pc := gcfg.PluginConfig{Type: gcfg.PluginTypeBinary, Path: path, PoolSize: max(1, item.Instances), MaxInstances: max(1, item.Instances), Functions: item.Functions, Args: []string{}, Environment: env}
	return pm.AddPlugin(item.Name, pc)
}

func (g *GoprocRuntime) RemovePlugin(name string) error {
	g.mu.RLock()
	pm := g.manager
	g.mu.RUnlock()
	if pm == nil {
		return fmt.Errorf("goproc manager not initialized")
	}
	return pm.RemovePlugin(name)
}

func (g *GoprocRuntime) RestartPlugin(name string) error {
	g.mu.RLock()
	pm := g.manager
	g.mu.RUnlock()
	if pm == nil {
		return fmt.Errorf("goproc manager not initialized")
	}
	return pm.RestartPlugin(name)
}

func (g *GoprocRuntime) StopPlugin(name string) error {
	g.mu.RLock()
	pm := g.manager
	g.mu.RUnlock()
	if pm == nil {
		return fmt.Errorf("goproc manager not initialized")
	}
	return pm.StopPlugin(name)
}

func (g *GoprocRuntime) StartPlugin(name string) error {
	g.mu.RLock()
	pm := g.manager
	g.mu.RUnlock()
	if pm == nil {
		return fmt.Errorf("goproc manager not initialized")
	}
	return pm.StartPlugin(name)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func cfgTimeouts(cfg *config.PluginsConfig) string {
	if cfg == nil {
		return ""
	}
	if len(cfg.Registry) == 0 {
		return ""
	}
	return cfg.Registry[0].Timeout
}

func envPairs() [][2]string {
	arr := make([][2]string, 0, 64)
	for _, s := range os.Environ() {
		i := strings.Index(s, "=")
		if i <= 0 {
			continue
		}
		k := s[:i]
		v := s[i+1:]
		arr = append(arr, [2]string{k, v})
	}
	return arr
}

// indirection for testing/mocking
var getenv = func(k string) string { return os.Getenv(k) }
