package main

import (
	"GoCode/go_study/book"
	"fmt"
)

func hello() {
	//函数内的局部变量作用域为当前的函数内
	age1 := 13
	fmt.Println("这是hello函数内的 ", age1, sex)
}

// 这是全局变量，全局变量定义可以不使用；必须使用var
var sex string = "男"

// 这是同时定义多个变量
var (
	b1 string = "b1"
	b2 string = "b2"
)

// 常量变量定义时必须赋值，且无法修改
const version string = "1.0.0"

func main() {
	var name1 string          //先声明
	name1 = "这是使用先定义再赋值再使用的值" //再赋值
	fmt.Println(name1)

	var name2 string = "声明并赋值"
	fmt.Println(name2)

	name3 := "简短声明使用:="
	fmt.Println(name3)

	var name4 = "使用var但是未声明类型"
	fmt.Println(name4)

	//局部变量，局部变量定义必须使用
	age1 := 12
	fmt.Println(age1)
	hello()

	//定义多个变量
	var a1, a2 = 1, 2
	fmt.Println(a1, a2)

	fmt.Println("b1的值为", b1, "b2的值为", b2)

	//夸包使用变量
	fmt.Println(book.Version)
	//fmt.Println(book.Name)
	fmt.Println(book.EditorBy)
}
