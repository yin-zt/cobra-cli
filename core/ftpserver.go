package core

import (
	log "github.com/cihub/seelog"
	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
	"github.com/yin-zt/cobra-cli/utils"
)

func (this *Common) Ftpserver(user, pass, host, path string, port int) {

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
		log.Errorf("Error starting ftp server:", err)
	}

}
