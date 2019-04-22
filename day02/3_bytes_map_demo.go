package main

import (
	"bytes"
	"fmt"
)

/**
Map函数的功能是：
	首先将字符转换为utf-8编码的字节序列，然后用mapping函数把
字符中每个Unicode字符转换为对应的字符，最后将映射结果存放到一个
新创建的字节切片中，并返回此新字节切片。

*/
func main() {
	s := []byte("同学们，上午好！")
	m := func(r rune) rune {
		if r == '上' {
			r = '下'
		}
		return r
	}
	fmt.Println(string(s))
	fmt.Println(string(bytes.Map(m, s)))
}
