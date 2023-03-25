package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(mysqlCmd)
	mysqlCmd.Flags().StringP("username", "u", "root", "./net-tools mysql -u root -p 123456")
	mysqlCmd.Flags().StringP("password", "p", "123456", "net-tools mysql -u root -p 123456")
}

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "Simulate a mysql client to connect to the database",
	Long:  "net-tools mysql 127.0.0.1:3306 -u root -p 123456",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("You have entered the wrong parameter, Usage: ./net-tools mysql 127.0.0.1:3306 -u root -p 123456")
		}
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalln(err)
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalln(err)
		}
		if err := cli.Util.MysqlPingCheck(args[0], username, password); err != nil {
			log.Fatalln(err)
		}
		log.Printf("mysql %s connection successful", args[0])
	},
}
