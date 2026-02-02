package main

import (
	"fmt"
	"sync"
	"time"
)

func shopping(name string, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	var i int
	if name == "tom" {
		i = 2
	} else {
		i = 1
	}
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Printf("%s 结束购物\n", name)
	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	wait.Add(3)
	startTime := time.Now()
	go shopping("alen", &wait)
	go shopping("tom", &wait)
	go shopping("jack", &wait)
	wait.Wait()
	fmt.Println("三人共花费：", time.Since(startTime))
}
