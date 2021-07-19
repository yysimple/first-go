// 主函数第一行的包名一定是 main 跟文件夹无关
package main

/**
import "fmt"
import "time"
*/
import (
	"fmt"
	"time"
)

// 主函数的 { 一定要跟函数名在一行
func main() {
	// 打印， 最后面的 ; 可加可不加
	fmt.Println("Hello, World!")
	time.Sleep(1 * time.Second)
}
