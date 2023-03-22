package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
)

func init() {
	RootCmd.AddCommand(redisCmd)
	redisCmd.Flags().StringP("password", "p", "", "./net-tools redis -p 123456")
}

var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Simulating a redis client connection to redis",
	Long:  "net-tools redis 123456",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./net-tools redis 127.0.0.1:6379 -p 123456")
		}
		host := args[0]
		if len(strings.Split(host, ":")) != 2 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./net-tools redis 127.0.0.1:6379 -p 123456")
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalln(err)
		}
		if err := cli.Util.Redischeck(host, password); err != nil {
			log.Fatalln(err)
		}
		log.Printf("redis %s connection successful", host)
	},
}
