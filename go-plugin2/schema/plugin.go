package schema

import "time"

const StatusFileName = "status.json"

type PluginStatus string

const (
	Enabled  PluginStatus = "enabled"
	Disabled PluginStatus = "disabled"
	Error    PluginStatus = "error"
)

type PluginInstance struct {
	Meta  PluginMeta
	State PluginState
}

type PluginMeta struct {
	Name    string
	Version string
	Entry   string
	Enabled bool
	Timeout time.Duration

	ConfigPath string `json:"-" yaml:"-"`
}

type PluginState struct {
	Status       string `json:"status"` // idle / running / disabled / error
	LastError    string `json:"last_error"`
	ExecCount    int    `json:"exec_count"`
	LastExecTime string `json:"last_exec_time"`
}
