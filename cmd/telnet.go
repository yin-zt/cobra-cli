package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yin-zt/cobra-cli/utils"
	"log"
	"os"
)

func init() {
	RootCmd.AddCommand(TelnetCmd)
	TelnetCmd.Flags().IntP("timeout", "t", 5, "--timeout")
}

var TelnetCmd = &cobra.Command{
	Use:   "telnet",
	Short: "Simulate the telnet command to detect the port of the target host",
	Long:  "Usage: ./cli telnet 127.0.0.1 3306",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmdlog.Error("You have entered the wrong parameter, Usage: cli telnet 127.0.0.1 3306")
			log.Fatalln("You have entered the wrong parameter, Usage: ./cli telnet 127.0.0.1 3306")
		}
		timeout, err := cmd.Flags().GetInt("timeout")
		if err != nil {
			cmdlog.Error(err)
		}
		host := args[0]
		port := args[1]
		if ok := utils.IsIp(host); !ok {
			log.Printf("Please enter the correct IP address:%v", host)
			os.Exit(1)
		}
		if err := cli.Util.TelnetCheck(fmt.Sprintf("%v:%v", host, port), timeout); err != nil {
			log.Fatalln(err)
		}
		log.Printf("telnet %s connection successful", host)
	},
}
