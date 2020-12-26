package main

import "fmt"

// func

func f1(string, string) string {
	return ""
}

func f2(name string, _ string) string {
	return name
}

func f3(name string, job string) string {
	return name + job
}

func f4(name, job string) (info string) {
	info = name + job
	return
}

// 声明函数，多返回值

// 变长函数，形参数量任意
func sum(vals ...int) int {
	return 0
}

// 闭包

// defer 延迟调用

// 函数变量(基础数据类型，复合数据类型)

func main() {
	// 匿名函数
	func(str string) {
		fmt.Println(str)
	}("hello, golang")

}
