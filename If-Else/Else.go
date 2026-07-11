// ============================================================
//  IF-ELSE DECISIONS IN GO
//  Go has NO ternary (? :) operator — all branching uses if/else
// ============================================================

package main

import ("fmt")

func main() {

    // ============================================================
    //  EXAMPLE 1: if / else
    //  Condition: 7 % 2 == 0  (is 7 even?)
    // ============================================================
    // The modulo operator (%) returns the remainder of division.
    // 7 % 2 = 1 (not 0), so the condition is false.
    // The 'else' branch runs instead of the 'if' branch.
    // ============================================================
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }
    // Output:
    //   7 is odd

    // ============================================================
    //  EXAMPLE 2: if only (no else)
    //  Condition: 8 % 4 == 0  (is 8 divisible by 4?)
    // ============================================================
    // When you only need to handle the true case, you can
    // omit the 'else' block entirely.
    // ============================================================
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }
    // Output:
    //   8 is divisible by 4

    // ============================================================
    //  EXAMPLE 3: Logical OR (||) operator
    //  Condition: 8 % 2 == 0  ||  7 % 2 == 0
    // ============================================================
    // The || operator returns true if EITHER side is true.
    // 8 % 2 == 0  -> true  (8 is even)
    // 7 % 2 == 0  -> false (7 is odd)
    // Since one side is true, the whole condition is true.
    // ============================================================
    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }
    // Output:
    //   either 8 or 7 are even

    // ============================================================
    //  EXAMPLE 4: if with short statement + else if chain
    //  Syntax: if statement; condition { ... }
    // ============================================================
    // Go allows a short statement (num := 9) before the condition.
    // The variable 'num' is scoped to the entire if/else chain.
    //
    // Flow:
    //   1. num := 9  (runs first)
    //   2. num < 0?  -> false (9 is not negative), skip to else if
    //   3. num < 10? -> true  (9 < 10), execute this block
    //   4. else block is skipped
    //
    // This pattern is cleaner than declaring num outside the if.
    // ============================================================
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
    // Output:
    //   9 has 1 digit

    // ============================================================
    //  SUMMARY: if/else Decision Table
    // ============================================================
    //  ┌──────────────────────┬──────────────────────────────────┐
    //  │ Form                 │ When to use                      │
    //  ├──────────────────────┼──────────────────────────────────┤
    //  │ if c { }             │ Single case, no alternative      │
    //  │ if c { } else { }    │ Two mutually exclusive branches   │
    //  │ if c { } else if { } │ Multiple cases (like switch)     │
    //  │ if s; c { }          │ Scoped variable + condition      │
    //  └──────────────────────┴──────────────────────────────────┘
    // ============================================================
}

