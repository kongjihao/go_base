package main

import "fmt"

// 结构体嵌套学习
type Person struct {
	name    string
	age     int
	address Address // 嵌套另外一个结构体Address，此处还可以使用嵌套匿名结构体
}

type Person1 struct {
	name    string
	age     int
	Address // 使用嵌套匿名结构体
}

type Address struct {
	province string
	city     string
}

func main() {
	// 试用一下
	var p Person
	p = Person{
		name: "Tom",
		age:  20,
		address: Address{
			province: "BeiJing",
			city:     "海淀",
		},
	}

	fmt.Printf("%+v", p)
}
