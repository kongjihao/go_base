package main

import "fmt"

func main() {
	// 整数类型
	num := 10
	fmt.Printf("十进制: %d\n", num)
	fmt.Printf("八进制: %o\n", num)
	fmt.Printf("十六进制（小写）: %x\n", num)
	fmt.Printf("十六进制（大写）: %X\n", num)
	fmt.Printf("二进制: %b\n", num)

	// 浮点数类型
	fnum := 3.14159
	fmt.Printf("一般格式: %f\n", fnum)
	fmt.Printf("科学计数法（小写 e）: %e\n", fnum)
	fmt.Printf("科学计数法（大写 E）: %E\n", fnum)
	fmt.Printf("自动选择格式（小写）: %g\n", fnum)
	fmt.Printf("自动选择格式（大写）: %G\n", fnum)

	// 字符和字符串类型
	char := 'A'
	str := "Hello"
	fmt.Printf("单个字符: %c\n", char)
	fmt.Printf("字符串: %s\n", str)
	fmt.Printf("带转义的字符串: %q\n", str)

	// 布尔类型
	isTrue := true
	fmt.Printf("布尔值: %t\n", isTrue)

	// 指针类型
	var p *int
	p = &num
	fmt.Printf("指针值（内存地址）: %p\n", p)

	// 宽度和精度控制
	fmt.Printf("宽度为 5 的整数: %5d\n", num)
	fmt.Printf("宽度为 8 小数点后 2 位的浮点数: %8.2f\n", fnum)
	fmt.Printf("宽度为 6 的字符串: %6s\n", str)

	// 输出变量类型
	fmt.Printf("变量 num 的类型: %T\n", num)
}
