package main

import (
	"fmt"
	"github.com/spf13/cast"
	"strconv"
)

func main() {
	salary := 5000
	salaryStr := strconv.Itoa(salary)
	fmt.Printf("%T salary=%d\n", salary, salary)
	fmt.Printf("%T salaryStr=%s\n", salaryStr, salaryStr)

	age := "23"
	ageInt, err := strconv.Atoi(age)
	fmt.Printf("%T age=%s\n", age, age)
	fmt.Printf("%T ageInt=%d err=%v\n", ageInt, ageInt, err)

	ageInt8 := int8(ageInt)
	fmt.Println(ageInt8)

	phoneNumber := "138001380001380013800013800138000"
	phoneNumberInt, err := strconv.Atoi(phoneNumber)
	fmt.Printf("%T phoneNumber=%s\n", phoneNumber, phoneNumber)
	fmt.Printf("%T phoneNumberInt=%d err=%v\n", phoneNumberInt, phoneNumberInt, err)

	age2 := "23"
	age2Int8 := cast.ToInt8(age2)
	fmt.Printf("%T age2=%s\n", age2, age2)
	fmt.Printf("%T age2Int8=%d\n", age2Int8, age2Int8)

	phoneNumber2 := "138001380001380013800013800138000"
	phoneNumber2Int := cast.ToInt(phoneNumber2)
	fmt.Printf("%T phoneNumber2=%s\n", phoneNumber2, phoneNumber2)
	fmt.Printf("%T phoneNumber2Int=%d\n", phoneNumber2Int, phoneNumber2Int)

	month := "07"
	monthInt8 := cast.ToInt8(month)
	fmt.Printf("%T month=%s\n", month, month)
	fmt.Printf("%T monthInt8=%d\n", monthInt8, monthInt8)

	month2 := "08"
	month2Int8 := cast.ToInt8(month2)
	fmt.Printf("%T month2=%s\n", month2, month2)
	fmt.Printf("%T month2Int8=%d\n", month2Int8, month2Int8)

	month2Int2, err := strconv.Atoi(month2)
	fmt.Printf("%T month2Int2=%d err=%v\n", month2Int2, month2Int2, err)
}
