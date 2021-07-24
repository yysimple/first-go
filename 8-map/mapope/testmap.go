package main

import "fmt"

func main() {
	fruitMap := make(map[string]string, 10)
	// 添加元素
	fruitMap["apple"] = "sweet and sour"
	fruitMap["banana"] = "sweet"
	fruitMap["yangmei"] = "sour"

	for k, v := range fruitMap {
		fmt.Println("key =", k, "value =", v)
	}

	fmt.Println("---------------")

	delete(fruitMap, "apple")
	for k, v := range fruitMap {
		fmt.Println("key =", k, "value =", v)
	}

	fmt.Println("---------------")

	fruitMap["yangmei"] = "very sour"
	for k, v := range fruitMap {
		fmt.Println("key =", k, "value =", v)
	}

}
