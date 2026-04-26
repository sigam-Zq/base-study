/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	pluginman "plugin-sche/core"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:          "run <plugin-name> [plugin-args...]",
	Short:        "执行指定插件",
	SilenceUsage: true,
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pluginName := args[0]
		pluginArgs := args[1:]

		overrideTimeout := time.Duration(0)
		if flag := cmd.Flag("timeout"); flag != nil && flag.Changed {
			overrideTimeout = time.Duration(timeout) * time.Second
		}

		// Define pluginDir, workDir, verbose, timeout here or ensure they are global.
		if verbose {
			fmt.Fprintf(os.Stderr, "DEBUG: Initializing PluginManager with pluginDir: %s, workDir: %s\n", pluginDir, workDir)

			fmt.Fprintf(os.Stderr, "DEBUG: Entering RunPlugin for plugin: %s\n", pluginName)
		}

		pm, err := pluginman.NewPluginManager(cmd.Context(), pluginDir, workDir)
		if err != nil {
			return err
		}

		if verbose {
			pm.OpenDebug()
		}
		output, err := pm.RunPlugin(pluginName, pluginArgs, overrideTimeout)
		cmd.Printf("%v", output)

		return err
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
