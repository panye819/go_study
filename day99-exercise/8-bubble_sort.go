package main

import "fmt"

func main() {
	var a = 10
	//&取地址运算符，通过运算符可以打印变量所在的内存地址
	fmt.Printf("%p\n", &a)
	//在基本类型前加*表示指针类型
	var p *int
	p = &a
	fmt.Println(*p)
	a = 20
	//*称为取值运算符
	fmt.Println(*p)

}
