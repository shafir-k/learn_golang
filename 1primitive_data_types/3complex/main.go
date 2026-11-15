package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println("================================================================================")
	fmt.Println("                        COMPREHENSIVE GOLANG COMPLEX NUMBERS GUIDE              ")
	fmt.Println("================================================================================")

	demonstrateBasicsAndMemory()
	demonstrateCreationAndComponents()
	demonstrateArithmeticOperations()
	demonstrateComparison()
	demonstrateConversionsAndParsing()
	demonstrateMathCmplxPackage()
	demonstrateFormatting()
	demonstratePracticalApplication()
}

// -----------------------------------------------------------------------------
// 1. BASICS & MEMORY CHARACTERISTICS
// -----------------------------------------------------------------------------
// - In Go, complex numbers are built-in primitive types representing numbers in
//   the form `a + bi`, where `a` is the real component, `b` is the imaginary
//   component, and `i` is the imaginary unit satisfying i² = -1.
//
// - Go provides two complex types:
//   1. `complex64` : Composed of two float32 values (float32 real + float32 imaginary).
//      - Size: 64 bits = 8 bytes.
//   2. `complex128`: Composed of two float64 values (float64 real + float64 imaginary).
//      - Size: 128 bits = 16 bytes.
//      - Default type for untyped complex literals (e.g. `c := 3 + 4i`).
//
// - ZERO VALUE:
//   - The zero value is `0 + 0i` (printed as `(0+0i)`).
//
// - MEMORY LAYOUT:
//   - Stored in memory as two contiguous IEEE-754 floating-point numbers.
//   - Real part is stored first, followed immediately by the imaginary part.
// -----------------------------------------------------------------------------
func demonstrateBasicsAndMemory() {
	fmt.Println("\n--- 1. BASICS & MEMORY CHARACTERISTICS ---")

	// Zero values
	var zero64 complex64
	var zero128 complex128
	fmt.Printf("Zero value of complex64:  %v\n", zero64)
	fmt.Printf("Zero value of complex128: %v\n", zero128)

	// Untyped complex literal defaults to complex128
	defaultComplex := 3 + 4i
	fmt.Printf("Type of '3 + 4i': %T\n", defaultComplex)

	// Memory sizes
	fmt.Printf("Size of complex64 in memory:  %2d bytes (two float32s: %d + %d bytes)\n",
		unsafe.Sizeof(zero64), unsafe.Sizeof(float32(0)), unsafe.Sizeof(float32(0)))
	fmt.Printf("Size of complex128 in memory: %2d bytes (two float64s: %d + %d bytes)\n",
		unsafe.Sizeof(zero128), unsafe.Sizeof(float64(0)), unsafe.Sizeof(float64(0)))
}

// -----------------------------------------------------------------------------
// 2. CREATION & COMPONENT EXTRACTION
// -----------------------------------------------------------------------------
// - Method 1: Imaginary Literals using the suffix `i`
//   - Examples: `3 + 4i`, `5.5 - 2.1i`, `1i`, `0i`
// - Method 2: The built-in `complex(real, imag)` function
//   - Both arguments MUST have matching floating-point types:
//     - `complex(float32, float32)` yields `complex64`.
//     - `complex(float64, float64)` yields `complex128`.
//
// - COMPONENT EXTRACTION:
//   - `real(c)`: extracts the real component.
//   - `imag(c)`: extracts the imaginary component.
//   - Return types match the constituent float types:
//     - If `c` is `complex64`, `real(c)` and `imag(c)` return `float32`.
//     - If `c` is `complex128`, `real(c)` and `imag(c)` return `float64`.
// -----------------------------------------------------------------------------
func demonstrateCreationAndComponents() {
	fmt.Println("\n--- 2. CREATION & COMPONENT EXTRACTION ---")

	// 1. Literal syntax with 'i'
	c1 := 5.0 + 12.0i
	fmt.Printf("Literal c1: %v (type: %T)\n", c1, c1)

	// Pure imaginary number
	pureImag := 7i
	fmt.Printf("Pure imaginary: %v (real: %v, imag: %v)\n", pureImag, real(pureImag), imag(pureImag))

	// 2. Built-in complex() function
	var r32, i32 float32 = 1.5, 2.5
	c64 := complex(r32, i32) // Returns complex64
	fmt.Printf("complex64 from float32s:  %v (type: %T)\n", c64, c64)

	var r64, i64 float64 = 10.25, -4.75
	c128 := complex(r64, i64) // Returns complex128
	fmt.Printf("complex128 from float64s: %v (type: %T)\n", c128, c128)

	// 3. Extracting components using real() and imag()
	fmt.Printf("c128 components -> real: %.2f (type %T), imag: %.2f (type %T)\n",
		real(c128), real(c128), imag(c128), imag(c128))
}

