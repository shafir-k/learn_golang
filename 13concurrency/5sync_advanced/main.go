package main

import "fmt"

func main() {
	// Advanced Synchronization: sync.Cond, sync.Once, and sync.Pool
	// - sync.Once guarantees a function runs exactly once across multiple concurrent callers.
	// - sync.Cond coordinates goroutines waiting for a specific arbitrary condition (Wait, Signal, Broadcast).
	// - sync.Pool provides an efficient, thread-safe cache of reusable objects to reduce GC pressure.
	fmt.Println("Advanced Synchronization: sync.Cond, sync.Once, and sync.Pool")
}
