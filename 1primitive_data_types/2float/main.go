package main

import "fmt"

func main() {
	//Float is a primitive data type in Golang
	//Float zero value is 0.0
	//Go provides float32 (32-bit IEEE-754) and float64 (64-bit IEEE-754)
	var myFloat float64
	fmt.Println(myFloat)
	myFloat = 3.14
	fmt.Println(myFloat)

	var myFloat32 float32 = 3.14159
	fmt.Println(myFloat32)
}
