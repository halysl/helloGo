package main

import (
	"fmt"
	"runtime"
	"time"
)

func printNum(pre string, num int) {
	fmt.Printf("%s: %d\n", pre, num)
}


func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func oneProcessTest() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println()
}

func defaultProcessTest() {
	go a()
	go b()
	time.Sleep(time.Second)
}