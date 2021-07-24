package main

import "fmt"

func main() {
	// 定义一个管道,存放int类型，类似于 java中的 ThreadLocal
	channel := make(chan int)

	go func() {
		defer fmt.Println("goroutine..程序结束..")
		fmt.Println("goroutine..正在运行..")
		// 表示往 channel中塞入 666， 可供其他协程使用
		channel <- 666
	}()

	// 这里是从管道中取出数据，这里注意的是 <-和channel之间不能有空格
	num := <-channel
	fmt.Println("num =", num)
	fmt.Println("main goroutine 结束...")

	/**
	其实这里会发现一个问题，就是最有总是 main函数结束，按理来说应该两者都有可能
	这是因为：
	1. 当main函数已经执行到管道取值处，即 num := <-channel，但是里面为空，就开始阻塞
	2. 而子进程则是把 值放到channel中channel <- 666，也是需要等待的，等main取出才会结束
	3. 因为子进程后续没有逻辑了，所以一般会优于main，提前结束
	*/
}
