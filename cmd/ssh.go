package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// -u username -p password -h hostIp -P port -c command
func init() {
	RootCmd.AddCommand(SshCmd)
	SshCmd.Flags().StringP("username", "u", "root", "--username")
	SshCmd.Flags().StringP("password", "p", "", "--password")
	SshCmd.Flags().StringP("port", "P", "22", "--port")
	SshCmd.Flags().StringP("command", "c", "pwd", "--command")
}

var SshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "support ssh figure host and exec command",
	Long:  "cli ssh 192.168.1.1 -u username -p password -P port -c comand ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmdlog.Error("You have entered the wrong parameter, Usage: ./cli ssh 192.168.1.1 -u username -p password -P port")
			log.Fatalln("You have entered the wrong parameter, Usage: ./cli ssh 192.168.1.1 -u username -p password -P port")
		}
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			cmdlog.Error(err)
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			cmdlog.Error(err)
		}
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			cmdlog.Error(err)
		}
		command, err := cmd.Flags().GetString("command")
		if err != nil {
			cmdlog.Error(err)
		}
		cli.Ssh(args[0], username, password, port, command)
	},
}
