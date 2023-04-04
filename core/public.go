package core

import (
	"bufio"
	"encoding/json"
	"os"
	"runtime"
	"strings"
)

func (this *Cli) StdinJson() (interface{}, string) {
	var lines []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		lines = append(lines, input.Text())
	}
	in := strings.Join(lines, "")
	var obj interface{}
	if err := json.Unmarshal([]byte(in), &obj); err != nil {
		corelog.Error(err, in)
		obj = nil
	}
	return obj, in
}

// IsWindows 判断是否为windows操作系统
func (this *Cli) IsWindows() bool {

	if "windows" == runtime.GOOS {
		return true
	}
	return false

}

// 在本地执行cmd列表里的命令，程序在linux下将会如下执行：
// linux: bash -c cmd1 cmd2 cmd3
func (this *Cli) ExecCmd(cmd []string, timeout int) string {

	var cmds []string

	if "windows" == runtime.GOOS {
		cmds = []string{
			"cmd",
			"/C",
		}
		for _, v := range cmd {
			cmds = append(cmds, v)
		}

	} else {
		cmds = []string{
			"/bin/bash",
			"-c",
		}
		for _, v := range cmd {
			cmds = append(cmds, v)
		}

	}
	result, _, _ := this.Exec(cmds, timeout, nil)
	return result
}
