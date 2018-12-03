package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Println(s1)
	fmt.Print("\n--------------------------------------\n")
	//slice的扩展
	//slice可以向后扩展，但不可以向前扩展
	s2 := s1[3:5]
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))
	//结果不是[5],而是[5,6],是从arr中取到s1[5]这个元素
	fmt.Println(s2)
}
