package main

import (
	"fmt"
	"math"
	"math/big"
	"math/bits"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println("================================================================================")
	fmt.Println("                        COMPREHENSIVE GOLANG INTEGERS GUIDE                     ")
	fmt.Println("================================================================================")

	demonstrateTypesAndMemory()
	demonstrateLiteralsAndBases()
	demonstrateArithmeticAndDivision()
	demonstrateBitwiseOperators()
	demonstrateOverflowAndUnderflow()
	demonstrateMathBitsPackage()
	demonstrateArbitraryPrecisionBigInt()
	demonstrateConversionsAndParsing()
	demonstrateFormatting()
}

// -----------------------------------------------------------------------------
// 1. INTEGER TYPES, SIZES, & ARCHITECTURE DEPENDENCE
// -----------------------------------------------------------------------------
// Go provides fixed-size, architecture-dependent, and alias integer types:
//
// 1. Signed Integers (Two's Complement):
//   - int8  : 8 bits  (-128 to 127)
//   - int16 : 16 bits (-32,768 to 32,767)
//   - int32 : 32 bits (-2,147,483,648 to 2,147,483,647)
//   - int64 : 64 bits (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807)
//
// 2. Unsigned Integers:
//   - uint8  : 8 bits  (0 to 255)
//   - uint16 : 16 bits (0 to 65,535)
//   - uint32 : 32 bits (0 to 4,294,967,295)
//   - uint64 : 64 bits (0 to 18,446,744,073,709,551,615)
//
// 3. Architecture-Dependent Types:
//   - int     : Same size as pointer/architecture word size (32 bits on 32-bit systems, 64 bits on 64-bit systems).
//   - uint    : Same size as pointer/architecture word size (32 bits on 32-bit systems, 64 bits on 64-bit systems).
//   - uintptr : Unsigned integer large enough to store uninterpreted bits of a pointer value.
//
// 4. Type Aliases:
//   - byte : Alias for uint8  (used for raw byte data)
//   - rune : Alias for int32  (used for Unicode code points)
//
// Zero value for all integer types is 0.
// -----------------------------------------------------------------------------
func demonstrateTypesAndMemory() {
	fmt.Println("\n--- 1. INTEGER TYPES, SIZES, & ARCHITECTURE ---")

	var (
		i8   int8
		i16  int16
		i32  int32
		i64  int64
		u8   uint8
		u16  uint16
		u32  uint32
		u64  uint64
		i    int
		u    uint
		uptr uintptr
		b    byte
		r    rune
	)

	fmt.Printf("%-10s | Size: %d byte(s) (%2d bits) | Range: [%d to %d]\n",
		"int8", unsafe.Sizeof(i8), unsafe.Sizeof(i8)*8, math.MinInt8, math.MaxInt8)
	fmt.Printf("%-10s | Size: %d byte(s) (%2d bits) | Range: [%d to %d]\n",
		"int16", unsafe.Sizeof(i16), unsafe.Sizeof(i16)*8, math.MinInt16, math.MaxInt16)
	fmt.Printf("%-10s | Size: %d byte(s) (%2d bits) | Range: [%d to %d]\n",
		"int32", unsafe.Sizeof(i32), unsafe.Sizeof(i32)*8, math.MinInt32, math.MaxInt32)
	fmt.Printf("%-10s | Size: %d byte(s) (%2d bits) | Range: [%d to %d]\n",
		"int64", unsafe.Sizeof(i64), unsafe.Sizeof(i64)*8, int64(math.MinInt64), int64(math.MaxInt64))

	fmt.Printf("%-10s | Size: %d byte(s) (%2d bits) | Range: [0 to %d]\n",
		"uint8", unsafe.Sizeof(u8), unsafe.Sizeof(u8)*8, math.MaxUint8)
	fmt.Printf("%-10s | Size: %d byte(s) (%2d bits) | Range: [0 to %d]\n",
		"uint16", unsafe.Sizeof(u16), unsafe.Sizeof(u16)*8, math.MaxUint16)
	fmt.Printf("%-10s | Size: %d byte(s) (%2d bits) | Range: [0 to %d]\n",
		"uint32", unsafe.Sizeof(u32), unsafe.Sizeof(u32)*8, uint32(math.MaxUint32))
	fmt.Printf("%-10s | Size: %d byte(s) (%2d bits) | Range: [0 to %d]\n",
		"uint64", unsafe.Sizeof(u64), unsafe.Sizeof(u64)*8, uint64(math.MaxUint64))

	fmt.Printf("%-10s | Size: %d byte(s) (%2d bits) | Architecture-dependent (int)\n",
		"int", unsafe.Sizeof(i), unsafe.Sizeof(i)*8)
	fmt.Printf("%-10s | Size: %d byte(s) (%2d bits) | Architecture-dependent (uint)\n",
		"uint", unsafe.Sizeof(u), unsafe.Sizeof(u)*8)
	fmt.Printf("%-10s | Size: %d byte(s) (%2d bits) | Architecture-dependent (uintptr)\n",
		"uintptr", unsafe.Sizeof(uptr), unsafe.Sizeof(uptr)*8)

	// Aliases
	fmt.Printf("byte is alias for uint8: %T (size %d bytes)\n", b, unsafe.Sizeof(b))
	fmt.Printf("rune is alias for int32: %T (size %d bytes)\n", r, unsafe.Sizeof(r))
}

