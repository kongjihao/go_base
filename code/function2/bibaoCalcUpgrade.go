package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	add := func(x int) int {
		base += x
		return base // 注意这里，区别于： return base + x
	}

	sub := func(x int) int {
		base -= x
		return base // 注意这里，区别于： return base - x
	}

	return add, sub
}

func main() {
	x, y := calc(100)
	res := x(200)
	fmt.Println("inner add: ", res)

	res2 := y(200)
	fmt.Println("inner sub: ", res2)
}
