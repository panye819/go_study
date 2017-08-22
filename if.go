package main

import (
	"fmt"
)

func main() {
	a := 10
	if a := 2; a > 1 {
		fmt.Println("a is not smaller than 1")
	} else {
		fmt.Println("a is  smaller than 1")
	}
	fmt.Println(a)
}
