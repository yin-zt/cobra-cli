package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

func GetArgsMap() map[string]string {

	return ParseArgs(strings.Join(os.Args, "$$$$"), "$$$$")

}

// ParseArgs 将传入字符串进行解析处理，传入字符串一般是[ -s$$$$xx$$$$-f$$$$filename$$$$--test$$$$aabb]等模式
func ParseArgs(args string, sep string) map[string]string {

	ret := make(map[string]string)

	var argv []string

	argv = strings.Split(args, sep)
	for i, v := range argv {
		if strings.HasPrefix(v, "-") && len(v) == 2 {
			if i+1 < len(argv) && !strings.HasPrefix(argv[i+1], "-") {
				ret[v[1:]] = argv[i+1]
			}
		}

	}
	for i, v := range argv {
		if strings.HasPrefix(v, "-") && len(v) == 2 {
			if i+1 < len(argv) && strings.HasPrefix(argv[i+1], "-") {
				ret[v[1:]] = "1"
			} else if i+1 == len(argv) {
				ret[v[1:]] = "1"
			}
		}

	}

	for i, v := range argv {
		if strings.HasPrefix(v, "--") && len(v) > 3 {
			if i+1 < len(argv) && !strings.HasPrefix(argv[i+1], "--") {
				ret[v[2:]] = argv[i+1]
			}
		}

	}
	for i, v := range argv {
		if strings.HasPrefix(v, "--") && len(v) > 3 {
			if i+1 < len(argv) && strings.HasPrefix(argv[i+1], "--") {
				ret[v[2:]] = "1"
			} else if i+1 == len(argv) {
				ret[v[2:]] = "1"
			}
		}

	}
	for k, v := range ret {
		if k == "file" && IsExist(v) {
			ret[k] = ReadFile(v)
		}
	}
	return ret

}

// ReadFile 将文件内容读取出来并返回
func ReadFile(path string) string {
	if IsExist(path) {
		fi, err := os.Open(path)
		if err != nil {
			return ""
		}
		defer fi.Close()
		fd, err := ioutil.ReadAll(fi)
		return string(fd)
	} else {
		return ""
	}
}
