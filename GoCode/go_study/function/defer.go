package main

import "fmt"

func main() {
	var name string = "alen"
	defer func() {
		fmt.Println(name)
	}()
	name = "tom"
	return
}
