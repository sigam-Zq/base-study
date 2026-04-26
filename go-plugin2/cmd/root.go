/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// 可以在所有命令中访问的变量
	pluginDir string
	workDir   string
	verbose   bool
	timeout   int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "plugin-sche",
	Short: "插件调度",
	Long:  "My application description",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if verbose {
		fmt.Fprintf(os.Stderr, "DEBUG: Entering cmd.Execute()\n")
	}
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(genPluginCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.plugin-sche.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// 持久性标志 - 所有子命令都可使用
	rootCmd.PersistentFlags().StringVarP(&pluginDir, "plugins", "p", "plugins", "配置文件路径")
	rootCmd.PersistentFlags().StringVarP(&workDir, "wrok-spcae", "w", "runtime", "运行状态文件")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "详细输出")
	rootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 60, "超时时间(秒)")
}
