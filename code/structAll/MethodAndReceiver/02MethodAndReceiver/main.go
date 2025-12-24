package main

import "fmt"

// 由于go原生要求 只能对自己写的代码包中的类型（type定义的）才能定义接受者类型的方法
// 所以如何能对所有类型都定义方法呢
// 答：可以通过 type 封装一层，go的原生类型

type Myint int

func (m Myint) SayHello() {
	fmt.Println("hello")
}

func main() {
	var m Myint
	fmt.Println(m) // 0
	m.SayHello()
}
