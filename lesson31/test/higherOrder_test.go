package test

import "testing"

func TestHigherOrder(t *testing.T) {
	t.Log(operation(1, 2, func(x, y int) int {
		return x + y
	}))

	sum := operation1()
	t.Log(sum(2, 3))
}

func operation(x, y int, operaFuc func(int, int) int) int {
	rst := operaFuc(x, y)
	return rst
}

func operation1() func(x, y int) int {
	sum := func(x, y int) int {
		return x + y
	}
	return sum
}
