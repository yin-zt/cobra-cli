package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(mysqlCmd)
	mysqlCmd.Flags().StringP("username", "u", "root", "./cli mysql -u root -p 123456 -c sql")
	mysqlCmd.Flags().StringP("password", "p", "123456", "cli mysql -u root -p 123456 -c sql")
	mysqlCmd.Flags().StringP("command", "c", "", "net-tools mysql -u root -p 123456 -c sql")
}

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Simulate a mysql client to connect to the database",
	Long:  "cli mysql 127.0.0.1:3306 -u root -p 123456",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./net-tools mysql 127.0.0.1:3306 -u root -p 123456 -c sql")
		}
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalln(err)
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalln(err)
		}
		rawsql, err := cmd.Flags().GetString("command")
		if err != nil {
			log.Fatalln(err)
		}
		if err := cli.MysqlPingCheck(args[0], username, password, rawsql); err != nil {
			log.Fatalln(err)
		}
		log.Printf("mysql %s exec sql successful", args[0])
	},
}
