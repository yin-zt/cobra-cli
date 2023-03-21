package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(UuidCmd)
}

var UuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "find out native-specific uuid or generate new one",
	Long:  "cli uuid or cli uuid new",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 || (len(args) == 1 && args[0] != "new") {
			cmdlog.Error("You have entered the wrong parameter, Usage: ./cli uuid or ./cli uuid new")
			log.Fatalln("You have entered the wrong parameter, Usage: ./cli uuid or ./cli uuid new")
		}
		if len(args) == 0 {
			fmt.Println(cli.Util.GetProductUUID())
		} else {
			fmt.Println(cli.Util.GetUUID())
		}
	},
}
