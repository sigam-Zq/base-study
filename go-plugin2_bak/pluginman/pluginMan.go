package pluginman

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"plugin-sche/schema"
	"runtime"
	"strings"
	"time"
)

// PluginManager coordinates plugin configuration and status.
type PluginManager struct {
	configManager *ConfigManager
	statusManager StatusManager
	workDir       string
	ctx           context.Context
}

// NewPluginManager creates a new PluginManager.
func NewPluginManager(ctx context.Context, pluginDir, workDir string) (*PluginManager, error) {
	statusPath := filepath.Join(workDir, schema.StatusFileName)

	cm := NewConfigManager(pluginDir)
	sm, err := NewFileStatusManager(statusPath)
	if err != nil {
		return nil, err
	}

	return &PluginManager{
		configManager: cm,
		statusManager: sm,
		workDir:       workDir,
		ctx:           ctx,
	}, nil
}

// GetAllPlugins retrieves all plugin configurations and merges them with their current status.
func (pm *PluginManager) GetAllPlugins() ([]*schema.PluginInstance, error) {
	configs, err := pm.configManager.ReadConfigs()
	if err != nil {
		return nil, err
	}

	statuses := pm.statusManager.GetAll()
	var instances []*schema.PluginInstance

	for _, meta := range configs {
		instance := &schema.PluginInstance{
			Meta: *meta,
		}

		if state, ok := statuses[meta.Name]; ok {
			instance.State = *state
		} else {
			// If no status is found, initialize with a default idle state.
			instance.State = schema.PluginState{
				Status: "idle",
			}
		}
		instances = append(instances, instance)
	}

	return instances, nil
}

// GeneratePlugin creates a new plugin's configuration and Go source file.
func (pm *PluginManager) GeneratePlugin(pluginName string, outputBaseDir string) error {
	// Create a subdirectory for the plugin
	pluginDir := filepath.Join(outputBaseDir, pluginName)
	if err := os.MkdirAll(pluginDir, 0755); err != nil {
		return fmt.Errorf("无法创建插件目录 %s: %w", pluginDir, err)
	}

	// Generate plugin.yaml
	pluginYamlContent := fmt.Sprintf(`name: %s
version: 1.0
entry: %s.go
timeout: 3s
enabled: true
`, pluginName, strings.ToLower(pluginName))

	yamlPath := filepath.Join(pluginDir, "plugin.yaml")
	if err := os.WriteFile(yamlPath, []byte(pluginYamlContent), 0644); err != nil {
		return fmt.Errorf("无法写入 plugin.yaml 文件到 %s: %w", yamlPath, err)
	}

	// Generate Go plugin file
	pluginGoContent := fmt.Sprintf(`package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"plugin-sche/schema"
)

// Plugin entry point
func main() {
	// Read request from stdin
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %%v\n", err)
		os.Exit(1)
	}

	var req schema.Request
	if err := json.Unmarshal(input, &req); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling request: %%v\n", err)
		os.Exit(1)
	}

	// Execute plugin logic
	resp := execute(req)

	// Write response to stdout
	output, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling response: %%v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(output))
}

// execute contains the actual plugin logic
func execute(req schema.Request) schema.Response {
	// TODO: Implement your plugin logic here
	fmt.Fprintf(os.Stderr, "Plugin %%s received request: %%+v\n", %q, req)

	return schema.Response{
		Status: "success",
		Data:   map[string]interface{}{fmt.Sprintf("message": "Hello from %%s!", %q)},
		Error:  "",
	}
}
`, pluginName, pluginName, pluginName)

	goPath := filepath.Join(pluginDir, fmt.Sprintf("%s.go", strings.ToLower(pluginName)))
	if err := os.WriteFile(goPath, []byte(pluginGoContent), 0644); err != nil {
		return fmt.Errorf("无法写入 Go 插件文件到 %s: %w", goPath, err)
	}

	return nil
}

