package main

import "fmt"

func main() {
	// Empty Interface, Type Assertions, and Type Switches
	// - The empty interface 'interface{}' (or alias 'any') can hold values of any type.
	// - Type assertion: val, ok := i.(ConcreteType) retrieves the underlying concrete value.
	// - Type switch: switch v := i.(type) executes branches based on the dynamic type.
	fmt.Println("Empty Interface, Type Assertions, and Type Switches")
}
