package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var a int64
	//a = 1
	//fmt.Printf("%T\t%d\n", a, a)
	//var b int8
	//b = int8(a)
	//fmt.Printf("%T\t%d\n", b, b)

	//var a float64
	//a = 3.1415926
	//fmt.Printf("%T\t%f\n", a, a)
	//var b float32
	//b = float32(a)
	//fmt.Printf("%T\t%f\n", b, b)
	//var c int64
	//c = int64(b)
	//fmt.Printf("%T\t%d\n", c, c)

	//var a string
	//a = "golang"
	//fmt.Printf("%T\t%s\n", a, a)
	//var b []byte
	//b = []byte(a)
	//fmt.Printf("%T\t%d\n", b, b)

	//var a bool
	//a = true
	//fmt.Printf("%T\t%t\n", a, a)
	//var b string
	//b = strconv.FormatBool(a)
	//fmt.Printf("%T\t%s\n", b, b)

	//var id interface{}
	//id = 1 // 参数 id 接收到的值为整型
	//fmt.Printf("%T\t%v\n", id, id)
	//// 需要使用字符串类型的变量 id 赋值给字符串类型的变量 uid
	//var uid string
	//value, ok := id.(string)
	//if ok {
	//	uid = value
	//}
	//fmt.Printf("%T\t%v\n", uid, uid)

	var id interface{}
	id = 1 // 参数 id 接收到的值为整型
	fmt.Printf("0-%T\t%v\n", id, id)
	// 需要使用字符串类型的变量 id 赋值给字符串类型的变量 uid
	var uid string
	switch val := id.(type) {
	case string:
		uid = val
		fmt.Printf("1-%T\t%v\n", uid, uid)
	case int:
		uid = strconv.Itoa(val)
		fmt.Printf("2-%T\t%v\n", uid, uid)
	default:
		fmt.Printf("3-%T\t%v\n", uid, uid)
	}
}
