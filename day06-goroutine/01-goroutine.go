package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("This is a new task!")
		time.Sleep(time.Second)
	}
}
func main() {
	go newTask()
	for {
		fmt.Println("This is main task....")
		time.Sleep(time.Second)
	}

}
