package main

import "fmt"

/**
切片指针作为函数参数
*/

func main() {
	var p *[]int
	fmt.Printf("%p\n", p)
	p = new([]int)
	fmt.Printf("%p\n", p)

	*p = append(*p, 1, 2, 3)
	fmt.Println(*p)
	for i := 0; i < len(*p); i++ {
		fmt.Println((*p)[i])
	}
	fmt.Println("==================")
	for i, v := range *p {
		fmt.Println(i, v)
	}

}
