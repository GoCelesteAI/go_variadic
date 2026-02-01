# Go Variadic Functions

Code examples from **Go Tutorial for Beginners #12 - Variadic Functions**

## Contents

- `variadic_basics.go` - Basic variadic syntax, sum function
- `variadic_mixed.go` - Mixed parameters (regular + variadic)
- `variadic_spread.go` - Spread operator for slices
- `variadic_real.go` - Real-world patterns (logging, min/max)

## Running the Examples

```bash
go run variadic_basics.go
go run variadic_mixed.go
go run variadic_spread.go
go run variadic_real.go
```

## Key Concepts

### Variadic Function
```go
// The ... makes nums variadic - it accepts any number of ints
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// Call with any number of arguments
sum()           // 0
sum(5)          // 5
sum(1, 2, 3)    // 6
```

### Spread Operator
```go
// Pass a slice to variadic function
numbers := []int{10, 20, 30}
sum(numbers...)  // Spreads slice into individual args

// Works with append too!
slice1 := []int{1, 2, 3}
slice2 := []int{4, 5, 6}
combined := append(slice1, slice2...)
```

### Mixed Parameters
```go
// Regular params come first, variadic must be last
func greetAll(greeting string, names ...string) {
    for _, name := range names {
        fmt.Println(greeting, name)
    }
}

greetAll("Hello", "Alice", "Bob", "Carol")
```

### Real-World Patterns
```go
// Custom logger (like fmt.Println)
func log(level string, messages ...interface{}) {
    fmt.Printf("[%s] ", level)
    fmt.Println(messages...)
}

// Min/max with at least one required value
func min(first int, rest ...int) int {
    minimum := first
    for _, v := range rest {
        if v < minimum {
            minimum = v
        }
    }
    return minimum
}
```

## Watch the Tutorial

[Go Tutorial for Beginners #12 - Variadic Functions](https://youtube.com/@CelesteAI)

## License

MIT
