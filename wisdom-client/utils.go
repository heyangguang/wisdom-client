package wisdom_client

import (
	"fmt"
	"os"
	"strings"
)

// 获取项目路径
func BaseDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("get project location failed, err: %v \n", err))
	}
	return dir
}

// 将host分隔开
func StrSeparate(str string) []string {
	strSlice := strings.Split(str, ":")
	return strSlice
}