// -----------------------------------------------------------------------------
// 2. INTEGER LITERALS & NUMBER BASES (Go 1.13+)
// -----------------------------------------------------------------------------
// Go supports integer literals in multiple number systems:
// - Decimal (Base 10)     : Standard numbers (e.g. 42)
// - Binary (Base 2)       : Prefix `0b` or `0B` (e.g. 0b101010)
// - Octal (Base 8)        : Prefix `0o` or `0O` (e.g. 0o52), or legacy `052`
// - Hexadecimal (Base 16) : Prefix `0x` or `0X` (e.g. 0x2A)
// - Underscore separators : `_` can be placed anywhere between digits for readability
// -----------------------------------------------------------------------------
func demonstrateLiteralsAndBases() {
	fmt.Println("\n--- 2. INTEGER LITERALS & NUMBER BASES ---")

	dec := 42
	bin := 0b101010
	oct := 0o52
	hex := 0x2A
	withSeparators := 1_000_000_000

	fmt.Printf("Decimal:               %d\n", dec)
	fmt.Printf("Binary     (0b101010): %d\n", bin)
	fmt.Printf("Octal      (0o52):     %d\n", oct)
	fmt.Printf("Hexadecimal(0x2A):     %d\n", hex)
	fmt.Printf("All represent same:    %t\n", dec == bin && bin == oct && oct == hex)
	fmt.Printf("With underscores:      %d\n", withSeparators)
}

// -----------------------------------------------------------------------------
// 3. ARITHMETIC OPERATORS & DIVISION BEHAVIOR
// -----------------------------------------------------------------------------
// - Supported operators:
//   - Addition (`+`), Subtraction (`-`), Multiplication (`*`)
//   - Integer Division (`/`): Truncates toward zero (e.g., 7 / 2 = 3, -7 / 2 = -3).
//   - Modulus (`%`): Remainder operator (e.g., 7 % 2 = 1, -7 % 2 = -1).
//     Notice: The sign of the result of `%` matches the dividend!
//
// - Strict typing:
//   - No implicit conversions! You cannot add an `int` to an `int64`.
//
// - Division by Zero:
//   - Dividing an integer by zero causes a runtime panic!
//
// -----------------------------------------------------------------------------
func demonstrateArithmeticAndDivision() {
	fmt.Println("\n--- 3. ARITHMETIC & INTEGER DIVISION ---")

	a := 17
	b := 5

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("Addition       (a + b) = %d\n", a+b)
	fmt.Printf("Subtraction    (a - b) = %d\n", a-b)
	fmt.Printf("Multiplication (a * b) = %d\n", a*b)
	fmt.Printf("Division       (a / b) = %d (integer division truncates towards zero)\n", a/b)
	fmt.Printf("Modulus        (a %% b) = %d (remainder)\n", a%b)

	// Negative integer division and modulus
	fmt.Printf("-17 /  5 = %d, -17 %%  5 = %d\n", -17/5, -17%5)
	fmt.Printf(" 17 / -5 = %d,  17 %% -5 = %d\n", 17/-5, 17%-5)

	// Division by zero warning:
	// zero := 0
	// _ = a / zero // RUNTIME PANIC: runtime error: integer divide by zero
	fmt.Println("Warning: Integer division by zero triggers a runtime panic.")

	// Strict typing demonstration:
	var small int32 = 100
	var large int64 = 200
	// sum := small + large // COMPILE ERROR: mismatched types int32 and int64
	sum := int64(small) + large
	fmt.Printf("Explicit conversion sum: %d (type %T)\n", sum, sum)
}

