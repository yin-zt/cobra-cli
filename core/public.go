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
