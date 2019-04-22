package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []byte("Golang")
	subslice1 := []byte("go")
	subslice2 := []byte("Go")

	fmt.Println(bytes.Contains(s, subslice1))
	fmt.Println(bytes.Contains(s, subslice2))
}
