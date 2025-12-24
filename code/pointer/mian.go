package main

import "fmt"

// 指针学习 v1.0
// func main() {
// 	var a int = 10
// 	var b *int = &a // *int 指针类型，&a 取地址，*取值，b 指针类型，*b 取值

// 	// a = 10, aAddr = 0x14000122018, b = 0x14000122018, bAddr = 0x14000128018
// 	// %v 打印变量真实值，%p 打印地址, %d 打印十进制, %x 打印十六进制, %T 打印类型
// 	fmt.Printf("a = %v, aAddr = %p, b = %v, bAddr = %p\n", a, &a, b, &b)

// 	// 指针取值，根据内存地址取值
// 	fmt.Println(*b)
// }

// 指针学习 v2.0
func main() {
	a := 1
	modify1(a)
	fmt.Println(a) // 1

	modify2(&a)
	fmt.Println(a) // 10

}

func modify1(x int) {
	x = 10
}

func modify2(x *int) {
	*x = 10
}

// 指针学习 v3.0
// // 空指针异常, 空map异常, 需注意
// func main() {
// 	var a *int
// 	// a = nil, 此时a 是一个空指针，因为没有指向任何变量，只是声明了变量，没有赋值，所以取值会报错
// 	// 需要先用new初始化，new返回一个指针，指向一个地址， func new(T) *T，a = new(int)
// 	*a = 10 // panic: runtime error: invalid memory address or nil pointer dereference
// 	fmt.Println(*a)

// 	var b map[string]int
// 	// b = nil, map 是一个引用类型，nil 代表没有指向任何变量，所以赋值会报错，需要先make初始化一个map类型
// 	// b = make(map[string]int, 10) // make 返回一个map类型, map容量为10, func make(t Type, size ...integerType) Type，make只用来初始化map、slice、channel，返回的是三个引用类型本身
// 	b["沙河娜扎"] = 10 // panic: assignment to entry in nil map
// 	fmt.Println(b)
// }
