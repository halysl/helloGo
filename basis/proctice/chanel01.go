package main

import (
	"fmt"
	"time"
)

func recv(c chan int) {
	for ret := range c {
	fmt.Println("接收成功：", ret)
	}
}

func useBufferChan(c1, c2 chan int) {
	for {
		if i, ok := <- c2;!ok{
			break
		} else {
            fmt.Println("i, ok 接受成功：", i)
		}
	}
	for i := range c1 {
		fmt.Println("range 接受成功：", i)
	}
}

func channelTest() {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	ch <- 10
	fmt.Println("发送成功")
}

func bufferChannelTest() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		ch1 <- i
		ch2 <- i * i
	}
	go useBufferChan(ch1, ch2)
	time.Sleep(time.Second)
}