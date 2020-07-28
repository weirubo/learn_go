package main

// 数组

import "fmt"

func main() {
	var arr [3]int
	fmt.Println("arr 数组的第一个元素是：", arr[0])
	fmt.Println("arr:", arr)

	var arr2 = [4]int{1, 2, 3, 4}
	fmt.Println("arr2:", arr2)

	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr3 的长度：", len(arr3))

	const (
		name int = iota
		email
	)
	var arr4 = [...]string{name: "lucy", email: "lucy@gmail.com"}
	fmt.Println(name, arr4[name])
	fmt.Println("arr4:", arr4[0])

	var arr5 = [...]int{99: 100}
	fmt.Println(arr5[0], arr5[1], arr5[99])

	var arrS = [2][3]int{{1, 2, 3}, {10, 20, 30}}
	fmt.Println("二维数组 arrS:", arrS)

	var arrS2 = [...][2][3]int{
		{
			{1, 2, 3},
			{4, 5, 6},
		},
		{
			{11, 12, 13},
			{14, 15, 16},
		},
	}
	fmt.Println("arrS2:", arrS2)

	var arr6 = [5]int{1, 2, 3, 4, 5}
	var arr7 = [...]int{1, 2, 3, 4, 5}
	var arr8 = [5]int{2, 3, 4, 5, 6}
	if arr6 == arr7 {
		fmt.Println("arr6 与 arr7 相等")
	} else if arr6 == arr8 {
		fmt.Println("arr6 与 arr8 相等")
	}

	for index, value := range arr6 {
		fmt.Printf("arr6[%d]=%d\n", index, value)
	}

	for _, value := range arr6 {
		fmt.Println("arr6=", value)
	}

	x, y := 1, 2
	a := [2]*int{&x, &y}
	p := &a
	fmt.Println("数组的元素为指针类型的指针数组：", a)
	fmt.Println("存储数组的内存地址的数组指针：", p)

	var arr9 = [3]int{1, 2, 3}
	var arr10 [3]int
	arr10 = arr9
	fmt.Println("arr9:", arr9)
	fmt.Println("arr10:", arr10)
	arr9[0] = 10
	fmt.Println("arr9:", arr9)
	fmt.Println("arr10:", arr10)
}
