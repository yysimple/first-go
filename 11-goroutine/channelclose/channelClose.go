package main

import "fmt"

func main() {
	channelClose := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			channelClose <- i
		}
		// 这里是关闭通道，关闭通道后，下面的 ok 就会变为false
		close(channelClose)
	}()

	/*for {
		// 这里往通道里面一直读数据，只要通道存在，ok一直为true
		if data, ok := <-channelClose; ok {
			fmt.Println("data =", data)
		}else {
			break
		}
	}*/

	// range 就会一直阻塞的去通道中读数据
	for data := range channelClose {
		fmt.Println("for range data ===> ", data)
	}

}
