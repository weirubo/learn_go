package main

import (
	"fmt"
	"time"
)

func main() {
	var m = make(map[int]string)
	fmt.Println(m)
	// 赋值
	m[0] = "a"
	m[1] = "b"
	fmt.Println(m)
	// 读取
	val := m[1]
	fmt.Println(val)
	// 删除
	delete(m, 1)
	fmt.Println(m)
	// 更新
	m[1] = "x"
	fmt.Println(m)

	//go func() {
	//	for {
	//		m[1] = "xx"
	//	}
	//}()

	go func() {
		for {
			_ = m[1]
		}
	}()

	go func() {
		for {
			delete(m, 1)
		}
	}()

	//// 并发读
	//for i := 0; i < 10; i++ {
	//	go fmt.Println(m[i])
	//}
	//
	//// 并发写
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		m[i] = "y"
	//		fmt.Println(m[1])
	//	}()
	//}
	time.Sleep(time.Second * 3)
}
