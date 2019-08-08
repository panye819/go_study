package main

import (
	"fmt"
	"time"
)

/*
	需求：
		计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中，最后显示出来。
		要求使用goroutine完成
*/

//思路：
//	1、编写一个函数，来计算各个数的阶乘，并放入到map中
//	2、我们启动的协程多个，统计的将结果放入到map中
//	3、map应该做出一个全局的

var (
	myMap = make(map[int]int, 10)
)

func test2(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}
func main() {
	for i := 1; i <= 200; i++ {
		go test2(i)
	}
	//休眠10秒钟
	time.Sleep(time.Second * 10)
	//输出结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
		time.Sleep(time.Second * 1)
	}
}
