package main

import (
	"fmt"
)

// go iterator func
func squared(s []int) func(func(x, i int) bool) {
	// closure that accepts our yeild function
	return func(yield func(x int, i int) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			squaredVal := s[i] * s[i]

			// if yield is false we stop our iterating and stop passing values to our squared function
			// Yield function that is supposed to process each squared value and decide whether to continue iterating or not.
			// If yield returns true, the loop continuee
			// If yield returns false, the loop stops early.
			if !yield(i, squaredVal) {
				return
			}
		}
	}
}

func main() {
	fmt.Println(fmt.Sprintf("hi"))

	nums := []int{1, 2, 3, 4, 5}

	// We are squaring the values in the nums slice
	for i, x := range squared(nums) {
		fmt.Printf("this is i: %d and this is x: %d\n", i, x)
	}

	// Call squared with a yield function
	squared(nums)(func(i int, x int) bool {
		fmt.Printf("this is i: %d and this is x: %d\n", i, x)
		// return true to continue, return false to stop early
		return false // b/c I returned false the loop ran only once
	})

	fmt.Println("----------------")

	x := 42
	fmt.Println("x: ", x)   // 42
	fmt.Println("&x: ", &x) // 0xc0000160b0

	fmt.Println("----------------")

	y := &x                                                // Assign the memory address of x to y
	fmt.Printf("y is a pointer to an int: %v\t%T\n", y, y) // 0xc0000160b0 *int (pointer to an int)
	fmt.Println("y: ", y)                                  // 0xc0000160b0
	fmt.Println("*y: ", *y)                                // Value of 42. Dereference the memory address of y to get the value of x
	fmt.Println("*&x: ", *&x)                              // Value of 42

	fmt.Println("----------------")

	*y = 43                 // Dereference on the left means "write to the pointed-to location", not read.
	fmt.Println("x: ", x)   // 43
	fmt.Println("y: ", y)   // 0xc0000160b0
	fmt.Println("*y: ", *y) // 43
}
