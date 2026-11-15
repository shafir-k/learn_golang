package main

import "fmt"

func main() {
	// Struct Embedding and Field Promotion in Go
	// - Go favors composition over inheritance using embedded (anonymous) fields.
	// - Fields and methods of embedded structs are 'promoted' to the outer struct.
	// - Outer struct can override promoted fields or methods while retaining access to the inner ones.
	fmt.Println("Struct Embedding and Field Promotion in Go")
}
