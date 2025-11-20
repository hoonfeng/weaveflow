package ext

import (
    "net/http"
    "time"
    "ifaceconf/internal/config"
    "ifaceconf/internal/core"
    "ifaceconf/internal/datasource"
    "ifaceconf/internal/plugin"
)

type Hooks struct {
    Auth   func(project *config.ProjectConfig, ic *config.InterfaceConfig, ds *datasource.Manager, prt plugin.Runtime, base http.HandlerFunc) http.HandlerFunc
    Before func(project *config.ProjectConfig, ic *config.InterfaceConfig, ds *datasource.Manager, prt plugin.Runtime, r *http.Request) error
    After  func(project *config.ProjectConfig, ic *config.InterfaceConfig, ds *datasource.Manager, prt plugin.Runtime, r *http.Request, res *core.ExecResult, dur time.Duration)
}

var cur Hooks

func SetHooks(h Hooks) { cur = h }

func GetHooks() Hooks { return cur }