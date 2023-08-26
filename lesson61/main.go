package main

import "fmt"

func main() {
	//m := min()
	//fmt.Println(m)
	//s := []int{1, 2, 3}
	//m := min(s...)
	//fmt.Println(m)
	//m := min(3.14, math.NaN(), 1.0)
	//fmt.Println(m)
	//m := max(2, math.NaN(), 1)
	//fmt.Println(m)
	//var x int
	//m := min(x)    // m == x
	//fmt.Println(m) // 0
	//var x, y int = 1, 2
	//m := min(x, y) // m is the smaller of x and y
	//fmt.Println(m)
	//m := max(x, y, 10)          // m is the larger of x and y but at least 10
	//c := min(1, 2.0, 3) // c == 1.0 (floating-point kind)
	//fmt.Printf("%T\t%v\n", c, c)
	//f := max(0, float32(x))     // type of f is float32
	//var s []string
	//_ = min(s...)               // invalid: slice arguments are not permitted
	//t := min("", "foo", "bar") // t == "" (string kind)
	//fmt.Println(t)

	//s := []int{1, 2, 3}
	//fmt.Printf("len=%d\t s=%+v\n", len(s), s)
	//clear(s)
	//fmt.Printf("len=%d\t s=%+v\n", len(s), s)

	//m := map[string]int{"go": 100, "php": 80}
	//fmt.Printf("len=%d\tm=%+v\n", len(m), m)
	//clear(m)
	//fmt.Printf("len=%d\tm=%+v\n", len(m), m)

	//var s []int
	//clear(s)

	//var m map[string]int
	//clear(m)

	d := []Data{
		{
			User:   map[int]string{1: "frank", 2: "lucy"},
			Salary: map[string]int{"frank": 1000, "lucy": 2000},
		},
	}
	fmt.Printf("d=%+v\n", d)
	clear(d)
	fmt.Printf("d=%+v\n", d)

	d1 := []Data1{
		{
			User:   "frank",
			Salary: 1000,
		},
	}
	fmt.Printf("d1=%+v\n", d1)
	clear(d1)
	fmt.Printf("d1=%+v\n", d1)
}

type Data struct {
	User   map[int]string
	Salary map[string]int
}

type Data1 struct {
	User   string
	Salary int
}
