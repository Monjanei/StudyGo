package main

import "fmt"

const (
	SucessCode       Code = 0
	CodeFailCode     Code = 1
	NetworkFaildCode Code = 2
)

type Code int

func (c Code) getMsg() string {
	switch c {
	case 0:
		return "成功"
	case 1:
		return "代码错误"
	case 2:
		return "网络错误"
	}
	return ""
}
func webServer(errorcode int) (code Code, msg string) {
	if errorcode == 2 {
		return NetworkFaildCode, CodeFailCode.getMsg()
	}
	if errorcode == 1 {
		return CodeFailCode, CodeFailCode.getMsg()
	}
	if errorcode == 0 {
		return SucessCode, SucessCode.getMsg()
	}
	return
}
func main() {
	fmt.Println(webServer(2))
}
