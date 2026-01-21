package main

import (
	"fmt"
	"sort"
)

func main() {
	var nameList []string = []string{"tom", "tim", "alen"}
	fmt.Println(nameList)

	var ageList []int
	ageList = []int{10, 20, 30}
	fmt.Println(ageList)

	var sexList []string
	fmt.Println(sexList)
	//fmt.Println(sexList[0])
	fmt.Println(sexList == nil)

	var nameList1 []string = []string{}
	fmt.Println(nameList1 == nil)

	nameList2 := make([]string, 0)
	fmt.Println(nameList2 == nil)

	ageList2 := make([]int, 3)
	ageList2 = append(ageList2, 5)
	fmt.Println(ageList2)

	ageList3 := make([]int, 0)
	ageList3 = append(ageList3, 7)
	fmt.Println(ageList3)

	array := [5]string{"tom", "tim", "alen", "jack", "jim"}
	slices := array[2:5]
	fmt.Println(slices)

	ints := []int{231, 234, 25, 12}
	fmt.Println(ints) //原封不动输出
	sort.Ints(ints)   //升序
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints))) //降序
	fmt.Println(ints)
}
