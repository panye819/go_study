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
	s := []byte("Google")
	sep := []byte("o")
	chars := "gle"
	var c byte = 'g'
	var r rune = 'l'
	fmt.Printf("切片%q在%q中的位置索引是：%d\n", sep, s, bytes.Index(s, sep))
	fmt.Printf("切片%q在%q中任一的位置索引是：%d\n", chars, s, bytes.IndexAny(s, chars))
	fmt.Printf("byte型字符%q在%q中任一的位置索引是：%d\n", c, s, bytes.IndexByte(s, c))
	fmt.Printf("切片%q第一个符合f()的字符索引是：%d\n", s, bytes.IndexFunc(s, f))
	fmt.Printf("rune型字符%q在%q中的位置索引是：%d\n", r, s, bytes.IndexRune(s, r))
	fmt.Printf("切片%q在%q中的最后位置索引是：%d\n", sep, s, bytes.LastIndex(s, sep))
	fmt.Printf("切片%q中的任一字符在%q中的最后位置索引是：%d\n", chars, s, bytes.LastIndexAny(s, chars))
	fmt.Printf("切片%q中最后一个符合f()的字符索引是：%d\n", s, bytes.LastIndexFunc(s, f))

}
func f(a rune) bool {
	if a > 'k' {
		return true
	} else {
		return false
	}
}
