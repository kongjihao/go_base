package main

import (
	"fmt"
)

// 数组，同一种数据类型的元素的集合
// 数组的长度是固定的，一旦初始化就不能改变，数组的长度一定是个常量
func main() {
	// 1. 声明一个长度为3的数组
	var arr [3]int
	arr[0] = 1       // 赋值
	arr[1] = 2       // 赋值
	arr[2] = 3       // 赋值
	fmt.Println(arr) // [1 2 3]

	// 2. 声明一个长度为3的数组并初始化
	var cityArray = [3]string{"北京", "上海", "广州"}
	fmt.Println(cityArray)

	// 3. 编译器推导出数组长度，声明一个长度为3的数组并初始化
	var arr2 = [...]int{1, 2, 3}     // 编译器推导出数组长度
	for i := 0; i < len(arr2); i++ { // 数组的遍历，for循环方法
		fmt.Println(arr2[i])
	}

	// 4. 声明一个长度为5的数组，前两个元素初始化为1和2，其他元素都为0
	var arr3 = [5]int{1, 2}
	fmt.Println(arr3)

	// 5. 数组的遍历，range方法
	for index, value1 := range arr3 {
		fmt.Printf("%d: %d\n", index, value1)
	}

	// 6. 多维数组，并遍历,
	// 注意点：多维数组只有最外层才可使用[...]其他层级不能用
	// var multiArr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var multiArr = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for _, v1 := range multiArr {
		for _, v2 := range v1 {
			fmt.Printf("%d \n", v2)
		}
	}

	// 7. 数组是值类型的，赋值时是值拷贝，修改一个数组不会影响另一个数组
	var a1 = [3]int{1, 2, 3}
	var a2 = a1
	a2[0] = 99
	fmt.Println("a1: ", a1) // a1:  [1 2 3]
	fmt.Println("a2: ", a2) // a2:  [99 2 3]
	f1(a1)
	fmt.Println("a1 user method f1(): ", a1) // a1 user method f1():  [1 2 3]

}

func f1(a [3]int) { // 数组是值类型，传递的是值拷贝
	a[0] = 100                               // 赋值
	fmt.Println("a inside method f1(): ", a) // a inside method f1():  [100 2 3]
}
