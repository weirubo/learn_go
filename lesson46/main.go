package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		// go func(param string) {
		// 	// fmt.Println(v)
		// 	fmt.Printf("val=%s pointer=%p\n", param, &param)
		// 	done <- true
		// }(v)

		param := v
		go func() {
			// fmt.Println(v)
			fmt.Printf("val=%s pointer=%p\n", param, &param)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}
