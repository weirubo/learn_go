package main

import (
	"fmt"
	"time"
)

// goroutine 和 channel

func main() {
	go func() {
		for i := 0; i < 50; i++ {
			fmt.Println("number:", i)
		}
	}()

	go func() {
		str := "hello,world"
		for index := range str {
			fmt.Println("byte:", string(str[index]))
		}
	}()

	ch := make(chan string, 11)
	go func() {
		str := "hello,world"
		for index := range str {
			ch <- string(str[index])
		}
	}()

	go func() {
		for ch != nil {
			select {
			case byte := <-ch:
				str := byte
				fmt.Println("str:", str)
				fmt.Printf("ch 的容量是%d，长度是%d。\n", cap(ch), len(ch))
			default:

			}
		}
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main.go func over")
}
