package main

import "fmt"

// type User struct {
// 	Id   int
// 	Name string
// }
//
// type option func(*User)
//
// func (u *User) Option(opts ...option) {
// 	for _, opt := range opts {
// 		opt(u)
// 	}
// }
//
// func WithId(id int) option {
// 	return func(u *User) {
// 		u.Id = id
// 	}
// }
//
// func WithName(name string) option {
// 	return func(u *User) {
// 		u.Name = name
// 	}
// }
//
// func main() {
// 	u1 := &User{}
// 	u1.Option(WithId(1))
// 	fmt.Printf("%+v\n", u1)
//
// 	u2 := &User{}
// 	u2.Option(WithId(1), WithName("frank"))
// 	fmt.Printf("%+v\n", u2)
// }

// type User struct {
// 	Id   int
// 	Name string
// }
//
// type option func(*User) interface{}
//
// func (u *User) Option(opts ...option) (id interface{}) {
// 	for _, opt := range opts {
// 		id = opt(u)
// 	}
// 	return id
// }
//
// func WithId(id int) option {
// 	return func(u *User) interface{} {
// 		prevId := u.Id
// 		u.Id = id
// 		return prevId
// 	}
// }
//
// func main() {
// 	u1 := &User{Id: 1}
// 	id := u1.Option(WithId(2))
// 	fmt.Println(id.(int))
// 	fmt.Printf("%+v\n", u1)
// }

// type User struct {
// 	Id   int
// 	Name string
// }
//
// type option func(*User) option
//
// func (u *User) Option(opts ...option) (prev option) {
// 	for _, opt := range opts {
// 		prev = opt(u)
// 	}
// 	return prev
// }
//
// func WithId(id int) option {
// 	return func(u *User) option {
// 		prevId := u.Id
// 		u.Id = id
// 		return WithId(prevId)
// 	}
// }
//
// func main() {
// 	u1 := &User{Id: 1}
// 	prev := u1.Option(WithId(2))
// 	fmt.Printf("%+v\n", u1)
// 	u1.Option(prev)
// 	fmt.Printf("%+v\n", u1)
// }

type User struct {
	Id    int
	Name  string
	Email string
}

type option func(*User)

func WithId(id int) option {
	return func(u *User) {
		u.Id = id
	}
}

func WithName(name string) option {
	return func(u *User) {
		u.Name = name
	}
}

func WithEmail(email string) option {
	return func(u *User) {
		u.Email = email
	}
}

func NewUser(opts ...option) *User {
	const (
		defaultId    = -1
		defaultName  = "guest"
		defaultEmail = "undefined"
	)
	u := &User{
		Id:    defaultId,
		Name:  defaultName,
		Email: defaultEmail,
	}

	for _, opt := range opts {
		opt(u)
	}
	return u
}

func main() {
	u1 := NewUser(WithName("frank"), WithId(1000000001))
	fmt.Printf("%+v\n", u1)
	u2 := NewUser(WithEmail("gopher@88.com"))
	fmt.Printf("%+v\n", u2)
	u3 := NewUser()
	fmt.Printf("%+v\n", u3)
}
