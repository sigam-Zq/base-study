package cmd

import (
	"context"
	"fmt"
	"path/filepath"
	pluginman "plugin-sche/core"

	"github.com/spf13/cobra"
)

var genPluginOutputDir string

// genPluginCmd represents the genPlugin command
var genPluginCmd = &cobra.Command{
	Use:   "genPlugin [plugin-name]",
	Short: "生成插件的配置文件和Go文件",
	Long: `genPlugin 命令用于在指定的目录下生成插件的配置文件 (plugin.yaml) 和Go语言的插件代码骨架。

示例:
  go-plugin2 genPlugin my-awesome-plugin -o ./plugins`,
	Args: cobra.ExactArgs(1), // 必须提供一个插件名称参数
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]

		finalOutputBaseDir := genPluginOutputDir
		if finalOutputBaseDir == "" {
			finalOutputBaseDir = pluginDir // 使用 rootCmd 中的全局 pluginDir 作为默认值
		}

		pm, err := pluginman.NewPluginManager(context.Background(), pluginDir, workDir)
		if err != nil {
			return fmt.Errorf("无法初始化插件管理器: %w", err)
		}

		if verbose {
			pm.OpenDebug()
		}

		if err := pm.GeneratePlugin(pluginName, finalOutputBaseDir); err != nil {
			return fmt.Errorf("生成插件失败: %w", err)
		}
		fmt.Printf("成功为 \"%s\" 插件生成文件到 %s\n", pluginName, filepath.Join(finalOutputBaseDir, pluginName))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genPluginCmd)
	// Here you will define your flags and configuration settings.
	genPluginCmd.Flags().StringVarP(&genPluginOutputDir, "output", "o", "", "指定插件的生成目录")
}
