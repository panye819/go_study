package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func(s string) {
		for i := 0; i < 5; i++ {
			fmt.Println(s)
		}
	}("world")
	for i := 0; i < 2; i++ {
		//让出时间片，让其他协程先执行,等其他协程执行完毕，再执行.
		runtime.Gosched()
		fmt.Println("Hello")
	}
}
