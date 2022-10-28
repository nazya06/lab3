package main

import (
	"context"
	"fmt"
)

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.
	generator := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			defer close(dst)
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	// Create a context that is cancellable.
	ctx, cancel := context.WithCancel(context.Background())

	ch := generator(ctx)
	for n := range ch {
		fmt.Println(n)
		if n == 5 {
			cancel()
		}
	}
}
