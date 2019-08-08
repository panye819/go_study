package main

import (
	"fmt"
	"runtime"
)

/*runtime包提供和go运行时环境的互操作，如控制go程的函数，它也包括用于reflect包的低层次类型
信息。
*/
func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("CpuNum=", cpuNum)

	//可以自己设置使用几个CPU，Go1.8以后程序默认运行在多个核上了，可以不用设置。
	runtime.GOMAXPROCS(cpuNum - 1)

	fmt.Println("OK!")

}
