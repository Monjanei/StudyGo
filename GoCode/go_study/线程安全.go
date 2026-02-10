package main

import (
	"fmt"
	"sync"
)

var sum int
var wait sync.WaitGroup
var lock sync.Mutex

func add() {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		sum++
	}
	lock.Unlock()
	wait.Done()
}
func sub() {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		sum--
	}
	lock.Unlock()
	wait.Done()
}

func main() {
	wait.Add(2)
	go add()
	go sub()
	wait.Wait()
	fmt.Println(sum)
}
