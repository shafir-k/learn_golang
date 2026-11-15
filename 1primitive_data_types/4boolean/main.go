package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println("================================================================================")
	fmt.Println("                         COMPREHENSIVE GOLANG BOOLEANS GUIDE                    ")
	fmt.Println("================================================================================")

	demonstrateBasicsAndMemory()
	demonstrateLogicalOperators()
	demonstrateShortCircuitEvaluation()
	demonstrateComparisonOperators()
	demonstrateControlFlow()
	demonstrateConversions()
	demonstrateBooleanIdioms()
	demonstrateFormatting()
}

// -----------------------------------------------------------------------------
// 1. BASICS & MEMORY CHARACTERISTICS
// -----------------------------------------------------------------------------
// - In Go, `bool` is a built-in primitive data type representing a logical truth value.
// - It can only hold one of two predeclared constant values: `true` or `false`.
// - The zero value of a `bool` variable is `false`.
// - MEMORY SIZE:
//   - A boolean in Go occupies 1 byte (8 bits) of memory, NOT 1 bit.
//   - Why 1 byte? Hardware architectures are byte-addressable; individual bits
//     cannot be addressed directly in memory without bitwise masking.
// - STRICT TYPE SYSTEM (NO "TRUTHY" / "FALSY"):
//   - Go does NOT support implicit conversions to or from boolean.
//   - Unlike C, C++, or JavaScript, `0`, `1`, `""`, or `nil` are NOT booleans.
//   - `if 1 { ... }` or `bool(1)` will fail at compile time.
// -----------------------------------------------------------------------------
func demonstrateBasicsAndMemory() {
	fmt.Println("\n--- 1. BASICS & MEMORY CHARACTERISTICS ---")

	// Zero value
	var uninitializedBool bool
	fmt.Printf("Zero value of bool: %t\n", uninitializedBool)

	// Explicit assignment
	isActive := true
	isComplete := false
	fmt.Printf("isActive: %t, isComplete: %t\n", isActive, isComplete)

	// Memory size of bool
	fmt.Printf("Size of bool in memory: %d byte(s)\n", unsafe.Sizeof(isActive))

	// Go's strict typing:
	// num := 1
	// if num { }      // COMPILE ERROR: non-boolean condition in if statement
	// b := bool(num)  // COMPILE ERROR: cannot convert num (type int) to type bool
	fmt.Println("Strict typing: integers and pointers cannot be coerced to bool implicitly.")
}

// -----------------------------------------------------------------------------
// 2. LOGICAL OPERATORS
// -----------------------------------------------------------------------------
// Go provides three logical operators for combining or negating boolean values:
// 1. `!`  : Logical NOT (unary)  - inverts the boolean value.
// 2. `&&` : Logical AND (binary) - returns true only if BOTH operands are true.
// 3. `||` : Logical OR  (binary) - returns true if AT LEAST ONE operand is true.
//
// Operator Precedence:
// - Highest: `!` (unary)
// - Middle:  `&&` (higher precedence than `||`)
// - Lowest:  `||`
//
// De Morgan's Laws:
// - `!(A && B)` is equivalent to `!A || !B`
// - `!(A || B)` is equivalent to `!A && !B`
// -----------------------------------------------------------------------------
func demonstrateLogicalOperators() {
	fmt.Println("\n--- 2. LOGICAL OPERATORS & TRUTH TABLES ---")

	t1, t2 := true, true
	f1, f2 := false, false

	// Logical NOT (!)
	fmt.Println("Logical NOT (!):")
	fmt.Printf("  !true  = %t\n", !t1)
	fmt.Printf("  !false = %t\n", !f1)

	// Logical AND (&&)
	fmt.Println("Logical AND (&&):")
	fmt.Printf("  true  && true  = %t\n", t1 && t2)
	fmt.Printf("  true  && false = %t\n", t1 && f1)
	fmt.Printf("  false && true  = %t\n", f1 && t1)
	fmt.Printf("  false && false = %t\n", f1 && f2)

	// Logical OR (||)
	fmt.Println("Logical OR (||):")
	fmt.Printf("  true  || true  = %t\n", t1 || t2)
	fmt.Printf("  true  || false = %t\n", t1 || f1)
	fmt.Printf("  false || true  = %t\n", f1 || t1)
	fmt.Printf("  false || false = %t\n", f1 || f2)

	// De Morgan's Laws demonstration
	a, b := true, false
	fmt.Println("De Morgan's Laws:")
	fmt.Printf("  !(a && b) == (!a || !b): %t\n", !(a && b) == (!a || !b))
	fmt.Printf("  !(a || b) == (!a && !b): %t\n", !(a || b) == (!a && !b))
}

