package main

import "fmt"

func main() {
	//中断式
	//var age int
	//fmt.Println("请输入你的年龄：")
	//fmt.Scan(&age)
	//if age <= 0 {
	//	fmt.Println("未出生")
	//	return
	//}
	//if age <= 18 {
	//	fmt.Println("未成年")
	//	return
	//}
	//if age <= 35 {
	//	fmt.Println("中年")
	//	return
	//}

	//嵌套式
	//var age int
	//fmt.Println("请输入你的年龄")
	//fmt.Scan(&age)
	//
	//if age <= 18 {
	//	if age <= 0 {
	//		fmt.Println("未出生")
	//	} else {
	//		fmt.Println("未成年")
	//	}
	//} else {
	//	if age <= 35 {
	//		fmt.Println("青年")
	//	} else {
	//		fmt.Println("中年")
	//	}
	//}

	//多条件判断式
	fmt.Println("请输入你的年龄")
	var age int
	fmt.Scan(&age)

	if age <= 0 {
		fmt.Println("未出生")
	}
	if age > 0 && age <= 18 {
		fmt.Println("未成年")
	}
	if age > 18 && age < 35 {
		fmt.Println("青年")
	}
	if age >= 35 {
		fmt.Println("中年")
	}
}
