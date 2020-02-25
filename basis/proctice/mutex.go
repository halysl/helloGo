package main

import (
	"fmt"
	"sync"
)

var x int64
var lock sync.Mutex

func add() {
	for i:= 0; i< 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func mutexTest() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}