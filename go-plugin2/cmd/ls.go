/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	pluginman "plugin-sche/core"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "查看当前插件列表",
	Long:  `便利当前的配置文件目录,查看插件列表信息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pm, err := pluginman.NewPluginManager(context.Background(), pluginDir, workDir)
		if err != nil {
			return err
		}
		ins, err := pm.GetAllPlugins()
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(w, "NAME\tVERSION\tENTRY\tENABLED\tTIMEOUT\tSTATUS")
		for _, plugin := range ins {
			if plugin == nil {
				continue
			}
			status := plugin.State.Status
			if plugin.State.Status == "error" && plugin.State.LastError != "" {
				errorInfo := plugin.State.LastError
				if plugin.State.LastExecTime != "" {
					errorInfo = fmt.Sprintf("%s at %s", errorInfo, plugin.State.LastExecTime)
				}
				status = fmt.Sprintf("%s (%s)", status, errorInfo)
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%t\t%s\t%s\n", plugin.Meta.Name, plugin.Meta.Version, plugin.Meta.Entry, plugin.Meta.Enabled, plugin.Meta.Timeout, status)
		}
		w.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
