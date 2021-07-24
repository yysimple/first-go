package main

import (
	"fmt"
)

func printArray(arr [4]int) {

	for index, value := range arr {
		fmt.Println("index =", index, "value =", value)
	}
	// 这里并不会修改原有数组的值，这里只是进行了一次浅拷贝，只是拷贝了值
	arr[1] = 1111
}

func main() {

	// 初始化长度为 10 的数组，类型对应的是 [10]int
	var arr1 [10]int
	arr2 := [10]int{1, 2, 3, 4}
	arr3 := [4]int{11, 22, 33, 44}

	// 这里的i相当于 数组的位置索引
	for i := range arr1 {
		fmt.Println("对应的元素 ==> ", arr1[i])
	}

	// 这里对应的 index 为索引位置，value为具体的值
	for index, value := range arr2 {
		fmt.Println("打印index =", index, "打印value =", value)
	}

	// 这里需要注意的是，该打印方法的形参类型必须和输入的实参一致，否则无法接收，即两者都应该是 arr [4]int
	printArray(arr3)

	fmt.Println("----------")

	for index, value := range arr3 {
		fmt.Println("index =", index, "value =", value)
	}

	fmt.Printf("arr1 type = %T\n", arr1)
	fmt.Printf("arr2 type = %T\n", arr2)
	fmt.Printf("arr3 type = %T\n", arr3)

}
