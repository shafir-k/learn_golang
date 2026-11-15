package main

import "fmt"

func main() {
	//Float32 is a primitive data type in Golang
	//Float32 zero value is 0.0
	//Float32 can hold values from 1.7e-38 to 1.7e+38
	//Float32 is a signed floating-point type that uses 32 bits of memory to store values. It can represent both positive and negative numbers, as well as zero. The range of values that can be stored in a float32 variable is from 1.7e-38 to 1.7e+38, inclusive. This means that the smallest value that can be stored in a float32 variable is 1.7e-38, and the largest value that can be stored is 1.7e+38.
	//Float32 is often used in situations where memory usage is a concern, such as in embedded systems or when working with large datasets. However, it is important to be aware of the limitations of float32 and to use larger floating-point types when necessary to avoid precision errors.
	var myFloat32 float32
	fmt.Println(myFloat32)
	myFloat32 = 3.14
	fmt.Println(myFloat32)

	//Float64 is a primitive data type in Golang
	//Float64 zero value is 0.0
	//Float64 can hold values from 5.0e-324 to 1.7e+308
	//Float64 is a signed floating-point type that uses 64 bits of memory to store values. It can represent both positive and negative numbers, as well as zero. The range of values that can be stored in a float64 variable is from 5.0e-324 to 1.7e+308, inclusive. This means that the smallest value that can be stored in a float64 variable is 5.0e-324, and the largest value that can be stored is 1.7e+308.
	//Float64 is often used in situations where precision is important, such as in scientific computing or when working with financial data. However, it is important to be aware of the limitations of float64 and to use larger floating-point types when necessary to avoid precision errors.
	var myFloat64 float64
	fmt.Println(myFloat64)
	myFloat64 = 3.141592653589793
	fmt.Println(myFloat64)

}
