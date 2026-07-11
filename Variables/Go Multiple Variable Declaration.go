// ============================================================
//  MULTIPLE VARIABLE DECLARATION IN GO
//  Declare several variables of the SAME type in one line
// ============================================================

package variables

import ("fmt")

func main() {
	// ============================================================
	//  Parallel Variable Declaration
	//  Syntax: var name1, name2, ... type = val1, val2, ...
	// ============================================================
	// This declares four int variables (a, b, c, d) in one line.
	// The values are assigned positionally:
	//   a = 1, b = 3, c = 5, d = 7
	//
	// This is equivalent to writing:
	//   var a int = 1
	//   var b int = 3
	//   var c int = 5
	//   var d int = 7
	//
	// All variables MUST share the same type.
	// The number of names on the left must equal the number of
	// values on the right.
	// ============================================================
	var a, b, c, d int = 1, 3, 5, 7

	// ============================================================
	//  Print each variable on its own line
	// ============================================================
	// Each call to fmt.Println prints one value followed by a
	// newline. There is no formatting — just the raw value.
	// ============================================================
	fmt.Println(a)   // Output: 1
	fmt.Println(b)   // Output: 3
	fmt.Println(c)   // Output: 5
	fmt.Println(d)   // Output: 7

	// ============================================================
	//  ALTERNATIVE PATTERNS
	// ============================================================
	// // Mixed types — use var block:
	// var (
	//     name   string = "Go"
	//     year   int    = 2024
	//     active bool   = true
	// )
	//
	// // Short declaration with := (same position rule):
	// x, y, z := "a", 42, true
	// ============================================================
}