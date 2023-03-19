package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yin-zt/cobra-cli/core"
	"log"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A tool to detect network, exec sql, exec command etc.",
	Long: `A tool to detect network, exec sql, exec command etc., 
and support detection of middleware such as redis,mysql,traceroute, etc., 
please use cli help for detailed usage`,
}

var (
	cli = core.NewCli()
)

//func init() {
//	net_tools.TelnetCmd.Flags().IntP("timeout", "t", 5, "-- timeout")
//}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
