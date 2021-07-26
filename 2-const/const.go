package main

import "fmt"

// 相当于定义了一个常量值
const name = "zs"

// Missing value in the const declaration
// const weight float64

// iota,必须和const一起使用，iota是按行数增加的
const (
	// iota = 0,然后依次往下递增
	age1 = iota
	// iota = 1
	age2
	// iota = 2
	age3
)

const (
	// 这里 a， b都在一行，所以 iota的值都是一样的
	a, b = iota + 1, iota + 2 // a = 1, b = 2
	c, d                      // c = 2, d = 3
	e, f                      // e = 3, f = 4

	g, h = iota * 1, iota * 2 // g = 3, h = 6
	i, j                      // i = 4, j = 8
)

func main() {
	fmt.Println(name)
	// Cannot assign to name,下面就是无法修改常量值
	// name = "ls"

	fmt.Println("age1 =", age1, " age2 =", age2, " age3 =", age3)

	fmt.Println("a =", a, " b =", b)
	fmt.Println("c =", c, " d =", d)
	fmt.Println("e =", e, " h =", h)
	fmt.Println("i =", i, " j =", j)
}
