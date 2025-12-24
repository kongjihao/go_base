package main

import "fmt"

// 函数学习
func main() {
	/*
		函数是完成特定功能的代码块, 可以提高代码的可读性和可维护性，减少代码冗余，
		在go语言中，函数是第一类对象，函数可以作为参数传递给其他函数, 也可以作为返回值
		创建函数注意点;
		1. 函数名区分大小写，函数名首字母大写, 外部才能访问（函数名首字母大写, 其他包可以调用）
		2. 函数名区分大小写，函数名小写, 只能在本包内使用（ 函数名首字母小写, 其他包不能调用）
		3. 在同一个包内，函数名不能重复
		4.函数名由字母、数字、下划线组成，不能以数字开头
		5.函数名遵循驼峰命名法
	*/

	/*
		函数声明
		func 函数名(形参列表) (返回值列表){
			// 函数体
		}
	*/

	// 调用函数
	// add := func(a, b int) {
	// 	fmt.Println("a+b=", a+b)
	// }
	sum := add(10, 20)
	fmt.Println("sum=", sum)

}

func add(a, b int) int {
	// return (a + b)
	sum := a + b
	return sum
}

func add2(a, b int) (sum int) {
	sum = a + b
	return sum // 这里的sum也可以省略不写，因为go语言是自动返回，但是不推荐
}

// 函数接收可变参数，在参数名后面加...表示可变参数，可变参数的函数在调用时可以传递任意个参数（包括什么也不传）,在函数内部通过for循环遍历
// 可变参数在函数体中是切片类型
// 可变参数类型必须一致，不可以混用，否则会报错
func sum2(nums ...int) int {
	fmt.Println("nums=", nums)
	fmt.Printf("type of nums: %T\n", nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 当固定参数和可变参数一起使用时，可变参数必须放在最后, 此时在调用函数sum3时，a参数必须传入，b参数可传可不传
func sum3(a int, b ...int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}

// go语言中函数可以返回多个值，多个返回值必须用括号括起来
func calc(a, b int) (sum, sub int) {
	// return a + b, a - b
	sum = a + b
	sub = a - b
	return
}
