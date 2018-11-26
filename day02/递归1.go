package main

import "fmt"

func test() int {
	i := 1
	sum := 0
	for i=1;i<100;i++ {
		sum += i
	}
	return sum
}

func main (){
	result := test()
	fmt.Println("The sum is ", result)
}