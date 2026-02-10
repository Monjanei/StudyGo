package main

import (
	"fmt"
	"sync"
)

func main() {
	var maps = sync.Map{} //使用sync的方法去定义一个map
	go func() {
		for {
			maps.Store(1, "alen") //使用map.Store添加一个key：1，value：alen
		}
	}()
	go func() {
		for {
			val, _ := maps.Load(1) //使用map.Load读取map，传的值为key，
			fmt.Println(val)
		}
	}()
	select {}
}
