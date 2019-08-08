package main

import (
	"fmt"
	"runtime"
)

/*runtime包提供和go运行时环境的互操作，如控制go程的函数，它也包括用于reflect包的低层次类型
信息。
*/
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
