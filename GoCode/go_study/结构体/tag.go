package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age,omitempty"`
	Passwd string `json:"-"`
}

func (u User) Info() {
	fmt.Println(u.Name, u.Age, u.Passwd)
}

func main() {
	user := User{Name: "alen", Age: 1, Passwd: "123456"}
	user.Info()
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))
}
