package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(sqliteCmd)
	sqliteCmd.Flags().StringP("filename", "f", "", "--filename")
	sqliteCmd.Flags().StringP("sqlstr", "s", "", "--sqlstr")
	sqliteCmd.Flags().StringP("table", "t", "", "--table")
}

var sqliteCmd = &cobra.Command{
	Use:   "sqlite",
	Short: "sqlite tool support to operate sqlite",
	Long:  "cli sqlite -f filename -t tablename -s sql",
	Run: func(cmd *cobra.Command, args []string) {
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			log.Fatalln(err)
		}
		sqlString, err := cmd.Flags().GetString("sqlstr")
		if err != nil {
			log.Fatalln(err)
		}
		table, err := cmd.Flags().GetString("table")
		if err != nil {
			log.Fatalln(err)
		}
		if err := cli.Util.OperateSqlite(filename, sqlString, table); err != nil {
			log.Fatalln(err)
		}
		log.Println("sqlite [default: localhost] exec sql successful")
	},
}
