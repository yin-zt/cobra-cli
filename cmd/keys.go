package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(KeysCmd)
}

var KeysCmd = &cobra.Command{
	Use:   "keys",
	Short: "find out the keys of map, map value from scan keyboard",
	Long:  "cli keys",
	Run: func(cmd *cobra.Command, args []string) {
		cli.Keys()
	},
}
