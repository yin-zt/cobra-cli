package core

import (
	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
	"github.com/yin-zt/cobra-cli/utils"
	"log"
)

func (this *Cli) Ftpserver(user, pass, host, path string, port int) {
	if path == "" {
		if v, err := utils.Home(); err == nil {
			path = v
		}
	}

	factory := &filedriver.FileDriverFactory{
		RootPath: path,
		Perm:     server.NewSimplePerm("user", "group"),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     port,
		Hostname: host,
		Auth:     &server.SimpleAuth{Name: user, Password: pass},
	}

	ftp := server.NewServer(opts)

	err := ftp.ListenAndServe()
	if err != nil {
		corelog.Errorf("Error starting ftp server:", err)
		log.Fatalln("Error starting ftp server:", err)
	}

}
