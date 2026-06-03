package biz

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin-sche/schema"
	"strings"

	"gopkg.in/yaml.v3"
)

// ReadPluginConfig 读取并解析单个插件配置文件。
func ReadPluginConfig(filePath string) (*schema.PluginMeta, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %v", err)
	}

	var content *schema.PluginMeta = &schema.PluginMeta{
		Enabled: true,
	}
	err = yaml.Unmarshal(data, content)
	if err != nil {
		return nil, fmt.Errorf("解析YAML失败: %v", err)
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("获取绝对路径失败: %v", err)
	}
	content.ConfigPath = absPath

	return content, nil
}

// ReadPluginConfigs 遍历目录并读取全部插件配置。
func ReadPluginConfigs(dirPath string) ([]*schema.PluginMeta, error) {
	r := make([]*schema.PluginMeta, 0)
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("访问路径 %s 失败: %v", path, err)
		}

		// 跳过目录
		if info.IsDir() {
			return nil
		}

		// 检查文件扩展名是否为yaml或yml
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".yaml" || ext == ".yml" {
			content, err := ReadPluginConfig(path)
			if err != nil {
				return fmt.Errorf("处理文件 %s 时出错: %v", path, err)
			}
			r = append(r, content)
		}

		return nil
	})

	return r, err
}
