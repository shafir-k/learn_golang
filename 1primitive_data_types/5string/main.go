package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

func main() {
	fmt.Println("================================================================================")
	fmt.Println("                         COMPREHENSIVE GOLANG STRINGS GUIDE                     ")
	fmt.Println("================================================================================")

	demonstrateBasicsAndLiterals()
	demonstrateImmutability()
	demonstrateBytesRunesAndUTF8()
	demonstrateIndexingAndSlicing()
	demonstrateIteration()
	demonstrateConversions()
	demonstrateConcatenationAndBuilder()
	demonstrateComparison()
	demonstrateStringsPackage()
	demonstrateUnicodePackage()
	demonstrateFormatting()
	demonstrateInternalStructureAndMemory()
}

// -----------------------------------------------------------------------------
// 1. BASICS & STRING LITERALS
// -----------------------------------------------------------------------------
// - In Go, a string is an immutable sequence of arbitrary bytes (typically UTF-8).
// - The zero value of a string is an empty string: "" (length = 0).
// - Go supports two types of string literals:
//  1. Interpreted String Literals (double quotes: "..."):
//     - Supports escape sequences (\n, \t, \", \\, \xHH, \uXXXX, \UXXXXXXXX).
//     - Cannot span multiple lines unless concatenated.
//  2. Raw String Literals (backticks: `...`):
//     - Can span multiple lines.
//     - Ignores escape characters (all characters are taken literally).
//     - Cannot contain backticks inside.
//     - Ideal for multiline text, JSON templates, regex patterns, and SQL queries.
//
// -----------------------------------------------------------------------------
func demonstrateBasicsAndLiterals() {
	fmt.Println("\n--- 1. BASICS & STRING LITERALS ---")

	// Zero value
	var emptyStr string
	fmt.Printf("Zero value of string: %q (len: %d)\n", emptyStr, len(emptyStr))

	// Interpreted string with escape sequences
	interpreted := "Hello,\tWorld!\nWelcome to \"Go\" programming.\x21"
	fmt.Println("Interpreted string literal:")
	fmt.Println(interpreted)

	// Raw string literal (multiline, preserves whitespace and ignores escape sequences)
	raw := `This is a raw string literal.
Escape characters like \n and \t are NOT processed here.
Great for JSON: {"name": "Alice", "role": "developer"}
Great for Regex: ^[a-zA-Z0-9_]+$`
	fmt.Println("Raw string literal:")
	fmt.Println(raw)
}

// -----------------------------------------------------------------------------
// 2. STRING IMMUTABILITY
// -----------------------------------------------------------------------------
// - Strings in Go are strictly IMMUTABLE.
// - Once created, the underlying bytes cannot be modified directly.
// - For example, `s[0] = 'h'` causes a compile-time error: "cannot assign to s[0]".
// - WHY IMMUTABILITY?
//  1. Safety & Concurrency: Strings can be shared across goroutines without mutexes.
//  2. Zero-copy slicing: Slicing a string produces a new string header pointing
//     to the same backing byte array without copying memory.
//  3. Hash keys: Safe to use as map keys because contents will never mutate.
//
// - HOW TO MODIFY A STRING:
//   - Convert to []byte (if dealing with ASCII/byte data) or []rune (for Unicode).
//   - Mutate the slice, then convert back to string.
//
// -----------------------------------------------------------------------------
func demonstrateImmutability() {
	fmt.Println("\n--- 2. STRING IMMUTABILITY & SAFE MODIFICATION ---")

	original := "Hello, World!"
	fmt.Println("Original string:", original)

	// s[0] = 'h' // COMPILE ERROR: cannot assign to original[0]

	// Modifying via []byte (suitable for ASCII characters)
	byteSlice := []byte(original)
	byteSlice[0] = 'h'
	byteSlice[7] = 'w'
	modifiedFromBytes := string(byteSlice)
	fmt.Println("Modified via []byte:", modifiedFromBytes)

	// Modifying via []rune (safe for Unicode characters)
	greeting := "Hello, 世界"
	runeSlice := []rune(greeting)
	runeSlice[7] = '友' // Replace '界' with '友'
	modifiedFromRunes := string(runeSlice)
	fmt.Println("Modified via []rune:", modifiedFromRunes)
}

