package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now().Unix()
	fmt.Println(nowTime)
	// NewTimerDemo()
	AfterFuncDemo()
}

func NewTimerDemo() {
	timer := time.NewTimer(2 * time.Second)
	select {
	case <-timer.C:
		currentTime := time.Now().Unix()
		fmt.Println(currentTime, "do something")
	}
}

func AfterFuncDemo() {
	time.AfterFunc(2*time.Second, func() {
		currentTime := time.Now().Unix()
		fmt.Println(currentTime, "do something")
	})
	time.Sleep(3 * time.Second)
}
