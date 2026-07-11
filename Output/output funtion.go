// ============================================================
//  OUTPUT FUNCTIONS IN GO
//  Printing to console using the fmt package
// ============================================================

package main
 
import ("fmt")

func main() {
	// ============================================================
	//  Variable Declaration with :=
	// ============================================================
	// Short variable declaration (:=) infers the type from the
	// value. Here both a and b are inferred as int.
	// Equivalent to: var a int = 10; var b int = 20
	// ============================================================
	a, b := 10, 20

	// ============================================================
	//  fmt.Println — Print values with spaces between arguments
	// ============================================================
	// Println accepts multiple arguments separated by commas.
	// It inserts spaces between arguments and adds a newline at
	// the end. The expression a + b is evaluated first, then the
	// result (30) is printed.
	// ============================================================
	fmt.Println("The sum of a and b is: ", a + b)
	// Output:
	//   The sum of a and b is:  30

	fmt.Println("Hello World!")
	// Output:
	//   Hello World!
}

// ============================================================
//  USER-DEFINED FUNCTION: add
//  Parameters: a int, b int   (both integers)
//  Return:     int            (sum of a and b)
// ============================================================
// This function is defined but NOT called in main() above.
// To use it, you would call: result := add(10, 20)
//
// Key Go syntax rules:
//   - Type follows the variable name (unlike C/Java)
//   - If two consecutive params share the same type, you can
//     write: add(a, b int) int
// ============================================================
func add(a int, b int) int {
	return a + b
}