// -----------------------------------------------------------------------------
// 3. BYTES VS RUNES VS CHARACTERS (UTF-8 ENCODING DEPTH)
// -----------------------------------------------------------------------------
// - A `byte` is an alias for uint8 (1 byte = 8 bits).
// - A `rune` is an alias for int32 (represents 1 Unicode Code Point).
// - UTF-8 is a variable-length encoding:
//   - ASCII characters: 1 byte (0 to 127)
//   - Accented / Greek / Cyrillic: 2 bytes
//   - Asian characters (CJK): 3 bytes
//   - Emojis / Historic symbols: 4 bytes
//
// - CRITICAL DISTINCTION:
//   - `len(s)` returns the number of BYTES, NOT characters/runes!
//   - `utf8.RuneCountInString(s)` returns the number of Unicode code points (runes).
//
// -----------------------------------------------------------------------------
func demonstrateBytesRunesAndUTF8() {
	fmt.Println("\n--- 3. BYTES VS RUNES VS CHARACTERS (UTF-8) ---")

	asciiText := "Golang"
	unicodeText := "Go 世界 🚀"

	fmt.Printf("Text: %q\n", asciiText)
	fmt.Printf("  Byte count len(): %d\n", len(asciiText))
	fmt.Printf("  Rune count:      %d\n", utf8.RuneCountInString(asciiText))

	fmt.Printf("Text: %q\n", unicodeText)
	fmt.Printf("  Byte count len(): %d (ASCII + 3-byte CJK + 4-byte emoji)\n", len(unicodeText))
	fmt.Printf("  Rune count:      %d\n", utf8.RuneCountInString(unicodeText))

	// Breakdown of individual runes and their byte lengths
	fmt.Println("Character breakdown for 'Go 世界 🚀':")
	for _, r := range unicodeText {
		fmt.Printf("  Rune: %-2c | Unicode: %-7U | Bytes occupied: %d\n", r, r, utf8.RuneLen(r))
	}
}

// -----------------------------------------------------------------------------
// 4. INDEXING AND SLICING (AND UTF-8 PITFALLS)
// -----------------------------------------------------------------------------
// - Indexing (`s[i]`):
//   - Accesses the BYTE at index i, returning uint8 (byte value), NOT the character.
//
// - Slicing (`s[low:high]`):
//   - Slices based on BYTE offsets, returning a new string.
//   - Substring shares the underlying byte array with the original string.
//
// - PITFALL:
//   - Slicing through the middle of a multi-byte UTF-8 character corrupts the
//     character, resulting in invalid UTF-8 bytes (displayed as  / \uFFFD).
//   - SOLUTION FOR UNICODE: Convert to []rune first before slicing.
//
// -----------------------------------------------------------------------------
func demonstrateIndexingAndSlicing() {
	fmt.Println("\n--- 4. INDEXING AND SLICING ---")

	msg := "Golang"
	// Byte indexing on ASCII
	fmt.Printf("msg: %s\n", msg)
	fmt.Printf("msg[0] (byte value): %d, as char: %c\n", msg[0], msg[0])
	fmt.Printf("msg[len-1]:          %d, as char: %c\n", msg[len(msg)-1], msg[len(msg)-1])

	// Slicing on ASCII
	fmt.Printf("msg[0:2]: %s\n", msg[0:2])
	fmt.Printf("msg[2:]:  %s\n", msg[2:])
	fmt.Printf("msg[:4]:  %s\n", msg[:4])

	// Slicing Unicode - The Gotcha!
	uniMsg := "Hello, 世界"
	fmt.Println("\nUnicode text:", uniMsg)
	// '世' starts at byte 7 and occupies 3 bytes (indices 7, 8, 9).
	// '界' occupies 3 bytes (indices 10, 11, 12).
	// Slicing at byte 8 splits '世' in half!
	corrupted := uniMsg[0:8]
	fmt.Printf("Dangerous byte slicing [0:8]: %s (notice corrupted last character)\n", corrupted)

	// Safe Unicode slicing via []rune
	runes := []rune(uniMsg)
	safeSlice := string(runes[0:8]) // First 8 characters (runes)
	fmt.Printf("Safe rune slicing [0:8]:       %s\n", safeSlice)
}

