package main

import "fmt"

func test1(a int) {
	a++
	fmt.Println(a)
}
func test2() func() int {
	var a int

	return func() int {
		a++
		return a
	}
}

func main() {
	a := 0
	for i := 0; i < 10; i++ {
		test1(a)
	}
	fmt.Println("-------------------------------")
	f := test2()
	//fmt.Printf("%T",f)
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
