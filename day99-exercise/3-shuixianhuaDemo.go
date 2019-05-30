package main

import "fmt"

func main() {
	i := 100
	for ; i < 1000; i++ {
		//百位
		a := i / 100
		//十位
		b := (i / 10) % 10
		//个位
		c := i % 10

		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}

}
