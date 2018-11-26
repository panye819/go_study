package main

import "fmt"

//func funcc(c int) {

//	fmt.Println("c = ", c)
//}

//func funcb(b int) {
//	funcc(b - 1)
//	fmt.Println("b = ", b)
//}

func test(a int) {
	if a == 1 {
		fmt.Println("a = ", a)
		return
	}
	test(a - 1)
	fmt.Println("a = ", a)
}

func main() {
	test(3)
	fmt.Println("main")
}
