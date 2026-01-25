package main

import "fmt"

func main() {
	//var sum int
	//for i := 1; i <= 100; i++ {
	//	sum = sum + i
	//	//sum += i
	//}
	//fmt.Println(sum)
	//
	//for i := 1; true; i++ {
	//	fmt.Println(time.Now())
	//	time.Sleep(time.Second)
	//}
	//
	//for true {
	//	fmt.Println(time.Now())
	//	time.Sleep(time.Second)
	//}
	//
	//for {
	//	fmt.Println(time.Now())
	//	time.Sleep(time.Second)
	//}

	//var sum int
	//var i int = 1
	//for i <= 100 {
	//	sum += i
	//	i++
	//}
	//fmt.Println(sum)

	//var sum int
	//var i int = 1
	//for {
	//	sum += i
	//	i++
	//	if i > 100 {
	//		break
	//	}
	//}
	//fmt.Println(sum)
	var List [3]string = [3]string{"hello", "world", "!"}
	for i := 0; i < len(List); i++ {
		fmt.Println(i, List[i])
	}

	var ageList []int = []int{12, 13, 14, 15}
	for i := 0; i < len(ageList); i++ {
		fmt.Println(i, ageList[i])
	}

	var nameList []string = []string{"alen", "hello", "world", "!"}
	for index, item := range nameList {
		fmt.Println(index, item)
	}
	var nameMap map[int]string = map[int]string{1001: "alen", 2003: "tom", 3: "jack"}
	for key, value := range nameMap {
		fmt.Println(key, value)
	}
}
