package main

import "fmt"

// 函数外的变量名都需要通过关键字var\const\type\func来声明
const (
	pi = 3.14
	e  = 2.7
	// 如果常量声明后面的变量没有赋值，则默认是上面紧挨着常量值一样
	n1
	n2
	/*
		iota是go语言的常量计数器，只能在常量的表达式中使用。
		iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
		使用iota能简化定义，在定义枚举时很有用。
	*/
)

// 变量
func main() {
	// 标准声明格式
	var age int
	// 声明一个变量并初始化变量值
	var name string = "Jerry"
	// 一次声明多个变量
	var age1, age2 int = 11, 22
	var isOk bool

	age = 19 + 9

	fmt.Println(age, name, isOk, age1, age2)

	// var c []int
	/*
		在Go语言中，一个未指定具体类型的整数类型默认为int，而不是int32或uint。因此，在定义切片var c []int时，切片中的元素类型默认是int，即有符号整数类型。
		具体来说：
			在64位系统上，int类型是64位有符号整数类型。
			在32位系统上，int类型是32位有符号整数类型。
		如果您想要使用无符号整数类型，可以显式地指定切片中的元素类型为uint，例如var c []uint。
		需要注意的是，Go语言中的整数类型默认是有符号的，如果需要无符号整数类型，需要显式地指定为uint类型。
	*/

	// 批量声明
	var (
		a int
		b string
		c bool
		d float32
	)

	a = 1
	b = "Tom"
	c = true
	d = 3.1415926

	fmt.Println(a, b, c, d)

	// 类型推导
	// 有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。
	var name3 = "kongjihao"
	var age4 = 25
	fmt.Println(name3, age4)

	// 短变量声明，只能在函数内部使用
	name4 := "chenyanjun"
	age5 := 25
	fmt.Println(name4, age5)

	// 匿名变量,比如一个函数返回值有两个，只想要其中一个，就可以使用匿名变量忽略不想要的值
	// 多用于占位，不占用内存
	_, _ = name4, age5

	// 常量
	const nihao = 1
	fmt.Println(pi, nihao, e, n1, n2)

}
