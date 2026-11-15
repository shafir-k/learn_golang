package main

import (
	"fmt"
	"unsafe"
)

// =============================================================================
// PACKAGE-LEVEL CONSTANTS
// =============================================================================
// - Constants can be declared at package level or inside functions.
// - Exported (PascalCase) vs Unexported (camelCase) applies:
//   - `ExportedAPIEndpoint`: accessible outside this package.
//   - `databaseTimeout`: accessible only within package main.
// - Unlike local variables, UNUSED CONSTANTS DO NOT CAUSE COMPILE ERRORS!
// =============================================================================

const (
	ExportedAPIEndpoint string = "https://api.example.com/v1"
	databaseTimeoutSec  int    = 30
	unusedConstant             = "I cause no compiler errors even if never used!"
)

// Custom type for type-safe enum pattern
type LogLevel int

// Defining enum constants using iota
const (
	LogLevelDebug LogLevel = iota // 0
	LogLevelInfo                  // 1
	LogLevelWarn                  // 2
	LogLevelError                 // 3
)

// Implementing fmt.Stringer interface for LogLevel
func (l LogLevel) String() string {
	switch l {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	default:
		return fmt.Sprintf("LogLevel(%d)", l)
	}
}

func main() {
	fmt.Println("================================================================================")
	fmt.Println("                        COMPREHENSIVE GOLANG CONSTANTS & IOTA GUIDE             ")
	fmt.Println("================================================================================")

	demonstrateBasicsAndCompileTime()
	demonstrateTypedVsUntyped()
	demonstrateArbitraryPrecision()
	demonstrateConstantExpressionsAndBuiltins()
	demonstrateImplicitRepetition()
	demonstrateIotaBasics()
	demonstrateIotaBitmasksAndSizes()
	demonstrateEnumWithCustomTypeAndStringer()
}

// -----------------------------------------------------------------------------
// 1. BASICS & COMPILE-TIME IMMUTABILITY
// -----------------------------------------------------------------------------
//   - In Go, constants are declared using the `const` keyword.
//   - They represent fixed values that are determined at COMPILE TIME and can NEVER
//     be changed at runtime.
//   - Only boolean, rune, integer, floating-point, complex, and string types can
//     be constants.
//   - NOTE: Arrays, slices, maps, structs, channels, and pointers CANNOT be constants!
//   - A constant cannot be assigned the result of a runtime function call:
//     e.g., `const now = time.Now()` is a COMPILE-TIME ERROR!
//
// -----------------------------------------------------------------------------
func demonstrateBasicsAndCompileTime() {
	fmt.Println("\n--- 1. BASICS & COMPILE-TIME IMMUTABILITY ---")

	const maxUsers = 1000
	const appName = "GoService"
	const isProduction = false

	fmt.Printf("Constant maxUsers:     %d\n", maxUsers)
	fmt.Printf("Constant appName:      %s\n", appName)
	fmt.Printf("Constant isProduction: %t\n", isProduction)

	// Immutability:
	// maxUsers = 2000 // COMPILE ERROR: cannot assign to maxUsers (neither addressable nor a variable)

	// Constants cannot take memory addresses:
	// ptr := &maxUsers // COMPILE ERROR: cannot take the address of maxUsers

	fmt.Println("Constants are immutable and not addressable (no pointer to const).")
}

