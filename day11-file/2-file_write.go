package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("1.txt", os.O_RDWR, 0666)
	if err == nil {
		fmt.Println("文件打开成功.....")
	} else {
		fmt.Println("文件打开失败。。。。 err= ", err)
		return
	}
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
		fmt.Println("文件已关闭")
	}()
	s, err := file.WriteString("Hello World")
	if err == nil {
		fmt.Println("写入成功")
		fmt.Println(s)
	} else {
		fmt.Println("写入失败，err: ", err)
	}
}
