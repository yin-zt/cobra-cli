package core

import (
	"fmt"
	"github.com/yin-zt/cobra-cli/utils"
	"regexp"
)

// Match 检查给定字符串满足正则的子串
func (this *Common) Match(matchStr, modeStr, outputStr string) {
	var is_all bool
	for i := 0; i < len(outputStr); i++ {
		if string(outputStr[i]) == "i" {
			modeStr = "(?i)" + modeStr
		}
		if string(outputStr[i]) == "a" || string(outputStr[i]) == "m" {
			is_all = true
		}
	}

	if reg, err := regexp.Compile(modeStr); err == nil {

		if is_all {
			ret := reg.FindAllString(matchStr, -1)
			if len(ret) > 0 {
				fmt.Println(utils.JsonEncodePretty(ret))
			} else {
				fmt.Println("")
			}
		} else {
			ret := reg.FindString(matchStr)
			fmt.Println(ret)
		}

	}
}
