package main

import "fmt"

// Problem 02 - Assign real numbers to appropriate types
// Create a program, which assigns a set of real number values to the appropriate type in Go.
// Print the number on the terminal. The numbers should be stored in the smallest type possible.

// The numbers:
// 14.435234
// 14.435234234234235

// https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go
// https://en.wikipedia.org/wiki/Single-precision_floating-point_format
// https://en.wikipedia.org/wiki/Double-precision_floating-point_format

func typeReal() {
	var a float32 = 14.435234
	var b float64 = 14.435234234234235

	fmt.Println(a)
	fmt.Println(b)
}
