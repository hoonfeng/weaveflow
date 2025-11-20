package plugin

import (
    "fmt"
    "ifaceconf/internal/config"
)

// 运行时接口 / Runtime interface
type Runtime interface {
    Call(plugin string, function string, params map[string]any) (any, error)
    Status() map[string]any
    SetEnabled(name string, enabled bool)
    AddPlugin(item config.PluginRegistryItem) error
    RemovePlugin(name string) error
    RestartPlugin(name string) error
    StopPlugin(name string) error
    StartPlugin(name string) error
}

// 空实现（占位） / No-op runtime (placeholder)
type NoopRuntime struct{}

func (n *NoopRuntime) Call(plugin string, function string, params map[string]any) (any, error) {
    return nil, fmt.Errorf("插件未启用或未实现 / plugin runtime not enabled")
}

func (n *NoopRuntime) Status() map[string]any { return map[string]any{"enabled": false} }

func (n *NoopRuntime) SetEnabled(name string, enabled bool) {}
func (n *NoopRuntime) AddPlugin(item config.PluginRegistryItem) error { return fmt.Errorf("plugin runtime not enabled") }
func (n *NoopRuntime) RemovePlugin(name string) error { return fmt.Errorf("plugin runtime not enabled") }
func (n *NoopRuntime) RestartPlugin(name string) error { return fmt.Errorf("plugin runtime not enabled") }
func (n *NoopRuntime) StopPlugin(name string) error { return fmt.Errorf("plugin runtime not enabled") }
func (n *NoopRuntime) StartPlugin(name string) error { return fmt.Errorf("plugin runtime not enabled") }
