package main

import (
	"fmt"
)

//panic函数返回的是让程序崩溃的错误
func TestA() {
	fmt.Println("func TestA")
}
func TestB() {
	fmt.Println("func TestB(): panic")
}

//在程序中直接调用panic后，程序会终止执行
func TestC() {
	panic("Panic")
}
func main() {
	TestA()
	TestC()
	TestB()

}
