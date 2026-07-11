package main

import ("fmt")

func main() {
	// ============================================================
	//  SWITCH STATEMENT
	//  Syntax: switch expression { case value: ... }
	// ============================================================
	// A switch statement evaluates an expression and compares it
	// against multiple cases. The first matching case executes.
	// If no case matches, the optional 'default' case runs.
	// ============================================================
	num := 2

	switch num {
	case 1:
		fmt.Println("num is one")
	case 2:
		fmt.Println("num is two")
	case 3:
		fmt.Println("num is three")
	default:
		fmt.Println("num is something else")
	}
	// Output:
	//   num is two

	// ============================================================
	//  SWITCH WITH MULTIPLE CASES
	// ============================================================
	// You can group multiple values in a single case using commas.
	// The first matching case executes, and the rest are skipped.
	// ============================================================
	day := "Saturday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}
	// Output:
	//   It's the weekend!
	
}
