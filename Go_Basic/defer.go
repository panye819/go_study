package main

import "fmt"

/*
	defer的执行方式类似其他语言中的析构函数，在函数体执行结束后按照调用顺序的相反顺序
	逐个执行，即使函数发生严重错误也会执行
	支持匿名函数的调用
	常用于资源清理、文件关闭、解锁及记录时间等操作
	通过与匿名函数配合可以在return之后修改函数计算结果
	如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时已经获得了拷贝，否则
	则是引用了某个变量的地址

	go没有异常机制，但有panic、recover模式来处理错误
	panic可以在任何地方引发，但recover只有在defer调用的函数中有效
*/
func main() {

	// fmt.Println("a")
	// defer fmt.Println("b")
	// defer fmt.Println("c")
	// for i := 0; i < 3; i++ {
	// 	defer func() {
	// 		fmt.Println(i)
	// 	}()
	A()
	B()
	C()

}
func A() {
	fmt.Println("Func A")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}
func C() {
	fmt.Println("Func C")
}
