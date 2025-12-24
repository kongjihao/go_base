package main

import "fmt"

// switch 语句
func main() {
	var a int = 20
	b := 20
	fmt.Println(b)

	// 1.switch 判断a的值
	switch a {
	case 1:
		println("a is 1")
	case 2:
		println("a is 2")
	default:
		println("a is default")
	}

	// 2.switch 判断a的值，如果满足条件就执行相应代码块
	switch b {
	case 1, 3, 5:
		println("b is odd number")
	case 2, 4, 6:
		println("b is even number")
	default:
		println("b is default")
	}

	// 3.switch 没有表达式，默认是true
	switch {
	case true:
		println("true")
	default:
		println("false")
	}

	// 4.switch 可以不带表达式，默认是true，相当于 if-else
	switch {
	case false:
		println("false")
	case true:
		println("true")
	default:
		println("default")
	}

	// 5.switch 中可以添加初始化语句，变量作用域只在 case 中有效
	switch x := a; x {
	case 1:
		println("x is 1")
	case 2:
		println("x is 2")
	default:
		println("x is default")
	}

	// 6.switch 中可以添加多个匹配条件，用逗号分隔
	switch y := a; y {
	case 1, 3, 5:
		println("y is odd number")
	case 2, 4, 6:
		println("y is even number")
	default:
		println("y is default") // y is default
	}

	// 7.switch 中可以添加fallthrough关键字，表示继续向下执行
	switch z := 1; z {
	case 1:
		println("z is 1") // z is 1
		fallthrough
	case 2:
		println("z is 2") // z is 2
	default:
		println("z is default")
	}

	// 8.switch 中可以添加break关键字，表示跳出switch语句
	switch w := 1; w {
	case 1:
		println("w is 1") // w is 1
		break
	case 2:
		println("w is 2")
	default:
		println("w is default")
	}

}
