package main

import "fmt"

func main() {
	fmt.Println("请输入你要进行的操作")
	fmt.Println("1.登录")
	fmt.Println("2.注册")
	fmt.Println("3.个人中心")
	//定义一个变量接受用户输入的值
	var index int
	fmt.Scan(&index)

	//定义一个map，key为选项==index，value为所对应的函数操作
	var menu = map[int]func(){
		1: login,
		2: register,
		3: userCenter,
	}
	//将menu这个map的value赋值给fun
	fun, ok := menu[index]
	//如果这个key不存在ok为false，如果值存在，ok为true
	if ok {
		fun()
	}

}

func login() {
	fmt.Println("登录")
}
func register() {
	fmt.Println("注册")
}
func userCenter() {
	fmt.Println("用户中心")
}
