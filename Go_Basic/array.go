package main

import (
	"fmt"
)

func main() {
	/*	var (
		a [2]int
		b [1]string
		c [2][3]int
	)*/
	d := [2][3]int{
		{1, 2, 3},
		{4, 5, 6}}
	/*	fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)*/
	fmt.Println(d)
}
