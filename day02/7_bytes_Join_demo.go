package main

import (
	"bytes"
	"fmt"
)

/**
Join函数的功能是用字节切片sep把s中的每个字节切片连接成一个
字节切片并返回。
*/
func main() {
	s := [][]byte{
		[]byte("你好"),
		[]byte("世界"),
	}
	sep := []byte(",")
	fmt.Println(string(bytes.Join(s, sep)))
	sep1 := []byte("&")
	fmt.Println(string(bytes.Join(s, sep1)))
	sep2 := []byte("@")
	fmt.Println(string(bytes.Join(s, sep2)))
}
