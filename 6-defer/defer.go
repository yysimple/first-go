package main

import "fmt"

func main() {
	// defer就是程序结束前执行，这里的话是 栈 的存储结构，所以先进后出
	/*defer fmt.Println("--- 别看我排在第一行，但是我最后执行 ---")
	defer fmt.Println("--- 别看我排在第二行，但是我倒数第二执行 ---")

	fmt.Println("--- 我第一个执行 ---")
	fmt.Println("--- 我第二个执行 ---")*/

	defer deferOne()
	defer deferTwo()
	defer deferThree()
}

func deferOne() {
	fmt.Println("--- 第一行，最后执行 ----")
}

func deferTwo() {
	fmt.Println("--- 第二行，倒数第二执行 ----")
}

func deferThree() {
	fmt.Println("--- 第三行，倒数第三执行 ----")
}