// -----------------------------------------------------------------------------
// 3. ARITHMETIC OPERATIONS
// -----------------------------------------------------------------------------
// Go natively supports all basic arithmetic operations on complex numbers:
// - Addition:       (a + bi) + (c + di) = (a + c) + (b + d)i
// - Subtraction:    (a + bi) - (c + di) = (a - c) + (b - d)i
// - Multiplication: (a + bi) * (c + di) = (ac - bd) + (ad + bc)i  (since i² = -1)
// - Division:       (a + bi) / (c + di) = ((ac + bd) + (bc - ad)i) / (c² + d²)
// - Unary Negation: -(a + bi) = (-a - bi)
//
// NOTE: Go is strongly typed; operands must have the exact same complex type.
// You cannot add a complex64 to a complex128 without explicit type conversion.
// -----------------------------------------------------------------------------
func demonstrateArithmeticOperations() {
	fmt.Println("\n--- 3. ARITHMETIC OPERATIONS ---")

	a := 3 + 2i
	b := 1 + 4i

	fmt.Printf("a = %v, b = %v\n", a, b)
	fmt.Printf("Addition       (a + b) = %v\n", a+b)
	fmt.Printf("Subtraction    (a - b) = %v\n", a-b)
	fmt.Printf("Multiplication (a * b) = %v\n", a*b)
	fmt.Printf("Division       (a / b) = %v\n", a/b)
	fmt.Printf("Unary Negation (-a)    = %v\n", -a)

	// Multiplying imaginary units: i * i = -1
	iVal := 0 + 1i
	fmt.Printf("i * i = %v (confirms i² = -1)\n", iVal*iVal)
}

// -----------------------------------------------------------------------------
// 4. COMPARISON & EQUALITY
// -----------------------------------------------------------------------------
// - Complex numbers support equality comparisons: `==` and `!=`.
//   - Two complex numbers are equal if and only if both their real parts AND
//     their imaginary parts are equal.
// - NO RELATIONAL OPERATORS:
//   - Complex numbers do NOT have a natural mathematical ordering.
//   - Therefore, `<`, `<=`, `>`, `>=` are NOT defined and cause compile errors!
//   - (e.g. `c1 < c2` is invalid code).
//   - If you need to compare complex numbers by magnitude, compare their absolute
//     values using `cmplx.Abs(c)`.
// -----------------------------------------------------------------------------
func demonstrateComparison() {
	fmt.Println("\n--- 4. COMPARISON & EQUALITY ---")

	c1 := 2 + 3i
	c2 := 2 + 3i
	c3 := 3 + 2i

	fmt.Printf("%v == %v: %t\n", c1, c2, c1 == c2)
	fmt.Printf("%v == %v: %t\n", c1, c3, c1 == c3)
	fmt.Printf("%v != %v: %t\n", c1, c3, c1 != c3)

	// Note on ordering:
	fmt.Println("Relational operators (<, <=, >, >=) are NOT supported for complex numbers.")
	fmt.Println("To order complex numbers, compare their magnitudes (cmplx.Abs).")
}

