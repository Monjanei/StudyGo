package main

import "fmt"

type Student1 struct {
	Name string
}

func (s Student1) SetName(name string) {
	s.Name = name
}

func (s *Student1) SetName2(name string) {
	s.Name = name
}

func main() {
	s1 := Student1{Name: "alen"}

	s1.SetName("tom")
	fmt.Println(s1.Name)
	s1.SetName2("jack")
	fmt.Println(s1.Name)
}
