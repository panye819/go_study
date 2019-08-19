package main

import (
	"fmt"
	"testing"
)

/*
使用TestMain作为初始化test，并且使用m.Run()来调用其他tests来完成一些需要初始化
操作的testing，比如数据库连接，文件打开，REST服务登录等

如果没有在TestMian中调用m.Run()，则除了TestMain以外的tests都不会被执行

*/
func Print1to20m() int {
	res := 0
	for i := 0; i <= 20; i++ {
		res += i
	}
	return res
}
func testPrint3(t *testing.T) {
	res := Print1to20m()
	fmt.Println("Hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}
func testPrint4(t *testing.T) {
	res := Print1to20m()
	res++
	if res != 211 {
		t.Errorf("Test print2 Failed")
	}
}
func TestALL2(t *testing.T) {
	t.Run("TestPrint", testPrint3)
	t.Run("TestPrint2", testPrint4)
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print1to20m()
	}

}
