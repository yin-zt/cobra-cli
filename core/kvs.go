package core

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Kvs 使用方法是：echo [k1, k2, k3] | gmd kvs  || echo '{"k1": "v1", "k2": "v2"}' | gmd kvs
// echo '{"aa": {"test":"helloworld"}, "bb": {"key1": "value"}}' | gmd kvs
// 作用是将输入的列表样式的字符串或者字典样式的字符串，转换为显示友好的字典样式
func (this *Common) Kvs() {
	obj, _ := this.StdinJson()
	var keys []string
	switch obj.(type) {
	case map[string]interface{}:
		for k, v := range obj.(map[string]interface{}) {
			switch v.(type) {
			case map[string]interface{}, []interface{}:
				if b, e := json.Marshal(v); e == nil {
					s := strings.Replace(string(b), "\\", "\\\\", -1)
					s = strings.Replace(s, "\"", "\\\"", -1)
					keys = append(keys, fmt.Sprintf(k+"=\"%s\"", s))
				}
			default:
				keys = append(keys, fmt.Sprintf(k+"=\"%s\"", v))
			}
		}
	case []interface{}:
		for i, v := range obj.([]interface{}) {

			keys = append(keys, fmt.Sprintf("a%d=\"%s\"", i, v))

		}
	}
	fmt.Println(strings.Join(keys, "\n"))
}
