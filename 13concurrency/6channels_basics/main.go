package main

import "fmt"

func main() {
	// Channels in Go: Unbuffered, Buffered, and Directional
	// - Channels are typed conduits for sending and receiving values between goroutines: make(chan T).
	// - Unbuffered channel: synchronous handshake (send blocks until receive, receive blocks until send).
	// - Buffered channel: asynchronous up to capacity (blocks only when buffer is full).
	// - Directional channels: send-only (chan<- T) and receive-only (<-chan T).
	fmt.Println("Channels in Go: Unbuffered, Buffered, and Directional")
}
