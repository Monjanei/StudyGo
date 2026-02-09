package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan1 = make(chan int)
var nameChan1 = make(chan string)
var doneChan = make(chan struct{})

func send(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 结束购物\n", name)

	moneyChan1 <- money
	nameChan1 <- name
	wait.Done()

}

func main() {
	var wait sync.WaitGroup
	startTime := time.Now()
	wait.Add(3)
	go send("alen", 10, &wait)
	go send("tom", 10, &wait)
	go send("jack", 10, &wait)

	go func() {
		defer close(moneyChan1)
		defer close(nameChan1)
		defer close(doneChan)
		wait.Wait()
	}()

	var moneyList []int
	var nameList []string

	for {
		select {
		case money := <-moneyChan1:
			moneyList = append(moneyList, money)
		case name := <-nameChan1:
			nameList = append(nameList, name)
		case <-doneChan:
			fmt.Println("购买完成", time.Since(startTime))
			fmt.Println(moneyList)
			fmt.Println(nameList)
			return
		}
	}

}