// -----------------------------------------------------------------------------
// 4. BITWISE OPERATORS (INCLUDING GO'S UNIQUE `&^` BIT CLEAR)
// -----------------------------------------------------------------------------
// Go provides a comprehensive set of bitwise operators:
// 1. `&`  : Bitwise AND
// 2. `|`  : Bitwise OR
// 3. `^`  : Bitwise XOR (as binary operator: a ^ b)
// 4. `^`  : Bitwise NOT (as unary operator: ^a, inverts all bits)
// 5. `&^` : Bit Clear (AND NOT) - A unique Go operator!
//   - `z = x &^ y` clears the bits in x that are set in y.
//   - Equivalent to `x & (^y)`.
//
// 6. `<<` : Left shift (multiplies by 2^n)
// 7. `>>` : Right shift (divides by 2^n)
//   - Signed integers perform arithmetic shift (preserves sign bit).
//   - Unsigned integers perform logical shift (zero-fills).
//
// -----------------------------------------------------------------------------
func demonstrateBitwiseOperators() {
	fmt.Println("\n--- 4. BITWISE OPERATORS & BIT CLEAR (&^) ---")

	var x uint8 = 0b1100 // 12 in decimal
	var y uint8 = 0b1010 // 10 in decimal

	fmt.Printf("x           = %04b (%2d)\n", x, x)
	fmt.Printf("y           = %04b (%2d)\n", y, y)
	fmt.Printf("x & y (AND) = %04b (%2d)\n", x&y, x&y)
	fmt.Printf("x | y (OR)  = %04b (%2d)\n", x|y, x|y)
	fmt.Printf("x ^ y (XOR) = %04b (%2d)\n", x^y, x^y)

	// Unary Bitwise NOT (^)
	fmt.Printf("^x    (NOT) = %08b (%2d)\n", ^x, ^x)

	// Go's Unique Bit Clear operator (&^)
	// For every bit in y that is 1, the corresponding bit in x is cleared to 0.
	bitClear := x &^ y
	fmt.Printf("x &^ y (BIT CLEAR) = %04b (%2d) (clears bits in x that are 1 in y)\n", bitClear, bitClear)

	// Shifts
	var val uint8 = 0b0000_1000 // 8
	fmt.Printf("val        = %08b (%d)\n", val, val)
	fmt.Printf("val << 2   = %08b (%d) (shifted left: multiplied by 4)\n", val<<2, val<<2)
	fmt.Printf("val >> 2   = %08b (%d) (shifted right: divided by 4)\n", val>>2, val>>2)

	// Arithmetic shift vs Logical shift
	var signedVal int8 = -8
	var unsignedVal uint8 = 248 // Same binary representation 11111000
	fmt.Printf("Signed int8 -8 >> 2:   %d (sign bit preserved)\n", signedVal>>2)
	fmt.Printf("Unsigned uint8 248 >> 2: %d (zero-filled)\n", unsignedVal>>2)
}

