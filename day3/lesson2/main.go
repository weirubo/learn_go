package main

import "fmt"

// slice

func main() {
	arr := [...]int{10, 11, 12, 13, 14}
	slice := arr[1:4]
	fmt.Printf("slice 的长度：%d，slice 的容量：%d\n", len(slice), cap(slice))
	fmt.Println("slice = ", slice)
	fmt.Printf("slice 的内存地址：%p\n", slice)
	fmt.Printf("arr[1] 的内存地址：%p\n", slice)

	s1 := make([]int, 2, 4)
	fmt.Printf("s1 的长度：%d，容量：%d\n", len(s1), cap(s1))
	s2 := make([]int, 2)
	fmt.Printf("s2 的长度：%d，容量：%d\n", len(s2), cap(s2))

	s3 := []int{0, 1, 2, 3}
	fmt.Printf("s3 的长度：%d，容量：%d\n", len(s3), cap(s3))
	s4 := []int{9: 9}
	fmt.Printf("s4 的长度：%d，容量：%d\n", len(s4), cap(s4))

	var sliceNil []int
	fmt.Printf("sliceNil 的长度：%d，容量：%d，类型：%T\n", len(sliceNil), cap(sliceNil), sliceNil)
	fmt.Println("sliceNil = ", sliceNil)

	s5 := make([]int, 0)
	fmt.Printf("s5 的长度：%d，容量：%d\n", len(s5), cap(s5))

	s6 := []int{}
	fmt.Printf("s6 的长度：%d，容量：%d\n", len(s6), cap(s6))

	s7 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	s8 := s7[1:3]
	fmt.Printf("s7 的长度：%d，容量：%d\n", len(s7), cap(s7))
	fmt.Printf("s8 的长度：%d，容量：%d\n", len(s8), cap(s8))
	s8[0] = 21
	fmt.Println("s8 = ", s8)
	fmt.Println("s7 = ", s7)

	s9 := [][]int{{1, 2, 3}, {11, 12, 13, 14}}
	fmt.Printf("s9 的长度：%d，容量：%d\n", len(s9), cap(s9))
	fmt.Printf("s9[0] 的长度：%d，容量：%d\n", len(s9[0]), cap(s9[0]))
	fmt.Printf("s9[1] 的长度：%d，容量：%d\n", len(s9[1]), cap(s9[1]))
	fmt.Println(s9)

	s10 := []int{0, 1, 2, 3, 4, 5}
	for index, value := range s10 {
		fmt.Printf("s10[%d] = %d\n", index, value)
	}

	s11 := []int{10, 11, 12, 13, 14}
	s12 := s11[1:3]
	fmt.Printf("s12 的长度：%d，容量：%d，地址：%p\n", len(s12), cap(s12), s12)
	fmt.Println("s12 = ", s12)
	s12 = append(s12, 21, 22, 23)
	fmt.Printf("s12 的长度：%d，容量：%d，地址：%p\n", len(s12), cap(s12), s12)
	fmt.Println("s12 = ", s12)

	s13 := []int{10, 11, 12}
	s14 := []int{20, 21, 22}
	fmt.Printf("s14 的长度：%d，容量：%d，地址：%p\n", len(s14), cap(s14), s14)
	fmt.Println("s14 = ", s14)
	s14 = append(s14, s13...)
	fmt.Printf("s14 的长度：%d，容量：%d，地址：%p\n", len(s14), cap(s14), s14)
	fmt.Println("s14 = ", s14)

	s15 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	s16 := s15[4:]
	newSlice := copy(s15[:4], s16)
	fmt.Printf("newSlice = %d，s15 = %v", newSlice, s15)
}
