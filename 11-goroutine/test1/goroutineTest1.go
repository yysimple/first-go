package main

import (
	"fmt"
	"time"
)

func newTask() {
	start := 0
	for {
		start++
		fmt.Println("new GoRoutine: start ===> ", start)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 这里就是使用了 go协程，就可以支持并发执行了
	go newTask()

	i := 0
	for {
		i++
		fmt.Println("main GoRoutine: i ===> ", i)
		time.Sleep(1 * time.Second)
	}
}
