package main

import "fmt"

func sayHello() {
	fmt.Printf("Hello World\n")
}

func param1(id int) {
	fmt.Println(id)
}

func param2(id, age int) {
	fmt.Println(id, age)
}

func param3(id int, name string) {
	fmt.Println(id, name)
}
func addNum(numberList ...int) {
	var sum int
	for _, item := range numberList {
		sum += item
	}
	fmt.Println(sum)
}

func r1(getName string) string {
	return getName
}

func r2(getList []int) (indexSum, itemSum int) {
	for index, item := range getList {
		indexSum = index + 1
		itemSum += item
	}
	return indexSum, itemSum
}

func r4() (val string, ok bool) {
	if 1 < 2 {
		val = "11"
	}
	return
}
func main() {
	sayHello()
	param1(3)
	param2(3, 4)
	param3(14, "alen")
	addNum(1, 2, 3, 4, 100)
	var name string = r1("hello")
	fmt.Println(name)

	var nameList []int = []int{2, 6, 3, 4, 10}
	fmt.Println(r2(nameList))
	index, sum := r2(nameList)
	fmt.Println(index, sum)
	fmt.Println(r4())
}
