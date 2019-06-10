package main

import "fmt"

func main() {
	arr := [5]int{11, 2, 333, 44, 5}
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		//fmt.Println(arr[i])
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("After ")
	fmt.Println(arr)
}
