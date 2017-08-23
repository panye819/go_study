package main

import (
	"fmt"
	// "time"
)

/*
并发concurrency
goroutine是由官方实现的超级“线程池”
每个实例4-5KB的栈内存占用和由于实现机制而大幅减少的创建和销毁开销，是制造Go号称高并发的根本原因

并发并不是并行
并发主要由切换时间片来实现“同时”运行，并行则是直接利用多核实现多线程的运行，但Go可以设置使用核数，
以发挥多核计算记得能力

Goroutine奉行通过通信来共享内存，而不是共享内存来通信。

Channel
channel是goroutine沟通的桥梁，大都是阻塞同步的
通过make创建，close关闭
channel是引用类型
可以使用for range来迭代不断操作channel
可以设置单向或双向通道
可以设置缓存大小，在未被填满前不会发生阻塞

Select
可处理一个或多个channel的发送与接收】
同时有多个可用的channel时按随机顺序处理
可用空的select来阻塞main函数
可设置超时
*/
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!!")
		c <- true
	}()
	<-c
	// time.Sleep(2 * time.Second)
}

// func Go() {
// 	fmt.Println("GO GO GO!!!")
// }
