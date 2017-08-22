package main

import "fmt"

const (
	/*
	   1、在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
	   2、使用相同的表达式不代表具有相同的值
	   3、iota是常量的计数器，从0开始，组中每定义一个常量自动递增1
	   4、通过初始化规则与iota可以打到枚举的效果
	   5、每遇到一个const关键字，iota就会重置为0
	*/

	a = 'A'
	b
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
