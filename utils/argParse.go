package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"runtime"
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

// WriteFile 把输入参数的内容变量写到文件中；如果存在文件，则先删除后创建；如果不存在则直接创建
func WriteFile(path string, content string) bool {
	var f *os.File
	var err error
	if IsExist(path) {
		err = os.Remove(path)
		if err != nil {
			return false
		}
		f, err = os.Create(path)
	} else {
		f, err = os.Create(path)
	}

	if err == nil {
		defer f.Close()
		if _, err = io.WriteString(f, content); err == nil {
			//log.Debug(err)
			return true
		} else {
			return false
		}
	} else {
		//log.Warn(err)
		return false
	}

}

// MD5File 输出文件内容的md5值
func MD5File(fn string) string {
	file, err := os.Open(fn)
	if err != nil {
		return ""
	}
	defer file.Close()
	md5 := md5.New()
	io.Copy(md5, file)
	return hex.EncodeToString(md5.Sum(nil))
}

// IsWindows 判断是否为windows操作系统
func IsWindows() bool {

	if "windows" == runtime.GOOS {
		return true
	}
	return false

}
