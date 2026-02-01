package main

import "fmt"

func sum(nums ...int) int {
  total := 0
  for _, n := range nums {
    total += n
  }
  return total
}

func main() {
  // Passing a slice to variadic function
  numbers := []int{10, 20, 30, 40, 50}

  // This wont work:
  // fmt.Println(sum(numbers)) // Error: cannot use []int as int

  // Use spread operator (...) to expand slice
  fmt.Println("Sum of slice:", sum(numbers...))

  moreNumbers := []int{1, 2, 3}
  fmt.Println("Sum:", sum(moreNumbers...))

  fmt.Println("\n=== append is variadic ===")
  slice1 := []int{1, 2, 3}
  slice2 := []int{4, 5, 6}

  // Append individual elements
  slice1 = append(slice1, 7, 8, 9)
  fmt.Println("After append elements:", slice1)

  // Append another slice using spread
  combined := append(slice1, slice2...)
  fmt.Println("Combined slices:", combined)
}
