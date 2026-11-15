package main

import "fmt"

func main() {
	// Advanced Maps & Comma-OK Idiom in Go
	// - Looking up a non-existent key returns the value type's zero value.
	// - Comma-ok idiom: val, exists := myMap[key] to safely check if a key is present.
	// - Iterating over maps with 'for k, v := range myMap' yields keys in random order.
	fmt.Println("Advanced Maps & Comma-OK Idiom in Go")
}
