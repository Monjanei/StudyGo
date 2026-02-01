package main

import "fmt"

// 自定义数据类型
type MyCode int

// 类型别名
type MyAliasCode = int

func (mc MyCode) Hello() {}

//func (mc MyAliasCode) Hello() {}

const (
	mycode      MyCode      = 1
	myAliasCode MyAliasCode = 2
)

func main() {
	fmt.Printf("%v,%T\n", mycode, mycode)
	fmt.Printf("%v,%T\n", myAliasCode, myAliasCode)
	var age int
	fmt.Scan(&age) //接受一个age的值
	//自定义类型的变量类型为MyCode，无法与int做比较
	//if age == mycode {
	//	fmt.Printf("no")
	//}
	//可以将mycode变量的类型转为int型，或将age的类型转换为MyCode类型
	if age == int(mycode) { //或者是if MyCode(age) == mycode
		fmt.Printf("yes")
	}
	//类型别名的变量类型还是int，所以可以直接和int对比较
	if age == myAliasCode {
		fmt.Printf("ok")
	}

}
