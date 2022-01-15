package main

import "fmt"

func main() {
	name := "frank"
	user := &User{
		Id:   1,
		Name: &name,
	}
	fmt.Println(user)
}

type User struct {
	Id   int
	Name *string
}

func (u *User) String() string {
	return fmt.Sprintf("{Id: %v, Name: %v}", u.Id, *u.Name)
}
