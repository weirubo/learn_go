package main

import (
	"fmt"
)

// method

type user struct {
	uid   int
	uname string
	age   uint16
}

func (u user) info() string {
	str := fmt.Sprintf("I am %s, My id is %d, and my age is %d.\n", u.uname, u.uid, u.age)
	return str
}

func (u *user) eat() string {
	str := fmt.Sprintf("I am %s, I am eating now.\n", u.uname)
	return str
}

func main() {
	u := user{
		uid:   1,
		uname: "lucy",
		age:   20,
	}
	userInfo := u.info()
	fmt.Println(userInfo)
	fmt.Println(u)

	u1 := &user{
		uid:   2,
		uname: "lily",
		age:   19,
	}
	eater := u1.eat()
	fmt.Println(eater)
	fmt.Println(*u1)
	fmt.Println(u1)
}
