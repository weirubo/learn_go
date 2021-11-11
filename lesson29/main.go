package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	intSlice := IntSlice([]int{9, 7, 5, 3, 1})
	fmt.Println(intSlice) // 排序前
	// sort.Sort(intSlice)
	// fmt.Println(intSlice) // 排序后
	sort.Ints(intSlice)
	fmt.Println(intSlice) // 使用 sort.Ints() 排序数据

	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
}
