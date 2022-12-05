package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	str := "hello"
	reverseStr := stringutil.Reverse(str)
	fmt.Println(reverseStr)
	toUpperStr := stringutil.ToUpper(str)
	fmt.Println(toUpperStr)
}
