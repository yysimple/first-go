package main

import "fmt"

func main() {
	// 1. 最简单的声明式，key为string类型，value为string类型，是个空map
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("map1 is null")
	}

	// 2. 使用make给map2分配初始化空间
	map1 = make(map[string]string, 10)
	map1["1"] = "one"
	map1["2"] = "two"
	fmt.Println("map1 =", map1)

	// 3. 简写声明
	map2 := make(map[string]string, 10)
	map2["1"] = "one"
	map2["2"] = "two"
	fmt.Println("map2 =", map2)

	// 4. 直接赋值
	map3 := map[string]string{
		// 这里注意，每行都要加上 , 结束
		"1": "one",
		"2": "two",
	}
	fmt.Println("map3 =", map3)

}
