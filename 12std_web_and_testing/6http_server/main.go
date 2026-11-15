package main

import "fmt"

func main() {
	// Building HTTP Servers and Routing in Go (Go 1.22+)
	// - http.HandleFunc registers route handlers; http.ListenAndServe starts the server.
	// - Go 1.22 enhanced ServeMux supports method matching ("GET /items/{id}") and wildcard parameters.
	// - r.PathValue("param") extracts URL parameters cleanly without external routers.
	fmt.Println("Building HTTP Servers and Routing in Go (Go 1.22+)")
}
