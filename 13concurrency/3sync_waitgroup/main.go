package main

import "fmt"

func main() {
	// Synchronizing Goroutines with sync.WaitGroup
	// - sync.WaitGroup waits for a collection of goroutines to finish.
	// - Add(n): increments the counter.
	// - Done(): decrements the counter (typically deferred inside goroutine).
	// - Wait(): blocks until the counter reaches zero.
	fmt.Println("Synchronizing Goroutines with sync.WaitGroup")
}
