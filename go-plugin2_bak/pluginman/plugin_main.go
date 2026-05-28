package pluginman

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"plugin-sche/biz"
// 	"plugin-sche/schema"
// 	"runtime"
// 	"time"
// )

// type PluginMan struct {
// 	ctx       context.Context
// 	pluginDir string
// 	workDir   string
// }

// func NewPluginMan(ctx context.Context, pluginDir string, wrokDir string) *PluginMan {
// 	if ctx == nil {
// 		ctx = context.Background()
// 	}

// 	return &PluginMan{
// 		ctx:       ctx,
// 		pluginDir: pluginDir,
// 		workDir:   wrokDir,
// 	}
// }

// func (pm *PluginMan) ReadConfigs() ([]*schema.PluginMeta, error) {
// 	return biz.ReadPluginConfigs(pm.pluginDir)
// }

// func (pm *PluginMan) ReadStatusFile() (map[string]schema.PluginState, error) {
// 	statusFile := filepath.Join(pm.workDir, "status.json")
// 	if _, err := os.Stat(statusFile); err == exec.ErrNotFound {
// 		pm.GenerateStatusFile()
// 	}
// 	return biz.ReadPluginStatusFile(statusFile)
// }

// func (pm *PluginMan) CofingsMergeStatus(configs []*schema.PluginMeta) ([]*schema.PluginInstance, error) {

// 	statusMap, err := pm.ReadStatusFile()
// 	if err != nil {
// 		return nil, err
// 	}

// 	re := make([]*schema.PluginInstance, len(configs))

// 	for i := range configs {
// 		re[i] = &schema.PluginInstance{
// 			Meta:  *configs[i],
// 			State: statusMap[configs[i].Name],
// 		}
// 	}

// 	return re, nil
// }

// func (pm *PluginMan) GenerateStatusFile() error {
// 	plugins, err := pm.ReadConfigs()
// 	if err != nil {
// 		return err
// 	}

// 	return biz.WritePluginStatusFile(filepath.Join(pm.workDir, schema.StatusFileName), plugins)
// }

// func (pm *PluginMan) SetPluginEnabled(name string, enabled bool) error {
// 	plugins, err := pm.ReadConfigs()
// 	if err != nil {
// 		return err
// 	}

// 	plugin, err := pm.findPluginByName(plugins, name)
// 	if err != nil {
// 		return err
// 	}

// 	if err := biz.UpdatePluginEnabledInConfig(plugin.ConfigPath, enabled); err != nil {
// 		return err
// 	}

// 	// 状态参数只存在于 yaml中
// 	// if err := biz.UpdatePluginEnabledInStatus(plugin.Name, enabled); err != nil {
// 	// 	return err
// 	// }

// 	return nil
// }

// func (pm *PluginMan) RunPlugin(name string, args []string, timeoutOverride time.Duration) ([]byte, error) {
// 	plugins, err := pm.ReadConfigs()
// 	if err != nil {
// 		return nil, err
// 	}

// 	plugin, err := pm.findPluginByName(plugins, name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	statusMap, err := pm.ensureStatusMap(plugins)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if !plugin.Enabled {
// 		runErr := fmt.Errorf("插件 %s 已禁用", plugin.Name)
// 		if saveErr := pm.finishPluginRun(statusMap, plugin, runErr); saveErr != nil {
// 			return nil, fmt.Errorf("%w; 状态写入失败: %v", runErr, saveErr)
// 		}
// 		return nil, runErr
// 	}

// 	commandName, commandArgs, err := pm.resolvePluginCommand(plugin, args)
// 	if err != nil {
// 		if saveErr := pm.finishPluginRun(statusMap, plugin, err); saveErr != nil {
// 			return nil, fmt.Errorf("%w; 状态写入失败: %v", err, saveErr)
// 		}
// 		return nil, err
// 	}

// 	if err := pm.markPluginRunning(statusMap, plugin); err != nil {
// 		return nil, err
// 	}

// 	runCtx := pm.ctx
// 	cancel := func() {}
// 	effectiveTimeout := pm.resolveTimeout(plugin, timeoutOverride)
// 	if effectiveTimeout > 0 {
// 		runCtx, cancel = context.WithTimeout(pm.ctx, effectiveTimeout)
// 	}
// 	defer cancel()

// 	cmd := exec.CommandContext(runCtx, commandName, commandArgs...)
// 	cmd.Dir = filepath.Dir(plugin.ConfigPath)
// 	output, runErr := cmd.CombinedOutput()

// 	if errors.Is(runCtx.Err(), context.DeadlineExceeded) {
// 		runErr = fmt.Errorf("插件 %s 执行超时: %s", plugin.Name, effectiveTimeout)
// 	} else if errors.Is(runCtx.Err(), context.Canceled) {
// 		runErr = fmt.Errorf("插件 %s 执行被取消", plugin.Name)
// 	}

// 	if err := pm.finishPluginRun(statusMap, plugin, runErr); err != nil {
// 		if runErr != nil {
// 			return output, fmt.Errorf("%w; 状态写入失败: %v", runErr, err)
// 		}
// 		return output, err
// 	}

// 	return output, runErr
// }

// func (pm *PluginMan) findPluginByName(plugins []*schema.PluginMeta, name string) (*schema.PluginMeta, error) {
// 	for _, plugin := range plugins {
// 		if plugin == nil {
// 			continue
// 		}
// 		if plugin.Name == name {
// 			return plugin, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("未找到插件 %s", name)
// }

