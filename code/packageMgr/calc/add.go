package calc

import "fmt"

// Name 定义一个全局变量
var Name = "沙河娜扎"

func Add(a, b int) int {
	return a + b
}

/*
init 函数特点：
init 函数在包导入的时候自动执行
init 函数没有参数也没有返回值

程序加载顺序：全局变量 --> init 函数 --> main 函数
*/
func init() {
	fmt.Println("calc add.go init 函数自动执行")
	fmt.Println("calc add.go Name = ", Name)
}
