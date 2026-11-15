package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println("================================================================================")
	fmt.Println("                        COMPREHENSIVE GOLANG FLOATS GUIDE                       ")
	fmt.Println("================================================================================")

	demonstrateBasicsAndMemory()
	demonstrateLiteralsAndScientificNotation()
	demonstrateArithmeticAndStrictTyping()
	demonstratePrecisionAndEpsilonComparison()
	demonstrateSpecialValues()
	demonstrateMathPackageFunctions()
	demonstrateArbitraryPrecisionBigFloat()
	demonstrateConversionsAndParsing()
	demonstrateFormatting()
}

// -----------------------------------------------------------------------------
// 1. BASICS & MEMORY CHARACTERISTICS (IEEE-754)
// -----------------------------------------------------------------------------
// - Go implements the IEEE-754 standard for floating-point arithmetic.
// - Go provides two built-in floating-point types:
//   1. `float32`: Single-precision (32 bits = 4 bytes)
//      - 1 sign bit, 8 exponent bits, 23 mantissa (fraction) bits.
//      - Provides ~7 decimal digits of precision.
//      - Range: ~1.4e-45 to ~3.4e+38.
//   2. `float64`: Double-precision (64 bits = 8 bytes)
//      - 1 sign bit, 11 exponent bits, 52 mantissa (fraction) bits.
//      - Provides ~15-17 decimal digits of precision.
//      - Range: ~4.9e-324 to ~1.8e+308.
//      - Default type for untyped floating-point literals (e.g. `f := 3.14`).
//
// - ZERO VALUES:
//   - Zero value for both types is 0.0 (or 0).
//   - IEEE-754 defines both +0.0 and -0.0, though +0.0 == -0.0 evaluates to true.
// -----------------------------------------------------------------------------
func demonstrateBasicsAndMemory() {
	fmt.Println("\n--- 1. BASICS & MEMORY CHARACTERISTICS ---")

	var f32 float32
	var f64 float64
	fmt.Printf("Zero value of float32: %v\n", f32)
	fmt.Printf("Zero value of float64: %v\n", f64)

	// Untyped float literal defaults to float64
	defaultFloat := 3.14159
	fmt.Printf("Type of '3.14159': %T\n", defaultFloat)

	// Memory sizes
	fmt.Printf("Size of float32 in memory: %d bytes (32 bits)\n", unsafe.Sizeof(f32))
	fmt.Printf("Size of float64 in memory: %d bytes (64 bits)\n", unsafe.Sizeof(f64))

	// Ranges from the math package
	fmt.Printf("float32 range: [%e, %e]\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Printf("float64 range: [%e, %e]\n", math.SmallestNonzeroFloat64, math.MaxFloat64)
}

// -----------------------------------------------------------------------------
// 2. FLOAT LITERALS & SCIENTIFIC NOTATION
// -----------------------------------------------------------------------------
// Go supports multiple representations for floating-point literals:
// 1. Standard decimal notation: `3.14`, `.5` (0.5), `10.` (10.0)
// 2. Scientific / Exponential notation (`e` or `E`):
//    - `1.5e3`  = 1.5 * 10^3 = 1500
//    - `2.5e-4` = 2.5 * 10^-4 = 0.00025
// 3. Hexadecimal floating-point (Go 1.13+):
//    - `0x1p-2` = 1.0 * 2^-2 = 0.25 ('p' specifies base-2 exponent)
// 4. Underscore digit separators (Go 1.13+):
//    - `1_000_000.500_001` for improved readability
// -----------------------------------------------------------------------------
func demonstrateLiteralsAndScientificNotation() {
	fmt.Println("\n--- 2. FLOAT LITERALS & SCIENTIFIC NOTATION ---")

	standard := 3.14159
	omittedLeadingZero := .75 // equivalent to 0.75
	scientificBig := 1.23e6   // 1.23 * 10^6 = 1,230,000
	scientificSmall := 4.5e-5 // 4.5 * 10^-5 = 0.000045
	hexFloat := 0x1.0p-3      // 1.0 * 2^-3 = 0.125
	withUnderscores := 1_234_567.890_123

	fmt.Printf("Standard:                %f\n", standard)
	fmt.Printf("Omitted leading zero:    %f\n", omittedLeadingZero)
	fmt.Printf("Scientific positive (e6):%f (%e)\n", scientificBig, scientificBig)
	fmt.Printf("Scientific negative (e-5):%f (%e)\n", scientificSmall, scientificSmall)
	fmt.Printf("Hex float (0x1.0p-3):    %f\n", hexFloat)
	fmt.Printf("With underscores:        %f\n", withUnderscores)
}

// -----------------------------------------------------------------------------
// 3. ARITHMETIC OPERATIONS & TYPE STRICTNESS
// -----------------------------------------------------------------------------
// - Supported arithmetic: Addition (`+`), Subtraction (`-`), Multiplication (`*`),
//   Division (`/`), and Unary Negation (`-`).
// - NO MODULUS OPERATOR (`%`):
//   - `%` is strictly for integer types in Go.
//   - For floating-point remainder, use `math.Mod(x, y)`.
// - STRICT TYPING:
//   - No implicit conversions! You cannot combine `float32` and `float64`,
//     nor `int` and `float64` without explicit conversion.
// -----------------------------------------------------------------------------
func demonstrateArithmeticAndStrictTyping() {
	fmt.Println("\n--- 3. ARITHMETIC OPERATIONS & STRICT TYPING ---")

	a := 10.5
	b := 2.5

	fmt.Printf("a = %.1f, b = %.1f\n", a, b)
	fmt.Printf("Addition       (a + b) = %.2f\n", a+b)
	fmt.Printf("Subtraction    (a - b) = %.2f\n", a-b)
	fmt.Printf("Multiplication (a * b) = %.2f\n", a*b)
	fmt.Printf("Division       (a / b) = %.2f\n", a/b)
	fmt.Printf("Unary Negation (-a)    = %.2f\n", -a)

	// Floating-point remainder using math.Mod (since a % b does not compile)
	remainder := math.Mod(a, b)
	fmt.Printf("math.Mod(10.5, 2.5)    = %.2f (10.5 = 4 * 2.5 + 0.5)\n", remainder)

	// Explicit type conversions required:
	var f32Val float32 = 3.5
	var f64Val float64 = 6.5
	// total := f32Val + f64Val // COMPILE ERROR: mismatched types float32 and float64
	total := float64(f32Val) + f64Val
	fmt.Printf("Explicit conversion sum: %.2f\n", total)
}

// -----------------------------------------------------------------------------
// 4. FLOATING-POINT IMPRECISION & EPSILON COMPARISON
// -----------------------------------------------------------------------------
// - THE BINARY FRACTION PITFALL:
//   - Computers represent numbers in binary (base-2). Just as 1/3 cannot be
//     represented precisely in base-10 (0.33333...), fractions like 0.1 and 0.2
//     cannot be represented precisely in base-2 floating-point.
//   - Result: `0.1 + 0.2` in float64 is `0.30000000000000004`, NOT `0.3`!
//
// - NEVER USE `==` FOR COMPUTED FLOATS:
//   - Direct equality checks frequently fail due to rounding inaccuracies.
//   - SOLUTION: Compare using an epsilon threshold (`|a - b| < epsilon`).
//
// - FINANCIAL WARNING:
//   - Never use float32/float64 for monetary transactions!
//   - Use integer cents (e.g., $10.50 as 1050 cents) or arbitrary-precision libraries.
// -----------------------------------------------------------------------------
func demonstratePrecisionAndEpsilonComparison() {
	fmt.Println("\n--- 4. IMPRECISION & EPSILON COMPARISON ---")

	x := 0.1
	y := 0.2
	sum := x + y
	expected := 0.3

	fmt.Printf("0.1 + 0.2 formatted with standard %%f:     %f\n", sum)
	fmt.Printf("0.1 + 0.2 formatted with high precision:  %.17f\n", sum)
	fmt.Printf("Direct equality check (sum == 0.3):       %t\n", sum == expected)

	// Correct approach: Epsilon comparison
	const epsilon = 1e-9
	isEqual := math.Abs(sum-expected) < epsilon
	fmt.Printf("Epsilon comparison (|sum - 0.3| < 1e-9):  %t (Safe and reliable)\n", isEqual)

	// Precision comparison between float32 and float64
	var p32 float32 = 1.0 / 3.0
	var p64 float64 = 1.0 / 3.0
	fmt.Printf("1.0/3.0 in float32 (~7 digits):  %.10f\n", p32)
	fmt.Printf("1.0/3.0 in float64 (~16 digits): %.18f\n", p64)
}

// -----------------------------------------------------------------------------
// 5. SPECIAL FLOATING-POINT VALUES (NaN & INFINITY)
// -----------------------------------------------------------------------------
// The IEEE-754 standard defines special non-finite representations:
// 1. Positive Infinity (`+Inf`) & Negative Infinity (`-Inf`):
//    - Arises from division of non-zero by zero (e.g. `1.0 / 0.0`), or overflow.
//    - Created explicitly via `math.Inf(1)` (+Inf) and `math.Inf(-1)` (-Inf).
//    - Check with `math.IsInf(f, 0)` (either), `math.IsInf(f, 1)`, `math.IsInf(f, -1)`.
// 2. Not-a-Number (`NaN`):
//    - Arises from undefined mathematical operations (e.g. `0.0 / 0.0`, `math.Sqrt(-1)`).
//    - Created explicitly via `math.NaN()`.
//    - Check with `math.IsNaN(f)`.
//    - CRITICAL PROPERTY: `NaN == NaN` is ALWAYS FALSE!
// -----------------------------------------------------------------------------
func demonstrateSpecialValues() {
	fmt.Println("\n--- 5. SPECIAL VALUES: NaN & INFINITY ---")

	// Infinity
	posInf := math.Inf(1)
	negInf := math.Inf(-1)
	fmt.Printf("Positive Infinity: %v, math.IsInf(posInf, 1): %t\n", posInf, math.IsInf(posInf, 1))
	fmt.Printf("Negative Infinity: %v, math.IsInf(negInf, -1): %t\n", negInf, math.IsInf(negInf, -1))

	// Division by zero in floats yields Infinity (unlike integers, which panic!)
	var zero float64 = 0.0
	divByZero := 10.0 / zero
	fmt.Printf("10.0 / 0.0 = %v (does NOT panic like integer division)\n", divByZero)

	// NaN (Not a Number)
	nan := math.NaN()
	invalidSqrt := math.Sqrt(-4.0) // Sqrt of negative real number yields NaN
	fmt.Printf("math.NaN(): %v\n", nan)
	fmt.Printf("math.Sqrt(-4.0): %v\n", invalidSqrt)
	fmt.Printf("math.IsNaN(invalidSqrt): %t\n", math.IsNaN(invalidSqrt))

	// The Golden Rule of NaN:
	fmt.Printf("Is NaN equal to itself (nan == nan)? %t (ALWAYS false by design!)\n", nan == nan)
}

// -----------------------------------------------------------------------------
// 6. THE `math` STANDARD LIBRARY PACKAGE
// -----------------------------------------------------------------------------
// Common floating-point functions in the standard `math` package:
// - Rounding:
//   - `math.Round(x)` : Round to nearest integer (half away from zero)
//   - `math.Floor(x)` : Round down toward negative infinity
//   - `math.Ceil(x)`  : Round up toward positive infinity
//   - `math.Trunc(x)` : Discard fractional part (round toward zero)
// - Powers & Roots:
//   - `math.Pow(x, y)`: x^y
//   - `math.Sqrt(x)`  : Square root
// - Absolute, Min, Max:
//   - `math.Abs(x)`
//   - `math.Max(x, y)`, `math.Min(x, y)`
// - Trigonometry & Logarithms:
//   - `math.Sin(x)`, `math.Cos(x)`, `math.Tan(x)`
//   - `math.Log(x)` (natural log), `math.Log10(x)`, `math.Exp(x)`
// -----------------------------------------------------------------------------
func demonstrateMathPackageFunctions() {
	fmt.Println("\n--- 6. MATH PACKAGE ESSENTIALS ---")

	val := 3.75
	negVal := -3.75

	// Rounding operations
	fmt.Printf("Value: %.2f\n", val)
	fmt.Printf("  math.Round(%.2f): %.0f\n", val, math.Round(val))
	fmt.Printf("  math.Floor(%.2f): %.0f\n", val, math.Floor(val))
	fmt.Printf("  math.Ceil(%.2f):  %.0f\n", val, math.Ceil(val))
	fmt.Printf("  math.Trunc(%.2f): %.0f\n", val, math.Trunc(val))

	fmt.Printf("Negative Value: %.2f\n", negVal)
	fmt.Printf("  math.Floor(%.2f): %.0f (moves toward -Inf)\n", negVal, math.Floor(negVal))
	fmt.Printf("  math.Trunc(%.2f): %.0f (moves toward zero)\n", negVal, math.Trunc(negVal))

	// Powers, roots, and trigonometry
	fmt.Printf("math.Pow(2.0, 8.0): %.0f\n", math.Pow(2.0, 8.0))
	fmt.Printf("math.Sqrt(81.0):    %.0f\n", math.Sqrt(81.0))
	fmt.Printf("math.Sin(π / 2):    %.1f\n", math.Sin(math.Pi/2))
	fmt.Printf("math.Max(4.5, 9.2): %.1f\n", math.Max(4.5, 9.2))
	fmt.Printf("math.Min(4.5, 9.2): %.1f\n", math.Min(4.5, 9.2))
}

// -----------------------------------------------------------------------------
// 7. ARBITRARY PRECISION WITH `math/big.Float`
// -----------------------------------------------------------------------------
// When standard float64 precision (~16 digits) is insufficient (e.g. cryptography,
// astronomy, financial calculations requiring exact decimal representations):
// - Use `math/big.Float`.
// - Precision can be configured to arbitrary bits (e.g. 128, 256, 512 bits).
// -----------------------------------------------------------------------------
func demonstrateArbitraryPrecisionBigFloat() {
	fmt.Println("\n--- 7. ARBITRARY PRECISION: math/big.Float ---")

	// In standard float64, 0.1 + 0.2 has precision artifacts:
	// With big.Float and exact rational initialization:
	f1 := new(big.Float).SetPrec(128).SetFloat64(0.1)
	f2 := new(big.Float).SetPrec(128).SetFloat64(0.2)
	sum := new(big.Float).Add(f1, f2)

	fmt.Printf("big.Float (128-bit prec) sum: %s\n", sum.Text('f', 20))

	// Exact string initialization
	bf1, _, _ := new(big.Float).SetPrec(256).Parse("0.1", 10)
	bf2, _, _ := new(big.Float).SetPrec(256).Parse("0.2", 10)
	exactSum := new(big.Float).Add(bf1, bf2)
	fmt.Printf("big.Float parsed from strings: %s\n", exactSum.Text('f', 10))
}

// -----------------------------------------------------------------------------
// 8. TYPE CONVERSIONS & PARSING WITH `strconv`
// -----------------------------------------------------------------------------
// - `strconv.FormatFloat(f, fmt, prec, bitSize)`:
//   - fmt: 'f' (decimal), 'e' (scientific), 'g' (compact), 'b' (binary)
//   - prec: number of digits (-1 uses fewest digits needed for exact round-trip)
//   - bitSize: 32 or 64
// - `strconv.ParseFloat(str, bitSize)`:
//   - Parses string into float64.
// - Conversions between int and float:
//   - `float64(i)`: int to float
//   - `int(f)`: float to int (TRUNCATES decimal portion, does NOT round!)
// -----------------------------------------------------------------------------
func demonstrateConversionsAndParsing() {
	fmt.Println("\n--- 8. TYPE CONVERSIONS & PARSING WITH STRCONV ---")

	pi := 3.1415926535

	// 1. strconv.FormatFloat
	strDecimal := strconv.FormatFloat(pi, 'f', 2, 64)
	strScientific := strconv.FormatFloat(pi, 'e', 3, 64)
	strCompact := strconv.FormatFloat(pi, 'g', -1, 64)
	fmt.Printf("FormatFloat 'f' (prec 2): %s\n", strDecimal)
	fmt.Printf("FormatFloat 'e' (prec 3): %s\n", strScientific)
	fmt.Printf("FormatFloat 'g' (compact):%s\n", strCompact)

	// 2. strconv.ParseFloat
	parsedFloat, err := strconv.ParseFloat("123.456", 64)
	if err == nil {
		fmt.Printf("strconv.ParseFloat(\"123.456\"): %.3f (type %T)\n", parsedFloat, parsedFloat)
	}

	// 3. Casting int <-> float
	originalInt := 42
	castToFloat := float64(originalInt)
	fmt.Printf("int -> float64: %d -> %.2f\n", originalInt, castToFloat)

	// WARNING: float to int truncates toward zero!
	floatValue := 9.99
	truncatedInt := int(floatValue)
	roundedInt := int(math.Round(floatValue))
	fmt.Printf("float64 9.99 -> int(9.99):        %d (TRUNCATED!)\n", truncatedInt)
	fmt.Printf("float64 9.99 -> int(Round(9.99)): %d (Properly rounded)\n", roundedInt)
}

// -----------------------------------------------------------------------------
// 9. FORMATTING FLOATS WITH `fmt`
// -----------------------------------------------------------------------------
// Common formatting verbs for floats:
// - `%f`   : Standard decimal notation (default precision is 6 digits)
// - `%.2f` : Decimal notation with explicit 2 decimal places
// - `%e`   : Scientific notation with lowercase 'e'
// - `%E`   : Scientific notation with uppercase 'E'
// - `%g`   : Compact format (uses %e for huge/tiny numbers, %f otherwise; strips trailing zeros)
// - `%8.2f`: Width 8, right-aligned, 2 decimal places
// - `%-8.2f`: Width 8, left-aligned, 2 decimal places
// - `%08.2f`: Width 8, zero-padded
// - `%T`   : Type name (e.g. float64)
// -----------------------------------------------------------------------------
func demonstrateFormatting() {
	fmt.Println("\n--- 9. FORMATTING WITH FMT ---")

	num := 123.456789

	fmt.Printf("%%f    (default):    %f\n", num)
	fmt.Printf("%%.2f  (2 decimals): %.2f\n", num)
	fmt.Printf("%%.4f  (4 decimals): %.4f\n", num)
	fmt.Printf("%%e    (scientific): %e\n", num)
	fmt.Printf("%%E    (scientific): %E\n", num)
	fmt.Printf("%%g    (compact):    %g\n", num)
	fmt.Printf("%%T    (type):       %T\n", num)

	// Width, alignment, and padding
	fmt.Printf("Right-aligned (%%8.2f): |%8.2f|\n", num)
	fmt.Printf("Left-aligned  (%%-8.2f):|%-8.2f|\n", num)
	fmt.Printf("Zero-padded   (%%08.2f):|%08.2f|\n", num)

	fmt.Println("\n================================================================================")
	fmt.Println("                         END OF FLOATS CONCEPTS GUIDE                           ")
	fmt.Println("================================================================================")
}
