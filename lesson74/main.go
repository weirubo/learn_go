package main

import (
	"fmt"
	"slices"
)

func main() {
	//s1 := []int{1, 2, 3}
	//s2 := []int{4, 5, 6}
	//s3 := []int{7, 8, 9}
	//s := slices.Concat(s1, s2, s3)
	////var s []int
	////s = append(s, s1...)
	////s = append(s, s2...)
	////s = append(s, s3...)
	//fmt.Println(s)

	//s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s2 := slices.Delete(s1, 2, 5)

	//s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s2 := slices.DeleteFunc(s1, func(i int) bool {
	//	return i%2 != 0
	//})

	//s1 := []int{1, 2, 2, 3, 4, 3, 5, 4}
	//s2 := slices.Compact(s1)
	//fmt.Printf("len=%d\tcap=%d\tval=%d\t\n", len(s1), cap(s1), s1)
	//fmt.Printf("len=%d\tcap=%d\tval=%d\t\n", len(s2), cap(s2), s2)

	//s1 := []string{"go", "php", "php", "java", "python"}
	//s2 := slices.CompactFunc(s1, strings.EqualFold)

	//s1 := []string{"php", "java", "go", "python"}
	//s2 := slices.Replace(s1, 1, 3, "rust")

	s1 := []string{"Go", "PHP", "Java", "Rust"}
	s2 := slices.Insert(s1, 5)
	fmt.Printf("len=%d\tcap=%d\tval=%#v\n", len(s1), cap(s1), s1)
	fmt.Printf("len=%d\tcap=%d\tval=%#v\n", len(s2), cap(s2), s2)
}
