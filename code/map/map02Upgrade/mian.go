package main

import "fmt"

// 学习复杂类型的map
func main() {
	// 1.定义一个map类型的slice

	// 注意点：使用make初始化slice的时候可以指定长度和容量，而使用make初始化map的时候，只能指定容量
	var mapSlice = make([]map[string]int, 8, 8) // 此时只是完成了slice的初始化，并没有分配内存，注意这里指定长度不能为0

	// 看一下内层Map是否被初始化
	fmt.Println(mapSlice[0] == nil) // true

	// 还需要完成内部map的初始化
	mapSlice[0] = make(map[string]int, 4)
	mapSlice[0]["name"] = 2
	fmt.Println(mapSlice) // [map[name:2] map[] map[] map[] map[] map[] map[] map[]]

	// 2. 定义一个值为slice类型的map
	var sliceMap = make(map[string][]int, 8) // 注意只初始化了map，没有分配内存
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		sliceMap["中国"] = make([]int, 8, 8)
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 96
		sliceMap["中国"][2] = 97
		sliceMap["中国"][3] = 98
	}
	fmt.Println(sliceMap) // map[中国:[100 96 97 98 0 0 0 0]]

}
