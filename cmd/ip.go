package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(IpCmd)
}

var IpCmd = &cobra.Command{
	Use:   "ip",
	Short: "find out local ip",
	Long:  "cli ip",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			cmdlog.Error("You have entered the wrong parameter, Usage: ./cli ip")
			log.Fatalln("You have entered the wrong parameter, Usage: ./cli ip")
		}

		fmt.Println(cli.Util.GetNetworkIP())
	},
}
