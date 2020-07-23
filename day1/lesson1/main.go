package main

import "fmt"

// 定义变量
var totalScore int = 678
var totalScore2 int
var totalScore3 = 345
var a, b, c int
var d, e, f int = 1, 2, 3
var g, h, i = 3.14, "hello", true
var myname = name("lucy")

func main() {
	fmt.Printf("totalScore=%d\n", totalScore)
	fmt.Printf("totalScore2=%d\n", totalScore2)
	fmt.Printf("totalScore3=%d\n", totalScore3)

	fmt.Printf("a=%d b=%d c=%d\n", a, b, c)
	fmt.Printf("d=%d e=%d f=%d\n", d, e, f)
	fmt.Printf("g=%g h=%s i=%t\n", g, h, i)

	fmt.Println(myname)

	totalScore4 := 123
	fmt.Printf("totalScore4=%d\n", totalScore4)

	x, y := 1, 2
	fmt.Printf("x=%d y=%d\n", x, y)

	z := 3
	p := &z
	fmt.Printf("z=%d\n", z)
	fmt.Printf("p=%p\n", p)
	*p = 4
	fmt.Printf("z=%d\n", z)
	fmt.Printf("p=%p\n", p)

	abc := 1
	a1 := &abc
	a2 := &abc
	if a1 == a2 {
		fmt.Println("a1 和 a2 相等，a1 和 a2 指向同一变量")
	}

	aa := 2
	bb := 2
	aa1 := &aa
	aa2 := &bb
	if aa1 == aa2 {
		fmt.Println("aa1 和 aa2 相等")
	} else {
		fmt.Println("aa1 和 aa2 不相等")
	}
}

func name(name string) string {
	var str string = "My name is " + name
	return str
}
