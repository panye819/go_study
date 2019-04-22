package main

import (
	"bytes"
	"fmt"
)

/**
Repeat函数的功能是把切片复制count次组合成一个新的字节切片并返回。

*/
func main() {
	s := []byte("na")
	count := 2
	fmt.Println("ba" + string(bytes.Repeat(s, count)))
}
