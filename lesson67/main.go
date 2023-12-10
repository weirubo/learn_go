package main

import (
	"fmt"
	"strings"
)

//func main() {
//	v := "true"
//	if s, err := strconv.ParseBool(v); err == nil {
//		fmt.Printf("%T, %v\n", s, s)
//	}
//}

//func main() {
//	v := "3.1415926535"
//	if s, err := strconv.ParseFloat(v, 32); err == nil {
//		fmt.Printf("%T, %v\n", s, s)
//	}
//	if s, err := strconv.ParseFloat(v, 64); err == nil {
//		fmt.Printf("%T, %v\n", s, s)
//	}
//	if s, err := strconv.ParseFloat("NaN", 32); err == nil {
//		fmt.Printf("%T, %v\n", s, s)
//	}
//	// ParseFloat is case insensitive
//	if s, err := strconv.ParseFloat("nan", 32); err == nil {
//		fmt.Printf("%T, %v\n", s, s)
//	}
//	if s, err := strconv.ParseFloat("inf", 32); err == nil {
//		fmt.Printf("%T, %v\n", s, s)
//	}
//	if s, err := strconv.ParseFloat("+Inf", 32); err == nil {
//		fmt.Printf("%T, %v\n", s, s)
//	}
//	if s, err := strconv.ParseFloat("-Inf", 32); err == nil {
//		fmt.Printf("%T, %v\n", s, s)
//	}
//	if s, err := strconv.ParseFloat("-0", 32); err == nil {
//		fmt.Printf("%T, %v\n", s, s)
//	}
//	if s, err := strconv.ParseFloat("+0", 32); err == nil {
//		fmt.Printf("%T, %v\n", s, s)
//	}
//}

//func main() {
//	v32 := "-354634382"
//	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
//		fmt.Printf("s:%T, %v\n", s, s)
//	}
//	s1, err := strconv.ParseInt(v32, 16, 32)
//	fmt.Println(s1, err)
//	if err == nil {
//		fmt.Printf("s1:%T, %v\n", s1, s1)
//	}
//
//	v64 := "-3546343826724305832"
//	if s2, err := strconv.ParseInt(v64, 10, 64); err == nil {
//		fmt.Printf("s2:%T, %v\n", s2, s2)
//	}
//	if s3, err := strconv.ParseInt(v64, 16, 64); err == nil {
//		fmt.Printf("s3:%T, %v\n", s3, s3)
//	}
//}

//func main() {
//	v := true
//	s := strconv.FormatBool(v)
//	fmt.Printf("%T, %v\n", s, s)
//}

//func main() {
//	v := 3.1415926535
//
//	s32 := strconv.FormatFloat(v, 'E', -1, 32)
//	fmt.Printf("%T, %v\n", s32, s32)
//
//	s64 := strconv.FormatFloat(v, 'E', -1, 64)
//	fmt.Printf("%T, %v\n", s64, s64)
//
//	// fmt.Println uses these arguments to print floats
//	fmt64 := strconv.FormatFloat(v, 'g', -1, 64)
//	fmt.Printf("%T, %v\n", fmt64, fmt64)
//}

//func main() {
//	v := int64(-42)
//
//	s10 := strconv.FormatInt(v, 10)
//	fmt.Printf("%T, %v\n", s10, s10)
//
//	s16 := strconv.FormatInt(v, 16)
//	fmt.Printf("%T, %v\n", s16, s16)
//}

//func main() {
//	fmt.Println(strings.Contains("seafood", "foo"))
//	fmt.Println(strings.Contains("seafood", "bar"))
//	fmt.Println(strings.Contains("seafood", ""))
//	fmt.Println(strings.Contains("", ""))
//}

//func main() {
//	fmt.Println(strings.ContainsAny("team", "i"))
//	fmt.Println(strings.ContainsAny("fail", "ui"))
//	fmt.Println(strings.ContainsAny("ure", "ui"))
//	fmt.Println(strings.ContainsAny("failure", "ui"))
//	fmt.Println(strings.ContainsAny("foo", ""))
//	fmt.Println(strings.ContainsAny("", ""))
//}

//func main() {
//	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
//}

//func main() {
//	fmt.Println(strings.ToUpper("Gopher"))
//}

//func main() {
//	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
//	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
//	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
//	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
//}

func main() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
}