// -----------------------------------------------------------------------------
// 2. TYPED VS. UNTYPED CONSTANTS
// -----------------------------------------------------------------------------
// - TYPED CONSTANTS:
//   - Explicitly declared with a type: `const TypedVal int = 10`
//   - Adhere strictly to Go's type system; cannot be assigned to variables of
//     differing types without explicit conversion.
//
// - UNTYPED CONSTANTS:
//   - Declared without a type: `const UntypedVal = 10`
//   - Has an abstract *kind* (untyped int, untyped float, untyped string, etc.)
//   - Highly flexible: can be assigned to ANY compatible type without explicit casting!
//   - Has a "default type" when assigned to `:=` or `any`:
//   - untyped int    -> int
//   - untyped float  -> float64
//   - untyped string -> string
//   - untyped bool   -> bool
//   - untyped rune   -> rune (int32)
//
// -----------------------------------------------------------------------------
func demonstrateTypedVsUntyped() {
	fmt.Println("\n--- 2. TYPED VS. UNTYPED CONSTANTS ---")

	// Typed constant
	const typedInt int = 42

	// Untyped constant
	const untypedVal = 42

	// Untyped constants can be used across multiple distinct types freely:
	var myInt int = untypedVal
	var myInt64 int64 = untypedVal
	var myFloat float64 = untypedVal
	var myComplex complex128 = untypedVal

	fmt.Printf("untypedVal assigned to int:        %d (%T)\n", myInt, myInt)
	fmt.Printf("untypedVal assigned to int64:      %d (%T)\n", myInt64, myInt64)
	fmt.Printf("untypedVal assigned to float64:    %.1f (%T)\n", myFloat, myFloat)
	fmt.Printf("untypedVal assigned to complex128: %v (%T)\n", myComplex, myComplex)

	// In contrast, typed constant cannot be assigned to another type without cast:
	// var invalid float64 = typedInt // COMPILE ERROR: cannot use typedInt (type int) as float64
	var validFloat = float64(typedInt) // Requires explicit cast
	fmt.Printf("typedInt cast to float64:          %.1f\n", validFloat)
}

// -----------------------------------------------------------------------------
// 3. ARBITRARY PRECISION OF UNTYPED CONSTANTS
// -----------------------------------------------------------------------------
//   - The Go specification guarantees that untyped numeric constants have at
//     least 256 bits of precision at compile time (effectively arbitrary precision).
//   - They can hold values that far exceed the limits of `uint64` or `float64`
//     during intermediate compile-time computations!
//   - Overflow errors only occur if the resulting constant is assigned to a
//     typed variable that is too small to hold it.
//
// -----------------------------------------------------------------------------
func demonstrateArbitraryPrecision() {
	fmt.Println("\n--- 3. ARBITRARY PRECISION OF UNTYPED CONSTANTS ---")

	// 1 << 100 exceeds int64's maximum value (which is 1 << 63 - 1)
	const hugeNumber = 1 << 100 // 1,267,650,600,228,229,401,496,703,205,376

	// We cannot assign hugeNumber directly to an int variable:
	// var overflow int = hugeNumber // COMPILE ERROR: constant 1267650600228229401496703205376 overflows int

	// But we CAN use hugeNumber in intermediate compile-time arithmetic!
	const result = hugeNumber >> 95 // Shifts by 95 bits, leaving 1 << 5 = 32
	var normalInt int = result      // Perfectly valid, because 32 fits in int!

	fmt.Printf("hugeNumber (1 << 100) >> 95 at compile time = %d (type %T)\n", normalInt, normalInt)

	// High precision floating-point constant calculation:
	const exactFraction = 1.0 / 3.0 // Stored with extreme compile-time precision
	var f64 float64 = exactFraction
	fmt.Printf("1.0/3.0 evaluated at compile time: %.18f\n", f64)
}

