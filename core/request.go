package core

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/yin-zt/cobra-cli/utils"
	"log"
)

var (
	body map[string]string
	req  *httplib.BeegoHTTPRequest
	html string
)

func (this *Common) Request(url, data string) {

	if data == "" {
		req = httplib.Get(url)
		if html, err = req.String(); err != nil {
			corelog.Error(err)
			log.Fatalln(err)
		}
		fmt.Println(utils.GBKToUTF(html))
		return
	}
	if err = json.Unmarshal([]byte(data), &body); err != nil {
		fmt.Println(err)
		return
	}
	req = httplib.Post(url)
	for k, v := range body {
		req.Param(k, v)
	}
	if v, ok := body["f"]; ok {
		if utils.IsExist(v) {
			req.PostFile("file", v)
		}
	}

	if html, err = req.String(); err != nil {
		corelog.Error(err)
		log.Fatalln(err)
	}
	fmt.Println(utils.GBKToUTF(html))
	return
}
