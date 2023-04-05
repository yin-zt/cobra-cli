package core

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/yin-zt/cobra-cli/config"
	"github.com/yin-zt/cobra-cli/utils"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// Daemon 使用方式：cli daemon  -s daemon|stop|start|status| restart
func (this *Cli) Daemon(operate string) {

	if operate == "daemon" {
		this.Run()
		return
	}
	if operate == "stop" {
		this.stop()
	}
	if operate == "status" {
		if !this.isRunning() {
			os.Remove(config.PID_FILE)
			fmt.Println("cli is not running")
		} else {
			fmt.Println(fmt.Sprintf("pid:%v", this.getPids()))
			fmt.Println("cli is running")
		}
		return
	}

	if operate == "restart" || operate == "start" {
		this.stop()
		args := []string{"daemon", "-s", "daemon"}
		cmd := exec.Command(os.Args[0], args...)
		go func() {
			cmd.Start()
			cmd.Wait()
		}()
		time.Sleep(time.Millisecond * 300)
		return
	}
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

// 判断uuids文件是否存在，如果存在，则使用文件内的uuid作为心跳检测\查询cli server端的etcd\处理命令
func (this *Cli) Run() {
	if utils.IsExist("uuids") {

		uuids := utils.ReadFile("uuids")
		config.BENCHMARK = true

		for i, uuid := range strings.Split(uuids, "\n") {
			uuid = strings.TrimSpace(uuid)
			//go this.Heartbeat(uuid)
			//go this.WatchEtcd(uuid)
			//go this.DealCommands(uuid)
			fmt.Println(fmt.Sprintf("No:%v UUID:%v", i, uuid))

		}

	} else {
		//go this.Heartbeat("")
		//go this.WatchEtcd("")
		for i := 0; i < 30; i++ {
			//go this.DealCommands("")
		}
	}

	// 获取进程的内存使用情况
	getMem := func() uint64 {
		var ms runtime.MemStats
		runtime.ReadMemStats(&ms)
		return ms.HeapAlloc / 1024 / 1024
	}

	// 初次执行 cli daemon -s daemon时会删除旧的pid文件，然后进入for循环检查
	// 不断检查cli进程内存使用量是否超过500M；
	// 不断检查cli daemon -s daemon 进程数是否超过1
	go func() {
		this.ExecCmd([]string{"cli shell -d system -f install_cli_sdk.py"}, 10)
		this.stop()
		for {
			mpids := make(map[string]string)

			pids := this.getPids()
			for _, pid := range pids {
				mpids[pid] = pid
			}
			// 获取当前进程的pid，并写入PID_FILE
			mpids[fmt.Sprintf("%d", os.Getpid())] = fmt.Sprintf("%d", os.Getpid())
			var pps []string
			for k, _ := range mpids {
				if strings.TrimSpace(k) != "" {
					pps = append(pps, k)
				}
			}
			if fp, err := os.OpenFile(config.PID_FILE, os.O_CREATE|os.O_RDWR, 0644); err == nil {
				fp.WriteString(fmt.Sprintf("%s", strings.Join(pps, "\n")))
				fp.Close()
			} else {
				fmt.Println(err.Error())
			}
			if getMem() > 500 {
				log.Warn("cli process suicide...")
				os.Exit(1)
			}
			this.isRunning()
			if utils.RandInt(0, 30) < 3 {
				this.killPython()
			}
			time.Sleep(time.Second * 60)

		}
	}()

	select {}
}

// kill python进程，且进程带有cli daemon
func (this *Cli) killPython() {
	this.ExecCmd([]string{"ps aux|grep python|grep 'cli daemon'|awk '{print $2}'|xargs -n 1 kill"}, 10)
}
