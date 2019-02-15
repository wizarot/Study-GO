package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	// 多个channels
	var channels [10]chan<- int
	// 建立 10个wroker,每个wroker监听一个channel
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
		// go worker(i, channels[i])
	}
	// 分发数据
	for j := 0; j < 10; j++ {
		// 每个channel中发送一个字母过去
		channels[j] <- 'a' + j
	}

	// 再来一遍分发
	for j := 0; j < 10; j++ {
		// 每个channel中发送一个字母过去
		channels[j] <- 'A' + j
	}

	time.Sleep(time.Millisecond * 2)
}

func createWorker(id int) chan<- int { //返回的channel是用来接收数据到channel里面的
	// 返回一个channel , 创建一个channel
	c := make(chan int)
	// 中间要使用一个协程
	// 创建一个go协程
	go func() {
		// 真正处理事情,是在这个协程里面
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	// 返回channel
	return c
}

func main() {
	chanDemo()
}
