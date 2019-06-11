package main

import "fmt"

/**
切片指针作为函数参数
*/

func test(s []int) {
	s = append(s, 4, 5, 6)
}

//切片指针作为函数参数是地址传递，形参可以改变实参的值
func test2(s *[]int) {
	*s = append(*s, 4, 5, 6)
}
func main() {
	s := []int{1, 2, 3}
	test(s)
	fmt.Println("1,2,3")
	test(s)
	fmt.Println(s)
	fmt.Println("------------------------------")
	test2(&s)
	fmt.Println(s)

}
