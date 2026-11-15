package main

import "fmt"

func main() {
	// The Init Function in Go
	// - init() functions run automatically when a package is initialized before main().
	// - Takes no arguments and returns nothing; a single file or package can have multiple init() functions.
	// - Blank import (import _ "pkg") runs a package's init() purely for side effects (e.g. database drivers).
	fmt.Println("The Init Function in Go")
}
