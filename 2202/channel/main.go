package main

import (
	"fmt"
	"time"
)

/**
通道有发送（send）、接收(receive）和关闭（close）三种操作:
	chan T          // 可以接收和发送类型为 T 的数据
	chan<- float64  // 只可以用来发送 float64 类型的数据
	<-chan int      // 只可以用来接收 int 类型的数据
	close(ch)       // 关闭通道

<-总是优先和最左边的类型结合：
	chan<- chan int    // 等价 chan<- (chan int)
	chan<- <-chan int  // 等价 chan<- (<-chan int)

关闭后的通道有以下特点：
    1.对一个关闭的通道再发送值就会导致panic。
    2.对一个关闭的通道进行接收会一直获取值直到通道为空。
    3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
    4.关闭一个已经关闭的通道会导致panic。
*/

func worker(done chan bool) {
	time.Sleep(time.Second * 2)
	// 通知任务已完成
	done <- true
}
func main() {
	done := make(chan bool, 1)
	go worker(done)
	// 等待任务完成
	<-done
	fmt.Println("finish")
}
