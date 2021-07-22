package anonymous

import "fmt"

func AnonymousTest() {
	fmt.Println("------ 我是匿名方法 ------")
}

func init() {
	fmt.Println("------ 我是匿名 init() 方法 ------")
}
