package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + ".go"
		}
		return name
	}
}

// 使用闭包做文件后缀名检测
func main() {
	f1 := makeSuffixFunc(".go")
	name := f1("test")
	fmt.Println(name)

	f2 := makeSuffixFunc(".avi")
	name2 := f2("test2.avi")
	fmt.Println(name2)
}