// -----------------------------------------------------------------------------
// 5. STRING ITERATION (LOOPS)
// -----------------------------------------------------------------------------
// - Method 1: Index loop (`for i := 0; i < len(s); i++`)
//   - Iterates BYTE by BYTE.
//   - Does NOT decode UTF-8; multi-byte characters are split across iterations.
//
// - Method 2: Range loop (`for byteIndex, runeValue := range s`)
//   - Decodes UTF-8 automatically, iterating RUNE by RUNE.
//   - KEY OBSERVATION: byteIndex does NOT increment by 1 for multi-byte characters!
//     It advances by the number of bytes that the rune occupies.
//   - Invalid UTF-8 bytes produce utf8.RuneError (\uFFFD).
//
// - Method 3: Converting to []rune for index-based character access.
// -----------------------------------------------------------------------------
func demonstrateIteration() {
	fmt.Println("\n--- 5. STRING ITERATION (LOOPS) ---")

	text := "Hi, 景!"

	fmt.Println("Method 1: Byte-by-byte loop (for i := 0; i < len(s); i++):")
	for i := 0; i < len(text); i++ {
		fmt.Printf("  Byte index %d: 0x%02X (%d)\n", i, text[i], text[i])
	}

	fmt.Println("Method 2: Rune-by-rune loop (for index, runeVal := range s):")
	for byteIndex, runeVal := range text {
		fmt.Printf("  Byte offset %2d: rune %-2c (Unicode: %U, size: %d bytes)\n",
			byteIndex, runeVal, runeVal, utf8.RuneLen(runeVal))
	}

	fmt.Println("Method 3: Converting to []rune for 0-based character indexing:")
	runes := []rune(text)
	for charIndex, r := range runes {
		fmt.Printf("  Character #%d: %c\n", charIndex, r)
	}
}

