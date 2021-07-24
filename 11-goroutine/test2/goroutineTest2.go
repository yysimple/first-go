package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// go的协程直接通过 go声明即可，这种是 匿名函数形式声明
	go func() {
		defer fmt.Println("A.defer")
		// 这里的话先执行 B 然后在 B.defer，这是协程里面的匿名函数
		func() {
			defer fmt.Println("B.defer")
			// 这里return， B 就不会被执行
			// return
			// 这里是直接退出程序 所以只会打印 两个defer
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	// 这里是传入参数的协程，且有返回值的，但是返回值目前是无法接收的，需要配个channel一起实现
	go func(a, b int) bool {
		fmt.Println("a =", a, "b =", b)
		return true
	}(10, 20)

	// 这里写个死循环，类似于项目跑起来后，主线程一直在
	for {
		time.Sleep(2 * time.Second)

	}
}
