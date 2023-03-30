package core

import (
	"fmt"
	"github.com/yin-zt/cobra-cli/utils"
)

func (this *Common) Keys() {
	obj, _ := this.StdinJson()
	fmt.Println(obj)
	var keys []string
	switch obj.(type) {
	case map[string]interface{}:
		for k, _ := range obj.(map[string]interface{}) {
			keys = append(keys, k)
		}
	}
	fmt.Println(utils.JsonEncodePretty(keys))
}
