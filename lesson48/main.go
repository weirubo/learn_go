package main

import (
	"errors"
	"log"
)

func main() {
	err := foo()
	if err != nil {
		log.Printf("err=%v\n", err)
		return
	}
}

func foo() (err error) {
	var code int
	if code, err = bar(); err != nil {
		log.Printf("code=%v err=%v\n", code, err)
		return // Compiler reports: err is shadowed during return
	}
	return nil
}

func bar() (int, error) {
	err := errors.New("this is bar's err")
	return 200, err
}
