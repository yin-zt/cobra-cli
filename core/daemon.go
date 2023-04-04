package core

import (
	"fmt"
	log "github.com/cihub/seelog"
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
		if this.IsWindows() {
			this.ExecCmd([]string{"taskkill.exe", "/F", "/PID", pid}, 10)
		} else {
			//this.util.ExecCmd(strings.Split(fmt.Sprintf("sudo kill -9 %s", pid)," "), 10)
			this.ExecCmd([]string{fmt.Sprintf("sudo kill -9 %s", pid)}, 10)
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

// 检查是否有多个 cli daemon -s daemon 进程运行，如果有则进程退出，退出状态码为：0
func (this *Cli) isRunning() bool {
	pids := this.getPids()
	count := 0
	for _, pid := range pids {
		if pid == "" {
			continue
		}
		if this.IsWindows() {
			pids := this.ExecCmd([]string{"tasklist", "/FI", fmt.Sprintf("PID eq %s", pid)}, 10)
			if strings.Index(pids, pid) > 0 {
				count = count + 1
			}
		} else {
			if utils.IsExist(fmt.Sprintf("/proc/%s", pid)) {
				count = count + 1
			}
		}
	}
	if count > 1 {
		log.Error("muti process is running", pids)
		log.Flush()
		os.Exit(0)
	}
	if count == 1 {
		return true
	}
	return false
}
