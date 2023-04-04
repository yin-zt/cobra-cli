package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(KvsCmd)
}

var KvsCmd = &cobra.Command{
	Use:   "kvs",
	Short: "receive input and print it in friendly",
	Long:  "cli kvs",
	Run: func(cmd *cobra.Command, args []string) {
		cli.Kvs()
	},
}
