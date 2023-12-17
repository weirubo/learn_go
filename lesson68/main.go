package main

import "fmt"

func main() {
	var arr1 [2]int
	var arr2 = [2]int{1, 2}
	var arr3 = [...]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(len(arr3))
	//fmt.Println(arr2 == arr3)
	//fmt.Println(arr3[3])
	//arr3[3] = 4
	Get(arr2)
	fmt.Printf("arr2=%p\n%d\n", &arr2, arr2)
	Store()
}

func Get(arr [2]int) {
	fmt.Printf("Get()=%p\n%d\n", &arr, arr)
}

func Store() {
	var arr [2]int
	for i := 0; i < 5; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)
}