// -----------------------------------------------------------------------------
// 3. SHORT-CIRCUIT EVALUATION
// -----------------------------------------------------------------------------
// Go evaluates logical expressions from left to right and stops as soon as the
// final outcome is guaranteed:
// - For `A && B`: If `A` is false, `B` is NEVER evaluated (result is already false).
// - For `A || B`: If `A` is true,  `B` is NEVER evaluated (result is already true).
//
// WHY SHORT-CIRCUIT MATTERS:
// 1. Performance: Skips expensive computations when unnecessary.
// 2. Safety / Guard Clauses: Prevents runtime panics (e.g. nil pointer dereferences
//    or index-out-of-range errors) by placing a safety check on the left.
// -----------------------------------------------------------------------------
func demonstrateShortCircuitEvaluation() {
	fmt.Println("\n--- 3. SHORT-CIRCUIT EVALUATION ---")

	// Helper functions with side-effects to show if they execute
	expensiveOperationFalse := func() bool {
		fmt.Println("    [Executed] expensiveOperationFalse()")
		return false
	}
	expensiveOperationTrue := func() bool {
		fmt.Println("    [Executed] expensiveOperationTrue()")
		return true
	}

	fmt.Println("Testing AND (&&) short-circuit when left is false:")
	if expensiveOperationFalse() && expensiveOperationTrue() {
		fmt.Println("  Condition passed")
	} else {
		fmt.Println("  Condition failed (notice expensiveOperationTrue was NOT called!)")
	}

	fmt.Println("\nTesting OR (||) short-circuit when left is true:")
	if expensiveOperationTrue() || expensiveOperationFalse() {
		fmt.Println("  Condition passed (notice expensiveOperationFalse was NOT called!)")
	}

	// Practical Safety Guard Idiom:
	fmt.Println("\nSafe Guard Pattern (Nil Pointer Protection):")
	var ptr *int = nil
	// Because of short-circuiting, *ptr is never evaluated if ptr == nil, preventing a panic!
	if ptr != nil && *ptr > 10 {
		fmt.Println("Pointer value is greater than 10")
	} else {
		fmt.Println("Safely evaluated without nil-pointer panic because ptr is nil.")
	}
}

// -----------------------------------------------------------------------------
// 4. COMPARISON OPERATORS YIELDING BOOLEANS
// -----------------------------------------------------------------------------
// Comparison operators evaluate expressions and return a boolean result:
// - Equality:    `==` (equal), `!=` (not equal)
// - Relational:  `<` (less), `<=` (less or equal), `>` (greater), `>=` (greater or equal)
//
// Comparable types include integers, floats, strings, booleans, pointers,
// channels, interfaces, and arrays/structs (if all their fields are comparable).
// -----------------------------------------------------------------------------
func demonstrateComparisonOperators() {
	fmt.Println("\n--- 4. COMPARISON OPERATORS ---")

	x, y := 10, 20
	fmt.Printf("%d == %d: %t\n", x, y, x == y)
	fmt.Printf("%d != %d: %t\n", x, y, x != y)
	fmt.Printf("%d <  %d: %t\n", x, y, x < y)
	fmt.Printf("%d <= %d: %t\n", x, y, x <= y)
	fmt.Printf("%d >  %d: %t\n", x, y, x > y)
	fmt.Printf("%d >= %d: %t\n", x, y, x >= y)

	// Comparing strings (lexicographical comparison)
	s1, s2 := "apple", "banana"
	fmt.Printf("%q < %q: %t\n", s1, s2, s1 < s2)

	// Comparing boolean values directly
	fmt.Printf("true == false: %t\n", true == false)
	fmt.Printf("true != false: %t\n", true != false)
}

// -----------------------------------------------------------------------------
// 5. CONTROL FLOW USAGE
// -----------------------------------------------------------------------------
// Booleans drive Go's control structures:
// 1. `if` statement: condition must evaluate to a boolean expression.
// 2. `if` with initialization: executes a short statement before testing the bool.
// 3. `for` loop: loops while boolean condition is true.
// 4. Tagless `switch`: evaluates boolean expressions in case clauses (idiomatic Go).
// -----------------------------------------------------------------------------
func demonstrateControlFlow() {
	fmt.Println("\n--- 5. CONTROL FLOW USAGE ---")

	score := 85

	// 1. Standard if-else
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or below")
	}

	// 2. If with short statement initialization
	// The variable `isValid` is scoped only to the if-else block
	if isValid := score > 0 && score <= 100; isValid {
		fmt.Printf("Valid score: %d\n", score)
	}

	// 3. For loop with boolean condition
	counter := 0
	for counter < 3 {
		fmt.Printf("  Loop iteration counter: %d\n", counter)
		counter++
	}

	// 4. Tagless switch (idiomatic alternative to long if-else if chains)
	// In a tagless switch, each case is an expression that evaluates to a bool.
	age := 25
	switch {
	case age < 13:
		fmt.Println("Category: Child")
	case age >= 13 && age < 20:
		fmt.Println("Category: Teenager")
	case age >= 20 && age < 65:
		fmt.Println("Category: Adult")
	default:
		fmt.Println("Category: Senior")
	}
}

