package main

import "fmt"

// 空接口：接口中没有定义任何方法的接口，也就是说，任何类型都实现了空接口。空接口一般不需要提前定义，直接使用即可。
/*
type xxx interface {
	// 空接口:接口中没有定义要求要实现的方法，空接口类型的变量可以存储任意类型的值
}
*/

/*
空接口的使用场景：
1. 函数的参数可以是任意类型的值
2. 返回值可以是任意类型的值
3. 切片可以存储任意类型的值
4. map 的 value 可以是任意类型的值， 空接口作为map的value

*/

func main() {
	var x interface{} // 定义了一个空接口类型的变量 x
	x = 10            // x 可以存储 int 类型的值
	x = "hello"       // x 可以存储 string 类型的值
	x = true          // x 可以存储 bool 类型的值
	// fmt.Println(x)    // true

	// map的空接口的使用
	// var m = make(map[string]interface{}, 20)
	// m["name"] = "张三"
	// m["age"] = 18
	// m["sex"] = true
	// m["hobby"] = []string{"足球", "篮球"}
	// fmt.Println(m) // map[age:18 name:张三]

	// 空接口类型断言
	ret := x.(bool)  // x 转为 int 类型
	fmt.Println(ret) // true

	ret1, ok := x.(int) // 类型断言，x 断言为 int 类型
	if ok {
		fmt.Println(ret1) // 10
	} else {
		fmt.Println("类型断言失败")
	}

	// 使用switch 进行接口断言
	switch v := x.(type) {
	case bool:
		fmt.Println("bool 类型: ", v) // true
	case int:
		fmt.Println("int 类型: ", v) // 10
	case string:
		fmt.Println("string 类型: ", v) // "xxx"
	default:
		fmt.Println("其他类型，猜不到了！")
	}

}
