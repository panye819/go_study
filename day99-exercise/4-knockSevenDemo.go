package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for ; i < 101; i++ {
		if i%7 == 0 || i%10 == 7 || i/10 == 7 {
			fmt.Println("敲桌子")
			time.Sleep(time.Duration(1) * time.Second)
		} else {
			fmt.Println(i)
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}