// resolveTimeout determines the effective timeout for a plugin run.
func (pm *PluginManager) resolveTimeout(plugin *schema.PluginMeta, timeoutOverride time.Duration) time.Duration {
	if timeoutOverride > 0 {
		return timeoutOverride
	}
	if plugin.Timeout > 0 {
		return plugin.Timeout
	}
	return 60 * time.Second // Default timeout
}

// resolvePluginCommand resolves the command and arguments to execute a plugin.
func (pm *PluginManager) resolvePluginCommand(plugin *schema.PluginMeta, args []string) (string, []string, error) {
	if plugin.Entry == "" {
		return "", nil, fmt.Errorf("插件 %s 未配置 entry", plugin.Name)
	}
	if plugin.ConfigPath == "" {
		return "", nil, fmt.Errorf("插件 %s 缺少配置路径", plugin.Name)
	}

	pluginBaseDir := filepath.Dir(plugin.ConfigPath)
	entryPath := filepath.Join(pluginBaseDir, plugin.Entry)
	mainGoPath := filepath.Join(pluginBaseDir, "main.go") // Assuming a main.go for Go plugins

	// On Windows, check for .exe first, then raw entry, then go run main.go
	if runtime.GOOS == "windows" {
		if pm.isRegularFile(entryPath + ".exe") {
			return entryPath + ".exe", args, nil
		}
		if pm.isRegularFile(entryPath) { // Could be a .bat, .cmd, or script without extension
			return entryPath, args, nil
		}
		if pm.isRegularFile(mainGoPath) {
			return "go", append([]string{"run", mainGoPath}, args...), nil
		}
	} else { // On Linux/macOS, check raw entry, then go run main.go
		if pm.isRegularFile(entryPath) {
			return entryPath, args, nil
		}
		if pm.isRegularFile(mainGoPath) {
			return "go", append([]string{"run", mainGoPath}, args...), nil
		}
	}

	return "", nil, fmt.Errorf("插件 %s 的可执行入口不存在: %s", plugin.Name, entryPath)
}

// RunPlugin executes a plugin by name.
func (pm *PluginManager) RunPlugin(name string, args []string, timeoutOverride time.Duration) ([]byte, error) {
	allPlugins, err := pm.GetAllPlugins() // Get all plugin instances (meta + state)
	if err != nil {
		return nil, err
	}

	var targetPlugin *schema.PluginInstance
	for _, p := range allPlugins {
		if p.Meta.Name == name {
			targetPlugin = p
			break
		}
	}

	if targetPlugin == nil {
		return nil, fmt.Errorf("未找到插件 %s", name)
	}

	if !targetPlugin.State.Enabled {
		runErr := fmt.Errorf("插件 %s 已禁用", targetPlugin.Meta.Name)
		// Update status as disabled
		if saveErr := pm.finishPluginRun(&targetPlugin.Meta, runErr); saveErr != nil {
			return nil, fmt.Errorf("%w; 状态写入失败: %v", runErr, saveErr)
		}
		return nil, runErr
	}

	commandName, commandArgs, err := pm.resolvePluginCommand(&targetPlugin.Meta, args)
	if err != nil {
		if saveErr := pm.finishPluginRun(&targetPlugin.Meta, err); saveErr != nil {
			return nil, fmt.Errorf("%w; 状态写入失败: %v", err, saveErr)
		}
		return nil, err
	}

	// Mark plugin as running
	if err := pm.markPluginRunning(&targetPlugin.Meta); err != nil {
		return nil, err
	}

	runCtx := pm.ctx
	cancel := func() {}
	effectiveTimeout := pm.resolveTimeout(&targetPlugin.Meta, timeoutOverride)
	if effectiveTimeout > 0 {
		runCtx, cancel = context.WithTimeout(pm.ctx, effectiveTimeout)
	}
	defer cancel()

	cmd := exec.CommandContext(runCtx, commandName, commandArgs...)
	cmd.Dir = filepath.Dir(targetPlugin.Meta.ConfigPath) // Execute in plugin's config directory
	output, runErr := cmd.CombinedOutput()

	// Check for context errors (timeout or cancellation)
	if errors.Is(runCtx.Err(), context.DeadlineExceeded) {
		runErr = fmt.Errorf("插件 %s 执行超时 (%s): %s", targetPlugin.Meta.Name, effectiveTimeout, runErr)
	} else if errors.Is(runCtx.Err(), context.Canceled) {
		runErr = fmt.Errorf("插件 %s 执行被取消: %s", targetPlugin.Meta.Name, runErr)
	}

	// Finalize plugin run status
	if err := pm.finishPluginRun(&targetPlugin.Meta, runErr); err != nil {
		if runErr != nil {
			return output, fmt.Errorf("%w; 状态写入失败: %v", runErr, err)
		}
		return output, err
	}

	return output, runErr
}

