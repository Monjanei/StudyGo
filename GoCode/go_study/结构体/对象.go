package main

import "fmt"

type Class struct { //定义一个Class结构体，含有Name一个成员
	Name string
}
type Student struct {
	Name string //Student结构体中也含有一个Name成员
	Age  int
	Class
}

func (s Student) Info() {
	fmt.Printf("%s今年%d岁班级在：%s\n", s.Name, s.Age, s.Class.Name) //在此处同名需要先加上是哪个结构体
}
func (s Student) Study() {
	fmt.Printf("%s 正在学习", s.Name)
}

func main() {
	c1 := Class{Name: "三年级"} //需要先给Class结构体创建一个对象，才能在下面赋值
	s1 := Student{Name: "alen", Age: 14, Class: c1}
	s1.Info()
}
