package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	// ToString
	val1 := 1
	fmt.Printf("Before val=%v type=%T\n", val1, val1)
	fmt.Printf("After val=%v type=%T\n", cast.ToString(val1), cast.ToString(val1))

	val2 := []int{1, 2, 3}
	fmt.Printf("Before val=%v type=%T\n", val2, val2)
	fmt.Printf("After val=%v type=%T\n", cast.ToStringSlice(val2), cast.ToStringSlice(val2))

	// ToInt
	// 如果不是数值字符串，转换后的值是类型零值
	val3 := "1"
	fmt.Printf("Before val=%v type=%T\n", val3, val3)
	fmt.Printf("After val=%v type=%T\n", cast.ToInt(val3), cast.ToInt(val3))

	// 如果切片元素不是数值字符串，转换后的值是类型零值
	val4 := []string{"1", "2", "3"}
	fmt.Printf("Before val=%v type=%T\n", val4, val4)
	fmt.Printf("After val=%v type=%T\n", cast.ToIntSlice(val4), cast.ToIntSlice(val4))

	// ToBool
	// 非零整型都是 true，非整型和零都是 false
	val5 := 1
	fmt.Printf("Before val=%v type=%T\n", val5, val5)
	fmt.Printf("After val=%v type=%T\n", cast.ToBool(val5), cast.ToBool(val5))

	// 切片元素必须是 int 类型
	val6 := []int{-1, 0, 1}
	fmt.Printf("Before val=%v type=%T\n", val6, val6)
	fmt.Printf("After val=%v type=%T\n", cast.ToBoolSlice(val6), cast.ToBoolSlice(val6))

	// 参数值必须是整型，或字符串数值
	val7 := 1636811715
	fmt.Printf("Before val=%v type=%T\n", val7, val7)
	fmt.Printf("After val=%v type=%T\n", cast.ToTime(val7), cast.ToTime(val7))

	a := 1
	fmt.Printf("val=%v type=%T\n", cast.ToString(a), cast.ToString(a))
	b := 3.14
	fmt.Printf("val=%v type=%T\n", cast.ToString(b), cast.ToString(b))
	c := "hello"
	fmt.Printf("val=%v type=%T\n", cast.ToString(c), cast.ToString(c))
	d := []byte("golang")
	fmt.Printf("val=%v type=%T\n", cast.ToString(d), cast.ToString(d))
	var e interface{} = "frank"
	fmt.Printf("val=%v type=%T\n", cast.ToString(e), cast.ToString(e))
	f := []int{1, 2, 3}
	fmt.Printf("val=%v type=%T\n", f, f)
	fmt.Printf("val=%v type=%T\n", cast.ToString(nil), cast.ToString(nil))

	v, err := cast.ToStringE([]int{1, 2, 3})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("val=%v type=%T\n", v, v)
}
