package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(RequestCmd)
	RequestCmd.Flags().StringP("url", "u", "http://hostip:port/", "--url")
	RequestCmd.Flags().StringP("data", "d", "", "--data")
}

var RequestCmd = &cobra.Command{
	Use:   "request",
	Short: "Supports http requests",
	Long:  "cli request -u url -d data",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatalln(err)
		}
		data, err := cmd.Flags().GetString("data")
		if err != nil {
			log.Fatalln(err)
		}
		cli.Util.Request(url, data)
	},
}
