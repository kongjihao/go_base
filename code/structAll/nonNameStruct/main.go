package main

import "fmt"

// 匿名结构体学习
// 在定义一些临时数据结构等场景下还可以使用匿名结构体。
func main() {
	// 临时变量定义，匿名结构体
	var user struct {
		name    string
		age     int
		married bool
	}

	user.name = "张三"
	user.age = 20
	user.married = false
	fmt.Printf("user: %#v\n", user)

}
