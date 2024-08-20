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
}