package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("1.txt")
	if err != nil {
		fmt.Println("Error！！！！")
		return
	} else {
		fmt.Println(fp)
	}

}
