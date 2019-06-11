package main

import "fmt"

type person1 struct {
	name string
	age  int
	sex  string
}

type student1 struct {
	*person1 //指针匿名字段
	id       int
	score    int
}

func main() {
	var stu student1
	//未使用以下new语句进行内存分配先会报如下错误
	//panic: runtime error: invalid memory address or nil pointer dereference
	stu.person1 = new(person1)
	stu.name = "Castle"
	stu.person1.name = "Bob"
	stu.id = 101
	stu.age = 30
	stu.score = 98
	fmt.Println(stu.name)
	fmt.Println(stu.id)
	fmt.Println(stu.score)
}
