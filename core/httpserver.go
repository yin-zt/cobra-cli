package core

import (
	"fmt"
	"github.com/yin-zt/cobra-cli/utils"
	"net/http"
)

func (this *Common) Httpserver(host, path string, port int) {
	defer corelog.Flush()

	if path == "" {
		if v, err := utils.Home(); err == nil {
			path = v
		}
	}

	// 设置http服务的根目录
	h := http.FileServer(http.Dir(path))
	fmt.Println(fmt.Sprintf("http server listen %s:%s", host, port))
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), h)

	if err != nil {
		corelog.Errorf("Error starting http server:", err)
		fmt.Println(err)
	}

}
