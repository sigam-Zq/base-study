/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	pluginman "plugin-sche/core"

	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:          "enable <plugin-name>",
	Short:        "启用指定插件",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pm, err := pluginman.NewPluginManager(cmd.Context(), pluginDir, workDir)
		if err != nil {
			return err
		}
		return pm.SetPluginEnabled(args[0], true)
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
