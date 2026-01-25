package main

import "fmt"

func main() {
	//打印1-10
	for i := 1; i <= 10; i++ {
		if i == 5 {
			//break
			continue
		}
		fmt.Printf("第%d次循环\t", i)
	}
}
