package main

import (
  "fmt"
  "strings"
)

// Custom logger with variadic
func log(level string, messages ...interface{}) {
  fmt.Printf("[%s] ", level)
  fmt.Println(messages...)
}

// Find minimum of any number of values
func min(first int, rest ...int) int {
  minimum := first
  for _, v := range rest {
    if v < minimum {
      minimum = v
    }
  }
  return minimum
}

// Find maximum of any number of values
func max(first int, rest ...int) int {
  maximum := first
  for _, v := range rest {
    if v > maximum {
      maximum = v
    }
  }
  return maximum
}

func main() {
  fmt.Println("=== fmt.Println is variadic ===")
  fmt.Println("one")
  fmt.Println("one", "two")
  fmt.Println("one", "two", "three", 4, 5.0, true)

  fmt.Println("\n=== Custom Logger ===")
  log("INFO", "Server started on port", 8080)
  log("WARN", "Memory usage high:", "85%")
  log("ERROR", "Connection failed:", "timeout after", 30, "seconds")

  fmt.Println("\n=== Min and Max ===")
  fmt.Println("min(5, 3, 8, 1, 9):", min(5, 3, 8, 1, 9))
  fmt.Println("max(5, 3, 8, 1, 9):", max(5, 3, 8, 1, 9))

  fmt.Println("\n=== String Joining ===")
  words := []string{"Go", "is", "awesome"}
  fmt.Println(strings.Join(words, " "))
}
