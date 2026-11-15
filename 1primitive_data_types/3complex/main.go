package main

import "fmt"

func main() {
	// Complex numbers in Go: complex64 and complex128
	// Built-in functions: complex(real, imag), real(c), imag(c)
	var c1 complex64 = complex(3, 4)
	var c2 complex128 = complex(5, 12)
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("real:", real(c1), "imag:", imag(c1))
}
