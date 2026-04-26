package core

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin-sche/schema"

	"gopkg.in/yaml.v3"
)

// ConfigManager is responsible for reading plugin configurations.
type ConfigManager struct {
	configDir string
}

// NewConfigManager creates a new ConfigManager.
func NewConfigManager(configDir string) *ConfigManager {
	return &ConfigManager{configDir: configDir}
}

// ReadConfigs reads all plugin configurations from the config directory.
func (c *ConfigManager) ReadConfigs() ([]*schema.PluginMeta, error) {
	var configs []*schema.PluginMeta

	err := filepath.WalkDir(c.configDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil // Skip directories for direct processing, will recurse
		}

		ext := filepath.Ext(d.Name())
		if ext != ".yaml" && ext != ".yml" {
			return nil
		}

		// Ensure it's not a root-level yaml if we expect them in subdirs
		// This check is important if plugin.yaml might exist directly in configDir
		// but we only want those in subdirectories.
		// For now, let's assume any .yaml in any subdir (or root) is a plugin config.
		// If a plugin.yaml should ONLY be in a plugin-named subdirectory,
		// we'd add more logic here to check parent directory name.

		data, err := os.ReadFile(path)
		if err != nil {
			// Log the error but continue processing other files
			fmt.Fprintf(os.Stderr, "Error reading plugin config file %s: %v\n", path, err)
			return nil
		}

		var meta schema.PluginMeta
		if err := yaml.Unmarshal(data, &meta); err != nil {
			// Log the error but continue processing other files
			fmt.Fprintf(os.Stderr, "Error unmarshalling plugin config from %s: %v\n", path, err)
			return nil
		}

		meta.ConfigPath = path // Store the full path to the config file
		configs = append(configs, &meta)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("遍历插件配置目录失败: %w", err)
	}

	return configs, nil
}

// UpdateConfig updates a specific plugin's configuration file.
func (c *ConfigManager) UpdateConfig(pluginName string, updateFunc func(meta *schema.PluginMeta) error) error {
	configs, err := c.ReadConfigs()
	if err != nil {
		return err
	}

	var targetMeta *schema.PluginMeta
	for _, meta := range configs {
		if meta.Name == pluginName {
			targetMeta = meta
			break
		}
	}

	if targetMeta == nil {
		return fmt.Errorf("未找到插件 %s 的配置", pluginName)
	}

	if err := updateFunc(targetMeta); err != nil {
		return err
	}

	data, err := yaml.Marshal(targetMeta)
	if err != nil {
		return fmt.Errorf("无法序列化插件配置: %w", err)
	}

	return os.WriteFile(targetMeta.ConfigPath, data, 0644)
}
