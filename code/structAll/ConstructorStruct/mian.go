package main

import "fmt"

// 结构体是值类型
type person struct {
	name, city string
	age        int
}

// 给结构体创建构造函数：构造一个结构体实例化的函数
/*
	Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。
	因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，
	所以该构造函数返回的是结构体指针类型。
*/
func initPerson(name, ctiy string, age int) *person {
	return &person{
		name: name,
		city: ctiy,
		age:  age,
	}
}
func main() {
	p := initPerson("Tom", "BeiJing", 18)
	fmt.Printf("*p = %#v\n", *p)
	fmt.Printf("p = %#v\n", p)

}