// -----------------------------------------------------------------------------
// 6. TYPE CONVERSIONS & STRING PARSING
// -----------------------------------------------------------------------------
// Since Go does not allow direct casting between bool and other primitive types:
// 1. `strconv.FormatBool(b bool) string`:
//    - Converts bool to string ("true" or "false").
// 2. `strconv.ParseBool(str string) (bool, error)`:
//    - Parses string to bool.
//    - Accepts: "1", "t", "T", "true", "TRUE", "True" -> true
//    - Accepts: "0", "f", "F", "false", "FALSE", "False" -> false
// 3. Idiomatic integer <-> boolean conversions:
//    - Bool to Int: `if b { return 1 } return 0`
//    - Int to Bool: `i != 0`
// -----------------------------------------------------------------------------
func demonstrateConversions() {
	fmt.Println("\n--- 6. TYPE CONVERSIONS & PARSING ---")

	// 1. strconv.FormatBool
	flag := true
	strFlag := strconv.FormatBool(flag)
	fmt.Printf("strconv.FormatBool(%t) = %q\n", flag, strFlag)

	// 2. strconv.ParseBool
	validInputs := []string{"true", "1", "T", "false", "0", "F"}
	fmt.Println("strconv.ParseBool test cases:")
	for _, input := range validInputs {
		parsed, err := strconv.ParseBool(input)
		if err == nil {
			fmt.Printf("  ParseBool(%q) = %t\n", input, parsed)
		}
	}

	// Invalid input handling
	_, err := strconv.ParseBool("yes")
	if err != nil {
		fmt.Printf("  ParseBool(\"yes\") returned error: %v\n", err)
	}

	// 3. Manual Int <-> Bool conversion helpers
	boolToInt := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}

	intToBool := func(i int) bool {
		return i != 0
	}

	fmt.Printf("boolToInt(true):  %d\n", boolToInt(true))
	fmt.Printf("boolToInt(false): %d\n", boolToInt(false))
	fmt.Printf("intToBool(42):    %t\n", intToBool(42))
	fmt.Printf("intToBool(0):     %t\n", intToBool(0))
}

// -----------------------------------------------------------------------------
// 7. COMMON BOOLEAN IDIOMS IN GO
// -----------------------------------------------------------------------------
// Go embraces several boolean idioms across its standard library and conventions:
// 1. The "comma-ok" idiom:
//    - Map lookups: `val, ok := myMap[key]` (ok is a bool signaling existence).
//    - Type assertions: `str, ok := val.(string)` (ok indicates success).
//    - Channel receives: `msg, ok := <-ch` (ok indicates channel is open).
// 2. Tri-state / Optional Booleans using pointer `*bool`:
//    - In JSON APIs and configuration objects, a boolean property may be:
//      a) Explicitly true
//      b) Explicitly false
//      c) Unset / omitted (represented as `nil`)
// -----------------------------------------------------------------------------
func demonstrateBooleanIdioms() {
	fmt.Println("\n--- 7. COMMON BOOLEAN IDIOMS ---")

	// 1. The Comma-OK Idiom with Maps
	userRoles := map[string]string{
		"Alice": "Admin",
		"Bob":   "Viewer",
	}

	// Check if key exists
	if role, ok := userRoles["Alice"]; ok {
		fmt.Printf("Alice exists in map! Role: %s\n", role)
	}

	if _, ok := userRoles["Charlie"]; !ok {
		fmt.Println("Charlie does NOT exist in map (ok == false).")
	}

	// 2. Comma-OK Idiom with Type Assertion
	var rawData any = "Hello, Go!"
	if strVal, ok := rawData.(string); ok {
		fmt.Printf("Type assertion successful: %q (len: %d)\n", strVal, len(strVal))
	}

	// 3. Tri-state boolean (*bool pointer)
	type Config struct {
		FeatureEnabled *bool `json:"feature_enabled,omitempty"`
	}

	var cfgDefault Config // FeatureEnabled is nil (unset / unspecified)
	fmt.Printf("Default config feature state: isNil=%t\n", cfgDefault.FeatureEnabled == nil)

	enabled := true
	cfgExplicit := Config{FeatureEnabled: &enabled}
	fmt.Printf("Explicit config feature state: isNil=%t, value=%t\n",
		cfgExplicit.FeatureEnabled == nil, *cfgExplicit.FeatureEnabled)
}

// -----------------------------------------------------------------------------
// 8. FORMATTING BOOLEANS WITH `fmt`
// -----------------------------------------------------------------------------
// Formatter verbs for boolean values:
// - `%t` : Word true or false (standard for booleans)
// - `%v` : Default value format (prints true or false)
// - `%T` : Type representation (prints "bool")
// -----------------------------------------------------------------------------
func demonstrateFormatting() {
	fmt.Println("\n--- 8. FORMATTING BOOLEANS WITH FMT ---")

	b := true
	fmt.Printf("Verb %%t: %t\n", b)
	fmt.Printf("Verb %%v: %v\n", b)
	fmt.Printf("Verb %%T: %T\n", b)

	fmt.Println("\n================================================================================")
	fmt.Println("                         END OF BOOLEAN CONCEPTS GUIDE                          ")
	fmt.Println("================================================================================")
}
