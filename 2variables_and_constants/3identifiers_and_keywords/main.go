package main

import (
	"fmt"
	"unicode"
)

// =============================================================================
// PACKAGE-LEVEL IDENTIFIERS (EXPORTED VS. UNEXPORTED)
// =============================================================================
// In Go, visibility (encapsulation) is defined entirely by the capitalization
// of the identifier's FIRST letter:
//
// 1. EXPORTED (Public):
//    - First character is an uppercase letter (e.g., `ExportedService`).
//    - Accessible from any package that imports this package.
//
// 2. UNEXPORTED (Package-Private):
//    - First character is lowercase or an underscore (e.g., `internalConfig`).
//    - Accessible ONLY within files of this same package (`main`).
// =============================================================================

// ExportedStruct is accessible outside this package
type ExportedStruct struct {
	ExportedField   string // Accessible from other packages
	unexportedField int    // Accessible only inside this package
}

// ExportedConstant is visible outside this package
const ExportedConstant = "I am public"

// unexportedConstant is visible only within package main
const unexportedConstant = "I am package-private"

func main() {
	fmt.Println("================================================================================")
	fmt.Println("                   COMPREHENSIVE GOLANG IDENTIFIERS & KEYWORDS GUIDE            ")
	fmt.Println("================================================================================")

	demonstrateIdentifierRules()
	demonstrateReservedKeywords()
	demonstratePredeclaredVsKeywords()
	demonstrateExportedVsUnexported()
	demonstrateBlankIdentifier()
	demonstrateNamingConventions()
	demonstrateLabelsAsIdentifiers()
}

// -----------------------------------------------------------------------------
// 1. IDENTIFIER LEXICAL RULES
// -----------------------------------------------------------------------------
// An identifier is an arbitrary name given to program entities such as:
// variables, constants, types, functions, packages, labels, struct fields, and methods.
//
// LEXICAL RULES (Go Language Specification):
//  1. Must begin with a Unicode letter (A-Z, a-z) or an underscore `_`.
//  2. Cannot begin with a digit (0-9) - e.g., `1item` is a compile error!
//  3. May contain subsequent Unicode letters, digits (0-9), and underscores.
//  4. Case-sensitive: `myVariable`, `MyVariable`, and `MYVARIABLE` are 3 distinct identifiers.
//  5. Unicode support: Go source code is UTF-8 encoded; non-ASCII letters
//     (such as Greek letters, accented characters, or Asian characters) are legal!
//
// -----------------------------------------------------------------------------
func demonstrateIdentifierRules() {
	fmt.Println("\n--- 1. IDENTIFIER LEXICAL RULES ---")

	// Standard ASCII identifiers
	userCount := 42
	_internalID := "A-123"
	item2 := "Widget"

	// Case sensitivity demonstration
	value := 10
	Value := 20
	VALUE := 30
	fmt.Printf("Case sensitivity: value=%d, Value=%d, VALUE=%d\n", value, Value, VALUE)
	fmt.Printf("Valid standard identifiers: %d, %s, %s\n", userCount, _internalID, item2)

	// Unicode identifiers (legal in Go, though ASCII is idiomatic in production)
	const π = 3.141592653589793
	résumé := "Software Engineer"
	変数 := "Variable in Japanese"

	fmt.Printf("Unicode identifier π:      %v\n", π)
	fmt.Printf("Unicode identifier résumé: %s\n", résumé)
	fmt.Printf("Unicode identifier 変数:   %s\n", 変数)
}

// -----------------------------------------------------------------------------
// 2. THE 25 RESERVED KEYWORDS
// -----------------------------------------------------------------------------
// Go is intentionally designed to be small and simple. It has only 25 reserved
// keywords that cannot be used as identifiers (variable names, types, functions, etc.):
//
//	Category              Keywords
//	-------------------   -----------------------------------------------------
//	Package & Imports   : package, import
//	Declarations        : func, type, var, const
//	Control Flow        : if, else, switch, case, default, for, range,
//	                      break, continue, fallthrough, goto, return
//	Concurrency         : go, select, chan
//	Flow Deferral       : defer
//	Composite Types     : struct, interface, map
//
// Attempting to declare a variable with any reserved keyword causes a compile error:
//
//	var default = 10 // COMPILE ERROR: syntax error: unexpected default
//
// -----------------------------------------------------------------------------
func demonstrateReservedKeywords() {
	fmt.Println("\n--- 2. THE 25 RESERVED KEYWORDS ---")

	keywords := []string{
		"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var",
	}

	fmt.Printf("Total reserved keywords in Go: %d\n", len(keywords))
	for i := 0; i < len(keywords); i += 5 {
		end := i + 5
		if end > len(keywords) {
			end = len(keywords)
		}
		fmt.Printf("  %-10s %-10s %-10s %-10s %-10s\n",
			keywords[i], keywords[i+1], keywords[i+2], keywords[i+3], keywords[i+4])
	}
}

