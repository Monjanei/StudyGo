package main

import "fmt"

func main() {
	/*
		nameList:数组名称
		第一个[3]string:长度为3、每个元素的类型为string类型
		第二个[3]string:表示元素的值类型
	*/
	var nameList [3]string = [3]string{"alen", "jack", "tom"}
	var nameList2 [3]int
	nameList2 = [3]int{21, 22, 23}

	fmt.Println(nameList)
	fmt.Println(nameList2)
	fmt.Println("数组第一个值为", nameList[0])
	fmt.Println("数组第一个值为", nameList[2])
	fmt.Println(nameList[len(nameList)-1]) //查看最后一个值

	nameList[0] = "tim"
	fmt.Println(nameList)

}
