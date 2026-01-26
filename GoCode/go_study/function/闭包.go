package main

import (
	"fmt"
	"time"
)

//func addAwait(sec int) func(numberList ...int) int {
//	return func(numberList ...int) int {
//		time.Sleep(time.Duration(sec) * time.Second)
//		var sum int
//		for _, i := range numberList {
//			sum += i
//		}
//		return sum
//	}
//}

func addAwait(sec int) func(nameList ...int) int {
	return func(nameList ...int) int {
		time.Sleep(time.Duration(sec) * time.Second)
		var sum int
		for _, i := range nameList {
			sum += i
		}
		return sum
	}

}
func main() {
	t1 := time.Now()
	sum := addAwait(3)(1, 2, 3)
	t2 := time.Since(t1)
	fmt.Println(t2)
	fmt.Println(sum)
}
