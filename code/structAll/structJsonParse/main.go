package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	ID   int    `json:"stuID"` // `json:"id"`为结构体的tag
	Name string `json:"stuName"`
}

type class struct {
	Name     string    `json:"className"`
	Students []student `json:"studentList" db:"students" xml:"students"`
}

// student 构造函数
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

func main() {
	c := class{
		Name:     "火箭101班",
		Students: make([]student, 0, 20),
	}

	for i := 0; i < 10; i++ {
		// 创建10个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c.Students = append(c.Students, tmpStu)
	}

	fmt.Printf("%#v\n", c)
	fmt.Println("-----------------------------------------\n")

	// Json 序列化: Go语言中的数据格式  --> 满足Json 格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("c Json序列化失败:", err)
		return
	} else {
		fmt.Printf("%T\n%s\n", data, data)
	}

	fmt.Println("-----------------------------------------\n")

	// Json 反序列化: 满足Json 格式的字符串  --> Go语言中的数据格式
	var c2 class
	err = json.Unmarshal(data, &c2)
	if err != nil {
		fmt.Println("c data Json反序列化失败:", err)
	} else {
		fmt.Printf("%T\n%#v\n", c2, c2)
	}

}
