package main

import "fmt"

func main() {
	fmt.Println("println 自带换行")
	fmt.Print("没有换行")
	name := "alen"
	fmt.Printf("%s格式化输出%s\n", name)       //该输出语句会有一个missing,因为第二个%s没有值
	fmt.Printf("%s格式化输出%s\n", name, "你好") //"你好"这个字符串传给了第二个%s
	fmt.Printf("这是整数类型%d \n", 3)
	fmt.Printf("这是浮点数类型%.4f \n", 3.1415926)
	fmt.Printf("这是字符串类型%s \n", "hello")
	fmt.Printf("打印变量的类型%T %T \n", "hello", 1)
	fmt.Printf("打印值：%v \n", "")
	fmt.Printf("打印值：%#v \n", "")

	var name1 string
	fmt.Print("请输入你的名称：")
	fmt.Scan(&name1)

	var age int
	fmt.Print("请输入你的年龄：")
	n, err := fmt.Scan(&age)
	fmt.Println(n, err, age)

	var age2 int
	fmt.Print("请输入你的虚岁")
	n1, err1 := fmt.Scan(&age2)
	fmt.Println(n1, err1, age2)
}
