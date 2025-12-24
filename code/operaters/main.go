package main

import "fmt"

// GO 运算符学习
func main() {
	// 1. 算数运算符
	var a, b int32 = 10, 20
	fmt.Println(a + b)          // 30
	fmt.Println(a - b)          // -10
	fmt.Println(a * b)          // 200
	fmt.Println(a / b)          // 0
	fmt.Println(float32(a / b)) // 0
	/*
		在Go语言中，当对两个整数进行除法运算时，结果会被截断为整数部分，即保留整数部分并丢弃小数部分。
		因此，在表达式a / b中，a和b都是int32类型，结果也会是int32类型，小数部分会被截断。在代码中，
		a除以b的结果为0，因为10 / 20的整数部分为0。
		然后，当将结果转换为float32类型时，已经是整数部分为0的值，转换为float32类型后仍然是0。
		因此，即使在表达式fmt.Println(float32(a / b))中进行了类型转换为float32，由于整数部分已经是0，因此结果仍然是0。
		如果您希望得到浮点数的结果，可以先将其中一个操作数转换为浮点数，例如：
	*/
	fmt.Println(float32(a) / float32(b)) // 0.5
	fmt.Println(5 % 2)                   //1
	a++
	b--
	fmt.Println(a) // 11
	fmt.Println(b) // 19

	// 2. 关系运算符
	var c, d int32 = 10, 20
	fmt.Println(c == d) // false
	fmt.Println(c != d) // true
	fmt.Println(c > d)  // false
	fmt.Println(c < d)  // true
	fmt.Println(c >= d) // false
	fmt.Println(c <= d) // true

	// 3. 逻辑运算符
	var e, f bool = true, false
	fmt.Println(e && f) // false
	fmt.Println(e || f) // true
	fmt.Println(!e)     // false

	// 4. 位运算符
	/*
		&	参与运算的两数各对应的二进位相与。（两位均为1才为1）
		|	参与运算的两数各对应的二进位相或。（两位有一个为1就为1）
		^	参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。（两位不一样则为1）
		«	左移n位就是乘以2的n次方。 “a«b"是把a的各二进位全部左移b位，高位丢弃，低位补0。
		»	右移n位就是除以2的n次方。 “a»b"是把a的各二进位全部右移b位。
	*/
	var g, h uint8 = 60, 13
	fmt.Println(g & h)  // 12   与
	fmt.Println(g | h)  // 61   或
	fmt.Println(g ^ h)  // 49   异或
	fmt.Println(g << 2) // 240  左移
	fmt.Println(g >> 2) // 15   右移
	// 下面是简单的位运算符的运算规则：
	fmt.Println(1 << 2) // 4   1*2^2  0001 -> 0100  (1向左移动2位)
	fmt.Println(4 >> 2) // 1   4/2^2  0100 -> 0001  (1向右移动2位)

	// 5. 赋值运算符
	var i, j int32 = 10, 20
	i += j
	fmt.Println(i) // 30

	var k, l int32 = 10, 20
	k -= l
	fmt.Println(k) // -10
}
