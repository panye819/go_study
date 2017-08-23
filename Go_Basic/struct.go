package main

import (
	"fmt"
)

/*
GO中的struct与C中的struct非常相似，并且Go没有class
使用type <Name> struct{}定义结构，名称遵循可见性规则
支持指向自身的指针类型成员
支持匿名结构，可用作成员或定义成员变量
匿名结构也可以用于map值
可以使用字面值对结构进行初始化
允许直接通过指针来读写结构成员
相同类型的成员可以直接进行拷贝赋值
支持==与！==比较运算符，但不支持>或<
支持匿名字段，本质上是定义了以某个类型名为名称的字段
嵌入结构作为匿名字段看起来像继承，但不是继承
可以使用匿名字段指针
*/

type person struct {
	Name string
	Age  int
}

func main() {
	a := &person{
		Name: "Tom",
		Age:  20,
	}
	// a.Name = "Joe"
	// a.Age = 19
	fmt.Println(a)
	A(a)
	B(a)
	fmt.Println(a)
}

func A(per *person) {
	per.Age = 13
	fmt.Println("A", per)
}
func B(per *person) {
	per.Age = 15
	fmt.Println("A", per)
}
