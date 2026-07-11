// ============================================================
//  CONSTANTS IN GO
//  Constants are fixed values that CANNOT be changed at runtime
// ============================================================
// Syntax: const CONSTNAME type = value
//
// Key rules:
//   - Use 'const' keyword (not 'var')
//   - Value must be known at compile time (no function calls)
//   - Reassignment is a COMPILE ERROR — constants are immutable
//   - Naming convention: PascalCase for exported, camelCase for
//     unexported (but ALL_CAPS is also common for constants)
//   - Types: numeric (int, float64), string, bool only
// ============================================================

package variables

// ============================================================
//  CONSTANT DECLARATION
// ============================================================
// Pi is an exported constant (capital P) — accessible from
// other packages. Its type is float64 and value is 3.14.
//
// Trying to reassign:  Pi = 2
// would produce:       cannot assign to Pi (neither addressable
//                      nor a map index expression)
// ============================================================
const Pi float64 = 3.14

// ============================================================
//  COMMON CONSTANT PATTERNS
// ============================================================
// // Multiple constants (parallel declaration):
// const (
//     StatusOK    = 200
//     StatusNotFound = 404
//     StatusError    = 500
// )
//
// // Iota — auto-incrementing constant generator:
// const (
//     Monday = iota  // 0
//     Tuesday        // 1
//     Wednesday      // 2
// )
// ============================================================