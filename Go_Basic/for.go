package main

import (
	"fmt"
)

func main() {
	a := "string"
	b := len(a)
	/*for {
		a++
		if a > 3 {
			break
		}
		fmt.Println(a)
	}*/
	/*	for a <= 3 {
		a++
		fmt.Println(a)
	}*/
	for i := 0; i < b; i++ {
		fmt.Println(string(a[i]))
	}

	fmt.Println("Over!!!")
}
