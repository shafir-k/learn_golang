package main

import "fmt"

func main() {
	// Strings, Runes, and Bytes in Go
	// - A string is a read-only slice of bytes ([]byte).
	// - A 'rune' is an alias for int32, representing a Unicode code point.
	// - Iterating with a standard index loops over raw bytes; 'for ... range' decodes UTF-8 runes.
	fmt.Println("Strings, Runes, and Bytes in Go")
}
