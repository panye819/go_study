package main

import (
	"fmt"
)

func main() {
	/*
		切片Slice

		其本身并不是数组，它指向底层的数组
		作为变长数组的替代方案，可以关联底层数组的局部或全部
		为引用类型
		可以直接创建或从底层数组获取生成
		使用len()获取元素个数，cap()获取容量
		一般使用make()创建
		如果多个slice指向相同底层数组，其中一个值得改变会影响全部

		make([]T,len,cap)
		其中cap可以省略，则和len的值相同
		len表示存数的元素个数，cap表示容量
	*/
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a1 := a[0:4]
	a2 := a[4:]
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println("...")
	a2[0] = 0
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println("...")
	fmt.Println(a)
	fmt.Println(".............")
	a3 := make([]int, 3)
	fmt.Println(a3)
}