// -----------------------------------------------------------------------------
// 5. TYPE CONVERSIONS & STRING PARSING
// -----------------------------------------------------------------------------
// - Widening / Narrowing conversions:
//   - `complex128(c64)`: widens float32 parts to float64 without precision loss.
//   - `complex64(c128)`: narrows float64 parts to float32 (may lose precision).
// - String conversions using `strconv` standard library:
//   - `strconv.FormatComplex(c, fmt, prec, bitSize)`: formats complex to string.
//   - `strconv.ParseComplex(s, bitSize)`: parses strings into complex128.
//     Accepts formats like: "1+2i", "(3-4i)", "5i", "-2.5", "+1.2E-3i"
// -----------------------------------------------------------------------------
func demonstrateConversionsAndParsing() {
	fmt.Println("\n--- 5. TYPE CONVERSIONS & STRING PARSING ---")

	// 1. Casting between complex types
	var c128 complex128 = 1.123456789012345 + 9.876543210987654i
	c64 := complex64(c128) // Narrowing
	fmt.Printf("Original complex128: %.15f\n", c128)
	fmt.Printf("Narrowed complex64:  %.15f (precision reduced to float32)\n", c64)

	// 2. Formatting to string with strconv.FormatComplex
	formatted := strconv.FormatComplex(c128, 'f', 2, 128)
	fmt.Printf("strconv.FormatComplex: %s\n", formatted)

	// 3. Parsing strings with strconv.ParseComplex
	testInputs := []string{"3+4i", "(5.5-2.1i)", "-7i", "42"}
	for _, str := range testInputs {
		parsed, err := strconv.ParseComplex(str, 128)
		if err != nil {
			fmt.Printf("  Error parsing %q: %v\n", str, err)
		} else {
			fmt.Printf("  strconv.ParseComplex(%q) = %v\n", str, parsed)
		}
	}
}

// -----------------------------------------------------------------------------
// 6. THE `math/cmplx` STANDARD LIBRARY PACKAGE
// -----------------------------------------------------------------------------
// The `math/cmplx` package provides mathematical functions for complex128 numbers:
// - Magnitude & Phase:
//   - `cmplx.Abs(c)`   : Computes magnitude / absolute value: |z| = √(a² + b²)
//   - `cmplx.Phase(c)` : Computes phase angle / argument: θ = atan2(b, a)
// - Polar Coordinates:
//   - `cmplx.Polar(c)` : Returns magnitude r and phase angle θ.
//   - `cmplx.Rect(r, θ)` : Constructs a complex number from polar coordinates r and θ.
// - Complex Conjugate:
//   - `cmplx.Conj(c)`  : Negates the imaginary part: a + bi -> a - bi
// - Powers & Roots:
//   - `cmplx.Sqrt(c)`  : Square root (e.g. √(-1) = 0 + 1i)
//   - `cmplx.Pow(x, y)`: Raises complex x to power y
// - Exponential & Trigonometric:
//   - `cmplx.Exp(c)`   : e^c (Demonstrates Euler's identity: e^(i*π) + 1 = 0)
//   - `cmplx.Log(c)`   : Natural logarithm
//   - `cmplx.Sin(c)`, `cmplx.Cos(c)`, `cmplx.Tan(c)`
// - Special Values:
//   - `cmplx.IsNaN(c)`, `cmplx.IsInf(c)`
// -----------------------------------------------------------------------------
func demonstrateMathCmplxPackage() {
	fmt.Println("\n--- 6. ADVANCED MATH WITH `math/cmplx` ---")

	z := 3 + 4i // Classic 3-4-5 right triangle

	// Magnitude |z| and Phase θ
	magnitude := cmplx.Abs(z) // √(3² + 4²) = 5
	phase := cmplx.Phase(z)   // angle in radians
	fmt.Printf("z = %v\n", z)
	fmt.Printf("  Magnitude cmplx.Abs(z):   %.2f\n", magnitude)
	fmt.Printf("  Phase angle in radians:   %.4f (%.2f degrees)\n", phase, phase*(180/math.Pi))

	// Complex Conjugate: a + bi -> a - bi
	conj := cmplx.Conj(z)
	fmt.Printf("  Conjugate cmplx.Conj(z):  %v\n", conj)
	// Property: z * conj(z) = |z|²
	fmt.Printf("  z * Conj(z) = %v (equals |z|² = %.0f)\n", z*conj, magnitude*magnitude)

	// Polar / Cartesian conversions
	r, theta := cmplx.Polar(z)
	reconstructed := cmplx.Rect(r, theta)
	fmt.Printf("  Polar form: r = %.2f, θ = %.4f rad\n", r, theta)
	fmt.Printf("  Rect(r, θ) reconstructed: %.2f\n", reconstructed)

	// Square root of negative numbers!
	sqrtNeg1 := cmplx.Sqrt(-1)
	sqrtNeg4 := cmplx.Sqrt(-4)
	fmt.Printf("  cmplx.Sqrt(-1): %v\n", sqrtNeg1)
	fmt.Printf("  cmplx.Sqrt(-4): %v\n", sqrtNeg4)

	// Euler's Identity: e^(i * π) + 1 = 0
	pi := complex(math.Pi, 0)
	iVal := 0 + 1i
	eulersIdentity := cmplx.Exp(iVal*pi) + 1
	fmt.Printf("  Euler's Identity e^(i*π) + 1: %v (approaches zero)\n", eulersIdentity)

	// Complex trigonometric functions
	fmt.Printf("  cmplx.Sin(z): %v\n", cmplx.Sin(z))
	fmt.Printf("  cmplx.Cos(z): %v\n", cmplx.Cos(z))
}

