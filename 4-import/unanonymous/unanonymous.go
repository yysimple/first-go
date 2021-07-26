package unanonymous

import "fmt"

func UnAnonymousTest() {
	fmt.Println("------ 我是有名字的方法 ------")
}

func init() {
	fmt.Println("------ 我是有名字的 init() 方法 ------")
}
