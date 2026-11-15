package main

import "fmt"

func main() {
	// The Go Scheduler and Goroutines
	// - A goroutine is an ultra-lightweight thread managed by the Go runtime (starting at ~2KB stack).
	// - Launched simply using the 'go' keyword before a function call.
	// - Go runtime scheduler uses the GMP model (Goroutines, Machines/OS threads, Processors/logical cores).
	fmt.Println("The Go Scheduler and Goroutines")
}
