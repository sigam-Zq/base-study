package biz

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"plugin-sche/schema"
)

var statusFilePath = filepath.Join("runtime", "status.json")

// WritePluginStatusFile 根据插件配置生成运行时状态文件。
func WritePluginStatusFile(workDirFile string, plugins []*schema.PluginMeta) error {
	statusMap := make(map[string]schema.PluginState, len(plugins))
	for _, plugin := range plugins {
		if plugin == nil || plugin.Name == "" {
			continue
		}

		statusMap[plugin.Name] = schema.PluginState{
			// Enabled:      plugin.Enabled,
			Status:       "idle",
			LastError:    "",
			ExecCount:    0,
			LastExecTime: "",
		}
	}

	if err := os.MkdirAll("runtime", 0o755); err != nil {
		return err
	}

	return SavePluginStatusFile(workDirFile, statusMap)
}

func ReadPluginStatusFile(filePathName string) (map[string]schema.PluginState, error) {
	data, err := os.ReadFile(filePathName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return map[string]schema.PluginState{}, nil
		}
		return nil, err
	}

	statusMap := map[string]schema.PluginState{}
	if len(data) == 0 {
		return statusMap, nil
	}

	if err := json.Unmarshal(data, &statusMap); err != nil {
		return nil, err
	}

	return statusMap, nil
}

func SavePluginStatusFile(workDirFile string, statusMap map[string]schema.PluginState) error {
	if err := os.MkdirAll(filepath.Dir(workDirFile), 0o755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(statusMap, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(statusFilePath, data, 0o644)
}
