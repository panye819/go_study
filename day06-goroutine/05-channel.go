package main

import (
	"fmt"
	"time"
)

/*
	1、channel的本质是一个数据结构-队列
	2、数据是先进先出（First in First Out FIFO）
	3、线程安全，多goroutine访问时，不需要枷锁，channel本身是线程安全的
	4、channel是有类型的，一个string的channel只能存放string类型数据
*/
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
