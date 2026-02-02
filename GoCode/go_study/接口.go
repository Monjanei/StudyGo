package main

import "fmt"

func dy(val any) {
	fmt.Println(val)
}

type SingInterface interface {
	Sing()
	GetAge()
	//GetSex()
}

// 定义一个Dog对象，有Name、Age两个属性
type Dog struct {
	Name string
	Age  int
}

// 给Dog对象绑定一个Sing方法
func (d Dog) Sing() {
	fmt.Printf("%s：库里库里酷酷库里\n", d.Name)
}

// 给Dog对象绑定一个GetAge方法
func (d Dog) GetAge() {
	fmt.Println(d.Age)
}

// 定义一个Cat对象，有Name、Age、Sex三个属性
type Cat struct {
	Name string
	Age  int
	Sex  string
}

// 给Cat对象绑定一个Sing方法
func (c Cat) Sing() {
	fmt.Printf("%s 叮咚鸡叮咚鸡\n", c.Name)
}

// 给Cat对象绑定一个GetAge方法
func (c Cat) GetAge() {
	fmt.Println(c.Age)
}

// 给Cat对象增加一个GetSex方法
func (c Cat) GetSex() {
	fmt.Println(c.Sex)
}
func common(c SingInterface) {
	switch animals := c.(type) {
	case Dog:
		fmt.Println(animals)
	case Cat:
		fmt.Println(animals)
	default:
		fmt.Println("其他类型")
	}
}
func main() {
	d1 := Dog{Name: "旺财", Age: 18}
	c1 := Cat{Name: "咪咪", Age: 17, Sex: "man"}
	common(d1)
	common(c1)
	dy(d1)
	dy(c1)
	i1 := 1
	s1 := "你好"
	f1 := 0.32
	dy(i1)
	dy(s1)
	dy(f1)
}
