package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(JoinCmd)
	JoinCmd.Flags().StringP("strings", "s", "-", "--strings")
	JoinCmd.Flags().StringP("word", "w", "", "--word")
}

var JoinCmd = &cobra.Command{
	Use:   "join",
	Short: "Simulate the join command to integate some substring",
	Long:  "cli join -s '-' -w '$$'",
	Run: func(cmd *cobra.Command, args []string) {
		joinStr, err := cmd.Flags().GetString("strings")
		if err != nil {
			cmdlog.Error(err)
		}
		joinWord, err := cmd.Flags().GetString("word")
		if err != nil {
			cmdlog.Error(err)
		}
		cli.Join(joinStr, joinWord)
	},
}
