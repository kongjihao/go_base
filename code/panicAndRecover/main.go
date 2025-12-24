package main

import "fmt"

func a() {
	fmt.Println("a")
}

func b() {
	/*
		使用panic/recover模式来处理错误。
		panic可以在任何地方引发，但recover只有在defer调用的函数中有效。
		1. recover()必须搭配defer使用。
		2. defer一定要在可能引发panic的语句之前定义。
	*/

	// defer 用于怀疑某段代码可能发生panic，如果发生panic，则执行recover()，并返回错误信息
	// 如果没有执行recover()，则程序panic 便不会向下执行了

	/*
		程序运行期间funcB中引发了panic导致程序崩溃，异常退出了。
		这个时候我们就可以通过recover将程序恢复回来，继续往后执行。
	*/
	defer func() {
		err := recover() // 此时recover() 获取的是panic的错误信息，赋值给err
		if err != nil {  // 判断err是否为空，不为空，则打印错误信息
			fmt.Println(err)
		}
	}() // 调用defer函数，在函数返回之前执行，可以用来释放资源,defer函数执行顺序，先进后出，注意加上最后的（）来执行匿名函数
	panic("panic in b")
}

func c() {
	fmt.Println("a")
}
func main() {
	a()
	b()
	c()
}
