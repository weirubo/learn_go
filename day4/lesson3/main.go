package main

import "fmt"

// panic recover

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	panic("exit...")
	fmt.Println("game over!")

	/* defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()

	defer func() {
		panic("exit2...")
	}()

	panic("exit...") */
}