// -----------------------------------------------------------------------------
// 4. COMPILE-TIME CONSTANT EXPRESSIONS & BUILT-INS
// -----------------------------------------------------------------------------
// Constant expressions are evaluated entirely by the compiler:
// - Operators: `+`, `-`, `*`, `/`, `%`, `&`, `|`, `^`, `<<`, `>>`, `&&`, `||`, `!`
// - Allowed built-in functions:
//   - `len()` and `cap()` on constant strings or fixed-size arrays
//   - `complex()`, `real()`, and `imag()` with constant arguments
//   - `unsafe.Sizeof()`, `unsafe.Alignof()`, `unsafe.Offsetof()`
//
// -----------------------------------------------------------------------------
func demonstrateConstantExpressionsAndBuiltins() {
	fmt.Println("\n--- 4. COMPILE-TIME CONSTANT EXPRESSIONS & BUILT-INS ---")

	const (
		message       = "Hello, Golang"
		messageLength = len(message) // Evaluated at compile time!
		intSize       = unsafe.Sizeof(int(0))
		complexVal    = complex(3.0, 4.0)
		realComponent = real(complexVal)
		mathFormula   = (10 * 2) + (100 / 4) - 5
	)

	fmt.Printf("len(\"Hello, Golang\") at compile time: %d\n", messageLength)
	fmt.Printf("unsafe.Sizeof(int(0)) at compile time: %d bytes\n", intSize)
	fmt.Printf("complex(3.0, 4.0) at compile time:     %v\n", complexVal)
	fmt.Printf("real(complexVal) at compile time:      %v\n", realComponent)
	fmt.Printf("mathFormula:                          %d\n", mathFormula)
}

// -----------------------------------------------------------------------------
// 5. IMPLICIT REPETITION IN CONSTANT BLOCKS
// -----------------------------------------------------------------------------
//   - In a parenthesized `const (...)` block, if an expression is omitted after
//     the first constant, Go automatically REPEATS the preceding expression and type!
//
// -----------------------------------------------------------------------------
func demonstrateImplicitRepetition() {
	fmt.Println("\n--- 5. IMPLICIT REPETITION IN CONSTANT BLOCKS ---")

	const (
		A = "StaticValue"
		B // Inherits expression = "StaticValue"
		C // Inherits expression = "StaticValue"
		D = 100
		E // Inherits expression = 100
	)

	fmt.Printf("A = %q, B = %q, C = %q\n", A, B, C)
	fmt.Printf("D = %d, E = %d\n", D, E)
}

