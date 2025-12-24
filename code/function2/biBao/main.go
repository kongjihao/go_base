package main

import "fmt"

// 闭包函数
// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包 = 函数+引用环境。
// 定义一个函数，他的返回值为一个函数，且这个函数中使用了外部函数的变量name
func a(name string) func() {
	return func() {
		fmt.Println("hello, " + name)
	}
}

func main() {
	// 调用 a() 函数，返回一个函数
	f := a("沙盒娜扎") // 此时f就是一个闭包
	f()            // 相当于执行了a函数内部的匿名函数，输出： hello, 沙盒娜扎

}
