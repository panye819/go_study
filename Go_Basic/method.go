package main

import (
	"fmt"
)

/*
method方法
Go中虽然没有class，但依旧有method
通过显示说明receiver来实现与某个类型的组合
只能为同一个包中的类型定义方法
receiver可以是类型的值或者指针
不存在方法重载
可以使用值或指针来调用方法，编译器会自动完成转换
从某种意义上来说，方法是函数的与发烫，因为receiver其实就是方法所接收的第一个参数
如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
类型别名不会拥有底层类型所附带的方法
方法可以调用结构中的非公开字段
*/

type A struct {
	Name string
}
type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)
}
func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A")
}
func (b B) Print() {
	b.Name = "BB"
	fmt.Println("B")
}
