package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(PingCmd)
	PingCmd.Flags().IntP("count", "c", 10, "--count")
	PingCmd.Flags().IntP("size", "l", 56, "--size")
}

var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Simulate the ping command to send icmp packets to the target host",
	Long:  "cli ping 192.168.1.1 -c 10 -l 1000",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./cli ping 192.168.1.1")
		}
		counts, err := cmd.Flags().GetInt("count")
		if err != nil {
			log.Fatalln(err)
		}
		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(size, counts)
		fmt.Println(args[0])
		cli.Util.Ping(args[0], size, counts)
	},
}
