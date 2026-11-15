package main

import "fmt"

func main() {
	// Concurrency Patterns: Pipelines, Fan-In, and Fan-Out
	// - Pipelines: series of stages connected by channels where each stage is a group of goroutines.
	// - Fan-Out: multiple goroutines read from the same channel to parallelize CPU or I/O work.
	// - Fan-In: combining multiple output channels into a single multiplexed channel.
	fmt.Println("Concurrency Patterns: Pipelines, Fan-In, and Fan-Out")
}
