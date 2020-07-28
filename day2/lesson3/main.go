package main

// 流程控制

import "fmt"

func main() {
	x := 1
	if x >= 5 {
		fmt.Println("x >= 5")
	} else if x < 5 && x > 0 {
		fmt.Println("x < 5 && x > 0")
	} else {
		fmt.Println("x <= 0")
	}

	if a, b := x+1, x+10; a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a >= b")
	}

	a, b, c, x := 2, 3, 4, 2
	switch x {
	case a, b:
		println("x == a || x == b")
		fallthrough
	case c:
		println("x == c")
	case 3:
		println("x == 3")
	default:
		println("x 不符合任何 case 条件")
	}

	switch y := 5; {
	case y > 5:
		fmt.Println("y > 5")
	case y > 0 && y <= 5:
		println("y > 0 && y <= 5")
	default:
		println("y = ", y)
	}

	for i := 0; i < 5; i++ {
		println("i = ", i)
	}

	j := 0
	for j < 5 {
		j++
		fmt.Println("j = ", j)
	}

	k := 0
	for {
		k++
		if k == 2 {
			continue
		}
		if k > 5 {
			break
		}
		fmt.Println("k = ", k)
	}

	str := "hello"
	for index := range str {
		fmt.Println(str[index])
	}

	l := 0
	for l < 10 {
	start:
		l++
		if l < 5 {
			goto start
		}
		fmt.Println("l = ", l)
	}
}
