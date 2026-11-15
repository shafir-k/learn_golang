package main

import "fmt"

func main() {
	// Code Blocks and Variable Shadowing in Go
	// - A variable declared in an inner block shadows a variable with the same name in an outer block.
	// - Shadowing with ':=' is a common source of subtle bugs in Go.
	// - Tools like 'govet -shadow' help detect accidental shadowing.
	fmt.Println("Code Blocks and Variable Shadowing in Go")
}
