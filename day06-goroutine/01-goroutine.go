package main

import (
	"fmt"
	"time"
)

/**
1、GO主线程：
	一个Go主线程上，可以起多个协程，协程就是轻量级的线程【编译器做优化】
2、Go协程的特点：
	有独立的栈空间
	共享程序堆空间
	调度由用户控制
	协程是轻量级的线程

主线程是一个物理线程池，直接左右在CPU上的，是重量级的，非常耗费CPU资源
协程是从主线程开启的，是轻量级的线程，是逻辑态。对资源消耗相对较小
Golang的协程机制是其重要的特点，可以轻松的开启上万个协程，其他编程语言是基于线程的，
开启过多的线程会耗费大量资源。

**/
func newTask() {
	for {
		fmt.Println("This is a new task!")
		time.Sleep(time.Second)
	}
}
func main() {
	go newTask()
	for {
		fmt.Println("This is main task....")
		time.Sleep(time.Second)
	}

}
