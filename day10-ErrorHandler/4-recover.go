package main

import "fmt"

func demo(i int) {
	var arr [10]int
	//通过匿名函数和recover()进行错误拦截
	defer func() {
		//recover()可以从panic异常中重新获取控制权，程序可以继续运行
		recover()
	}()
	//如果传递超过数组下标值，会导致数组下标越界，导致程序崩溃
	//recover()一定要在错误出现之前进行拦截
	arr[i] = 100
	fmt.Println(arr)
}

func demo1() {
	fmt.Println("程序继续运行。。。。。。")
}
func main() {
	demo(10)
	demo1()
}
