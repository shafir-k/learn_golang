package main

import (
	"fmt"
)

// =============================================================================
// PACKAGE-LEVEL VARIABLES (GLOBAL TO THE PACKAGE)
// =============================================================================
// - Variables declared outside functions exist at package scope.
// - Accessible across all functions in all files belonging to the same package.
// - Naming convention determines visibility (Exported vs Unexported):
//   - Starts with UPPERCASE (e.g., `ExportedConfig`): Public / Exported outside package.
//   - Starts with lowercase (e.g., `packageCounter`): Private / Unexported.
// - NOTE: Short variable declaration (`:=`) is NOT allowed at package level!
// - Package-level variables are NOT required to be used (no "declared and not used" error).
// =============================================================================

// Exported package variable (accessible to external packages)
var ExportedAppName string = "GoVariablesDeepDive"

// Unexported package variable (accessible only within package main)
var packageCounter int = 100

// Grouped/Factored package-level declarations (idiomatic for grouping related state)
var (
	ServerHost string = "localhost"
	ServerPort int    = 8080
	IsDebug    bool   = true
)

func main() {
	fmt.Println("================================================================================")
	fmt.Println("                        COMPREHENSIVE GOLANG VARIABLES GUIDE                    ")
	fmt.Println("================================================================================")

	demonstrateDeclarationForms()
	demonstrateZeroValues()
	demonstrateShortDeclaration()
	demonstrateMultipleAssignments()
	demonstrateScopeAndShadowing()
	demonstrateBlankIdentifier()
	demonstratePointersAndAddress()
	demonstrateEscapeAnalysisConcept()
}

// -----------------------------------------------------------------------------
// 1. VARIABLE DECLARATION FORMS
// -----------------------------------------------------------------------------
// Go offers several ways to declare variables using the `var` keyword:
//
// 1. `var name type`:
//   - Declares a variable with an explicit type, initialized to its ZERO VALUE.
//
// 2. `var name type = value`:
//   - Declares with an explicit type and an initial value.
//
// 3. `var name = value`:
//   - Type is omitted; Go automatically INFERS the type from the right-hand value.
//
// 4. Grouped declaration `var (...)`:
//   - Visually groups related variables, improving readability.
//
// -----------------------------------------------------------------------------
func demonstrateDeclarationForms() {
	fmt.Println("\n--- 1. VARIABLE DECLARATION FORMS ---")

	// 1. Declaration with explicit type (defaults to zero value)
	var age int
	fmt.Printf("Explicit type, zero value: age = %d (type %T)\n", age, age)

	// 2. Declaration with explicit type and initial value
	var maxScore int = 100
	fmt.Printf("Explicit type and value:  maxScore = %d (type %T)\n", maxScore, maxScore)

	// 3. Type inference with var (Go deduces type from value)
	var inferredGreeting = "Hello, Gopher!" // Inferred as string
	var inferredRatio = 0.618               // Inferred as float64
	fmt.Printf("Inferred string: %q (type %T)\n", inferredGreeting, inferredGreeting)
	fmt.Printf("Inferred float:  %v (type %T)\n", inferredRatio, inferredRatio)

	// 4. Factored / Grouped variable declaration block
	var (
		firstName string = "Ada"
		lastName  string = "Lovelace"
		birthYear int    = 1815
	)
	fmt.Printf("Grouped declaration: %s %s, born in %d\n", firstName, lastName, birthYear)
}

// -----------------------------------------------------------------------------
// 2. ZERO VALUES (DEFAULT INITIALIZATION)
// -----------------------------------------------------------------------------
//   - In Go, variables are NEVER left uninitialized with random memory garbage
//     (unlike C or C++).
//   - If a variable is declared without an explicit initial value, Go guarantees
//     it is initialized to its type's "zero value":
//   - Numeric types (int, float, etc.) : 0 / 0.0
//   - Booleans                          : false
//   - Strings                           : "" (empty string)
//   - Pointers, Slices, Maps, Channels  : nil
//   - Functions, Interfaces             : nil
//   - Structs                           : each field gets its respective zero value
//   - Arrays                            : each element gets its respective zero value
//
// -----------------------------------------------------------------------------
func demonstrateZeroValues() {
	fmt.Println("\n--- 2. ZERO VALUES ---")

	var (
		defaultInt     int
		defaultFloat   float64
		defaultBool    bool
		defaultString  string
		defaultPointer *int
		defaultSlice   []int
		defaultMap     map[string]int
	)

	fmt.Printf("Zero value of int:      %d\n", defaultInt)
	fmt.Printf("Zero value of float64:  %v\n", defaultFloat)
	fmt.Printf("Zero value of bool:     %t\n", defaultBool)
	fmt.Printf("Zero value of string:   %q\n", defaultString)
	fmt.Printf("Zero value of pointer:  %v\n", defaultPointer)
	fmt.Printf("Zero value of slice:    %v (isNil: %t)\n", defaultSlice, defaultSlice == nil)
	fmt.Printf("Zero value of map:      %v (isNil: %t)\n", defaultMap, defaultMap == nil)

	// Zero value for structs
	type User struct {
		Name   string
		Age    int
		Active bool
	}
	var defaultUser User
	fmt.Printf("Zero value of struct:   %+v\n", defaultUser)
}

