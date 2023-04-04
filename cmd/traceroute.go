package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(traceRouteCmd)
}

var traceRouteCmd = &cobra.Command{
	Use:   "traceroute",
	Short: "Simulate traceroute command to detect routes",
	Long:  "cli traceroute 192.168.1.1",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./cli traceroute 192.168.1.1")
		}
		host := args[0]
		cli.Trace(host)
	},
}
