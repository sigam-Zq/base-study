package biz

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

// UpdatePluginEnabledInConfig 修改插件 YAML 中的 enabled 字段。
func UpdatePluginEnabledInConfig(filePath string, enabled bool) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return err
	}

	if len(root.Content) == 0 {
		root.Content = []*yaml.Node{{Kind: yaml.MappingNode, Tag: "!!map"}}
	}

	mappingNode := root.Content[0]
	enabledValue := strconv.FormatBool(enabled)

	for idx := 0; idx+1 < len(mappingNode.Content); idx += 2 {
		keyNode := mappingNode.Content[idx]
		valueNode := mappingNode.Content[idx+1]
		if keyNode.Value != "enabled" {
			continue
		}

		valueNode.Kind = yaml.ScalarNode
		valueNode.Tag = "!!bool"
		valueNode.Value = enabledValue
		return writeYAMLNode(filePath, &root)
	}

	mappingNode.Content = append(mappingNode.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: "enabled"},
		&yaml.Node{Kind: yaml.ScalarNode, Tag: "!!bool", Value: enabledValue},
	)

	return writeYAMLNode(filePath, &root)
}

// UpdatePluginEnabledInStatus 修改运行时状态中的启用状态。
// func UpdatePluginEnabledInStatus(pluginName string, enabled bool) error {
// 	statusMap, err := ReadPluginStatusFile()
// 	if err != nil {
// 		return err
// 	}

// 	state, exists := statusMap[pluginName]
// 	if !exists {
// 		state = schema.PluginState{}
// 	}

// 	state.Enabled = enabled
// 	if enabled {
// 		state.Status = "idle"
// 		state.LastError = ""
// 	} else {
// 		state.Status = "disabled"
// 		state.LastError = ""
// 	}

// 	statusMap[pluginName] = state
// 	return SavePluginStatusFile(statusMap)
// }

func writeYAMLNode(filePath string, node *yaml.Node) error {
	data, err := yaml.Marshal(node)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0o644)
}
