package main

import "fmt"

func main() {
	fmt.Println("aaaaaaaaaaaaaaaa")
	//defer 函数调用
	defer fmt.Println("bbbbbbbbbbbbbbbb")
	fmt.Println("cccccccccccccccc")
}
