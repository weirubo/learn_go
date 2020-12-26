package main

import "fmt"

// struct

// Admin sturct
type Admin struct {
	person User
	level  string
}

// User struct
type User struct {
	ID   int
	Name string
	age  int
}

func main() {
	var bob User
	fmt.Println("bob = ", bob)
	lucy := User{
		ID:   1,
		Name: "lucy",
		age:  18,
	}
	fmt.Println("lucy = ", lucy)

	lily := User{2, "lily", 17}
	fmt.Println("lily = ", lily)

	joy := Admin{
		person: User{
			ID:   10001,
			Name: "joy",
			age:  28,
		},
		level: "senior",
	}
	fmt.Println("joy = ", joy)
}
