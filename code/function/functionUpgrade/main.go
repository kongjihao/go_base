// package functionupgrade  // package command-line-arguments is not a main package
package main

import "fmt"

var a int = 10

func testGloab() {
	/*
		1. a变量先在自己内层函数中找
		2. 找不到就到函数外层中找全局变量
		3. 找不到就报错
	*/
	a := 20
	fmt.Println(a)
}

func testInner() {
	// i 变量是在for循环中，所以只能在for循环中使用
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}
	// fmt.Println(i) // 此时会报错
}

// 函数进阶之变量作用域
func main() {
	// 1. 函数变量作用域
	testGloab() // print = 10
	testInner() // inner函数中i变量作用域为for循环

	// 2. 匿名函数
	func() {
		println("Hello World!")
	}() // （）为执行函数，输出Hello World!

	// 3. 函数作为参数
	// 3.1
	abc := testInner        // 函数赋值给变量
	fmt.Printf("%T\n", abc) // func()类型
	abc()                   // 函数调用
	// 3.2
	func(a int, b int) { // 函数定义
		println(a + b)
	}(10, 20) // 函数调用  print = 30

	// 4. 函数作为变量
	func1 := func(a int, b int) { // 函数定义
		println(a + b)
	}

	func1(10, 20) // 函数调用  print = 30

	// 5. 闭包，反回了一个函数
	func3 := func(a int) func() { // 函数定义
		return func() { // 匿名函数，闭包，返回了一个函数
			println(a)
		}
	}(10)

	func3() // 函数调用  print = 10

}
