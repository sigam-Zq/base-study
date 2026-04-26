/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	pluginman "plugin-sche/core"

	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:          "disable <plugin-name>",
	Short:        "禁用指定插件",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pm, err := pluginman.NewPluginManager(cmd.Context(), pluginDir, workDir)
		if err != nil {
			return err
		}
		return pm.SetPluginEnabled(args[0], false)
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
