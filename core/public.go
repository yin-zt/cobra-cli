package core

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

func (this *Common) StdinJson() (interface{}, string) {
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
