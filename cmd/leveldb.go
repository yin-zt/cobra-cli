package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(LevelDbCmd)
	LevelDbCmd.Flags().StringP("operate", "o", "put", "--operate")
	LevelDbCmd.Flags().StringP("key", "k", "hello", "--key")
	LevelDbCmd.Flags().StringP("value", "v", "world", "--value")
	LevelDbCmd.Flags().StringP("dir", "d", "", "--dir")
}

var LevelDbCmd = &cobra.Command{
	Use:   "leveldb",
	Short: "operate leveldb, such set update delete etc.",
	Long:  "cli leveldb -o set -k key -v value",
	Run: func(cmd *cobra.Command, args []string) {
		operation, err := cmd.Flags().GetString("operate")
		if err != nil {
			log.Fatalln(err)
		}
		key, err := cmd.Flags().GetString("key")
		if err != nil {
			log.Fatalln(err)
		}
		value, err := cmd.Flags().GetString("value")
		if err != nil {
			log.Fatalln(err)
		}
		path, err := cmd.Flags().GetString("dir")
		if err != nil {
			log.Fatalln(err)
		}
		cli.Util.LevelOperate(operation, key, value, path)
	},
}
