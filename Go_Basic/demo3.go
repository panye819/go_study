package main

import "fmt"

func main(){
    //数值型默认值是0
    var a float32
    fmt.Println(a)
    //bool值得默认值是false
    var b bool
    fmt.Println(b)
    //字符串的默认值是空字符串
    var s string
    fmt.Println(s)
    var a1 int = 65
    b1 := string(a1)
    fmt.Println(b1)
}

