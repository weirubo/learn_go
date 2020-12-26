package main

// 常量

import "fmt"

const name string = "golang"

// PI π
const PI = 3.14

const (
	x = 1
	y = "hello"
	z = true
)

const (
	a = 1
	b
	c = 2
	d
)

// Weekday
const (
	Sunday int = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	aa int = iota
	bb
	cc
	dd = 10
	ee
	ff = iota
	gg
	hh
)

const username, age = "lucy", 17
const user = username

func main() {
	fmt.Printf("name = %s type:%T\n", name, name)
	fmt.Printf("PI = %g type:%T\n", PI, PI)
	fmt.Printf("x = %d type:%T\n", x, x)
	fmt.Printf("y = %s type:%T\n", y, y)
	fmt.Printf("z = %t type:%T\n", z, z)
	fmt.Printf("a = %d type:%T\n", a, a)
	fmt.Printf("b = %d type:%T\n", b, b)
	fmt.Printf("c = %d type:%T\n", c, c)
	fmt.Printf("d = %d type:%T\n", d, d)
	fmt.Printf("Sunday = %d type:%T\n", Sunday, Sunday)
	fmt.Printf("Monday = %d type:%T\n", Monday, Monday)
	fmt.Printf("Tuesday = %d type:%T\n", Tuesday, Tuesday)
	fmt.Printf("Wednesday = %d type:%T\n", Wednesday, Wednesday)
	fmt.Printf("Thursday = %d type:%T\n", Thursday, Tuesday)
	fmt.Printf("Friday = %d type:%T\n", Friday, Friday)
	fmt.Printf("Saturday = %d type:%T\n", Saturday, Saturday)
	fmt.Printf("aa = %d type:%T\n", aa, aa)
	fmt.Printf("bb = %d type:%T\n", bb, bb)
	fmt.Printf("cc = %d type:%T\n", cc, cc)
	fmt.Printf("dd = %d type:%T\n", dd, dd)
	fmt.Printf("ee = %d type:%T\n", ee, ee)
	fmt.Printf("ff= %d type:%T\n", ff, ff)
	fmt.Printf("gg = %d type:%T\n", gg, gg)
	fmt.Printf("hh = %d type:%T\n", hh, hh)
	fmt.Printf("username = %s, age = %d\n", username, age)
	fmt.Printf("user = %s\n", user)
}
