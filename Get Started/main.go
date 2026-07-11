// ============================================================
//  PACKAGE OVERVIEW
// ============================================================
// package main  -> special package that tells Go to compile
//                  this as an executable (not a reusable library)
// import "fmt"  -> fmt = format — Go's standard I/O package for
//                  printing, formatting, and reading data
// func main()   -> mandatory entry point; execution starts here
// ============================================================

package main

import ("fmt")

func main() {
	// ============================================================
	//  fmt.Println — Print Line
	// ============================================================
	// Prints the given text to the console, followed by a newline.
	// The string "Hello World!" is passed as an argument and
	// displayed in the terminal when the program runs.
	// ============================================================
	fmt.Println("Hello World!")
	// Output:
	//   Hello World!
}

