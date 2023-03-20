package utils

import (
	"bufio"
	"encoding/json"
	"github.com/henrylee2cn/mahonia"
	"os"
	"strings"
)

// StdinJson 使用方法是：echo "helloworld \n wocao" | cli StdinJson  //但不会有输出
// 作用获取用户输出，并将输入内容连成字符串返回，并不直接调用，而是由其他函数调用
func StdinJson() (interface{}, string) {

	var lines []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		lines = append(lines, input.Text())
	}
	in := strings.Join(lines, "")
	var obj interface{}
	if err := json.Unmarshal([]byte(in), &obj); err != nil {
		utilslog.Error(err, in)
		obj = nil
	}
	return obj, in
}

// GBKToUTF 作用是将GBK编码的字符串转换为UTF-8编码的字符串
func GBKToUTF(str string) string {
	decoder := mahonia.NewDecoder("GBK")
	if decoder != nil {
		if str, ok := decoder.ConvertStringOK(str); ok {
			return str
		}
	}
	return str
}
