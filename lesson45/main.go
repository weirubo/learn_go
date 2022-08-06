package main

import "fmt"

func main() {
	var name string = "frank"
	fmt.Printf("val:%s || ptr:%p\n", name, &name)
	name, age := "lucy", 18
	fmt.Printf("val:%s || ptr:%p\n", name, &name)
	fmt.Println(age)
	if name != "lily" {
		name := "lily"
		fmt.Printf("val:%s || ptr:%p\n", name, &name)
	}
}
