package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccccccccccccccccccc")
	//return
	runtime.Goexit()
	fmt.Println("dddddddddddddddddddd")
}
func main() {
	go func() {
		fmt.Println("aaaaaaaaaaaaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbbbbbbbbb")
	}()
	for {
	}
}
