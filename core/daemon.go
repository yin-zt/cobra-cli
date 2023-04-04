package core

import (
	"fmt"
	"github.com/yin-zt/cobra-cli/config"
	"github.com/yin-zt/cobra-cli/utils"
	"os"
	"strings"
)

func (this *Cli) Daemon(operate string) {
	_ = operate
}

// stop 作用是停止cli进程
func (this *Cli) stop() {
	// 获取到cli daemon -s daemon 的进程pid
	pids := this.getPids()
	for _, pid := range pids {
		if pid == "" {
			continue
		}
		// 如果pid不为空，且系统为windows，则调用 "cmd -c taskkill.exe /F /PID pid" kill调进程
		kMap := map[string]string{}
		if this.IsWindows() {
			this.ExecCmd([]string{"taskkill.exe", "/F", "/PID", pid}, 10, kMap)
		} else {
			//this.util.ExecCmd(strings.Split(fmt.Sprintf("sudo kill -9 %s", pid)," "), 10)
			this.ExecCmd([]string{fmt.Sprintf("sudo kill -9 %s", pid)}, 10, kMap)
		}
		// 移除记录cli daemon -s daemon 进程pid的文件
		os.Remove(config.PID_FILE)
	}
}

//读取存取进程pid的文件(/var/lib/cli/cli.pid)获取到进程pid
func (this *Cli) getPids() []string {
	//读取存取进程pid的文件(/var/lib/cli/cli.pid)获取到进程pid
	pid := utils.ReadFile(config.PID_FILE)
	return strings.Split(strings.TrimSpace(pid), "\n")
}
