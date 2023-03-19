package utils

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

// IsExist 方法用于判断文件是否存在
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// Home 返回当前用户的家目录
func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	return homeUnix()
}

// homeWindows获取windows用户家目录
func homeWindows() (string, error) {
	// 获取家目录磁盘盘符 C:   内置的环境变量
	drive := os.Getenv("HOMEDRIVE")
	// 获取用户家目录  内置的环境变量
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		// 获取用户家目录，内置环境变量
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

// homeUnix获取linux系统当前用户家目录
func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}
