package main

import "fmt"

func main() {
	f := fmt.Println
	f(a())
	f(b())
}

func a() int {
	i := 0
	defer func() {
		i += 1
		fmt.Println("a defer:", i)
	}()
	return i
}

func b() (i int) {
	i = 0
	defer func() {
		i += 1
		fmt.Println("b defer:", i)
	}()
	return i
}

func c() *int {
	i := 0
	return &i
}

func d() (i *int) {
	return
}
