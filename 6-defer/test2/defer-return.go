package main

import "fmt"

// 这里注意在同一个包下面只能有一个main方法
func main() {
	// defer方法是在return之后在执行的
	deferAndReturn()
}

func deferFunc() int {
	fmt.Println("--- defer ---")
	return 0
}

func returnFunc() int {
	fmt.Println("--- return ---")
	return 0
}

func deferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}
