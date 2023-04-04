package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(FtpServerCmd)
	FtpServerCmd.Flags().StringP("username", "u", "root", "--username")
	FtpServerCmd.Flags().StringP("password", "p", "hello@world", "--password")
	FtpServerCmd.Flags().StringP("path", "d", "", "--path")
	FtpServerCmd.Flags().StringP("host", "i", "0.0.0.0", "--host")
	FtpServerCmd.Flags().IntP("port", "P", 2121, "--port")
}

var FtpServerCmd = &cobra.Command{
	Use:   "ftpserver",
	Short: "deploy ftp server on the figure host",
	Long:  "cli ftpserver -u username -p password -i hostIP -P port -d path",
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatalln(err)
		}
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalln(err)
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalln(err)
		}
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalln(err)
		}
		hostIp, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatalln(err)
		}
		cli.Ftpserver(username, password, hostIp, path, port)
	},
}
