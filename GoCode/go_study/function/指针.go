package main

import "fmt"

//func Copy(fname *string) {
//	fmt.Println(fname)
//	*fname = "tom"
//}
//
//func main() {
//
//	var name string = "alen"
//	Copy(&name)
//	fmt.Printf("%p\n", &name)
//	fmt.Println(name)
//}

func copy(fname *string) { //定义一个函数，他的参数为引用类型 *表示，表示他要接受的值是一个内存地址
	fmt.Println(fname) //此时fname存的值是name的内存地址
	*fname = "tom"     //修改 fname 指针指向的内存中的值，因为fname和name指向同一个内存地址的值，所以修改这个值为tom，name的值也会修改为tom
}

func main() {
	var name string = "alen"
	fmt.Printf("%p,%v\n", &name, name) //将name的内存地址打印
	copy(&name)                        //将变量name的内存地址的值传给函数
	fmt.Printf("%p,%v\n", &name, name) //将name的内存地址打印

}
