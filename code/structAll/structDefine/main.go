package main

import "fmt"

// 定义结构体
/*
type structName struct {
	field1 type1
	field2 type2
	...
}
*/
// 定义 person 结构体，因为person也是type类型的，所以person也是一个类型（自定义类型）
type person struct {
	name, sex string // go语言中内存对齐的好处
	age       int
	married   bool
}

func main() {
	// 结构体的实例化，只有当实例化结构体后才会真正给其分配内存
	var p1 person

	p1.name = "Tom"
	p1.sex = "Male"
	p1.age = 20
	p1.married = false

	fmt.Printf("p1# = %#v\n", p1)
	fmt.Printf("p1+ = %+v\n", p1)

}
