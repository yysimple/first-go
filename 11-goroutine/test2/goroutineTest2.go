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

	// 这里写个死循环，类似于项目跑起来后，主线程一直在
	for {
		time.Sleep(2 * time.Second)

	}
}
