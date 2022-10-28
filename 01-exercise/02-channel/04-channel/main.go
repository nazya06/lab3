package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction
func ping(out chan<- string) {
	// send message on ch1
	out <- "ping"
}

func pong(in <-chan string, out chan<- string) {
	// recv message on ch1
	message := <-in
	message = message + " pong"
	// send it on ch2
	out <- message
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine ping and pong
	go ping(ch1)
	go pong(ch1, ch2)

	// recv message on ch2
	message := <-ch2

	fmt.Println(message)
}
