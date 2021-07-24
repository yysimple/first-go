package main

import (
	"fmt"
	"time"
)

func main() {
	bufferChannel := make(chan int, 3)

	fmt.Println("len =", len(bufferChannel), "cap =", cap(bufferChannel))

	/**
		这里有几种场景需要说下：
		1. 当子go程放入 3 个数据、main接收3个数据的时候，子go程可提前结束，两秒之后，main读取完之后，也会正常结束，不会发生阻塞情况
		2. 子go，存入4个值、main只接收3个值时，这种情况下会发生阻塞，而且程序执行结果会有很多种，但是最后的 len 一定会是 1，因为
	还有一个值一直无法被接收
		3. 子go程为3个值，main接收4个值，这种情况会出现异常，死锁 fatal error: all goroutines are asleep - deadlock!
	这是因为main一直在等待第四个值的出现，但是程序中已经没有子go能提供新的值了，就会出现死锁；
		4. 子go为4，main为4的情况下，这种情况下一定会发生阻塞的，子go和main都有可能发生阻塞，这种情况下程序也会出现很多中情况，
	但是最后的len一定是 0，数据都会被接收的

	-- 总结：
		1. 当通道满了，就会发生阻塞；
		2. 当通道是空的，读数据也会发生阻塞；
	*/
	go func() {
		defer fmt.Println("子go程结束...")
		for i := 0; i < 4; i++ {
			bufferChannel <- i
			fmt.Println("子go程正在运行,往 bufferChannel 中存入值 ===> ", i)
		}
	}()

	// 这里睡眠2s可以看到，因为使用的是缓存通道，所以子go程不需要阻塞
	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-bufferChannel
		fmt.Println("从缓存管道里读出来的数据：===> ", num)
	}
	fmt.Println("len =", len(bufferChannel), "cap =", cap(bufferChannel))

	fmt.Println("main GoRoutine finish...")
}
