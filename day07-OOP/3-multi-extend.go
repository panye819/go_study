package main

import "fmt"

type TestA struct {
	name string
	id   int
}

type TestB struct {
	TestA
	sex string
	age int
}

//结构体不能嵌套本结构体，但可以嵌套结构体指针
type TestC struct {
	TestB
	score int
}

func main() {
	var s TestC
	s.name = "Castle"
	s.id = 201
	s.sex = "male"
	s.age = 20
	s.score = 10

	fmt.Println(s)

}
