package main

import "fmt"

func main() {
	// Go's Call-By-Value and the Slices/Maps Mystery
	// - Go is strictly call-by-value: everything is passed as a copy.
	// - Slices, maps, and channels contain internal pointers to backing data structures.
	// - When a slice header is copied by value, both copies point to the same underlying array.
	fmt.Println("Go's Call-By-Value and the Slices/Maps Mystery")
}
