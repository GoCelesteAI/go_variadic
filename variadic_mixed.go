package main

import "fmt"

// Mixed parameters: regular + variadic
// Variadic must be the last parameter
func greetAll(greeting string, names ...string) {
  for _, name := range names {
    fmt.Println(greeting, name)
  }
}

// Another example with multiple regular params
func formatMessage(prefix string, separator string, items ...string) string {
  result := prefix
  for i, item := range items {
    if i > 0 {
      result += separator
    }
    result += item
  }
  return result
}

func main() {
  fmt.Println("=== Greeting Everyone ===")
  greetAll("Hello", "Alice", "Bob", "Carol")

  fmt.Println("\n=== Greeting with Hi ===")
  greetAll("Hi", "Dave", "Eve")

  fmt.Println("\n=== Format Message ===")
  msg := formatMessage("Items: ", ", ", "apple", "banana", "cherry")
  fmt.Println(msg)
}
