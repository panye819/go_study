package main

import "fmt"

type washer interface {
	wash()
	dry()
}

type Haier struct {
	name  string
	price float64
	mode  string
}

func (h Haier) wash() {
	fmt.Println("海尔洗衣机能洗衣服。。。")
}

func (h Haier) dry() {
	fmt.Println("海尔洗衣机自带甩干。。。")
}
func main() {
	var a washer
	h1 := Haier{
		name:  "小神童",
		price: 998.98,
		mode:  "全自动",
	}
	fmt.Printf("%T\n", h1)
	a = h1
	fmt.Println(a)
	fmt.Println("========================")
	a.wash()
	a.dry()
}
