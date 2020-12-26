package main

import "fmt"

// map

func main() {
	map1 := make(map[string]int)
	fmt.Println("map1 = ", map1)
	fmt.Printf("map1 的长度：%d，类型：%T，地址：%p\n", len(map1), map1, map1)

	map2 := map[string]int{}
	fmt.Println("map2 = ", map2)
	fmt.Printf("map2 的长度：%d，类型：%T，地址：%p\n", len(map2), map2, map2)

	var map3 map[string]int
	fmt.Println("map3 = ", map3)
	fmt.Printf("map3 的长度：%d，类型：%T，地址：%p\n", len(map3), map1, map3)

	if map1 == nil {
		fmt.Println("map1 是 nil")
	} else if map2 == nil {
		fmt.Println("map2 是 nil")
	} else if map3 == nil {
		fmt.Println("map3 是 nil")
	} else {
		fmt.Println("都不是 nil")
	}

	map2["One"] = 1
	map2["Two"] = 2
	map2["Three"] = 3
	map2["Four"] = 4
	map2["Five"] = 5
	fmt.Println("map2 = ", map2)
	fmt.Printf("map2 的长度：%d，类型：%T，地址：%p\n", len(map2), map2, map2)
	fmt.Println("=====分割线=====")

	if value, ok := map2["Two"]; !ok {
		fmt.Println("map2 中不存在键值为 Two 的元素")
	} else {
		fmt.Printf("value = %d，ok = %T\n", value, ok)
	}
	fmt.Println("=====分割线=====")

	for key, value := range map2 {
		fmt.Printf("第一次遍历：key = %s，value = %d\n", key, value)
	}
	fmt.Println("=====分割线=====")
	for key, value := range map2 {
		fmt.Printf("第二次遍历：key = %s，value = %d\n", key, value)
	}
	fmt.Println("=====分割线=====")

	delete(map2, "Three")
	for key, value := range map2 {
		fmt.Printf("删除键值为 Three 的 map2遍历：key = %s，value = %d\n", key, value)
	}
	fmt.Println("=====分割线=====")

	map2["One"] = 11
	fmt.Printf("map2 的键值为 One 的元素：%d\n", map2["One"])

}
