package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")

}
func person1() {
	Printer("hello")
	ch <- 666

}

func person2() {
	<-ch
	Printer("world")

}
func main() {
	go person1()
	go person2()
	for {
	}
}
