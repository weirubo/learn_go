package main

import "fmt"

func main() {
	//name := GetUserName("frank")
	//fmt.Printf("name:%s\n", name)
	//if name != "gopher" {
	//	return
	//}
	//defer fmt.Println("this is a defer call")
	var count int64
	defer func(data int64) {
		fmt.Println("defer:", data)
	}(count + 1)
	count = 100
	fmt.Println("main:", count)
}

func GetUserName(name string) string {
	return name
}
