package main

import "fmt"

func main() {
	var s0 []int
	fmt.Println(s0)
	if s0 == nil {
		fmt.Println("ok")
	}
	s0 = append(s0, 1)
	fmt.Println(s0)
	var s1 = []int{1, 2, 3}
	fmt.Println(s1)

	var s2 []int = make([]int, 3)
	fmt.Println(s2)
	var s3 []int = make([]int, 3, 5)
	fmt.Println(s3)

	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	//arr[2] = 3
}
