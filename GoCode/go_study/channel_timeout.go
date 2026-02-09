package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{})

func event() {
	fmt.Println("event start")
	time.Sleep(3 * time.Second)
	fmt.Println("event done")
	close(done)
}
func main() {
	go event()
	select {
	case <-done:
		fmt.Println("event执行完成")
	case <-time.After(2 * time.Second):
		fmt.Println("event执行超时")
	}
}
