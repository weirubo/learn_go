package main

// 基础数据类型

import "fmt"

var a int
var b int32
var c rune
var d byte
var e uint8
var f float32
var g float64
var h bool
var i string

func main() {
	fmt.Printf("a=%d type:%T\n", a, a)
	fmt.Printf("b=%d type:%T\n", b, b)
	fmt.Printf("c=%d type:%T\n", c, c)
	fmt.Printf("d=%d type:%T\n", d, d)
	fmt.Printf("e=%d type:%T\n", e, e)
	fmt.Printf("f=%g type:%T\n", f, f)
	fmt.Printf("g=%g type:%T\n", g, g)
	fmt.Printf("h=%t type:%T\n", h, h)
	fmt.Printf("i=%s type:%T\n", i, i)

	if c == b {
		fmt.Println("rune 类型是 int32 的类型别名，二者可以互换使用")
	}

	if d == e {
		fmt.Println("byte 类型是 uint8 的类型别名，二者可以互换使用")
	}

	i = "hello golang"
	fmt.Printf("字符串 i 的长度：%d\n", len(i))
	fmt.Printf("字符串 i 的第一个字符：%c\n", i[0])
	j := i[0:5]
	fmt.Printf("子字符串 j 的值：%s\n", j)
	k := i[5:]
	fmt.Printf("子字符串 k 的值：%s\n", k)
	l := i[:5]
	fmt.Printf("子字符串 l 的值：%s\n", l)
	m := i[:]
	fmt.Printf("子字符串 m 的值：%s\n", m)
	n := "hello "
	o := "world"
	p := n + o
	fmt.Printf("连接字符串 p 的值：%s\n", p)
}
