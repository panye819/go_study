package main

import "fmt"

/**
切片指针

*/
func main() {
	var slice []int = []int{1, 2, 3, 4, 5}

	var p *[]int
	//将指针变量与数组建立关系
	p = &slice
	//可以通过指针间接操作数组
	fmt.Println(p)

	//切片名本身就是一个地址
	fmt.Printf("%p\n", slice)
	fmt.Printf("%p\n", p)
	//直接使用指针[下标]操作数组元素
	//fmt.Println(p[0])
}
