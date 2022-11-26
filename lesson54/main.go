package main

import "fmt"

func main() {
	//fmt.Println(Min(2, 3))
	//fmt.Println(Min(2.2, 3.3))
	//fmt.Println(MinFloat64(2.2, 3.3))
	//fmt.Println(GMin(2, 3))
	//fmt.Println(GMin(2.2, 3.3))
	//fmt.Println(GMin[int](2, 3))
	//fmt.Println(GMin[float64](2.2, 3.3))

	//x, y := 11, 2
	//x, y := 2.2, 3.3
	//res := MinAny(x, y)
	//fmt.Println(res)

	//salary := &MinSalary[int]{
	//	salary: 1000,
	//}
	//salary := &MinSalary[float64]{
	//	salary: 1000.88,
	//}
	//salary := &MinSalary{
	//	salary: 1000,
	//}
	//fmt.Printf("%+v\n", salary)

	s := &Salary[int]{}
	x, y := 2, 3
	fmt.Println(s.Min(x, y))
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MinFloat64(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func GMin[T int | float64](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// 泛型函数
//func GMin[T Ordered](x, y T) T {
//	if x < y {
//		return x
//	}
//	return y
//}

// 泛型类型
type Ordered interface {
	~int | ~float64
}

func MinAny(x, y interface{}) interface{} {
	if x, ok := x.(int); ok {
		if y, ok := y.(int); ok {
			if x < y {
				return x
			}
		}
	}
	if x, ok := x.(float64); ok {
		if y, ok := y.(float64); ok {
			if x < y {
				return x
			}
		}
	}
	return y
}

type MinSalary[T int | float64] struct {
	salary T
}

type Salary[T int | float64] struct {
	money T
}

func (s *Salary[T]) Min(x, y T) T {
	if x < y {
		return x
	}
	return y
}

//func (s *Salary[T]) Min[T1 int](x, y T) T {
//	if x < y {
//		return x
//	}
//	return y
//}