// -----------------------------------------------------------------------------
// 6. TYPE CONVERSIONS
// -----------------------------------------------------------------------------
// - string <-> []byte:
//   - []byte(s) creates a mutable copy of the underlying bytes.
//   - string(b) creates an immutable string from the byte slice.
//
// - string <-> []rune:
//   - []rune(s) decodes UTF-8 bytes into int32 code points.
//   - string(r) encodes runes into UTF-8 bytes.
//
// - Numbers / Booleans <-> string using the `strconv` standard package:
//   - Itoa / Atoi (Integer <-> ASCII string)
//   - FormatInt / ParseInt (with specific base and bit sizes)
//   - FormatFloat / ParseFloat
//   - FormatBool / ParseBool
//
// - Converting byte/rune to string: `string('A')`, `string(runeVal)`
// -----------------------------------------------------------------------------
func demonstrateConversions() {
	fmt.Println("\n--- 6. TYPE CONVERSIONS ---")

	// 1. String and []byte
	str := "Gophers"
	bytes := []byte(str)
	bytes[0] = 'g'
	strFromBytes := string(bytes)
	fmt.Printf("string -> []byte -> string: %s -> %v -> %s\n", str, bytes, strFromBytes)

	// 2. String and []rune
	uni := "Go 日本語"
	runes := []rune(uni)
	fmt.Printf("string -> []rune: %q -> %v\n", uni, runes)
	fmt.Printf("[]rune -> string: %s\n", string(runes))

	// 3. Single byte or rune to string
	r := '★'
	fmt.Printf("Rune '★' to string: %s\n", string(r))

	// 4. strconv package conversions
	// Int <-> String
	intVal := 2026
	strFromInt := strconv.Itoa(intVal)
	parsedInt, _ := strconv.Atoi(strFromInt)
	fmt.Printf("strconv.Itoa(%d) = %q | strconv.Atoi(%q) = %d\n", intVal, strFromInt, strFromInt, parsedInt)

	// Binary / Hex conversion
	hexStr := strconv.FormatInt(255, 16)
	parsedHex, _ := strconv.ParseInt(hexStr, 16, 64)
	fmt.Printf("FormatInt(255, 16) = %q | ParseInt(%q, 16) = %d\n", hexStr, hexStr, parsedHex)

	// Float <-> String
	floatVal := 3.14159
	strFromFloat := strconv.FormatFloat(floatVal, 'f', 2, 64)
	parsedFloat, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("FormatFloat(%.5f) = %q | ParseFloat(\"3.14\") = %.2f\n", floatVal, strFromFloat, parsedFloat)

	// Bool <-> String
	boolVal := true
	strFromBool := strconv.FormatBool(boolVal)
	parsedBool, _ := strconv.ParseBool("true")
	fmt.Printf("FormatBool(%t) = %q | ParseBool(\"true\") = %t\n", boolVal, strFromBool, parsedBool)
}

// -----------------------------------------------------------------------------
// 7. STRING CONCATENATION & PERFORMANCE
// -----------------------------------------------------------------------------
// - 1. Operator `+` / `+=`:
//   - Convenient for joining a few static strings.
//   - In loops, repeated `+` causes O(N^2) allocations because strings are immutable;
//     a new string is allocated and previous data is copied on every iteration.
//
// - 2. `fmt.Sprintf`:
//   - Great for rich formatting with mixed types, but has reflection overhead.
//
// - 3. `strings.Join(slice, sep)`:
//   - Best when you already have a slice of strings. Precalculates total length
//     and performs a single allocation.
//
// - 4. `strings.Builder`:
//   - The IDIOMATIC, HIGH-PERFORMANCE way to build strings incrementally.
//   - Writes directly to an internal byte buffer without unnecessary allocations.
//   - `Grow(n)` preallocates capacity to avoid dynamic resizing.
//   - `String()` returns the constructed string without copying memory.
//   - NOTE: Do not copy a `strings.Builder` by value after writing to it!
//
// -----------------------------------------------------------------------------
func demonstrateConcatenationAndBuilder() {
	fmt.Println("\n--- 7. STRING CONCATENATION & STRINGS.BUILDER ---")

	// 1. Using '+' operator
	s1 := "Hello"
	s2 := "Golang"
	concat := s1 + ", " + s2 + "!"
	fmt.Println("Concatenation with '+':", concat)

	// 2. Using fmt.Sprintf
	formatted := fmt.Sprintf("%s has %d letters and score %.1f", "Go", 2, 99.5)
	fmt.Println("Concatenation with fmt.Sprintf:", formatted)

	// 3. Using strings.Join
	parts := []string{"Fast", "Concurrent", "Garbage-Collected"}
	joined := strings.Join(parts, " | ")
	fmt.Println("strings.Join:", joined)

	// 4. Using strings.Builder (Recommended for loops)
	var builder strings.Builder

	// Preallocate memory if approximate size is known (huge performance gain)
	builder.Grow(64)

	builder.WriteString("Building: [")
	for i := 1; i <= 5; i++ {
		if i > 1 {
			builder.WriteString(", ")
		}
		builder.WriteString(strconv.Itoa(i))
	}
	builder.WriteByte(']')
	builder.WriteRune(' ')
	builder.WriteRune('✨')

	result := builder.String()
	fmt.Println("strings.Builder result:", result)
	fmt.Printf("Builder length: %d bytes, capacity: %d bytes\n", builder.Len(), builder.Cap())
}

