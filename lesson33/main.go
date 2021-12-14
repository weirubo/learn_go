package main

import (
	"fmt"
)

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("err = ", err)
	// 	}
	// }()
	user := new(User)
	// user := &User{}
	userInfo := user.GetInfo()
	fmt.Println(userInfo)
	if userInfo.Age >= 18 {
		fmt.Println("this is a man")
	}

	// code := []string{"php", "golang"}
	// fmt.Printf("len:%d cap:%d val:%s \n", len(code), cap(code), code)
	// fmt.Println(code[2])

	// var ch chan int
	// close(ch)
	// ch = make(chan int, 1)
	// ch <- 1
	// close(ch)
	// close(ch)
	// ch <- 2

	// var m map[string]int
	// m = make(map[string]int)
	// go func(map[string]int) {
	// 	for {
	// 		m["php"] = 80
	// 	}
	// }(m)
	//
	// go func(map[string]int) {
	// 	for {
	// 		_ = m["php"]
	// 	}
	// }(m)
	//
	// time.Sleep(time.Second)

	var name interface{} = "frank"
	// a, ok := name.(int)
	// fmt.Println(a, ok)
	a := name.(int)
	fmt.Println(a)
}

type User struct {
	Name string
	Age  int
}

func (u *User) GetInfo() (data *User) {
	data = &User{
		Name: "frank",
		Age:  18,
	}
	return data
}
