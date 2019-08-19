package main

import (
	"fmt"
	"testing"
)

/*
使用t.Run来执行subtests可以做到控制test输出以及test的顺序


*/
func TestPrint3(t *testing.T) {
	t.SkipNow() //跳过测试
	t.Run("a1", func(t *testing.T) { fmt.Println("a1") })
	t.Run("a2", func(t *testing.T) { fmt.Println("a2") })
	t.Run("a3", func(t *testing.T) { fmt.Println("a3") })
}

func TestPrint4(t *testing.T) {
	t.Run("t1", func(t *testing.T) { fmt.Println("t1") })
	t.Run("t3", func(t *testing.T) { fmt.Println("t3") })
	t.Run("t2", func(t *testing.T) { fmt.Println("t2") })
}
