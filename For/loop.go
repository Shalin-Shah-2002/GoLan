package main

import ("fmt")

func main() {
	// ============================================================
	//  TYPE 1: STANDARD THREE-PART FOR LOOP
	//  Syntax:  for init; condition; post { ... }
	// ============================================================
	// This is the classic C-style for loop. It has three parts:
	//   - init (i := 0)     -> runs once before the loop starts
	//   - condition (i < 5) -> checked before each iteration;
	//                          loop continues while this is true
	//   - post (i++)        -> runs after each iteration body
	//
	// Execution flow:
	//   1. init (i = 0)
	//   2. condition -> true?  ->  enter body
	//   3. print, then post (i++)
	//   4. condition -> true?  ->  enter body
	//   ... repeat until condition is false
	// ============================================================
	for i := 0; i < 5; i++ {
		fmt.Println("The value of i is: ", i)
	}
	// Output:
	//   The value of i is:  0
	//   The value of i is:  1
	//   The value of i is:  2
	//   The value of i is:  3
	//   The value of i is:  4

	// ============================================================
	//  TYPE 2: FOR LOOP WITH break
	//  break  -> immediately exits the innermost loop
	// ============================================================
	// The 'break' keyword terminates the loop entirely and
	// execution resumes at the first statement after the loop.
	// Here, when j == 3, the loop stops — so only 0, 1, 2 print.
	// ============================================================
	for j := 0; j < 5; j++ {
		if j == 3 {
			break   // <-- exits the for loop entirely
		}
		fmt.Println("The value of j is: ", j)
	}
	// Output:
	//   The value of j is:  0
	//   The value of j is:  1
	//   The value of j is:  2

	// ============================================================
	//  TYPE 3: FOR LOOP WITH continue
	//  continue -> skips rest of current iteration, goes to post
	// ============================================================
	// The 'continue' keyword skips the remaining body of the
	// current iteration and jumps to the post statement (i++),
	// then rechecks the condition. Here k == 2 is skipped entirely.
	// ============================================================
	for k := 0; k < 5; k++ {
		if k == 2 {
			continue   // <-- skips the print below for k=2
		}
		fmt.Println("The value of k is: ", k)
	}
	// Output:
	//   The value of k is:  0
	//   The value of k is:  1
	//   The value of k is:  3
	//   The value of k is:  4

	// ============================================================
	//  TYPE 4: CONDITION-ONLY LOOP (while-style)
	//  Syntax:  for condition { ... }
	// ============================================================
	// Go has no 'while' keyword. Instead, you drop the init and
	// post parts, keeping only the condition. This behaves exactly
	// like a while loop in other languages.
	//
	// Flow:
	//   1. check condition (m < 5) -> if false, exit
	//   2. execute body
	//   3. go back to step 1
	// The variable m is declared and incremented manually inside
	// the body, giving full control over the loop variable.
	// ============================================================
	m := 0
	for m < 5 {
		fmt.Println("The value of m is: ", m)
		m++   // manual increment — without this, the loop runs forever
	}
	// Output:
	//   The value of m is:  0
	//   The value of m is:  1
	//   The value of m is:  2
	//   The value of m is:  3
	//   The value of m is:  4

	// ============================================================
	//  TYPE 5: INFINITE LOOP WITH break
	//  Syntax:  for { ... }
	// ============================================================
	// An empty 'for' clause (no init, no condition, no post) runs
	// forever. It must be terminated with 'break' (or 'return' or
	// 'panic'). This pattern is useful for:
	//   - server event loops (accepting connections)
	//   - worker queues (polling for jobs)
	//   - interactive menus (waiting for "quit" input)
	//
	// Here we manually break when n reaches 5 to avoid hanging.
	// ============================================================
	n := 0
	for {
		if n == 5 {
			break   // <-- must have some exit condition!
		}
		fmt.Println("The value of n is: ", n)
		n++
	}
	// Output:
	//   The value of n is:  0
	//   The value of n is:  1
	//   The value of n is:  2
	//   The value of n is:  3
	//   The value of n is:  4

	// ============================================================
	//  TYPE 6: RANGE LOOP
	//  Syntax:  for index, value := range collection { ... }
	// ============================================================
	// 'range' iterates over slices, arrays, maps, strings, or
	// channels. It returns two values per iteration:
	//   - index (position in the collection, 0-based)
	//   - value (copy of the element at that index)
	//
	// Key behaviors:
	//   - The value is a COPY — modifying it does NOT change the
	//     original collection
	//   - Use _ (blank identifier) to discard either return value
	//   - For maps: range returns (key, value) pairs in random order
	//   - For strings: range iterates over runes, not bytes
	// ============================================================
	nums := []string{"a", "b", "c"}
	for idx, val := range nums {
		fmt.Printf("Index: %d, Value: %s\n", idx, val)
	}
	// Output:
	//   Index: 0, Value: a
	//   Index: 1, Value: b
	//   Index: 2, Value: c

	// ============================================================
	//  SUMMARY: For Loop Variants in Go
	// ============================================================
	//  ┌────────────────────┬──────────────────────────────────┐
	//  │ Form               │ Use case                         │
	//  ├────────────────────┼──────────────────────────────────┤
	//  │ for init; c; post  │ Counted iteration (classic for)  │
	//  │ for condition      │ While-style loop                 │
	//  │ for { }            │ Infinite loop (+ break to exit)  │
	//  │ for _, v := range  │ Iterate over collections         │
	//  └────────────────────┴──────────────────────────────────┘
	// ============================================================
}