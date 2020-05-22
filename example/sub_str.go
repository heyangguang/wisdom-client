package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "127.0.0.1:3306:mysql-test"
	newStr := strings.Split(str, ":")
	fmt.Println(newStr)
}