// -----------------------------------------------------------------------------
// 3. PREDECLARED IDENTIFIERS (AND THE SHADOWING PITFALL)
// -----------------------------------------------------------------------------
// Go provides many built-in identifiers declared in the "universe block":
// - Types      : int, float64, string, bool, byte, rune, error, any, etc.
// - Constants  : true, false, iota, nil
// - Functions  : len, cap, make, new, append, copy, delete, panic, recover, etc.
//
// CRITICAL DISTINCTION:
//   - Predeclared identifiers are NOT reserved keywords!
//   - They CAN be shadowed (redeclared) in local scopes!
//   - SHADOWING GOTCHA: While syntactically legal, shadowing predeclared identifiers
//     is considered a dangerous anti-pattern because it hides the built-in type/function.
//
// -----------------------------------------------------------------------------
func demonstratePredeclaredVsKeywords() {
	fmt.Println("\n--- 3. PREDECLARED IDENTIFIERS & SHADOWING ---")

	// Legal standard usage
	standardSlice := make([]int, 0, 5)
	fmt.Printf("Standard built-in len(): %d\n", len(standardSlice))

	// Demonstration of shadowing in an isolated inner scope:
	{
		// Shadowing the built-in function `len` with a local variable:
		len := 42
		fmt.Printf("  [Inside Block] Shadowed 'len' as an integer: %d\n", len)
		// Now len("hello") would fail inside this block because 'len' is an int!

		// Shadowing the built-in constant `true`:
		true := false
		fmt.Printf("  [Inside Block] Shadowed 'true' as false: %t\n", true)
	}

	// Back outside the block, the built-in identifiers are intact:
	fmt.Printf("Back outside block: len(\"Go\") = %d, true = %t\n", len("Go"), true)
}

// -----------------------------------------------------------------------------
// 4. EXPORTED VS. UNEXPORTED IDENTIFIERS
// -----------------------------------------------------------------------------
// Go replaces traditional access modifiers (`public`, `private`, `protected`)
// with a simple character case rule:
//
// 1. Starts with UPPERCASE letter -> Exported (visible across packages).
// 2. Starts with lowercase letter or `_` -> Unexported (package-private).
//
// This rule applies uniformly across:
// - Functions & Methods (`CalculateTotal` vs `calculateTotal`)
// - Types (`User` vs `user`)
// - Struct Fields (`ID` vs `internalID`)
// - Variables & Constants (`MaxRetry` vs `maxRetry`)
// -----------------------------------------------------------------------------
func demonstrateExportedVsUnexported() {
	fmt.Println("\n--- 4. EXPORTED VS. UNEXPORTED IDENTIFIERS ---")

	instance := ExportedStruct{
		ExportedField:   "Public value",
		unexportedField: 101, // Accessible here because we are inside package main
	}

	fmt.Printf("ExportedStruct instance: %+v\n", instance)
	fmt.Printf("Package-level exported const:   %s\n", ExportedConstant)
	fmt.Printf("Package-level unexported const: %s\n", unexportedConstant)

	// Check if identifier name is exported programmatically using unicode package:
	isExported := func(name string) bool {
		if len(name) == 0 {
			return false
		}
		r := []rune(name)[0]
		return unicode.IsUpper(r)
	}

	testNames := []string{"ServerPort", "timeout", "DBConn", "_hidden"}
	for _, name := range testNames {
		fmt.Printf("  Identifier %-12q is exported: %t\n", name, isExported(name))
	}
}

