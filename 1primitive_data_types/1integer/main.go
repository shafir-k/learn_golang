package main

import "fmt"

func main() {

	//Int is a primitive data type in Golang
	//Int zero value is 0
	//Int can hold values from -2147483648 to 2147483647
	//Int is a signed integer type that uses 32 bits of memory to store values. It can represent both positive and negative numbers, as well as zero. The range of values that can be stored in an int variable is from -2147483648 to 2147483647, inclusive. This means that the smallest value that can be stored in an int variable is -2147483648, and the largest value that can be stored is 2147483647.
	//Int is often used in situations where memory usage is not a concern, such as in general-purpose programming or when working with small arrays of integers. However, it is important to be aware of the limitations of int and to use larger integer types when necessary to avoid overflow or underflow errors.
	var myInt int
	fmt.Println(myInt)
	myInt = 42
	fmt.Println(myInt)

	//Int8 is a primitive data type in Golang
	//Int8 zero value is 0
	//Int8 can hold values from -128 to 127
	//Int8 is a signed integer type that uses 8 bits of memory to store values. It can represent both positive and negative numbers, as well as zero. The range of values that can be stored in an int8 variable is from -128 to 127, inclusive. This means that the smallest value that can be stored in an int8 variable is -128, and the largest value that can be stored is 127.
	//Int8 is often used in situations where memory usage is a concern, such as in embedded systems or when working with large arrays of integers. However, it is important to be aware of the limitations of int8 and to use larger integer types when necessary to avoid overflow or underflow errors.
	var myInt8 int8
	fmt.Println(myInt8)
	myInt8 = 127
	fmt.Println(myInt8)

	//Int16 is a primitive data type in Golang
	//Int16 zero value is 0
	//Int16 can hold values from -32768 to 32767
	//Int16 is a signed integer type that uses 16 bits of memory to store values. It can represent both positive and negative numbers, as well as zero. The range of values that can be stored in an int16 variable is from -32768 to 32767, inclusive. This means that the smallest value that can be stored in an int16 variable is -32768, and the largest value that can be stored is 32767.
	//Int16 is often used in situations where memory usage is a concern, such as in embedded systems or when working with large arrays of integers. However, it is important to be aware of the limitations of int16 and to use larger integer types when necessary to avoid overflow or underflow errors.
	var myInt16 int16
	fmt.Println(myInt16)
	myInt16 = 32767
	fmt.Println(myInt16)

	//Int32 is a primitive data type in Golang
	//Int32 zero value is 0
	//Int32 can hold values from -2147483648 to 2147483647
	//Int32 is a signed integer type that uses 32 bits of memory to store values. It can represent both positive and negative numbers, as well as zero. The range of values that can be stored in an int32 variable is from -2147483648 to 2147483647, inclusive. This means that the smallest value that can be stored in an int32 variable is -2147483648, and the largest value that can be stored is 2147483647.
	//Int32 is often used in situations where memory usage is not a concern, such as in general-purpose programming or when working with small arrays of integers. However, it is important to be aware of the limitations of int32 and to use larger integer types when necessary to avoid overflow or underflow errors.
	var myInt32 int32
	fmt.Println(myInt32)
	myInt32 = 2147483647
	fmt.Println(myInt32)

	//Int64 is a primitive data type in Golang
	//Int64 zero value is 0
	//Int64 can hold values from -9223372036854775808 to 9223372036854775807
	//Int64 is a signed integer type that uses 64 bits of memory to store values. It can represent both positive and negative numbers, as well as zero. The range of values that can be stored in an int64 variable is from -9223372036854775808 to 9223372036854775807, inclusive. This means that the smallest value that can be stored in an int64 variable is -9223372036854775808, and the largest value that can be stored is 9223372036854775807.
	//Int64 is often used in situations where memory usage is not a concern, such as in general-purpose programming or when working with small arrays of integers. However, it is important to be aware of the limitations of int64 and to use larger integer types when necessary to avoid overflow or underflow errors.
	var myInt64 int64
	fmt.Println(myInt64)
	myInt64 = 9223372036854775807
	fmt.Println(myInt64)

	//Uint is a primitive data type in Golang
	//Uint zero value is 0
	//Uint can hold values from 0 to 4294967295
	//Uint is an unsigned integer type that uses 32 bits of memory to store values. It can represent only non-negative numbers, including zero. The range of values that can be stored in a uint variable is from 0 to 4294967295, inclusive. This means that the smallest value that can be stored in a uint variable is 0, and the largest value that can be stored is 4294967295.
	//Uint is often used in situations where memory usage is not a concern, such as in general-purpose programming or when working with small arrays of integers. However, it is important to be aware of the limitations of uint and to use larger integer types when necessary to avoid overflow errors.
	var myUint uint
	fmt.Println(myUint)
	myUint = 4294967295
	fmt.Println(myUint)

	//Uint8 is a primitive data type in Golang
	//Uint8 zero value is 0
	//Uint8 can hold values from 0 to 255
	//Uint8 is an unsigned integer type that uses 8 bits of memory to store values. It can represent only non-negative numbers, including zero. The range of values that can be stored in a uint8 variable is from 0 to 255, inclusive. This means that the smallest value that can be stored in a uint8 variable is 0, and the largest value that can be stored is 255.
	//Uint8 is often used in situations where memory usage is a concern, such as in embedded systems or when working with large arrays of integers. However, it is important to be aware of the limitations of uint8 and to use larger integer types when necessary to avoid overflow errors.
	var myUint8 uint8
	fmt.Println(myUint8)
	myUint8 = 255
	fmt.Println(myUint8)

	//Uint16 is a primitive data type in Golang
	//Uint16 zero value is 0
	//Uint16 can hold values from 0 to 65535
	//Uint16 is a unsigned integer type that uses 16 bits of memory to store values. It can represent only non-negative numbers, including zero. The range of values that can be stored in a uint16 variable is from 0 to 65535, inclusive. This means that the smallest value that can be stored in a uint16 variable is 0, and the largest value that can be stored is 65535.
	//Uint16 is often used in situations where memory usage is a concern, such as in embedded systems or when working with large arrays of integers. However, it is important to be aware of the limitations of uint16 and to use larger integer types when necessary to avoid overflow errors.
	var myUint16 uint16
	fmt.Println(myUint16)
	myUint16 = 65535
	fmt.Println(myUint16)

	//Uint32 is a primitive data type in Golang
	//Uint32 zero value is 0
	//Uint32 can hold values from 0 to 4294967295
	//Uint32 is a unsigned integer type that uses 32 bits of memory to store values. It can represent only non-negative numbers, including zero. The range of values that can be stored in a uint32 variable is from 0 to 4294967295, inclusive. This means that the smallest value that can be stored in a uint32 variable is 0, and the largest value that can be stored is 4294967295.
	//Uint32 is often used in situations where memory usage is not a concern, such as in general-purpose programming or when working with small arrays of integers. However, it is important to be aware of the limitations of uint32 and to use larger integer types when necessary to avoid overflow errors.
	var myUint32 uint32
	fmt.Println(myUint32)
	myUint32 = 4294967295
	fmt.Println(myUint32)

	//Uint64 is a primitive data type in Golang
	//Uint64 zero value is 0
	//Uint64 can hold values from 0 to 18446744073709551615
	//Uint64 is a unsigned integer type that uses 64 bits of memory to store values. It can represent only non-negative numbers, including zero. The range of values that can be stored in a uint64 variable is from 0 to 18446744073709551615, inclusive. This means that the smallest value that can be stored in a uint64 variable is 0, and the largest value that can be stored is 18446744073709551615.
	//Uint64 is often used in situations where memory usage is not a concern, such as in general-purpose programming or when working with small arrays of integers. However, it is important to be aware of the limitations of uint64 and to use larger integer types when necessary to avoid overflow errors.
	var myUint64 uint64
	fmt.Println(myUint64)
	myUint64 = 18446744073709551615
	fmt.Println(myUint64)

	//rune is a primitive data type in Golang , alias for int32
	//rune zero value is 0
	//rune can hold values from -2147483648 to 2147483647
	//rune is a signed integer type that uses 32 bits of memory to store values. It can represent both positive and negative numbers, as well as zero. The range of values that can be stored in a rune variable is from -2147483648 to 2147483647, inclusive. This means that the smallest value that can be stored in a rune variable is -2147483648, and the largest value that can be stored is 2147483647.
	//rune is often used in situations where memory usage is not a concern, such as in general-purpose programming or when working with small arrays of integers. However, it is important to be aware of the limitations of rune and to use larger integer types when necessary to avoid overflow or underflow errors.
	var myRune rune
	fmt.Println(myRune)
	myRune = 2147483647
	fmt.Println(myRune)

	//byte is a primitive data type in Golang , alias for uint8
	//byte zero value is 0
	//byte can hold values from 0 to 255
	//byte is an unsigned integer type that uses 8 bits of memory to store values. It can represent only non-negative numbers, including zero. The range of values that can be stored in a byte variable is from 0 to 255, inclusive. This means that the smallest value that can be stored in a byte variable is 0, and the largest value that can be stored is 255.
	//byte is often used in situations where memory usage is a concern, such as in embedded systems or when working with large arrays of integers. However, it is important to be aware of the limitations of byte and to use larger integer types when necessary to avoid overflow errors.
	var myByte byte
	fmt.Println(myByte)
	myByte = 255
	fmt.Println(myByte)
}
