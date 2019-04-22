package main

import (
	"bytes"
	"fmt"
)

/**
Runes函数的功能是把字符s转换成utf-8编码的字节序列，并返回对应的Unicode切片。
*/
func main() {
	s := []byte("中华人民共和国")
	r := bytes.Runes(s)
	fmt.Printf("转换前字符%q长度：%d 字节\n", string(s), len(s))
	fmt.Printf("转换后字符%q长度：%d 字节\n", string(r), len(r))

}
