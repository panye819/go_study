package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("Usage: XXXX file.")
	}
	filename := list[1]

	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("error is ", err)
		return
	}
	fmt.Println("Name = ", info.Name())
	fmt.Println("Size = ", info.Size())
}
