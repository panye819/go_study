package main

import (
	"fmt"
)

/*
选择语句switch

可以使用任何类型或表达式作为条件语句
不需要写break，一旦条件符合自动终止
如希望之星下一个case，需要使用fallthrough语句
支持一个初始化表达式（可以是并行方式），右侧需跟分号
左大括号必须和条件语句在同一行
*/

func main() {
	// a := 3
	/*	switch a {
		case 0:
			fmt.Println("a=0")
		case 1:
			fmt.Println("a=1")

		}*/
	/*
		switch {
		case a >= 0:
			fmt.Println("a=0")
			fallthrough
		case a >= 1:
			fmt.Println("a=1")
		}*/
	switch a := 1; {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	}
	// fmt.Println(a)
}
