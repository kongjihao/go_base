package main

import "fmt"

type person struct {
	name, city string
	age        int
}

// 结构体指针
func main() {
	var p = new(person) // 自动分配内存空间，p 是一个指针，指向 person 结构体，类型推导 p 的类型为 *person
	fmt.Printf("p = %#v\n", p)
	// 这种赋值不好看也麻烦，所以go语言提供了一个语法糖
	// (*p).name = "Tom"
	// (*p).age = 18
	// (*p).city = "BeiJing"

	//用go语言提供的语法糖
	p.name = "Tom"
	p.age = 18
	p.city = "BeiJing"

	// 也可以通过这种方式为p 赋值，取结构体的地址进行实例化，但是这种写法不推荐
	// p3 = &person{
	// 	name: "Tom",
	// 	age:  18,
	// 	city: "BeiJing",
	// }

	fmt.Printf("p2 = %#v\n", p)

}
