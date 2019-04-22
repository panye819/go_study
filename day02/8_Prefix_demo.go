package main

import (
	"bytes"
	"fmt"
)

/**
HasPrefix函数和HasSuffix函数的功能可以检查字节切片的前后缀，
并返回一个布尔型的检查结果。
*/
func main() {
	s := []byte("recover")
	prefix := []byte("re")
	suffix := []byte("ed")
	fmt.Println(bytes.HasPrefix(s, prefix))
	fmt.Println(bytes.HasSuffix(s, suffix))
}
