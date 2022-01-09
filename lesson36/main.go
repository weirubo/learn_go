package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.baidu2022.com/robots.txt")
	// fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
	user := GetUser()
	// fmt.Println(user)
	// if user != nil {
	// 	fmt.Println(user.Id)
	// }
	user.Login()

	var userData map[int]*User
	// fmt.Println(userData[1].Name)
	if val, ok := userData[1]; ok {
		fmt.Println(val.Name)
	}

}

func GetUser() (user *User) {
	user = new(User)
	// user = &User{}
	return
}

type User struct {
	Id   int
	Name string
}

// func (u User) Login() {
//
// }

func (u *User) Login() {

}
