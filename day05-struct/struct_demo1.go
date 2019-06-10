package main

import "fmt"

type person struct {
	id    int
	name  string
	score int
	sex   string
}

func test(s person) {
	//fmt.Println(s.name)
	//fmt.Println(s.score)
	//fmt.Println(s.sex)
	s.name = "James"
}

func test1(stu []person) {
	//所有的切片都属于地址传递
	stu[0].name = "Bravo"
}
func test2(stu [2]person) {
	//所有的数组都属于值传递
	stu[0].name = "Houston"
}
func main() {
	stu := person{101, "Castle", 90, "male"}
	//fmt.Println(stu)
	/**
	值传递：形参和实参是不同的存储区域，修改不会影响其他的值
	地址传递：形参和实参是相同的存储区域，修改会影响其他的值

	*/
	//结构体作为函数参数是值传递
	test(stu)
	fmt.Println(stu)
	fmt.Println("------------------------------------")
	//结构体切片
	stus := []person{{100, "John", 80, "male"}, {101, "Ball", 90, "male"}}
	stus2 := [2]person{{100, "John", 80, "male"}, {101, "Ball", 90, "male"}}
	stus = append(stus, person{103, "Cole", 99, "male"})
	fmt.Println(stus)
	fmt.Println("test1()")
	test1(stus)
	fmt.Println(stus)
	fmt.Println(stus2)
	fmt.Println("test2()")
	test2(stus2)
	fmt.Println(stus2)
}
