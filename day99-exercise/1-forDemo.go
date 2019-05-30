package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum: ", sum)
	sum2 := 0
	for i := 0; i <= 100; i += 2 {
		if i%2 == 0 {
			sum2 += i
		}
	}
	fmt.Println("sum2: ", sum2)
	sum3 := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			sum3 += i
		}
	}
	fmt.Println("sum3: ", sum3)
}
