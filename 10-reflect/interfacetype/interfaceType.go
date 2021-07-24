package main

import "fmt"

// interface{} 是个万能类型，可以理解为 java中的泛型
func allType(arg interface{}) {
	fmt.Println("can receive all type ...")
	fmt.Println("arg =", arg)

	// 断言机制，其实就是类型比较，先当于 java中的 instanceof
	v, status := arg.(string)
	if !status {
		fmt.Println("the type is not string...")
	} else {
		fmt.Printf("the type is = %T, the vaule = %v\n", v, v)
	}
}

type Book struct {
	name string
}

func main() {
	book := Book{name: "许三观卖血记"}
	allType(book)
	fmt.Println("---------------------------")
	allType("i'm string")
	fmt.Println("---------------------------")
	allType(100)
	fmt.Println("---------------------------")
	allType(179.5)
}
