package main

import "fmt"

func main() {
	s := []int{0, 1}
	for _, v := range s {
		s = append(s, v)
	}
	fmt.Printf("s=%v\n", s)

	// s := []int{0, 1}
	// for i := 0; i < len(s); i++ {
	// 	s = append(s, s[i])
	// }
	// fmt.Printf("s=%v\n", s)
}
