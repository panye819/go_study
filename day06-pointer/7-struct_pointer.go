package main

import "fmt"

/**
切片指针作为函数参数
*/

type Student struct {
	id   int
	name string
	age  int
	sex  string
}

func main() {
	var stu Student = Student{101, "哆啦A梦", 100, "男"}
	var p *Student
	p = &stu
	(*p).name = "超人"
	p.age = 30
	fmt.Println(stu)
}
