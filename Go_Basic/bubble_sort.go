package main

import (
	"fmt"
)

func main() {
	a := [...]int{14, 3, 65, 1, 99}
	fmt.Println("Before sort the array is :")
	fmt.Println(a)
	var b = len(a)
	// var c int
	for i := 0; i < b; i++ {
		for j := i + 1; j < b; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	fmt.Println("After sort the array is :")
	fmt.Println(a)
}
