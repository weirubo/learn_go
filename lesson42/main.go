package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	name := "frank"
	age := 18
	gender := "male"
	detail, err := CallUserCenter(name, age, gender)
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return
	}
	fmt.Printf("name:%s\nage:%d\ngender:%s\n", detail.name, detail.age, detail.gender)
	fmt.Printf("%s\n", "********************")
	name2 := "lucy"
	age2 := 17
	gender2 := "female"
	salary2 := 5000

	detail2, err := CallUserCenter(name2, age2, gender2, salary2)
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return
	}
	fmt.Printf("name:%s\nage:%d\ngender:%s\nsalary:%d\n", detail2.name, detail2.age, detail2.gender, detail2.salary)
}

func CallUserCenter(name string, age int, gender string, args ...interface{}) (detail *User, err error) {
	if len(args) == 0 {
		detail, err = NewUserUsecase().Create(name, age, gender)
	} else {
		detail, err = NewUserUsecase().Create(name, age, gender, args[0])
	}
	if err != nil {
		return
	}
	return
}

type User struct {
	name   string
	age    int
	gender string
	salary int
}

type UserUsecase interface {
	Create(name string, age int, gender string, args ...interface{}) (*User, error)
}

type userUsecase struct {
}

func NewUserUsecase() UserUsecase {
	return &userUsecase{}
}

func (u *userUsecase) Create(name string, age int, gender string, args ...interface{}) (detail *User, err error) {
	if name == "" || age < 0 || gender == "" {
		err = errors.New("param illegal")
	}
	if len(args) == 0 {
		detail = &User{
			name:   name,
			age:    age,
			gender: gender,
		}
	} else {
		detail = &User{
			name:   name,
			age:    age,
			gender: gender,
			salary: args[0].(int),
		}
	}
	return
}
