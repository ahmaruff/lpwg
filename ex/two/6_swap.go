package main

import "fmt"

// Problem 06 - Swap values of two numbers
// Declare two variables a and b with values 14 and 32. Swap their values such that a contains 32 and b contains 14.
//
// Example:
// Before swapping:
// a = 14
// b = 32
//
// After swapping
// a = 32
// b = 14

func swap(a int, b int) (int, int) {
	a, b = b, a
	return a, b
}

func printSwap() {
	fmt.Println(swap(14, 32))
}
