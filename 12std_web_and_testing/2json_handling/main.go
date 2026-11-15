package main

import "fmt"

func main() {
	// JSON Marshal, Unmarshal, and Struct Tags in Go
	// - json.Marshal converts Go structs to JSON bytes; json.Unmarshal decodes JSON bytes into Go values.
	// - Struct tags (e.g. `json:"name,omitempty"`) customize JSON field serialization.
	// - json.Decoder and json.Encoder stream JSON directly to/from io.Reader and io.Writer.
	fmt.Println("JSON Marshal, Unmarshal, and Struct Tags in Go")
}
