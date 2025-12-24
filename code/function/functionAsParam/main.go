package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calculate(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	fmt.Println(calculate(10, 5, add))
	fmt.Println(calculate(10, 5, sub))
}