// -----------------------------------------------------------------------------
// 5. THE BLANK IDENTIFIER (`_`)
// -----------------------------------------------------------------------------
// The underscore `_` is a special identifier known as the "blank identifier":
// - It acts as an anonymous, write-only placeholder.
// - It discards values and cannot be read (e.g. `val := _` is a compile error).
//
// COMMON USES:
//  1. Discarding unwanted multiple return values.
//  2. Ignoring index/value in `for range` loops.
//  3. Side-effect imports: `import _ "net/http/pprof"`.
//  4. Compile-time interface implementation verification:
//     `var _ fmt.Stringer = (*MyType)(nil)`
//  5. Discarding unused imports or variables during active debugging.
//
// -----------------------------------------------------------------------------
type MockStringer struct{}

func (m MockStringer) String() string {
	return "MockStringer verified"
}

func demonstrateBlankIdentifier() {
	fmt.Println("\n--- 5. THE BLANK IDENTIFIER (_) ---")

	// 1. Discarding multiple returns
	divMod := func(n, d int) (int, int) {
		return n / d, n % d
	}
	quotient, _ := divMod(10, 3) // Discards remainder
	fmt.Printf("Quotient: %d (remainder discarded)\n", quotient)

	// 2. Ignoring index in a range loop
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Print("Iterating without index: ")
	for _, name := range names {
		fmt.Printf("%s ", name)
	}
	fmt.Println()

	// 3. Compile-time interface implementation check:
	// If MockStringer does not implement fmt.Stringer, this line fails at compile time!
	var _ fmt.Stringer = (*MockStringer)(nil)
	fmt.Println("Compile-time interface verification succeeded: MockStringer implements fmt.Stringer")
}

// -----------------------------------------------------------------------------
// 6. GO NAMING CONVENTIONS & IDIOMS (EFFECTIVE GO)
// -----------------------------------------------------------------------------
// Go has strong conventions established in "Effective Go" and "Go Code Review Comments":
//
// 1. MixedCaps / CamelCase:
//   - Go avoids snake_case (`user_id` is unidiomatic).
//   - Use `camelCase` for unexported and `PascalCase` for exported.
//
// 2. Acronyms & Initialisms:
//   - Keep acronyms consistent in casing!
//   - Write `userID` (or `UserID`), NOT `userId` or `UserId`.
//   - Write `httpServer`, `parseURL`, `apiClient`, `xmlDocument`.
//
// 3. Name Length Proportional to Scope:
//   - Short scope -> short name: `i` for loops, `r` for reader, `err` for error, `ctx` for context.
//   - Large scope (package-level) -> descriptive name: `DefaultHTTPClientTimeout`.
//
// 4. Getters and Setters:
//   - Do NOT prefix getters with "Get":
//     Use `obj.Owner()` instead of `obj.GetOwner()`.
//     Use `obj.SetOwner(val)` for the setter.
//
// -----------------------------------------------------------------------------
func demonstrateNamingConventions() {
	fmt.Println("\n--- 6. GO NAMING CONVENTIONS & IDIOMS ---")

	// Correct casing for initialisms:
	userID := "usr_9981"
	apiURL := "https://api.github.com"
	httpStatusCode := 200

	fmt.Printf("Idiomatic initialisms: userID=%s, apiURL=%s, status=%d\n",
		userID, apiURL, httpStatusCode)

	fmt.Println("Conventions summary:")
	fmt.Println("  - Use camelCase / PascalCase, never snake_case.")
	fmt.Println("  - Consistent acronym casing: 'jsonDoc', 'parseXML', 'fetchHTML'.")
	fmt.Println("  - Name length proportional to scope (short loop index 'i', clear package var).")
	fmt.Println("  - Getters omit 'Get' prefix: user.Name(), not user.GetName().")
}

// -----------------------------------------------------------------------------
// 7. LABELS AS IDENTIFIERS
// -----------------------------------------------------------------------------
// - Labels are identifiers used as targets for `break`, `continue`, and `goto`.
// - Labels have their OWN separate namespace:
//   - A label name will not conflict with a variable or function of the same name.
//
// - Labels are scoped to the function in which they are declared.
// -----------------------------------------------------------------------------
func demonstrateLabelsAsIdentifiers() {
	fmt.Println("\n--- 7. LABELS AS IDENTIFIERS ---")

	// OuterLoop is a label identifier for breaking out of nested loops
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Printf("Breaking out of OuterLoop at i=%d, j=%d\n", i, j)
				break OuterLoop
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}

	fmt.Println("\n================================================================================")
	fmt.Println("                  END OF IDENTIFIERS & KEYWORDS CONCEPTS GUIDE                  ")
	fmt.Println("================================================================================")
}
