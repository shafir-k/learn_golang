package main

import "fmt"

func main() {
	// Custom JSON Marshalling and Unmarshalling in Go
	// - Types can implement json.Marshaler (MarshalJSON() ([]byte, error)).
	// - Types can implement json.Unmarshaler (UnmarshalJSON([]byte) error).
	// - Useful for custom date formats, dynamic typing, and hiding sensitive fields.
	fmt.Println("Custom JSON Marshalling and Unmarshalling in Go")
}
