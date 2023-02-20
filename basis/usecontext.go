package basis

import (
	"context"
	"fmt"
	"time"
)

func UseContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			duration := 500*time.Millisecond
			select {
			case <-ctx.Done():
				fmt.Println("handle", ctx.Err())
				return
			case <-time.After(duration):
				fmt.Println("process request with", duration)
			}
		}
	}()

	time.Sleep(1601*time.Millisecond)
	cancel()
	time.Sleep(601*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {

}
