package main

import "fmt"

func main() {
	a := make([]int, 0, 5)
	a = append(a, 1)
	b := append(a, 2)
	c := append(a, 3)
	fmt.Printf("v=%v || p=%p\n", a, &a)
	fmt.Printf("v=%v || p=%p\n", b, &b)
	fmt.Printf("v=%v || p=%p\n", c, &c)
}
