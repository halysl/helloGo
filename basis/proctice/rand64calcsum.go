package main

import (
	"context"
	"fmt"
	"math/rand"
)

func randInt64(ctx context.Context, c chan<- uint64) {
	for {
		c <- rand.Uint64()
	}
}

func complexCalc(ctx context.Context, out chan<- int, in <-chan uint64) {
	defer wg.Done()
	s := <- in
	sum := uint64(0)
	for ;s != 0; {
        sum += s % 10
        s /= 10
	}
	out <- int(sum)
}

func rand64CalcSum() {
	ch1 := make(chan uint64)
	ch2 := make(chan int, 24)
	ctx, cancel := context.WithCancel(context.Background())
	go randInt64(ctx, ch1)
	for i:= 0; i<24; i++ {
		wg.Add(1)
		go complexCalc(ctx, ch2, ch1)
	}
	wg.Wait()
	close(ch2)
	for i := range ch2 {
		fmt.Println(i)
	}
	cancel()
}
