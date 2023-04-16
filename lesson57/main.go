package main

import "fmt"

func main() {
	// 内置函数
	// append
	//s := []int{1, 2, 3}
	//fmt.Printf("%p %d\n", s, s)
	//s1 := append(s, 4)
	//fmt.Printf("%p %d\n", s, s)
	//fmt.Printf("%p %d\n", s1, s1)

	// copy
	//src := []string{"go", "vue"}
	//dst := make([]string, 3)
	//n := copy(dst, src)
	//fmt.Printf("%s %d\n", dst, n)

	//str := "hello"
	//bs := make([]byte, 5)
	//n := copy(bs, str)
	//fmt.Printf("%s %d\n", bs, n)

	// map
	//var m map[int]string
	//fmt.Println(m)
	//delete(m, 0)
	//fmt.Println(m)
	//m1 := make(map[int]string)
	//fmt.Println(m1)
	//delete(m1, 0)
	//fmt.Println(m1)
	//m2 := make(map[int]string, 2)
	//m2[0] = "hello"
	//m2[1] = "world"
	//fmt.Println(m2)
	//delete(m2, 0)
	//fmt.Println(m2)

	// len
	//arr := [3]int{1, 2, 3}
	//fmt.Println(arr)
	//fmt.Println(len(arr))
	//var arr1 *[3]int
	//fmt.Println(arr1)
	//fmt.Println(len(arr1))
	//var s []int
	//fmt.Println(len(s))
	//s = []int{1, 2, 3}
	//fmt.Println(len(s))
	//var m map[int]string
	//fmt.Println(len(m))
	//m = make(map[int]string)
	//m[0] = "hello"
	//fmt.Println(len(m))
	//str := "frank"
	//fmt.Println(len(str))
	//var c chan int
	//fmt.Println(c)
	//fmt.Println(len(c))
	//c = make(chan int)
	//fmt.Println(len(c))

	// cap
	var arr [3]int
	fmt.Println(arr)
	fmt.Println(cap(arr))
	var arr1 *[3]int
	fmt.Println(arr1)
	fmt.Println(cap(arr1))
	var s []string
	fmt.Println(s)
	fmt.Println(cap(s))
	s = make([]string, 1)
	s[0] = "go"
	fmt.Println(s)
	fmt.Println(cap(s))
	var c chan int
	fmt.Println(c)
	fmt.Println(cap(c))
}
