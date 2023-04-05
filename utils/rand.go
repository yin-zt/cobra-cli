package utils

import (
	random "math/rand"
	"time"
)

// RandInt 生成一个在某个区间内的随机整数
func RandInt(min, max int) int {
	r := random.New(random.NewSource(time.Now().UnixNano()))
	if min >= max {
		return max
	}
	return r.Intn(max-min) + min
}
