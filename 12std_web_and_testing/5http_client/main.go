package main

import "fmt"

func main() {
	// Making HTTP Requests with Go's HTTP Client
	// - http.Get and http.Post for simple HTTP calls; http.Client for custom timeouts and configurations.
	// - Always close resp.Body (e.g. defer resp.Body.Close()) to prevent resource leaks.
	// - Use http.NewRequestWithContext to attach cancellation and deadlines.
	fmt.Println("Making HTTP Requests with Go's HTTP Client")
}