// -----------------------------------------------------------------------------
// 5. INTEGER OVERFLOW & UNDERFLOW
// -----------------------------------------------------------------------------
// - RUNTIME BEHAVIOR:
//   - Go does NOT panic or throw exceptions when integer overflow/underflow occurs
//     at runtime!
//   - Integers wrap around silently using two's complement modular arithmetic.
//   - Example: MaxInt8 (127) + 1 wraps to MinInt8 (-128).
//   - Example: uint8(0) - 1 wraps to 255.
//
// - COMPILE-TIME CONSTANTS:
//   - Untyped numeric constants have arbitrary precision during compilation.
//   - If a constant expression overflows the target type during assignment,
//     the compiler raises an error (e.g. `var x int8 = 128` fails to compile).
//
// -----------------------------------------------------------------------------
func demonstrateOverflowAndUnderflow() {
	fmt.Println("\n--- 5. OVERFLOW & UNDERFLOW ---")

	// Signed 8-bit overflow & underflow
	var maxI8 int8 = math.MaxInt8 // 127
	var minI8 int8 = math.MinInt8 // -128

	fmt.Printf("maxI8:               %d\n", maxI8)
	fmt.Printf("maxI8 + 1 (overflow): %d (wraps around to MinInt8!)\n", maxI8+1)

	fmt.Printf("minI8:               %d\n", minI8)
	fmt.Printf("minI8 - 1 (underflow):%d (wraps around to MaxInt8!)\n", minI8-1)

	// Unsigned 8-bit underflow
	var zeroU8 uint8 = 0
	fmt.Printf("uint8(0) - 1:        %d (wraps to 255)\n", zeroU8-1)

	// Safe addition check pattern to detect overflow before it happens:
	a := math.MaxInt64 - 5
	b := 10
	if b > 0 && a > math.MaxInt64-b {
		fmt.Println("Safe overflow guard: a + b would overflow int64!")
	}
}

// -----------------------------------------------------------------------------
// 6. THE `math/bits` STANDARD LIBRARY PACKAGE
// -----------------------------------------------------------------------------
// The `math/bits` package provides optimized, architecture-accelerated bit
// manipulation functions:
// - `bits.OnesCount` / `bits.OnesCount64`: Counts the number of set bits (popcount)
// - `bits.LeadingZeros` / `bits.TrailingZeros`: Counts zero bits before/after data
// - `bits.RotateLeft` / `bits.Reverse`: Bit rotations and reversals
// - `bits.Add64`, `bits.Mul64`, `bits.Sub64`: Arithmetic with carry and overflow flags
// -----------------------------------------------------------------------------
func demonstrateMathBitsPackage() {
	fmt.Println("\n--- 6. BIT MANIPULATION WITH `math/bits` ---")

	var n uint64 = 0b0000_1101 // 13 in decimal

	fmt.Printf("n = %064b\n", n)
	fmt.Printf("bits.OnesCount64(n):  %d (population count / Hamming weight)\n", bits.OnesCount64(n))
	fmt.Printf("bits.TrailingZeros64: %d\n", bits.TrailingZeros64(n))
	fmt.Printf("bits.LeadingZeros64:  %d\n", bits.LeadingZeros64(n))
	fmt.Printf("bits.Len64(n):        %d (minimum bits required to represent n)\n", bits.Len64(n))

	// Rotate bits
	rotated := bits.RotateLeft64(n, 2)
	fmt.Printf("bits.RotateLeft64(2): %d (%064b)\n", rotated, rotated)

	// Arithmetic with carry (Add64)
	sum, carry := bits.Add64(math.MaxUint64, 1, 0)
	fmt.Printf("bits.Add64(MaxUint64, 1) -> sum: %d, carry: %d\n", sum, carry)
}

// -----------------------------------------------------------------------------
// 7. ARBITRARY PRECISION WITH `math/big.Int`
// -----------------------------------------------------------------------------
// When standard int64 (max ~9.22 × 10^18) is too small (e.g. cryptography,
// large factorials, blockchain tokens):
// - Use `math/big.Int`.
// - Can grow to consume as much memory as the system provides.
// -----------------------------------------------------------------------------
func demonstrateArbitraryPrecisionBigInt() {
	fmt.Println("\n--- 7. ARBITRARY PRECISION: math/big.Int ---")

	// Calculating 50 Factorial (50!)
	// 50! easily overflows int64
	factorial := big.NewInt(1)
	for i := int64(2); i <= 50; i++ {
		factorial.Mul(factorial, big.NewInt(i))
	}
	fmt.Printf("50! (Factorial of 50):\n%s\n", factorial.String())

	// Large numbers initialized from string
	bigNum1, _ := new(big.Int).SetString("123456789012345678901234567890", 10)
	bigNum2, _ := new(big.Int).SetString("987654321098765432109876543210", 10)
	bigSum := new(big.Int).Add(bigNum1, bigNum2)
	fmt.Printf("big.Int sum: %s\n", bigSum.String())
}

