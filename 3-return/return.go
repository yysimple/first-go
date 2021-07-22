package main

import "fmt"

// 返回一个值
func returnOne(a int) int {
	fmt.Println("------- returnOne ---------")
	fmt.Println("返回一个值,a =", a)
	c := 1
	return c
}

// 返回两个值,这是匿名的方式，就是按顺序返回
func returnTwo(a int) (int, int) {
	fmt.Println("------- returnTwo ---------")
	fmt.Println("返回两个值，匿名的,a =", a)
	// 按顺序返回的，接收的时候也是按顺序接收的
	return a * 1, a * 2
}

// 返回两个值，是带名字的
func returnThree(a int) (param1 int, param2 int) {
	fmt.Println("------- returnThree ---------")
	fmt.Println("返回两个值，带有名字的的,a =", a)
	// 可以对形参进行赋值
	param1 = 1
	param2 = 2
	// 这里最后还是按顺序返回的：（2， 1）
	return param2, param1
}

// 当两个形参都是 int 类型的时候 只需要写一个 int
func returnFour(a int) (p1, p2 int) {
	fmt.Println("------- returnFour ---------")
	fmt.Println("返回两个值，只需要写一个 int 即可 ,a =", a)
	return 1, 2
}

func main() {
	one := returnOne(1)
	fmt.Println("one =", one)

	multi1, multi2 := returnTwo(2)
	fmt.Println("multi1 =", multi1, " multi2 =", multi2)

	param1, param2 := returnThree(3)
	fmt.Println("param1 =", param1, " param2 =", param2)

	p1, p2 := returnFour(4)
	fmt.Println("p1 =", p1, " p2 =", p2)
}
