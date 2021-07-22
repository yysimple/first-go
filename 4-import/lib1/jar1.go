package lib1

import "fmt"

func JarTest1() {
	fmt.Println("------ 我是 jar-1 ------")
}

// init() 方法会优先执行
func init() {
	fmt.Println("------ 我是 jar-1 的 init() 方法 ------")
}
