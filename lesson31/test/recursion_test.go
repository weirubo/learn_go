package test

import "testing"

func TestRecursion(t *testing.T) {
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	t.Log(fib(7))
}
