package main

var r *int

func main() {
	sum(1, 2)
}

func sum(a, b int) *int {
	res := a + b
	return &res
}
