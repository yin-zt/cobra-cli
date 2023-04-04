package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(LenCmd)
}

var LenCmd = &cobra.Command{
	Use:   "len",
	Short: "Simulate the len command to figure out the length of object ",
	Long:  "echo '{\"key1\": \"val1\", \"key2\": \"val2\"}' | cli len",
	Run: func(cmd *cobra.Command, args []string) {
		cli.Len()
	},
}
