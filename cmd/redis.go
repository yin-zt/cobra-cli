package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
)

func init() {
	RootCmd.AddCommand(redisCmd)
	redisCmd.Flags().StringP("password", "p", "", "./cli redis -p 123456 -c command")
	redisCmd.Flags().StringP("command", "c", "", "./cli redis -p 123456 -c command")
}

var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Simulating a redis client connection to redis",
	Long:  "cli redis -p 123456 -c command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./cli redis 127.0.0.1:6379 -p 123456 -c command")
		}
		host := args[0]
		if len(strings.Split(host, ":")) != 2 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./cli redis 127.0.0.1:6379 -p 123456 -c command")
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalln(err)
		}
		commands, err := cmd.Flags().GetString("command")

		if err := cli.Redisexec(host, password, commands); err != nil {
			log.Fatalln(err)
		}
		log.Printf("redis %s exec comamnd successfully", host)
	},
}
