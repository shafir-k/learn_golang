package main

import "fmt"

func main() {
	//String is a primitive data type in Golang
	//String zero value is an empty string ""
	//Strings are immutable sequences of bytes (typically UTF-8 encoded)
	var myString string
	fmt.Println(myString)
	myString = "Hello Golang"
	fmt.Println(myString)
}
