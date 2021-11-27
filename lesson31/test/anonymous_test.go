package test

import "testing"

func TestAnonymous(t *testing.T) {
	// func () {
	// 	t.Log("Hello World")
	// }()
	// func (a, b int) int {
	// 	return a + b
	// }(1,2)
	sum := func(a, b int) int {
		return a + b
	}
	t.Log(sum(1, 2))
}