// -----------------------------------------------------------------------------
// 6. `iota` - THE CONSTANT GENERATOR
// -----------------------------------------------------------------------------
//   - `iota` is a predeclared identifier used in `const` blocks to generate
//     successive untyped integer values.
//   - KEY RULES:
//     1. `iota` starts at `0` on the first line of a `const` block.
//     2. `iota` increments by `1` for each subsequent specification line.
//     3. It resets to `0` whenever the keyword `const` appears.
//     4. Can be used in mathematical expressions (e.g. `iota * 10`, `1 << iota`).
//     5. Skipping values: Use the blank identifier `_ = iota` to ignore values.
//
// -----------------------------------------------------------------------------
func demonstrateIotaBasics() {
	fmt.Println("\n--- 6. IOTA CONSTANT GENERATOR ---")

	// 1. Basic 0-indexed enumeration (Days of week)
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Printf("Days: Sun=%d, Mon=%d, Tue=%d, Wed=%d, Thu=%d, Fri=%d, Sat=%d\n",
		Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	// 2. Skipping zero with blank identifier (useful when 0 represents an invalid state)
	const (
		_      = iota // Skip 0
		First         // 1
		Second        // 2
		Third         // 3
	)
	fmt.Printf("Skipped zero: First=%d, Second=%d, Third=%d\n", First, Second, Third)

	// 3. Custom formula with iota (e.g., status codes or step increments)
	const (
		StatusInit    = (iota * 10) + 100 // (0 * 10) + 100 = 100
		StatusRunning                     // (1 * 10) + 100 = 110
		StatusPaused                      // (2 * 10) + 100 = 120
		StatusStopped                     // (3 * 10) + 100 = 130
	)
	fmt.Printf("Custom formula: Init=%d, Running=%d, Paused=%d, Stopped=%d\n",
		StatusInit, StatusRunning, StatusPaused, StatusStopped)

	// 4. Multiple iota on the same line
	// Note: iota only increments once per specification line!
	const (
		c0, c1 = iota, iota + 10 // iota is 0: c0 = 0, c1 = 10
		c2, c3                   // iota is 1: c2 = 1, c3 = 11
	)
	fmt.Printf("Multi-iota per line: c0=%d, c1=%d | c2=%d, c3=%d\n", c0, c1, c2, c3)
}

// -----------------------------------------------------------------------------
// 7. BITMASKS & SIZES USING IOTA AND BIT SHIFTS
// -----------------------------------------------------------------------------
// The combination of `iota` and the left shift operator `<<` is one of Go's
// most famous idioms for creating bitmasks and digital storage multipliers:
//
// 1. Digital Storage Units (KB, MB, GB, TB, PB):
//   - `1 << (10 * iota)` shifts by 10 bits (multiply by 1024) per line.
//
// 2. Bitwise Permission Flags:
//   - `1 << iota` creates distinct bit flags: 1, 2, 4, 8, 16...
//
// -----------------------------------------------------------------------------
func demonstrateIotaBitmasksAndSizes() {
	fmt.Println("\n--- 7. BITMASKS & SIZES WITH IOTA ---")

	// Digital storage unit multipliers
	const (
		_  = 1 << (10 * iota) // 1 << (10 * 0) = 1 (ignored)
		KB                    // 1 << (10 * 1) = 1,024
		MB                    // 1 << (10 * 2) = 1,048,576
		GB                    // 1 << (10 * 3) = 1,073,741,824
		TB                    // 1 << (10 * 4) = 1,099,511,627,776
	)

	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)
	fmt.Printf("1 TB = %d bytes\n", TB)

	// Bitwise flag permissions
	const (
		PermissionRead    = 1 << iota // 1 << 0 = 0001 (1)
		PermissionWrite               // 1 << 1 = 0010 (2)
		PermissionExecute             // 1 << 2 = 0100 (4)
	)

	// Combining permissions using Bitwise OR (|)
	userRole := PermissionRead | PermissionWrite
	fmt.Printf("\nPermissionRead:    %04b (%d)\n", PermissionRead, PermissionRead)
	fmt.Printf("PermissionWrite:   %04b (%d)\n", PermissionWrite, PermissionWrite)
	fmt.Printf("PermissionExecute: %04b (%d)\n", PermissionExecute, PermissionExecute)
	fmt.Printf("User Role:         %04b (%d)\n", userRole, userRole)

	// Checking permissions using Bitwise AND (&)
	hasRead := (userRole & PermissionRead) != 0
	hasWrite := (userRole & PermissionWrite) != 0
	hasExec := (userRole & PermissionExecute) != 0
	fmt.Printf("Has Read: %t | Has Write: %t | Has Execute: %t\n", hasRead, hasWrite, hasExec)
}

// -----------------------------------------------------------------------------
// 8. TYPE-SAFE ENUMS & THE STRINGER INTERFACE
// -----------------------------------------------------------------------------
//   - In Go, enums are created by defining a custom type on top of an integer type
//     combined with `const` and `iota`.
//   - BENEFIT: Compile-time type safety preventing accidental assignment of raw ints.
//   - STRINGER PATTERN:
//   - By implementing `String() string` on the custom enum type, formatting
//     functions (like `fmt.Println` and `fmt.Printf("%v")`) print descriptive names
//     instead of magic numbers!
//
// -----------------------------------------------------------------------------
func demonstrateEnumWithCustomTypeAndStringer() {
	fmt.Println("\n--- 8. TYPE-SAFE ENUMS & STRINGER PATTERN ---")

	var currentLevel LogLevel = LogLevelWarn

	// Prints human-readable "WARN" because LogLevel implements fmt.Stringer
	fmt.Printf("Current log level: %v (underlying integer: %d, type: %T)\n",
		currentLevel, int(currentLevel), currentLevel)

	// Function receiving the enum type ensures compile-time safety
	logMessage := func(level LogLevel, msg string) {
		fmt.Printf("[%s] %s\n", level, msg)
	}

	logMessage(LogLevelInfo, "Application started successfully.")
	logMessage(LogLevelError, "Database connection failed!")

	fmt.Println("\n================================================================================")
	fmt.Println("                        END OF CONSTANTS & IOTA CONCEPTS GUIDE                  ")
	fmt.Println("================================================================================")
}
