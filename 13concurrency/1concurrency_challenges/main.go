package main

import "fmt"

func main() {
	// Concurrency Challenges: Deadlocks, Livelocks, and Starvation
	// - Race condition: multiple goroutines access shared memory concurrently with at least one write.
	// - Deadlock: two or more goroutines are waiting on each other indefinitely.
	// - Starvation: a goroutine is perpetually denied necessary resources to make progress.
	// - Go's philosophy: 'Do not communicate by sharing memory; instead, share memory by communicating.'
	fmt.Println("Concurrency Challenges: Deadlocks, Livelocks, and Starvation")
}
