package main

import "fmt"

func main() {
	// Packages, Exports, and Imports in Go
	// - Every Go file belongs to a package declared on the first line.
	// - Capitalized identifiers are public (exported); lowercase identifiers are package-private.
	// - Package alias imports allow resolving name conflicts: import mypkg "path/to/pkg".
	fmt.Println("Packages, Exports, and Imports in Go")
}
