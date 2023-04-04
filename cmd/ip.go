package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

func init() {
	RootCmd.AddCommand(IpCmd)
}

var IpCmd = &cobra.Command{
	Use:   "ip",
	Short: "find out local ip",
	Long:  "cli ip",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 || (len(args) == 1 && args[0] != "a") {
			cmdlog.Error("You have entered the wrong parameter, Usage: ./cli ip or ./cli ip a")
			log.Fatalln("You have entered the wrong parameter, Usage: ./cli ip or ./cli ip a")
		}
		if len(args) == 0 {
			fmt.Println(cli.GetNetworkIP())
		} else {
			var result string
			result = strings.Join(cli.GetAllIps(), "\n")
			fmt.Println(result)
		}
	},
}
