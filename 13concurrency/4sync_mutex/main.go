package main

import "fmt"

func main() {
	// Mutual Exclusion with sync.Mutex and sync.RWMutex
	// - sync.Mutex provides Lock() and Unlock() to guard critical sections and shared memory.
	// - sync.RWMutex allows multiple readers simultaneously (RLock/RUnlock) but exclusive writers (Lock/Unlock).
	// - Always defer mu.Unlock() right after acquiring a lock.
	fmt.Println("Mutual Exclusion with sync.Mutex and sync.RWMutex")
}
