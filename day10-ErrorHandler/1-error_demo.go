package main

import (
	"errors"
	"fmt"
)

func test(a, b int) (value int, err error) {
	if b == 0 {
		err = errors.New("0不能作为除数")
		return
	} else {
		value = a / b
		return
	}

}

func main() {
	value, err := test(10, 0)
	//fmt.Println(value)
	if err == nil {
		fmt.Println(value)
	} else {
		fmt.Println(err)
	}
}