// -----------------------------------------------------------------------------
// 3. SHORT VARIABLE DECLARATION (`:=`)
// -----------------------------------------------------------------------------
// - Syntax: `name := value`
// - RULES & CHARACTERISTICS:
//  1. Can ONLY be used INSIDE functions (cannot be used at package level).
//  2. Type is inferred automatically from the expression on the right.
//  3. Inferred defaults:
//     - Integer literals become `int`
//     - Floating-point literals become `float64`
//     - Boolean literals become `bool`
//     - String literals become `string`
//  4. Redeclaration rule:
//     - At least ONE variable on the left side of `:=` must be newly declared!
//     - Other existing variables in the same scope will simply be REASSIGNED.
//
// - WHEN TO USE `:=` vs `var`:
//   - Use `:=` for short-lived, local variables inside functions when an initial
//     value is known (idiomatic Go).
//   - Use `var` when:
//     1. Declaring package-level variables.
//     2. Desiring the zero value without an initial value (`var buf bytes.Buffer`).
//     3. Explicitly specifying a type different from the inferred default
//     (e.g., `var id int64 = 100`, instead of `id := int64(100)`).
//
// -----------------------------------------------------------------------------
func demonstrateShortDeclaration() {
	fmt.Println("\n--- 3. SHORT VARIABLE DECLARATION (:=) ---")

	// Standard short declaration
	language := "Go"
	version := 1.27
	isModern := true

	fmt.Printf("Language: %s (type %T)\n", language, language)
	fmt.Printf("Version:  %v (type %T)\n", version, version)
	fmt.Printf("Modern:   %t (type %T)\n", isModern, isModern)

	// Redeclaration using `:=`
	// Here, `language` already exists in this scope, but `creator` is NEW.
	// This is legal because at least ONE variable (`creator`) is newly declared.
	language, creator := "Golang", "Robert Griesemer, Rob Pike, Ken Thompson"
	fmt.Printf("Updated language: %s, Creators: %s\n", language, creator)

	// Both existing on := causes a compile error:
	// language, version := "Go 2", 2.0 // COMPILE ERROR: no new variables on left side of :=
}

// -----------------------------------------------------------------------------
// 4. MULTIPLE VARIABLE ASSIGNMENTS & TUPLE UNPACKING
// -----------------------------------------------------------------------------
//   - Go supports declaring and assigning multiple variables simultaneously.
//   - Multiple variables can be of the SAME type or DIFFERENT types.
//   - Enables idiomatic in-place variable swapping without temporary variables:
//     `x, y = y, x`
//
// -----------------------------------------------------------------------------
func demonstrateMultipleAssignments() {
	fmt.Println("\n--- 4. MULTIPLE VARIABLE ASSIGNMENT & SWAPPING ---")

	// Multiple variables of the same explicit type
	var x, y int = 10, 20
	fmt.Printf("Same type with var: x = %d, y = %d\n", x, y)

	// Multiple variables of heterogeneous types with type inference
	var name, rating, active = "Golang", 9.8, true
	fmt.Printf("Mixed types with var: %s (%T), %v (%T), %t (%T)\n",
		name, name, rating, rating, active, active)

	// Multiple variables with short declaration `:=`
	city, population := "Zurich", 434000
	fmt.Printf("Mixed types with :=: %s (%d)\n", city, population)

	// Idiomatic variable swap (no temporary variable needed!)
	a, b := 1, 99
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("After swap:  a = %d, b = %d\n", a, b)
}

