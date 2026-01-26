package main

import "fmt"

func main() {
	var addSum = func() string {
		return "alen1"
	}
	fmt.Println(addSum())

	var SetName = func(name string) {
		fmt.Println(name)
	}
	SetName("alen2")
}
