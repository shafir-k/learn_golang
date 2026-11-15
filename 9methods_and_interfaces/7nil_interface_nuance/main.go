package main

import "fmt"

func main() {
	// The Nil Interface Nuance in Go
	// - An interface value consists of two components: dynamic type and dynamic value (type, value).
	// - An interface is nil ONLY if both its type and value are nil.
	// - An interface holding a typed nil pointer (e.g., (*MyStruct, nil)) is NOT nil!
	fmt.Println("The Nil Interface Nuance in Go")
}
