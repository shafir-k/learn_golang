package main

import "fmt"

func main() {
	// Value vs. Pointer Method Receivers in Go
	// - Value receiver: operates on a copy of the caller; mutations do not affect the original.
	// - Pointer receiver: operates directly on the caller; can mutate fields and avoids copying.
	// - Rule of thumb: if any method on a type needs a pointer receiver, make all methods use pointer receivers.
	fmt.Println("Value vs. Pointer Method Receivers in Go")
}
