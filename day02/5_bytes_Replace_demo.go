package main

import (
	"bytes"
	"fmt"
)

/**
Replace函数的功能是返回字节切片的一个副本，并把前n个不重叠的子切片替换复制count次组合成一个新的字节切片并返回,
如果n<0,则不限制替换的数量。
*/
func main() {
	s := []byte("Google")
	old := []byte("o")
	news := []byte("O")
	n := 1
	fmt.Println(string(bytes.Replace(s, old, news, n)))
	fmt.Println(string(bytes.Replace(s, old, news, -1)))
}
