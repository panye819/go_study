package main

import (
	"bytes"
	"fmt"
	"unicode"
)

/**

 */
func main() {
	s := []byte("Hello,world!")
	s1 := []byte("HELLO,WORLD!")
	s2 := []byte("hello,world!")

	fmt.Println(string(bytes.Title(s)))
	fmt.Println(string(bytes.ToTitle(s)))
	fmt.Println(string(bytes.ToTitleSpecial(unicode.AzeriCase, s)))
	fmt.Println("-------------------------------------------")
	fmt.Println(string(bytes.ToLower(s1)))
	fmt.Println(string(bytes.ToLowerSpecial(unicode.AzeriCase, s)))
	fmt.Println("-------------------------------------------")
	fmt.Println(string(bytes.ToUpper(s2)))
	fmt.Println(string(bytes.ToUpperSpecial(unicode.AzeriCase, s2)))
}
