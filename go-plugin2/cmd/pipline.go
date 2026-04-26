/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	pluginman "plugin-sche/core"
	"plugin-sche/schema"
	"time"

	"github.com/spf13/cobra"
)

var data string

// piplineCmd represents the pipline command
var piplineCmd = &cobra.Command{
	Use:   "pipline  [plugin-name]  [plugin-name] ...",
	Short: "数据按照顺序依次执行",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		pm, err := pluginman.NewPluginManager(cmd.Context(), pluginDir, workDir)
		if err != nil {
			return err
		}

		if verbose {
			pm.OpenDebug()
		}
		dataCur := []string{data}
		for _, pluginName := range args {

			output, err := pm.RunPlugin(pluginName, dataCur, time.Second*time.Duration(timeout))
			if err != nil {
				return err
			}
			cmd.Printf("output %v \n", output)
			var req *schema.Request = &schema.Request{Data: map[string]interface{}{
				"data": output.Data,
			}}
			reqByte, err := json.Marshal(req)
			if err != nil {
				return err
			}
			dataCur = []string{string(reqByte)}
			// if len(output) > 0 {
			// 	cmd.Print(string(output))
			// }
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(piplineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// piplineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// piplineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	piplineCmd.Flags().StringVarP(&data, "data", "d", "{'data':{'foo':'bar'}}", "输入数据")
}
