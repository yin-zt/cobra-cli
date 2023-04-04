package core

import (
	"fmt"
	"github.com/yin-zt/cobra-cli/utils"
	"log"
	"net/http"
)

func (this *Cli) Httpserver(host, path string, port int) {
	defer corelog.Flush()

	if path == "" {
		if v, err := utils.Home(); err == nil {
			path = v
		}
	}

	// 设置http服务的根目录
	h := http.FileServer(http.Dir(path))
	fmt.Println(fmt.Sprintf("http server listen %s:%d", host, port))
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), h)

	if err != nil {
		corelog.Errorf("Error starting http server:", err)
		log.Fatalln(err)
	}

}
