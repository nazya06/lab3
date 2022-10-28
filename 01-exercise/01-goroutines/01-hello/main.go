package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("function call")

	// goroutine with anonymous function
	go func() {
		fun("anonymous function")
	}()

	// goroutine with function value call
	f := fun
	go f("function value call")

	// wait for goroutines to end
	fmt.Println("waiting for goroutines..")
	time.Sleep(1 * time.Second)

	fmt.Println("done..")
}
