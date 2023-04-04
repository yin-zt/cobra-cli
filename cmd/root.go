package cmd

import (
	"github.com/astaxie/beego/httplib"
	"github.com/spf13/cobra"
	"github.com/yin-zt/cobra-cli/core"
	"github.com/yin-zt/cobra-cli/utils"
	"log"
	"net/http"
	"time"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A tool to detect network, exec sql, exec command etc.",
	Long: `A tool to detect network, exec sql, exec command etc., 
and support detection of middleware such as redis,mysql,traceroute, etc., 
please use cli help for detailed usage`,
}

var (
	cli    = core.NewCli()
	cmdlog = utils.GetLog()
)

func init() {
	defer cmdlog.Flush()
	cmdlog.Info("success to init seelog ")
	initHttpLib()
}

// initHeepLib 函数用于初始化httplib的设置
func initHttpLib() {

	defaultTransport := &http.Transport{
		DisableKeepAlives:   true,
		Dial:                httplib.TimeoutDialer(time.Second*15, time.Second*300),
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}
	settins := httplib.BeegoHTTPSettings{
		UserAgent:        "Go-FastDFS",
		ConnectTimeout:   15 * time.Second,
		ReadWriteTimeout: 120 * time.Second,
		Gzip:             true,
		DumpBody:         true,
		Transport:        defaultTransport,
	}
	httplib.SetDefaultSetting(settins)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		cmdlog.Errorf("%s", err)
		log.Fatalln(err)
	}
}
