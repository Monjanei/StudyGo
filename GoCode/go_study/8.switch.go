package main

import "fmt"

func main() {
	var age int
	fmt.Printf("请输入你的年龄：")
	fmt.Scan(&age)
	switch {
	case age <= 0:
		fmt.Println("未出生")
	case age <= 18:
		fmt.Println("未成年")
		fallthrough
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")
	}

	var week int
	fmt.Printf("请输入星期几：")
	fmt.Scan(&week)
	switch week {
	case 1, 2, 3, 4:
		fmt.Println("😭")
	case 5:
		fmt.Println("😊")
	case 6, 7:
		fmt.Println("非常😊")
	}
}
