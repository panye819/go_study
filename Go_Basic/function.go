package main

import (
	"fmt"
)

func main() {
	A(1, 2, 3, 4, 5, 6, 7)
	s1 := []int{1, 2, 3, 4}
	B(s1)
	fmt.Println(s1)
	a := 1
	c(&a)
	fmt.Println(a)
	fmt.Println("-------------")

	a1 := func() {
		fmt.Println("Func A")
	}
	a1()

}
func A(a ...int) {
	fmt.Println(a)
}
func B(s []int) {
	s[0] = 5
	s[1] = 6
	s[2] = 7
	s[3] = 8
	fmt.Println(s)
}
func c(a *int) {
	*a = 2
	fmt.Println(*a)
}
