# 🐹 Go Learning Docs

I'm learning Go and making my own docs because the official ones are **so boring**.

This repo is my personal reference — every concept explained in plain language with runnable examples and detailed comments.

---

## 📂 Project Structure

```
GoLan/
├── README.md
├── Get Started/
│   └── main.go              # Hello World + package/import basics
├── Variables/
│   ├── Declare Variables.go         # var vs :=, data types
│   ├── Go Multiple Variable Declaration.go  # Parallel declaration
│   └── Constants.go                 # const, iota, immutability
├── Output/
│   └── output funtion.go    # fmt.Println, functions
├── If-Else/
│   └── Else.go              # if, else, else if, logical operators
├── For/
│   └── loop.go              # All 4 for-loop forms + break/continue
└── Array /
    └── arr.go               # 1D arrays, 2D arrays, sparse init
```

---

## 📘 Topics Covered

### 1. Getting Started — [`Get Started/main.go`](Get%20Started/main.go)

**Concepts:** package main, import, func main(), fmt.Println

```go
package main
import "fmt"
func main() {
    fmt.Println("Hello World!")
}
```

| Concept | Explanation |
|---------|-------------|
| `package main` | Tells Go to build an executable (not a library) |
| `import "fmt"` | Imports the format package for printing |
| `func main()` | Entry point — execution starts here |
| `fmt.Println()` | Prints text + newline to console |

---

### 2. Variables — [`Variables/Declare Variables.go`](Variables/Declare%20Variables.go)

**Concepts:** var keyword, := short declaration, type inference, basic data types

```go
var shalin int = 25    // explicit type
name := "Shalin"       // inferred as string
```

**Data types covered:**

| Type | Example | Description |
|------|---------|-------------|
| `int` | `42` | Whole numbers |
| `float64` | `3.14` | Decimal numbers |
| `string` | `"hello"` | Text (double quotes only) |
| `bool` | `true` | true or false |

**var vs := comparison:**

| | `var` | `:=` |
|---|---|---|
| Scope | package or function | function only |
| Type | explicit or inferred | always inferred |
| Zero values | supported | not supported |

---

### 3. Multiple Variables — [`Variables/Go Multiple Variable Declaration.go`](Variables/Go%20Multiple%20Variable%20Declaration.go)

**Concepts:** parallel declaration, same-type grouping

```go
var a, b, c, d int = 1, 3, 5, 7
```

Values are assigned positionally. All variables must share the same type.

---

### 4. Constants — [`Variables/Constants.go`](Variables/Constants.go)

**Concepts:** const keyword, immutability, compile-time evaluation

```go
const Pi float64 = 3.14
// Pi = 2  // COMPILE ERROR — constants cannot change
```

**Key rules:**
- Value must be known at compile time
- Reassignment causes a compile error
- Only primitive types: numeric, string, bool

**Bonus patterns (documented in comments):**
- Parallel const blocks
- `iota` auto-incrementing generator

---

### 5. Output — [`Output/output funtion.go`](Output/output%20funtion.go)

**Concepts:** fmt.Println, multi-argument printing, user-defined functions

```go
a, b := 10, 20
fmt.Println("The sum is: ", a+b)
```

**Function syntax in Go:**
```go
func add(a int, b int) int {
    return a + b
}
```

Type comes **after** the variable name (unlike C/Java). If consecutive params share a type, you can write `a, b int`.

---

### 6. If-Else — [`If-Else/Else.go`](If-Else/Else.go)

**Concepts:** if, else, else if, `||` operator, short-statement if

```go
if 7%2 == 0 {
    fmt.Println("even")
} else {
    fmt.Println("odd")
}

if num := 9; num < 10 {
    fmt.Println(num, "has 1 digit")
}
```

**All forms:**

| Form | Use case |
|------|----------|
| `if c { }` | Single case |
| `if c { } else { }` | Two branches |
| `if c { } else if { } else { }` | Multiple branches |
| `if s; c { }` | Scoped variable + condition |

**Note:** Go has NO ternary operator (`? :`). Use if-else everywhere.

---

### 7. For Loops — [`For/loop.go`](For/loop.go)

**Concepts:** All 4 loop forms, break, continue, range

**Loop variants:**

| Form | Example | Use case |
|------|---------|----------|
| Three-part | `for i := 0; i < 5; i++` | Counted iteration |
| Condition-only | `for m < 5` | While-style loop |
| Infinite | `for { }` | Server loops / until break |
| Range | `for i, v := range slice` | Iterate collections |

```go
// Standard
for i := 0; i < 5; i++ { }

// While-style
m := 0
for m < 5 { m++ }

// Infinite + break
for { if done { break } }

// Range over slice
for idx, val := range []string{"a", "b"} { }
```

**Flow control:**
- `break` — exits the loop entirely
- `continue` — skips to next iteration

---

### 8. Arrays — [`Array /arr.go`](Array%20/arr.go)

**Concepts:** 1D arrays, 2D arrays, literal init, sparse init, len(), zero values

```go
var a [5]int              // [0 0 0 0 0]
b := [5]int{1, 2, 3, 4, 5}
c := [...]int{100, 3: 400, 500}   // [100 0 0 400 500]
```

**Key properties:**

| Property | Behavior |
|----------|----------|
| Fixed size | Set at compile time — never changes |
| Value type | Assignment COPIES all elements |
| Zero values | All elements auto-initialized to 0 |
| `len()` | Returns compile-time length |
| Comparison | `arr1 == arr2` allowed (same type only) |

**2D arrays:**
```go
var twoD [2][3]int
twoD = [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

---

## 🚀 How to Run

```bash
# From the GoLan directory:
go run "Get Started/main.go"
go run For/loop.go
go run "If-Else/Else.go"
go run "Array /arr.go"
go run "Output/output funtion.go"
go run "Variables/Go Multiple Variable Declaration.go"
```

> **Note:** Files in the `Variables/` directory use `package variables` (not `package main`), so they're library packages meant for reading/reference, not standalone execution.

---

## 📝 Why This Exists

The official Go docs are technically correct but **painfully dry**. This repo is my attempt to document what I learn in a way that's:

- **Readable** — plain language, no jargon for the sake of it
- **Runnable** — every example is a real .go file you can execute
- **Annotated** — detailed inline comments explaining every line
- **Visual** — ASCII diagrams, tables, and expected output included

If you're also learning Go, feel free to use this as a reference. Contributions and suggestions are welcome!
