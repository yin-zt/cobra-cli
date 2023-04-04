package cmd

import (
	"github.com/spf13/cobra"
)

// -d path -f file -t 12 -u -a -x
func init() {
	RootCmd.AddCommand(ShellCmd)
	ShellCmd.Flags().StringP("dir", "d", "", "--dir")
	ShellCmd.Flags().StringP("file", "f", "", "--file")
	ShellCmd.Flags().IntP("timeout", "t", 10, "--timeout")
}

var ShellCmd = &cobra.Command{
	Use:   "shell",
	Short: "support excute command",
	Long:  "cli shell -d path -f file -t 12 -u -a -x",
	Run: func(cmd *cobra.Command, args []string) {
		cli.Shell()
	},
}
