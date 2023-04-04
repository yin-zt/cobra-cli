package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(MatchCmd)
	MatchCmd.Flags().StringP("string", "s", "", "--string")
	MatchCmd.Flags().StringP("mode", "m", "", "--mode")
	MatchCmd.Flags().StringP("output", "o", "i", "--output")
}

var MatchCmd = &cobra.Command{
	Use:   "match",
	Short: "Simulate the regex command",
	Long:  "cli match -s 'hell(i)45oworld' -m '[\\d+]+' -o 'i'",
	Run: func(cmd *cobra.Command, args []string) {
		matchStr, err := cmd.Flags().GetString("string")
		if err != nil {
			cmdlog.Error(err)
		}
		matchMode, err := cmd.Flags().GetString("mode")
		if err != nil {
			cmdlog.Error(err)
		}
		matchOutput, err := cmd.Flags().GetString("output")
		if err != nil {
			cmdlog.Error(err)
		}
		cli.Match(matchStr, matchMode, matchOutput)
	},
}
