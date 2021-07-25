package main

import (
	"fmt"
	"time"
)

func mutil(channelRead, channelStop chan int) {
	var x, y = 1, 1
	for {
		select {
		/**
		这里会一直向 通道中加入数据
		1. 当 读取操作正在阻塞时，这里会 先塞入 case channelRead <- x = 1 到通道，然后对应 for 0 读出来的x就是1
		2. 当第二次循环，for 1 时，因为通道里面没有数据，所以又阻塞，等待 x输入到通道，
		此时 x = y 然后再 case channelRead <- x = 1，此时读出来还是 1
		3. 第三就开始累加，当最后 for 5 读取数据之后，便不会再读取数据，而case里面也不会去存入数据，这个时候 for i执行完
		之后，想 channelStop 中写入任意值，然后 <-channelStop 这个就可以读出数据，执行该case的逻辑，退出程序
		*/
		case channelRead <- x:
			x = y
			y = x + y
			//fmt.Println("case 1 x ==>", x)
			time.Sleep(1 * time.Second)
		case <-channelStop:
			fmt.Println("stop")
			return
		}
	}
}

func main() {
	channelRead := make(chan int)
	channelStop := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			// 这里在执行的时候会进入阻塞状态的
			fmt.Println("管道可读 ==> ", <-channelRead)
		}

		time.Sleep(2 * time.Second)

		// 读完上面那个管道数据之后，将随意一个值放入到此通道中，使得case <-channelStop 这里可以读出数据，执行者这里的逻辑
		channelStop <- 10
	}()

	// 因为上面循环时会处于阻塞状态，所以，这里会同步执行，然后在根据不同的case做不同的逻辑
	mutil(channelRead, channelStop)
}
