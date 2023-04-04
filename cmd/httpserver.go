package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(HttpServerCmd)
	HttpServerCmd.Flags().StringP("path", "d", "", "--path")
	HttpServerCmd.Flags().StringP("host", "i", "0.0.0.0", "--host")
	HttpServerCmd.Flags().IntP("port", "P", 8080, "--port")
}

var HttpServerCmd = &cobra.Command{
	Use:   "httpserver",
	Short: "deploy http server on the figure ip",
	Long:  "cli httpserver -i hostIP -p port -d path",
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			cmdlog.Error(err)
		}
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			cmdlog.Error(err)
		}
		hostIp, err := cmd.Flags().GetString("host")
		if err != nil {
			cmdlog.Error(err)
		}
		cli.Httpserver(hostIp, path, port)
	},
}
