package core

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/yin-zt/cobra-cli/utils"
	"io"
	"log"
	"os"
	"os/user"
	"runtime"
	"strings"
)

// GetUUID 获取随机生成的UUID
func (this *Common) GetUUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	id := utils.MD5(base64.URLEncoding.EncodeToString(b))
	return fmt.Sprintf("%s-%s-%s-%s-%s", id[0:8], id[8:12], id[12:16], id[16:20], id[20:])

}

// GetProductUUID 获取本节点的UUID
func (this *Common) GetProductUUID() string {

	if "windows" == runtime.GOOS {
		uuid := this.windowsProductUUID()
		return uuid
	}

	filename := "/usr/local/cli/machine_id"
	uuid := utils.ReadFile(filename)
	if utils.IsExist("/usr/local/cli/") {
		os.Mkdir("/usr/local/cli/", 0744)
	}
	if uuid == "" {
		uuid := this.GetUUID()
		utils.WriteFile(filename, uuid)
	}
	return strings.Trim(uuid, "\n ")

}

// windowsProductUUID方法是在windows系统下先判断是否在用户家目录下存在.machine_id文件
// 如果存在，则判断里面是否存在本机的uuid，如果没有则获取本机uuid再写入文件中
func (this *Common) windowsProductUUID() string {
	user, err := user.Current()
	if err != nil {
		corelog.Debug(err)
		log.Fatalln(err)
	}

	filename := user.HomeDir + "/.machine_id"
	var uuid string
	if !utils.IsExist(filename) {
		uuid = this.GetUUID()
		utils.WriteFile(filename, uuid)
		return uuid
	}

	uuid = utils.ReadFile(filename)

	if uuid == "" {
		uuid = this.GetUUID()
		utils.WriteFile(filename, uuid)
		return uuid
	}

	return strings.Trim(uuid, "\n")
}
