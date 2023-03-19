package core

import (
	"crypto/tls"
	"github.com/astaxie/beego/httplib"
	"time"
)

type Common struct {
}

type Cli struct {
	Util *Common
	//Conf *config.Config
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