// -----------------------------------------------------------------------------
// 8. TYPE CONVERSIONS & PARSING WITH `strconv`
// -----------------------------------------------------------------------------
// - `strconv.Itoa(i int) string` : Int to ASCII string.
// - `strconv.Atoi(s string) (int, error)` : String to Int.
// - `strconv.FormatInt(i int64, base int) string` : Int to string in base 2, 8, 10, 16.
// - `strconv.ParseInt(s string, base int, bitSize int) (int64, error)` : String to int64.
// - `strconv.FormatUint`, `strconv.ParseUint` for unsigned integers.
// -----------------------------------------------------------------------------
func demonstrateConversionsAndParsing() {
	fmt.Println("\n--- 8. TYPE CONVERSIONS & PARSING WITH STRCONV ---")

	val := 255

	// 1. Itoa and Atoi
	strVal := strconv.Itoa(val)
	parsedVal, _ := strconv.Atoi(strVal)
	fmt.Printf("strconv.Itoa(%d) = %q | strconv.Atoi(%q) = %d\n", val, strVal, strVal, parsedVal)

	// 2. FormatInt into different bases
	fmt.Printf("FormatInt in base 2  (binary): %s\n", strconv.FormatInt(int64(val), 2))
	fmt.Printf("FormatInt in base 8  (octal):  %s\n", strconv.FormatInt(int64(val), 8))
	fmt.Printf("FormatInt in base 16 (hex):    %s\n", strconv.FormatInt(int64(val), 16))

	// 3. ParseInt from different bases
	parsedHex, _ := strconv.ParseInt("ff", 16, 64)
	parsedBin, _ := strconv.ParseInt("11111111", 2, 64)
	fmt.Printf("ParseInt(\"ff\", 16) = %d\n", parsedHex)
	fmt.Printf("ParseInt(\"11111111\", 2) = %d\n", parsedBin)
}

// -----------------------------------------------------------------------------
// 9. FORMATTING INTEGERS WITH `fmt`
// -----------------------------------------------------------------------------
// Integer formatting verbs:
// - `%d`  : Decimal (Base 10)
// - `%b`  : Binary (Base 2)
// - `%o`  : Octal (Base 8)
// - `%#o` : Octal with 0o prefix
// - `%x`  : Hexadecimal (lowercase)
// - `%X`  : Hexadecimal (uppercase)
// - `%#x` : Hexadecimal with 0x prefix
// - `%c`  : Character corresponding to Unicode code point
// - `%+d` : Always display sign (+ or -)
// - `%6d` : Right-aligned with width 6
// - `%-6d`: Left-aligned with width 6
// - `%06d`: Zero-padded with width 6
// - `%T`  : Type of variable
// -----------------------------------------------------------------------------
func demonstrateFormatting() {
	fmt.Println("\n--- 9. FORMATTING WITH FMT ---")

	n := 255

	fmt.Printf("%%d  (decimal):     %d\n", n)
	fmt.Printf("%%b  (binary):      %b\n", n)
	fmt.Printf("%%o  (octal):       %o\n", n)
	fmt.Printf("%%#o (octal prefix): %#o\n", n)
	fmt.Printf("%%x  (hex lower):   %x\n", n)
	fmt.Printf("%%X  (hex upper):   %X\n", n)
	fmt.Printf("%%#x (hex prefix):  %#x\n", n)
	fmt.Printf("%%c  (as char):     %c\n", 65) // 'A'
	fmt.Printf("%%+d (signed):      %+d\n", n)
	fmt.Printf("%%6d (width 6):     |%6d|\n", n)
	fmt.Printf("%%-6d(left-aligned):|%-6d|\n", n)
	fmt.Printf("%%06d(zero-padded): |%06d|\n", n)
	fmt.Printf("%%T  (type):        %T\n", n)

	fmt.Println("\n================================================================================")
	fmt.Println("                        END OF INTEGERS CONCEPTS GUIDE                          ")
	fmt.Println("================================================================================")
}
