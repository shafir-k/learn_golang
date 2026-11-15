package main

import "fmt"

func main() {
	// Preventing Goroutine Leaks in Go
	// - A goroutine leak occurs when a goroutine remains blocked indefinitely (e.g. sending to unread channel).
	// - Leaked goroutines remain in memory and never get garbage collected.
	// - Standard mitigation: provide a cancellation signal via a 'done' channel (<-chan struct{}) or context.Context.
	fmt.Println("Preventing Goroutine Leaks in Go")
}
