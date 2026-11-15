package main

import "fmt"

func main() {
	// io.Reader and io.Writer in Go
	// - io.Reader: Read(p []byte) (n int, err error) — stream input abstraction.
	// - io.Writer: Write(p []byte) (n int, err error) — stream output abstraction.
	// - io.Copy transfers data directly from any Reader to any Writer efficiently.
	fmt.Println("io.Reader and io.Writer in Go")
}
