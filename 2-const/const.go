package main

import "fmt"

// 相当于定义了一个常量值
const name = "zs"

const (
	// iota = 0,然后依次往下递增
	age1 = iota
	// iota = 1
	age2
	// iota = 2
	age3
)

func main() {
	fmt.Println(name)
	// Cannot assign to name,下面就是无法修改常量值
	// name = "ls"

	fmt.Println("age1 = ", age1, "age2 = ", age2, "age3 = ", age3)
}
