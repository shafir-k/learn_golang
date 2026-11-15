package main

import "fmt"

func main() {
	// Comparing Errors with errors.Is and errors.As
	// - errors.Is(err, target) checks if any error in the wrap chain matches target.
	// - errors.As(err, &customTarget) finds the first error in the chain that matches customTarget's type.
	// - Never compare wrapped errors using '==' directly; always prefer errors.Is/errors.As.
	fmt.Println("Comparing Errors with errors.Is and errors.As")
}
