package main

import "fmt"

/**
切片指针作为函数参数
*/

func main() {
	var arr [3]*int
	a := 10
	b := 20
	c := 30

	arr[0] = &a
	arr[1] = &b
	arr[2] = &c

	fmt.Println(arr)
	fmt.Println("==================================")
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)

}
