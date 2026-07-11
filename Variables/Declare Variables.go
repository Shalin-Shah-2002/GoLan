// ============================================================
//  VARIABLE DECLARATION IN GO
//  Go is statically typed — variable types are checked at
//  compile time. There are multiple ways to declare variables.
// ============================================================

package variables

// ============================================================
//  GO DATA TYPES (Overview)
// ============================================================
//  ┌──────────┬──────────────────────────────────────────────┐
//  │ Type     │ Description                                  │
//  ├──────────┼──────────────────────────────────────────────┤
//  │ int      │ Whole numbers (positive/zero/negative)       │
//  │ float32  │ Decimal numbers (single precision)           │
//  │ float64  │ Decimal numbers (double precision)           │
//  │ string   │ Text — always wrapped in double quotes ""    │
//  │ bool     │ Logical — only true or false                 │
//  └──────────┴──────────────────────────────────────────────┘
// ============================================================

// ============================================================
//  METHOD 1: var keyword with explicit type
//  Syntax:   var variablename type = value
// ============================================================
// The 'var' keyword is followed by the variable name, then
// the type, then = and the value. You can also omit the value
// (it defaults to the zero-value: 0 for int, "" for string,
// false for bool, 0.0 for float).
//
// Usage: when you want to be explicit about the type, or when
// you're declaring a package-level variable (:= is not allowed
// outside functions).
// ============================================================
var shalin int = 25   // Explicit: var name type = value

// ============================================================
//  METHOD 2: Short declaration with :=
//  Syntax:   variablename := value
// ============================================================
// The := operator declares AND assigns in one step. Go infers
// the type from the value on the right side.
//
// Rules:
//   - CANNOT be used outside of functions (package-level only
//     accepts var or const)
//   - At least ONE variable on the left must be new (no pure
//     reassignment with :=)
//   - Preferred style inside functions for brevity
//
// Example (inside a function):
//   name := "Shalin"       // inferred as string
//   age  := 25             // inferred as int
//   pi   := 3.14           // inferred as float64
//   ok   := true           // inferred as bool
// ============================================================
// (see accompanying files for function-scoped examples)

// ============================================================
//  COMPARISON: var vs :=
// ============================================================
//  ┌──────────────┬──────────────────────┬──────────────────────┐
//  │              │ var                  │ :=                   │
//  ├──────────────┼──────────────────────┼──────────────────────┤
//  │ Scope        │ package or function  │ function only        │
//  │ Type         │ explicit or inferred │ always inferred      │
//  │ Zero value   │ supported            │ not supported        │
//  │ Reassign     │ no (new variable)    │ no (new variable)    │
//  └──────────────┴──────────────────────┴──────────────────────┘
// ============================================================