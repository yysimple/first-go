package main

// 工程 first-go 需要放在bin目录下面的 xxx/.../src的下面
import (
	// 这个是匿名包的使用，一般用于 调用该 类里面的 init() 方法
	_ "first-go/4-import/anonymous"
	"first-go/4-import/lib1"
	// 这是为导进来的包加上名字
	haveName "first-go/4-import/unanonymous"
	// 这里的包如果导入了，但是不使用是会报错的
	"first-go/4-import/lib2"
	// 引入 context的上下文环境, 不需要在调用方法前面加任何调用者
	. "first-go/4-import/context"

	"fmt"
)

func main() {
	lib2.JarTest2()
	lib1.JarTest1()
	// 这里有个注意的点，方法名必须是 大写字母开头，才能使用，类似于 java 的 public
	// anonymous.anonymousTest()

	// 这是为导进来的包加上名字
	haveName.UnAnonymousTest()

	Context()

}

/**
 	1. init(),方法是在其他方法之前执行的
	2. 根据导入包的顺序执行，这里 lib1、lib2，两者按导入包的顺序执行init(),然后在执行main的init方法
*/
func init() {
	fmt.Println("------ 我是 main 的 init() 方法 ------")
}
