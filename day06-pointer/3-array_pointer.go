package main

import "fmt"

/**
数组指针

*/
func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	var p *[5]int
	//将指针变量与数组建立关系
	p = &arr
	//可以通过指针间接操作数组
	fmt.Println(p)
	fmt.Printf("%p\n", p)
	//直接使用指针[下标]操作数组元素
	fmt.Println(p[0])
}
