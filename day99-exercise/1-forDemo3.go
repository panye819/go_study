package main

import "fmt"

func main() {

	for i := 0; i < 6; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
