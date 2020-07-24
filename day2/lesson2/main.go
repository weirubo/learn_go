package main

import "fmt"

// + - * / %

// ++ --

// = += -= *= /= %=

// > >= < <= == !=

// && || !

// & *

func main() {
	a := 10
	b := 5
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
	fmt.Println("a % b = ", a%b)

	a++
	b--
	fmt.Println("a++:", a)
	fmt.Println("b--:", b)

	aa := 50
	aa += 10
	fmt.Println("aa = ", aa)

	bb := 50
	bb -= 10
	fmt.Println("bb = ", bb)

	cc := 50
	cc *= 10
	fmt.Println("cc = ", cc)

	dd := 50
	dd /= 10
	fmt.Println("dd = ", dd)

	ff := 50
	ff %= 10
	fmt.Println("ff = ", ff)

	if a > 5 && b > 5 {
		fmt.Println("a > 5 且 b > 5")
	}

	if a > 5 || b > 5 {
		fmt.Println("a > 5 或 b > 5")
	}

	if a > 5 && b >= 5 {
		fmt.Println("a > 5 且 b >= 5")
	}

	if a == b {
		fmt.Println("a == b")
	}

	if a != b {
		fmt.Println("a != b")
	}

	boolean := false
	if !boolean {
		fmt.Println("非假既真")
	}

	x := 50
	y := &x
	*y++
	z := *y
	fmt.Println("y = ", y)
	fmt.Println("x = ", x)
	fmt.Println("z = ", z)
}
