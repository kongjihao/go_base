package main

import "fmt"

// 切片内存分配
func main() {
	// 声明一个切片，此时不会分配内存地址
	var a []int32
	// 声明一个切片，并初始化，此时会分配内存地址
	var b = []int32{}
	var c = make([]int32, 0) // 等价于 var c = []int32{}  等价于 c := make([]int32, 0)

	if a == nil {
		println("a is nil") // a is nil
	} else {
		println("a is not nil")
	}

	if b == nil {
		println("b is nil")
	} else {
		println("b is not nil") // b is not nil
	}

	if c == nil {
		println("c is nil")
	} else {
		println("c is not nil") // c is not nil
	}

	fmt.Println("a length and capacity:", len(a), cap(a)) // a length and capacity: 0 0
	fmt.Println("b length and capacity:", len(b), cap(b)) // b length and capacity: 0 0
	fmt.Println("c length and capacity:", len(c), cap(c)) // c length and capacity: 0 0

}
