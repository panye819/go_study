package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}

type Student struct {
	person
	id    int
	score int
}

func main() {
	var stu Student = Student{person{"张三丰", 190, "男"}, 001, 100}
	fmt.Println(stu)
}