// -----------------------------------------------------------------------------
// 8. STRING COMPARISON & EQUALITY
// -----------------------------------------------------------------------------
//   - Equality: `==` and `!=` compare byte-for-byte equality.
//   - Lexicographical order: `<`, `<=`, `>`, `>=` compare lexicographically
//     based on byte values (ASCII / UTF-8 code point order).
//   - Case-Insensitive Equality: `strings.EqualFold(s1, s2)`
//   - Performs proper Unicode case-folding (handles English and international characters).
//   - Much faster and cleaner than `strings.ToLower(s1) == strings.ToLower(s2)`,
//     as it does not allocate new strings.
//   - `strings.Compare(a, b)`:
//   - Returns 0 if a == b, -1 if a < b, +1 if a > b.
//   - Note: Standard operators `<` and `==` are preferred in idiomatic Go.
//
// -----------------------------------------------------------------------------
func demonstrateComparison() {
	fmt.Println("\n--- 8. STRING COMPARISON & EQUALITY ---")

	a := "apple"
	b := "banana"
	c := "apple"

	fmt.Printf("%q == %q: %t\n", a, c, a == c)
	fmt.Printf("%q != %q: %t\n", a, b, a != b)
	fmt.Printf("%q <  %q: %t (lexicographical order)\n", a, b, a < b)

	// Case-sensitive comparison
	lower := "golang"
	upper := "GOLANG"
	fmt.Printf("%q == %q: %t\n", lower, upper, lower == upper)

	// Case-insensitive comparison via strings.EqualFold
	fmt.Printf("strings.EqualFold(%q, %q): %t\n", lower, upper, strings.EqualFold(lower, upper))

	// strings.Compare
	fmt.Printf("strings.Compare(%q, %q): %d\n", a, b, strings.Compare(a, b))
}

// -----------------------------------------------------------------------------
// 9. THE `strings` STANDARD LIBRARY PACKAGE
// -----------------------------------------------------------------------------
// The `strings` package provides dozens of high-performance utility functions:
// - Searching: Contains, ContainsAny, HasPrefix, HasSuffix, Index, LastIndex, Count
// - Splitting: Split, SplitN, Fields (whitespace-aware)
// - Transformations: ToUpper, ToLower, Repeat, Replace, ReplaceAll, Map
// - Trimming: Trim, TrimSpace, TrimPrefix, TrimSuffix, TrimLeft, TrimRight
// -----------------------------------------------------------------------------
func demonstrateStringsPackage() {
	fmt.Println("\n--- 9. STRINGS STANDARD PACKAGE UTILITIES ---")

	sample := "  Go is expressive, concise, clean, and efficient. Go is fun!  "

	// Searching & Checking
	fmt.Printf("Contains 'clean':        %t\n", strings.Contains(sample, "clean"))
	fmt.Printf("ContainsAny 'xyz':       %t\n", strings.ContainsAny(sample, "xyz"))
	fmt.Printf("HasPrefix (after trim):  %t\n", strings.HasPrefix(strings.TrimSpace(sample), "Go"))
	fmt.Printf("HasSuffix (after trim):  %t\n", strings.HasSuffix(strings.TrimSpace(sample), "fun!"))
	fmt.Printf("Index of 'expressive':   %d\n", strings.Index(sample, "expressive"))
	fmt.Printf("LastIndex of 'Go':       %d\n", strings.LastIndex(sample, "Go"))
	fmt.Printf("Count occurrences of 'Go': %d\n", strings.Count(sample, "Go"))

	// Trimming
	trimmed := strings.TrimSpace(sample)
	fmt.Printf("TrimSpace:               %q\n", trimmed)
	fmt.Printf("TrimPrefix 'Go is ':     %q\n", strings.TrimPrefix(trimmed, "Go is "))
	fmt.Printf("TrimSuffix '!':          %q\n", strings.TrimSuffix(trimmed, "!"))
	fmt.Printf("Trim cutset '! .':       %q\n", strings.Trim(trimmed, "! ."))

	// Transformations
	fmt.Printf("ToUpper:                 %q\n", strings.ToUpper("hello world"))
	fmt.Printf("ToLower:                 %q\n", strings.ToLower("HELLO WORLD"))
	fmt.Printf("Repeat 'Go! ':           %q\n", strings.Repeat("Go! ", 3))
	fmt.Printf("Replace (first 1):       %q\n", strings.Replace(trimmed, "Go", "Golang", 1))
	fmt.Printf("ReplaceAll:              %q\n", strings.ReplaceAll(trimmed, "Go", "Golang"))

	// Custom mapping via strings.Map (e.g. rot13 or casing filter)
	rot13 := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		}
		return r
	}, "Hello")
	fmt.Printf("strings.Map (rot13):     %q -> %q\n", "Hello", rot13)

	// Splitting
	csvData := "apple,banana,orange,grape"
	fruits := strings.Split(csvData, ",")
	fmt.Printf("Split by comma:          %#v\n", fruits)

	// Fields (splits by any whitespace: space, tab, newline)
	whitespaceText := "The   quick\tbrown\nfox"
	words := strings.Fields(whitespaceText)
	fmt.Printf("Fields (by whitespace):  %#v\n", words)
}

