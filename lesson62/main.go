package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{"lucy": 17, "lily": 18}
	m2 := map[string]int{"lucy": 17, "lily": 18}
	m3 := map[string]int{"lucy": 18, "lily": 17}
	fmt.Println(maps.Equal(m1, m2))
	fmt.Println(maps.Equal(m1, m3))

	//maps.EqualFunc()
}

/**
BinarySearch searches for target in a sorted slice and returns the position where target is found,
or the position where target would appear in the sort order;
it also returns a bool saying whether the target is really found in the slice.
The slice must be sorted in increasing order.
*/
//names := []string{"Alice", "Bob", "Vera"}
//n, found := slices.BinarySearch(names, "Vera")
//fmt.Println("Vera:", n, found)
//n, found = slices.BinarySearch(names, "Bill")
//fmt.Println("Bill:", n, found)

///**
//BinarySearchFunc works like BinarySearch, but uses a custom comparison function.
//The slice must be sorted in increasing order, where "increasing" is defined by cmp.
//cmp should return 0 if the slice element matches the target, a negative number if the slice element precedes the target,
//or a positive number if the slice element follows the target. cmp must implement the same ordering as the slice,
//such that if cmp(a, t) < 0 and cmp(b, t) >= 0, then a must precede b in the slice.
//*/
//type Person struct {
//	Name string
//	Age  int
//}
//people := []Person{
//	{"Alice", 55},
//	{"Bob", 24},
//	{"Gopher", 13},
//}
//n, found := slices.BinarySearchFunc(people, Person{"Bob", 0}, func(a, b Person) int {
//	return cmp.Compare(a.Name, b.Name)
//})
//fmt.Println("Bob:", n, found)
