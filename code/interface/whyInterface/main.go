package main

// 接口学习
// 接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
// 接口只有规范，而没有实现。接口是一种类型，一种抽象的类型。
// 每个接口类型由任意个方法签名组成，接口的定义格式如下：

/*
在Go语言中接口（interface）是一种类型，一种抽象的类型。相较于之前章节中讲到的那些具体类型（字符串、切片、结构体等）
更注重“我是谁”，接口类型更注重“我能做什么”的问题。接口类型就像是一种约定——概括了一种类型应该具备哪些方法，
在Go语言中提倡使用面向接口的编程方式实现解耦。
*/

// 为什么需要接口

type dog struct{}

func (d dog) speak() {
	println("wang wang ~~")
}

type cat struct{}

func (c cat) speak() {
	println("miao miao ~~")
}

type person struct {
	name string
}

func (p person) speak() {
	println("a~a~~")
}

// 接口不管你是什么类型，只管你要实现什么方法。只要他实现了这个接口，我就可以打他
// 定义一个类型，一个抽象类型，只要实现了speak()这个方法的类型，就都可以称为是speaker类型
// 也可以叫做多态？

// 定义一个接口
/*
type 接口类型名称 interface {
	方法名1(参数1) 返回值1
	方法名2(参数2) 返回值2
	方法名3(参数3) 返回值3
}
*/

type speaker interface {
	speak()
}

// 定义一个 打 函数，不管传进来什么类型参数我都要打他, 打他他就会叫
func da(arg speaker) {
	arg.speak()
}

func main() {
	cat := cat{}
	da(cat)

	dog := dog{}
	da(dog)

	person := person{name: "tom"}
	da(person)
}
