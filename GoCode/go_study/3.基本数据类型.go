package main

import "fmt"

func main() {
	var u8 uint8 = 255 //存无符号位的正整数
	// 0 0 0 0 0 0 0 0 = 2^8-1 = 255

	var a1 int8 = -127
	// 0   0 0 0 0 0 0 0
	//第一位表示符号

	fmt.Println(u8, a1)

	var f1 float32 = 1.1
	var f2 float64 = 1.2
	fmt.Println(f1, f2)

	var c1 byte = 'a'
	var c2 int8 = 97
	fmt.Printf("%d,%c\n", c1, c2)

	var c3 rune = '傻'
	var c4 int32 = 20667
	fmt.Printf("%d,%c\n", c3, c4)

	fmt.Println("制表符1\t制表符2")
	fmt.Println("\"转义\"")
	fmt.Println(`hello
word
,,,

`)

	var b1 bool = true
	var b2 bool = false
	fmt.Println(b1, b2)

	var q int
	var w float32
	var e string
	var r rune
	var t bool
	fmt.Println(w, q, w, e, r, t)

}