// -----------------------------------------------------------------------------
// 10. UNICODE & `unicode/utf8` PACKAGE
// -----------------------------------------------------------------------------
// - `unicode/utf8` provides UTF-8 validation, encoding, and decoding functions.
// - `unicode` provides character classification functions (letter, digit, space, etc.).
// -----------------------------------------------------------------------------
func demonstrateUnicodePackage() {
	fmt.Println("\n--- 10. UNICODE & UNICODE/UTF8 PACKAGES ---")

	validStr := "Hello, 世界 🌍"
	invalidStr := string([]byte{0xff, 0xfe, 0xfd}) // Invalid UTF-8 byte sequence

	// UTF-8 Validation
	fmt.Printf("Is %q valid UTF-8?   %t\n", validStr, utf8.ValidString(validStr))
	fmt.Printf("Is invalidStr valid UTF-8? %t\n", utf8.ValidString(invalidStr))

	// Decode first rune and last rune
	firstRune, firstRuneSize := utf8.DecodeRuneInString(validStr)
	lastRune, lastRuneSize := utf8.DecodeLastRuneInString(validStr)
	fmt.Printf("First rune: %c (size: %d bytes)\n", firstRune, firstRuneSize)
	fmt.Printf("Last rune:  %c (size: %d bytes)\n", lastRune, lastRuneSize)

	// Character classification using `unicode` package
	testChars := []rune{'A', '9', ' ', 'ñ', '世', '!', '\n'}
	fmt.Println("Inspecting runes with `unicode` package:")
	for _, ch := range testChars {
		fmt.Printf("  Rune: %-2q | Letter: %-5t | Digit: %-5t | Space: %-5t | Upper: %-5t | Lower: %-5t\n",
			ch,
			unicode.IsLetter(ch),
			unicode.IsDigit(ch),
			unicode.IsSpace(ch),
			unicode.IsUpper(ch),
			unicode.IsLower(ch),
		)
	}
}

