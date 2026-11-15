package main

import "fmt"

func main() {
	// Panic and Recover in Go
	// - panic() terminates normal execution, unwinds the stack, and runs deferred functions.
	// - recover() captures a panic value inside a deferred function, preventing program crash.
	// - Panics should be reserved for truly exceptional, unrecoverable programmer errors.
	fmt.Println("Panic and Recover in Go")
}
