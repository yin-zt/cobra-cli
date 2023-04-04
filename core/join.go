package core

import (
	"fmt"
	"strings"
)

// Join 使用方法是 echo '["aa", "bb", "cc"]' | gmd join -s "-" -w "GG"
// 作用是将输出的字符串列表使用指定字符拼接，-w表示列表元素拼接前在其前后端加上val，-t参数目前没有作用
func (this *Cli) Join(joinS, joinW string) {
	obj, _ := this.StdinJson()
	sep := joinS
	wrap := joinW
	trim := ""
	//argv := utils.GetArgsMap()
	var lines []string
	switch obj.(type) {
	case []interface{}:
		for _, v := range obj.([]interface{}) {
			if trim != "" {
				if v == nil || fmt.Sprintf("%s", v) == "" {
					continue
				}
			}
			if wrap != "" {
				lines = append(lines, fmt.Sprintf("%s%s%s", wrap, v, wrap))
			} else {
				lines = append(lines, fmt.Sprintf("%s", v))
			}
		}
	}
	fmt.Println(strings.Join(lines, sep))
}
