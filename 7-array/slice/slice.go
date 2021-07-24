package main

import "fmt"

func printArray(arr []int) {
	// _ 这个表示是匿名的数据，不需要索引时可以这样定义
	for _, v := range arr {
		fmt.Println("打印的值：", v)
	}

	// 这里进行的深拷贝，拷贝了一块引用地址
	arr[1] = 100
}

// 四种定义slice的方式
func fourDefineSlice() {
	// 1. 最普通的声明方式,长度是 4
	slice1 := []int{1, 2, 3, 4}
	fmt.Printf("type = %T\n", slice1)
	fmt.Printf("slice1: len = %d, slice1 = %v\n", len(slice1), slice1)

	// 2. 初始化slice，但是没有分配固定空间
	var slice2 []int
	fmt.Printf("type = %T\n", slice2)
	fmt.Printf("slice2: len = %d, slice2 = %v\n", len(slice2), slice2)

	// 3. make的作用就是，前面传入的是类型，后面时初始化大小即分配了默认的空间，默认值都是0
	var slice3 = make([]int, 4)
	fmt.Printf("type = %T\n", slice3)
	fmt.Printf("slice3: len = %d, slice3 = %v\n", len(slice3), slice3)

	// 4. 简写方式，初始化分配四个空间, 这种方式用的比较多
	slice4 := make([]int, 4)
	fmt.Printf("type = %T\n", slice4)
	fmt.Printf("slice4: len = %d, slice4 = %v\n", len(slice4), slice4)
}

func main() {

	slice1 := []int{1, 2, 3, 4}
	fmt.Printf("type = %T\n", slice1)

	printArray(slice1)

	fmt.Println("---------------")

	for _, v := range slice1 {
		fmt.Println("打印的值：", v)
	}

	fmt.Println("------- 初始化slice的方式 --------")

	fourDefineSlice()
}
