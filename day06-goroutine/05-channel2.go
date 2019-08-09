package main

import "fmt"

func main() {
	//创建一个可以存放三个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2、看看intChan是什么

	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	fmt.Println()
	//3、向管道中写入数据
	intChan <- 10
	num := 211
	intChan <- num

	//4、看看管道的长度和容量
	fmt.Printf("intChan的长度=%d intChan本身的容量=%d\n", len(intChan), cap(intChan))

}
