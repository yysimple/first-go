package main

import "fmt"

func main() {
	noUsePoint()
	usePoint()
}

func noUsePoint() {
	a := 1
	// 传递值
	changeValue(a)
	// 打印的值是不会变的，a = 1
	fmt.Println("值传递：a =", a)
}

func usePoint() {
	a := 1
	// 传递指针
	changePoint(&a)
	// 这里a = 10
	fmt.Println("指针传递：a =", a)
}

// 这里只需要传入值即可 a = 1
func changeValue(a int) {
	a = 10
}

// 这种方式需要传递进来的是引用地址 0x0000844 之类的
func changePoint(a *int) {
	*a = 10
}
