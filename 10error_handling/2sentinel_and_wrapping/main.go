package main

import "fmt"

func main() {
	// Sentinel Errors and Error Wrapping/Unwrapping in Go
	// - Sentinel errors: exported package-level constants/vars representing known failures (e.g. io.EOF).
	// - Error wrapping: fmt.Errorf("context: %w", err) wraps errors with context while preserving root causes.
	// - errors.Unwrap(err) extracts the underlying cause.
	fmt.Println("Sentinel Errors and Error Wrapping/Unwrapping in Go")
}
