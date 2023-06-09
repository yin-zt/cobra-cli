package core

import (
	"crypto/tls"
	"github.com/astaxie/beego/httplib"
	"github.com/takama/daemon"
	"github.com/yin-zt/cobra-cli/utils"
	"time"
)

var (
	corelog = utils.GetLog()
)

type Common struct {
}

type Daemon struct {
	daemon.Daemon
}

type Cli struct {
	Util *Common
	//Conf *config.Config
	_daemon *Daemon
}

func init() {
	defer corelog.Flush()
	corelog.Info("success to init seelog: corelog")
}

func NewCli() *Cli {

	var (
		NCli *Cli
		Util = &Common{}
	)

	setting := httplib.BeegoHTTPSettings{
		UserAgent:        "beegoServer",
		ConnectTimeout:   60 * time.Second,
		ReadWriteTimeout: 60 * time.Second,
		Gzip:             true,
		DumpBody:         true,
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	}

	httplib.SetDefaultSetting(setting)

	NCli = &Cli{Util: Util}
	return NCli

}