// -----------------------------------------------------------------------------
// 7. FORMATTING COMPLEX NUMBERS WITH `fmt`
// -----------------------------------------------------------------------------
// Formatting verbs for complex numbers:
// - `%v`    : Default format: `(3+4i)`
// - `%f`    : Floating-point format applied to both parts: `(3.000000+4.000000i)`
// - `%.2f`  : Precision-specified float: `(3.00+4.00i)`
// - `%e`    : Scientific notation: `(3.000000e+00+4.000000e+00i)`
// - `%g`    : Compact representation
// - `%T`    : Type name (`complex128` or `complex64`)
// -----------------------------------------------------------------------------
func demonstrateFormatting() {
	fmt.Println("\n--- 7. FORMATTING WITH FMT ---")

	c := 3.14159265 + 2.71828182i

	fmt.Printf("%%v    (default):    %v\n", c)
	fmt.Printf("%%f    (float):      %f\n", c)
	fmt.Printf("%%.2f  (2 decimals): %.2f\n", c)
	fmt.Printf("%%e    (scientific): %e\n", c)
	fmt.Printf("%%g    (compact):    %g\n", c)
	fmt.Printf("%%T    (type):       %T\n", c)
}

// -----------------------------------------------------------------------------
// 8. REAL-WORLD APPLICATION EXAMPLE: ELECTRICAL IMPEDANCE (AC CIRCUITS)
// -----------------------------------------------------------------------------
// In electrical engineering and signal processing, alternating current (AC)
// circuits use complex numbers to represent impedance:
//   Z = R + jX
// where:
//   - R is resistance in ohms (real part)
//   - X is reactance in ohms (imaginary part)
//
// Ohm's law in AC circuits:
//   Voltage V = Current I * Impedance Z
// -----------------------------------------------------------------------------
func demonstratePracticalApplication() {
	fmt.Println("\n--- 8. REAL-WORLD APPLICATION: AC CIRCUIT IMPEDANCE ---")

	// Current I = 2 + 1i Amperes
	current := 2.0 + 1.0i

	// Impedance Z = 4 + 3i Ohms (4 ohms resistance, 3 ohms inductive reactance)
	impedance := 4.0 + 3.0i

	// Calculate Voltage V = I * Z
	voltage := current * impedance

	fmt.Printf("Current   I = %v A\n", current)
	fmt.Printf("Impedance Z = %v Ω\n", impedance)
	fmt.Printf("Voltage   V = I * Z = %v V\n", voltage)
	fmt.Printf("Voltage Magnitude |V|: %.2f Volts\n", cmplx.Abs(voltage))
	fmt.Printf("Phase Angle:            %.2f degrees\n", cmplx.Phase(voltage)*(180/math.Pi))

	fmt.Println("\n================================================================================")
	fmt.Println("                     END OF COMPLEX NUMBERS CONCEPTS GUIDE                      ")
	fmt.Println("================================================================================")
}
