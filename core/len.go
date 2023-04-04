package core

import (
	"fmt"
)

// Len 使用方法是：echo "aabbcc" | gmd len  或者  echo '{"key1": "val1", "key2": "val2"}' | gmd len
// 作用是返回输出字符串长度或者类字典数据中key的个数
func (this *Cli) Len() {
	obj, in := this.StdinJson()
	switch obj.(type) {
	case []interface{}:
		fmt.Println(len(obj.([]interface{})))
		return

	case map[string]interface{}:
		i := 0
		for _ = range obj.(map[string]interface{}) {
			i = i + 1
		}
		fmt.Println(i)
		return
	}

	fmt.Println(len(in))
}
