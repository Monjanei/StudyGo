package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int)

func shopping1(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	var i int
	if name == "tom" {
		i = 2
	} else {
		i = 1
	}
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Printf("%s 结束购物\n", name)
	moneyChan <- money
	wait.Done()

}

func main() {
	var wait sync.WaitGroup
	wait.Add(3)
	startTime := time.Now()
	go shopping1("alen", 2, &wait)
	go shopping1("tom", 7, &wait)
	go shopping1("jack", 4, &wait)
	go func() {
		wait.Wait()
		close(moneyChan)
	}()
	var moneyList = make([]int, 0)
	for money := range moneyChan {
		fmt.Println(money)
		moneyList = append(moneyList, money)
	}

	fmt.Println("三人共花费：", time.Since(startTime))
	fmt.Println(moneyList)
}