// -----------------------------------------------------------------------------
// 5. VARIABLE SCOPE & SHADOWING
// -----------------------------------------------------------------------------
// Go has block-level lexical scoping:
// 1. Package scope: Declared at root of file outside functions.
// 2. Function scope: Declared within a function body.
// 3. Block scope: Declared within `{ ... }` blocks (if, for, switch).
//
// VARIABLE SHADOWING:
//   - If you declare a variable in an inner block with the SAME name as an outer
//     variable, the inner variable SHADOWS (hides) the outer variable.
//   - Changes to the inner shadowed variable do NOT affect the outer variable!
//   - SHADOWING GOTCHA: Often accidentally introduced with `:=` inside an `if` block.
//
// -----------------------------------------------------------------------------
func demonstrateScopeAndShadowing() {
	fmt.Println("\n--- 5. VARIABLE SCOPE & SHADOWING ---")

	outerVar := "I am OUTER"
	fmt.Println("Initial outerVar:", outerVar)

	// New inner block scope
	{
		// Shadowing outerVar with a new local declaration
		outerVar := "I am INNER (shadowed)"
		innerOnlyVar := "I only exist in this block"

		fmt.Println("Inside block:", outerVar)
		fmt.Println("Inside block:", innerOnlyVar)
	}

	// Back in function scope - the original outerVar was completely unaffected!
	fmt.Println("Outside block outerVar:", outerVar)
	// fmt.Println(innerOnlyVar) // COMPILE ERROR: undefined: innerOnlyVar

	// Shadowing gotcha in conditional statements:
	count := 10
	if count := 0; count == 0 {
		// This 'count' is scoped only to the if/else block!
		fmt.Printf("Inside 'if' condition: count = %d (shadows outer count)\n", count)
	}
	fmt.Printf("Outside 'if' statement: count = %d (unmodified outer count)\n", count)
}

// -----------------------------------------------------------------------------
// 6. THE BLANK IDENTIFIER (`_`) & UNUSED VARIABLES RULE
// -----------------------------------------------------------------------------
// - UNUSED VARIABLES RULE:
//   - In Go, a LOCAL variable (declared inside a function) that is never used
//     causes a COMPILE-TIME ERROR (`x declared and not used`).
//   - Go enforces this to keep codebases clean and avoid dead code/allocations.
//
// - THE BLANK IDENTIFIER `_`:
//   - Acts as a write-only discard target.
//   - Used to ignore unwanted return values from multiple-return functions.
//   - Can be used during development to silence the "unused variable" error.
//
// -----------------------------------------------------------------------------
func demonstrateBlankIdentifier() {
	fmt.Println("\n--- 6. THE BLANK IDENTIFIER (_) ---")

	// Helper function returning (result, error-flag)
	calculate := func() (int, bool) {
		return 42, true
	}

	// We only care about the result, discarding the second return value with `_`
	result, _ := calculate()
	fmt.Printf("Result kept: %d (second value discarded with '_')\n", result)

	// Silencing unused variables during development / debugging:
	tempVar := "I might use this later"
	_ = tempVar // Tells compiler: "I acknowledge this variable; do not fail compilation"
}

// -----------------------------------------------------------------------------
// 7. POINTERS AND VARIABLE ADDRESSABILITY
// -----------------------------------------------------------------------------
// - Every variable resides at a specific memory location.
// - `&x` : Address-of operator (returns a pointer `*T` holding memory address of x).
// - `*p` : Dereference operator (accesses or modifies the value at that address).
// - Value semantics vs Pointer semantics:
//   - Assigning `b = a` copies the value.
//   - Passing a pointer allows mutating the original variable.
//
// -----------------------------------------------------------------------------
func demonstratePointersAndAddress() {
	fmt.Println("\n--- 7. POINTERS & VARIABLE ADDRESSABILITY ---")

	original := 100

	// Get memory address with &
	var ptr *int = &original

	fmt.Printf("Variable original value:   %d\n", original)
	fmt.Printf("Variable memory address:   %p\n", &original)
	fmt.Printf("Pointer holds address:     %p\n", ptr)
	fmt.Printf("Dereferencing pointer (*p): %d\n", *ptr)

	// Mutating the variable through its pointer
	*ptr = 250
	fmt.Printf("After *ptr = 250, original = %d (value updated in place!)\n", original)
}

// -----------------------------------------------------------------------------
// 8. LIFETIME & MEMORY ALLOCATION (STACK VS. HEAP & ESCAPE ANALYSIS)
// -----------------------------------------------------------------------------
//   - In languages like C, returning a pointer to a local stack variable causes
//     undefined behavior / dangling pointer bug.
//   - IN GO, THIS IS COMPLETELY SAFE:
//   - Go uses ESCAPE ANALYSIS during compilation.
//   - If a variable's reference "escapes" the function's stack frame (e.g., returned
//     as a pointer), the Go compiler automatically allocates it on the HEAP!
//   - If the variable does not escape, it is efficiently allocated on the STACK.
//
// -----------------------------------------------------------------------------
func createCounter() *int {
	count := 42 // Local variable
	// In C, returning &count would be a fatal bug!
	// In Go, escape analysis detects this and safely allocates 'count' on the heap.
	return &count
}

func demonstrateEscapeAnalysisConcept() {
	fmt.Println("\n--- 8. ESCAPE ANALYSIS & LIFETIME ---")

	counterPtr := createCounter()
	fmt.Printf("Safely accessed heap-allocated variable from returned pointer: %d\n", *counterPtr)

	fmt.Println("\n================================================================================")
	fmt.Println("                        END OF VARIABLES CONCEPTS GUIDE                         ")
	fmt.Println("================================================================================")
}