// ResetPluginStatus resets all plugin states to their default based on current configurations.
func (pm *PluginManager) ResetPluginStatus() error {
	// First, clear existing states
	if err := pm.statusManager.Reset(); err != nil {
		return fmt.Errorf("无法重置状态管理器: %w", err)
	}

	// Load all plugins, which will trigger the status manager to create default states
	// for any plugin configs that don't have existing states.
	_, err := pm.GetAllPlugins()
	if err != nil {
		return fmt.Errorf("无法加载所有插件以初始化状态: %w", err)
	}

	// Explicitly save the newly initialized states
	return pm.statusManager.Save()
}

// findPluginByName finds a plugin by its name from a list of PluginMeta.
func (pm *PluginManager) findPluginByName(plugins []*schema.PluginMeta, name string) (*schema.PluginMeta, error) {
	for _, plugin := range plugins {
		if plugin == nil {
			continue
		}
		if plugin.Name == name {
			return plugin, nil
		}
	}
	return nil, fmt.Errorf("未找到插件 %s", name)
}

// isRegularFile checks if a file exists and is a regular file.
func (pm *PluginManager) isRegularFile(filePath string) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// defaultPluginState creates a default PluginState for a given PluginMeta.
func (pm *PluginManager) defaultPluginState(plugin *schema.PluginMeta) *schema.PluginState {
	return &schema.PluginState{
		Status:       "idle",
		LastError:    "",
		ExecCount:    0,
		LastExecTime: "",
		Enabled:      plugin.Enabled, // Initialize enabled status from config
	}
}

// currentState retrieves the current state of a plugin, creating a default one if it doesn't exist.
func (pm *PluginManager) currentState(plugin *schema.PluginMeta) *schema.PluginState {
	state, exists := pm.statusManager.Get(plugin.Name)
	if !exists {
		state = pm.defaultPluginState(plugin)
		// We don't save immediately here, as it might be part of a larger update
	}
	return state
}

// markPluginRunning updates the plugin's state to "running".
func (pm *PluginManager) markPluginRunning(plugin *schema.PluginMeta) error {
	return pm.statusManager.UpdatePluginState(plugin.Name, func(state *schema.PluginState) {
		state.Status = "running"
		state.LastError = "" // Clear last error when running
	})
}

// finishPluginRun updates the plugin's state after it finishes running.
func (pm *PluginManager) finishPluginRun(plugin *schema.PluginMeta, runErr error) error {
	return pm.statusManager.UpdatePluginState(plugin.Name, func(state *schema.PluginState) {
		state.ExecCount++
		state.LastExecTime = time.Now().Format(time.RFC3339)

		if runErr != nil {
			state.Status = "error"
			state.LastError = runErr.Error()
		} else {
			state.Status = "idle"
			state.LastError = ""
		}
	})
}

// SetPluginEnabled sets the enabled status of a plugin.
func (pm *PluginManager) SetPluginEnabled(name string, enabled bool) error {
	// Update the plugin.yaml file
	err := pm.configManager.UpdateConfig(name, func(meta *schema.PluginMeta) error {
		meta.Enabled = enabled
		return nil
	})
	if err != nil {
		return fmt.Errorf("无法更新插件配置: %w", err)
	}

	// Update the status.json file
	return pm.statusManager.UpdatePluginState(name, func(state *schema.PluginState) {
		state.Enabled = enabled
	})
}
