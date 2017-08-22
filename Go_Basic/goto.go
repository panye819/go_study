package main

import (
	"fmt"
)

/*
跳转语句：
goto,break,continue

三个语法都可以配合标签使用
标签名区分大小写，若不使用会造成编译错误
break与continue配合标签可用于多层循环的跳出
goto是调整执行位置，与其他2个语句配合标签的结果并不相同
*/
func main() {
LABEL1:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i > 3 {
			fmt.Println("reach the limit......")
			break LABEL1
		}
	}
	fmt.Println("end of loop!!")
}
