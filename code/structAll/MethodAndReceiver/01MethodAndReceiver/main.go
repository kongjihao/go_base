package main

import "fmt"

// 方法的定义

// Person是一个向外暴露的结构体
type Person struct {
	Name string
	Age  int
}

// newPerson 为 Person 类型（结构体）的构造函数
func newPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age, // 不要忘记最后一项后也要加逗号
	}
}

// SayHello 是为Person 类型（结构体）定义方法
func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s.\n", p.Name)
}

// Dream 是为Person 类型（结构体）定义方法
func (p Person) Dream() {
	fmt.Printf("%s's dream is to be a famous golang developer.\n", p.Name)
}

// SetAge 是使用 Person指针类型的接收者，来实现修改年龄的方法 ，指针类型在接收者方法中做的修改是会对原对象生效的
func (p *Person) SetAge(newAge int) {
	(*p).Age = newAge
	fmt.Println("SetAge method is called.")
}

// SetAge2 是使用Person值类型的接收者，来实现修改年龄的方法，值类型在接收者方法中做的修改是不会对原对象生效的
func (p Person) SetAge2(newAge int) {
	p.Age = newAge
	fmt.Println("SetAge2 method is called.")
}

func main() {
	// 调用Person类型的构造函数，实例化对象
	p := newPerson("Jack", 18)
	// 调用方法
	(*p).SayHello()
	(*p).Dream()
	// 也可以写成下面这样，go语言提供的语法糖
	p.SayHello()
	p.Dream()

	// 调用 Person指针类型 的方法
	(*p).SetAge(21)
	fmt.Println("Age2 is: ", (*p).Age)
	(*p).SetAge2(25)
	fmt.Println("Age3 is: ", (*p).Age)

	// 什么时候使用指针类型的接收者？
	/*
		1.  需要修改接收者中的值
		2.  接收者是拷贝代价比较大的大对象
		3.  保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

	*/
}
