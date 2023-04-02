package core

import (
	"fmt"
	"github.com/yin-zt/cobra-cli/utils"
	"regexp"
	"strconv"
	"strings"
)

// Jq 使用方法是echo '{"tt":"helloworld", "bb": "fufu"}' | gmd jq
// json解析器，-k json中的key，嵌套时使用逗号分隔，如 -k data,rows,ip
func (this *Common) Jq(module string, action string) {

	data, _ := this.StdinJson()
	if data == nil {
		return
	}
	key := ""
	var obj interface{}
	argv := utils.GetArgsMap()
	if v, ok := argv["k"]; ok {
		key = v
	}
	var ks []string
	if strings.Contains(key, ",") {
		ks = strings.Split(key, ",")
	} else {
		ks = strings.Split(key, ".")
	}

	obj = data

	// 解析传入data的接口字符串中，是否为字典类型，如果是则查找是否字典中有key这个键，有则返回key对应的值；
	// 如果不是字典类型，则返回nil
	ParseDict := func(obj interface{}, key string) interface{} {
		switch obj.(type) {
		case map[string]interface{}:
			if v, ok := obj.(map[string]interface{})[key]; ok {
				return v
			}
		}
		return nil

	}

	// 判断传入的obj接口类型的值是否为列表样式，如果是则继续判断key是数字样式还是字符串样式，
	// 如果是数字样式则，返回列表中索引下标为key的值；如果是字符串样式，则返回遍历列表中所有元素()
	ParseList := func(obj interface{}, key string) interface{} {
		var ret []interface{}
		switch obj.(type) {
		case []interface{}:
			if ok, _ := regexp.MatchString("^\\d+$", key); ok {
				i, _ := strconv.Atoi(key)
				return obj.([]interface{})[i]
			}

			for _, v := range obj.([]interface{}) {
				switch v.(type) {
				case map[string]interface{}:
					if key == "*" {
						for _, vv := range v.(map[string]interface{}) {
							ret = append(ret, vv)
						}
					} else {
						if vv, ok := v.(map[string]interface{})[key]; ok {
							ret = append(ret, vv)
						}
					}
				case []interface{}:
					if key == "*" {
						for _, vv := range v.([]interface{}) {
							ret = append(ret, vv)
						}
					} else {
						ret = append(ret, v)
					}
				}
			}
		}
		return ret
	}
	if key != "" {
		for _, k := range ks {
			switch obj.(type) {
			case map[string]interface{}:
				obj = ParseDict(obj, k)
			case []interface{}:
				obj = ParseList(obj, k)
			}
		}
	}

	switch obj.(type) {
	case map[string]interface{}, []interface{}:
		fmt.Println(utils.JsonEncodePretty(obj))
	default:
		fmt.Println(obj)
	}

}

// Json_val 使用方法是 echo '{"tt":"helloworld", "bb": "fufu"}' | cli json_val
// 直接调用jq方法
func (this *Common) Json_val(module string, action string) {
	this.Jq(module, action)
}
