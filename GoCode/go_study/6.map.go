package main

import "fmt"

func main() {
	var userMap map[int]string = map[int]string{
		1: "alen",
		2: "jack",
		3: "",
	}
	fmt.Println(userMap) //打印map

	userMap[4] = "tom"
	fmt.Println(userMap)
	userMap[1] = "hello"
	fmt.Println(userMap)

	fmt.Printf("%#v\n", userMap[5])
	value, ok := userMap[5]
	fmt.Println(value, ok)

	value1, ok1 := userMap[3]
	fmt.Println(value1, ok1)

	fmt.Println(userMap)
	delete(userMap, 3)
	fmt.Println(userMap)

	var ageMap map[string]int = make(map[string]int)
	ageMap["alen"] = 18
	fmt.Println(ageMap)

	var userList []int = make([]int, 3)
	fmt.Println(userList)

}
