package main

import "fmt"

/*
	空接口:
		可以接收所有值
*/

func showType(a interface{}) {
	fmt.Printf("type:%T\n", a)
}
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string , value is %v\n", v)
	case int:
		fmt.Printf("x is a int , value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool , value is %v\n", v)
	default:
		fmt.Println("unsupport type!!")
	}
}
func main() {
	var x interface{}
	x = 100
	fmt.Println(x)
	showType(x)
	x = false
	fmt.Println(x)
	showType(x)
	x = struct {
		name string
	}{name: "花花"}
	fmt.Println(x)
	showType(x)
	//switch
	justifyType(100)
	justifyType(struct {
	}{})
	s := "haha"
	justifyType(&s)

}
