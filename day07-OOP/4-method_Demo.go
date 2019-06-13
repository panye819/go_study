package main

import "fmt"

type stu struct {
	name string
	age  int
	sex  string
}

type Stu struct {
	name string
	age  int
	sex  string
}

func (s Stu) PrintInfo() {
	fmt.Println(s.name)
	fmt.Println(s.age)
	fmt.Println(s.sex)
}
func main() {
	var s = Stu{"刘姥姥", 54, "女"}
	s.PrintInfo()

}
