package main

import (
	"fmt"
)

// interface

type person struct {
	name string
	age  uint16
}

type animal struct {
	name string
	age  uint16
}

type skiller interface {
	run() string
	eat() string
}

func (p *person) run() string {
	str := fmt.Sprintf("My name is %s, I am %d years old. Now I am running.\n", p.name, p.age)
	return str
}

func (p *person) eat() string {
	str := fmt.Sprintf("My name is %s, I am %d years old. Now I am eating.\n", p.name, p.age)
	return str
}

func (a *animal) run() string {
	str := fmt.Sprintf("My name is %s, I am %d years old. Now I am running.\n", a.name, a.age)
	return str
}

func (a *animal) eat() string {
	str := fmt.Sprintf("My name is %s, I am %d years old. Now I am eating.\n", a.name, a.age)
	return str
}

func do(s skiller) string {
	str := s.run()
	str = s.eat()
	return str
}

func main() {
	p1 := new(person)
	p1.name = "lucy"
	p1.age = 18
	str1 := do(p1)
	fmt.Println(str1)

	a1 := new(animal)
	a1.name = "kity"
	a1.age = 2
	str2 := do(a1)
	fmt.Println(str2)
}
