package main

import "fmt"

// 函数外的变量名都需要通过关键字var\const\type\func来声明
const (
	pi = 3.14
	// e  = 2.7
	// 如果常量声明后面的变量没有赋值，则默认是上面紧挨着常量值一样
	m1
	m2
	/*
		iota是go语言的常量计数器，只能在常量的表达式中使用。
		iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
		使用iota能简化定义，在定义枚举时很有用。
	*/
)

const (
	n1 = iota
	n2
	n3 = 100
	n4
	n5 = iota
)

const n6 = iota

// 定义 单位
const (
	_  = iota
	KB = 1 << (10 * iota) // 1 << (10 * iota) = 1 << 10 = 1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 //iota = 0,  1,2
	c, d                      //iota = 1,  2,3
	e, f                      //iota = 1,  3,4
)

func main() {
	fmt.Println("验证iota关键字能力：")
	fmt.Println(n1, n2, n3, n4, n5, n6)
	fmt.Println(a, b, c, d, e, f)
}
