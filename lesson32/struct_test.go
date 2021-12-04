package lesson32

import (
	"testing"
)

type Sayer interface {
	Say() string
}

type Cat struct {
}

func (c Cat) Say() string {
	return "miaow"
}

type Dog struct {
}

func (d Dog) Say() string {
	return "woof woof"
}

type Horse struct {
}

func (h Horse) Say() string {
	return "neigh"
}

func TestSay(t *testing.T) {
	c := Cat{}
	// t.Log("Cat say:", c.Say())
	//
	d := Dog{}
	// t.Log("Dog say:", d.Say())

	h := Horse{}

	animals := []Sayer{c, d, h}
	for _, a := range animals {
		t.Log("say:", a.Say())
	}
}