// -----------------------------------------------------------------------------
// 11. STRING FORMATTING WITH `fmt`
// -----------------------------------------------------------------------------
// Common formatting verbs for strings and characters:
// - %s : Unquoted string
// - %q : Quoted string with escaped special characters
// - %+q: Quoted string with non-ASCII escaped to \uXXXX
// - %x : Hexadecimal representation of each byte (lowercase)
// - %X : Hexadecimal representation of each byte (uppercase)
// - %c : Formats integer/rune as Unicode character
// - %U : Formats rune as Unicode code point (e.g. U+0041)
// - %T : Type of the variable
// - Padding & width: %10s (right-aligned), %-10s (left-aligned)
// -----------------------------------------------------------------------------
func demonstrateFormatting() {
	fmt.Println("\n--- 11. STRING FORMATTING WITH FMT ---")

	word := "Hello, 世界"

	fmt.Printf("%%s  (plain):        %s\n", word)
	fmt.Printf("%%q  (quoted):       %q\n", word)
	fmt.Printf("%%+q (escaped uni):  %+q\n", word)
	fmt.Printf("%%x  (hex bytes):    %x\n", word)
	fmt.Printf("%%X  (HEX bytes):    %X\n", word)
	fmt.Printf("%%T  (data type):    %T\n", word)

	// Rune formatting
	r := '世'
	fmt.Printf("Rune %%c (character):  %c\n", r)
	fmt.Printf("Rune %%U (Unicode):    %U\n", r)
	fmt.Printf("Rune %%+q:             %+q\n", r)

	// Alignment and width
	item := "Go"
	fmt.Printf("Right-aligned (width 10): |%10s|\n", item)
	fmt.Printf("Left-aligned  (width 10): |%-10s|\n", item)
}

// -----------------------------------------------------------------------------
// 12. INTERNAL STRUCTURE, MEMORY, & `strings.Clone`
// -----------------------------------------------------------------------------
//   - How a string is stored in memory:
//     Internally, a string is represented as a 2-word header:
//     struct {
//     Data unsafe.Pointer // Pointer to underlying byte array
//     Len  int            // Length in bytes
//     }
//   - On 64-bit systems, unsafe.Sizeof("") is 16 bytes (8 bytes pointer + 8 bytes len).
//   - On 32-bit systems, unsafe.Sizeof("") is 8 bytes (4 bytes pointer + 4 bytes len).
//
// - SUBSTRING MEMORY SHARING & MEMORY LEAK PITFALL:
//   - When you slice a string `sub := largeStr[0:10]`, `sub` points to the SAME
//     underlying byte array as `largeStr`.
//   - If `largeStr` was a 50MB file and you only keep `sub` (10 bytes), the entire
//     50MB byte array CANNOT be garbage collected as long as `sub` is in memory!
//
// - SOLUTION: `strings.Clone(s)` (Introduced in Go 1.18)
//   - Allocates a new, dedicated copy of the sliced string bytes.
//   - Allows the large original backing array to be garbage collected.
//
// -----------------------------------------------------------------------------
func demonstrateInternalStructureAndMemory() {
	fmt.Println("\n--- 12. INTERNAL STRUCTURE & STRINGS.CLONE ---")

	sample := "Golang Internals"

	// Size of string header
	fmt.Printf("Size of string header: %d bytes (pointer + length)\n", unsafe.Sizeof(sample))

	// Substring memory sharing demonstration
	sub := sample[0:6] // "Golang"
	fmt.Printf("Original string: %q, Substring: %q\n", sample, sub)

	// Memory leak scenario explanation:
	// If sample was a huge string (e.g. read from a large file):
	hugeString := strings.Repeat("A", 1_000_000) // 1 Megabyte string
	smallSlice := hugeString[0:10]               // Retains the whole 1MB in memory!

	// strings.Clone creates a clean, independent copy without referencing hugeString's backing array:
	safeCopy := strings.Clone(smallSlice)
	fmt.Printf("Cloned substring: %q (safe from retaining 1MB memory)\n", safeCopy)

	fmt.Println("\n================================================================================")
	fmt.Println("                         END OF STRING CONCEPTS GUIDE                           ")
	fmt.Println("================================================================================")
}