// func (pm *PluginMan) ensureStatusMap(plugins []*schema.PluginMeta) (map[string]schema.PluginState, error) {
// 	statusMap, err := biz.ReadPluginStatusFile(filepath.Join(pm.workDir, schema.StatusFileName))
// 	if err != nil {
// 		return nil, err
// 	}

// 	changed := false
// 	for _, plugin := range plugins {
// 		if plugin == nil || plugin.Name == "" {
// 			continue
// 		}

// 		// state, exists := statusMap[plugin.Name]
// 		// if !exists {
// 		// 	statusMap[plugin.Name] = defaultPluginState(plugin)
// 		// 	changed = true
// 		// 	continue
// 		// }

// 		// if state.Enabled != plugin.Enabled {
// 		// 	state.Enabled = plugin.Enabled
// 		// 	statusMap[plugin.Name] = state
// 		// 	changed = true
// 		// }
// 	}

// 	if changed {
// 		if err := biz.SavePluginStatusFile(filepath.Join(pm.workDir, schema.StatusFileName), statusMap); err != nil {
// 			return nil, err
// 		}
// 	}

// 	return statusMap, nil
// }

// func (pm *PluginMan) markPluginRunning(statusMap map[string]schema.PluginState, plugin *schema.PluginMeta) error {
// 	state := pm.currentState(statusMap, plugin)
// 	// state.Enabled = plugin.Enabled
// 	state.Status = "running"
// 	state.LastError = ""
// 	statusMap[plugin.Name] = state

// 	return biz.SavePluginStatusFile(filepath.Join(pm.workDir, schema.StatusFileName), statusMap)
// }

// func (pm *PluginMan) finishPluginRun(statusMap map[string]schema.PluginState, plugin *schema.PluginMeta, runErr error) error {
// 	state := pm.currentState(statusMap, plugin)
// 	// state.Enabled = plugin.Enabled
// 	state.ExecCount++
// 	state.LastExecTime = time.Now().Format(time.RFC3339)

// 	if runErr != nil {
// 		state.Status = "error"
// 		state.LastError = runErr.Error()
// 	} else {
// 		state.Status = "idle"
// 		state.LastError = ""
// 	}

// 	statusMap[plugin.Name] = state
// 	return biz.SavePluginStatusFile(filepath.Join(pm.workDir, schema.StatusFileName), statusMap)
// }

// func (pm *PluginMan) currentState(statusMap map[string]schema.PluginState, plugin *schema.PluginMeta) schema.PluginState {
// 	state, exists := statusMap[plugin.Name]
// 	if !exists {
// 		return defaultPluginState(plugin)
// 	}

// 	return state
// }

// func (pm *PluginMan) resolveTimeout(plugin *schema.PluginMeta, timeoutOverride time.Duration) time.Duration {
// 	if timeoutOverride > 0 {
// 		return timeoutOverride
// 	}
// 	if plugin.Timeout > 0 {
// 		return plugin.Timeout
// 	}

// 	return 60 * time.Second
// }

// func (pm *PluginMan) resolvePluginCommand(plugin *schema.PluginMeta, args []string) (string, []string, error) {
// 	if plugin.Entry == "" {
// 		return "", nil, fmt.Errorf("插件 %s 未配置 entry", plugin.Name)
// 	}
// 	if plugin.ConfigPath == "" {
// 		return "", nil, fmt.Errorf("插件 %s 缺少配置路径", plugin.Name)
// 	}

// 	pluginBaseDir := filepath.Dir(plugin.ConfigPath)
// 	entryPath := filepath.Join(pluginBaseDir, plugin.Entry)
// 	mainGoPath := filepath.Join(pluginBaseDir, "main.go")

// 	if runtime.GOOS == "windows" {
// 		if filepath.Ext(entryPath) != "" && isRegularFile(entryPath) {
// 			return entryPath, args, nil
// 		}

// 		if isRegularFile(entryPath + ".exe") {
// 			return entryPath + ".exe", args, nil
// 		}

// 		if isRegularFile(mainGoPath) {
// 			return "go", append([]string{"run", mainGoPath}, args...), nil
// 		}

// 		if isRegularFile(entryPath) {
// 			return entryPath, args, nil
// 		}
// 	} else {
// 		if isRegularFile(entryPath) {
// 			return entryPath, args, nil
// 		}
// 		if isRegularFile(entryPath + ".exe") {
// 			return entryPath + ".exe", args, nil
// 		}
// 		if isRegularFile(mainGoPath) {
// 			return "go", append([]string{"run", mainGoPath}, args...), nil
// 		}
// 	}

// 	return "", nil, fmt.Errorf("插件 %s 的可执行入口不存在: %s", plugin.Name, entryPath)
// }

// func defaultPluginState(plugin *schema.PluginMeta) schema.PluginState {
// 	return schema.PluginState{
// 		// Enabled:      plugin.Enabled,
// 		Status:       "idle",
// 		LastError:    "",
// 		ExecCount:    0,
// 		LastExecTime: "",
// 	}
// }

// func isRegularFile(filePath string) bool {
// 	info, err := os.Stat(filePath)
// 	if err != nil {
// 		return false
// 	}

// 	return !info.IsDir()
// }
