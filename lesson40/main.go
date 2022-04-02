package main

import (
	"fmt"
	"math"
)

func main() {
	// var hello = func(name string) { fmt.Printf("hello %s\n", name) }
	// hello("gopher")

	// area := Circle(5, areaOfCircle)
	// fmt.Printf("area:%.2f\n", area)
	//
	// perimeter := Circle(5, perimeterOfCircle)
	// fmt.Printf("perimeter:%.2f\n", perimeter)

	calcArea := CircleCalc("area")
	fmt.Println(calcArea(5))
	calcPerimeter := CircleCalc("perimeter")
	fmt.Println(calcPerimeter(5))
}

// func Circle(r float64, op func(float64) float64) float64 {
// 	return op(r)
// }
//
// func areaOfCircle(r float64) float64 {
// 	return math.Pi * r * r
// }
//
// func perimeterOfCircle(r float64) float64 {
// 	return 2 * math.Pi * r
// }

func CircleCalc(s string) func(float64) float64 {
	switch s {
	case "area":
		return func(r float64) float64 {
			return math.Pi * r * r
		}
	case "perimeter":
		return func(r float64) float64 {
			return 2 * math.Pi * r
		}
	default:
		return nil
	}
}
