package main

import "fmt"

// Variadic function - accepts any number of integers
func sum(nums ...int) int {
  total := 0
  for _, n := range nums {
    total += n
  }
  return total
}

// The variadic parameter is received as a slice
func printNumbers(nums ...int) {
  fmt.Printf("Type: %T\n", nums)
  fmt.Printf("Length: %d\n", len(nums))
  fmt.Println("Values:", nums)
}

func main() {
  // Call with different number of arguments
  fmt.Println("sum():", sum())
  fmt.Println("sum(5):", sum(5))
  fmt.Println("sum(1, 2):", sum(1, 2))
  fmt.Println("sum(1, 2, 3, 4, 5):", sum(1, 2, 3, 4, 5))

  fmt.Println("\nVariadic parameter is a slice:")
  printNumbers(10, 20, 30)
}
