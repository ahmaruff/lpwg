package main

import "fmt"

// Problem 05 - Type conversions
// Convert the following values to the type specified & print them out:
// float32 with value 3.14 -> int
// float64 with value 12.3456789 -> float32
// int16 with value 1234 -> int8
// int16 with value 1234 -> uint8

func typeConv() {
	var a float32 = 3.14
	var conv_a int = int(a)
	var b float64 = 12.3456789
	var conv_b = float32(b)

	var c int16 = 1234
	var conv_c int8 = int8(c)
	var conv_d uint8 = uint8(c)

	fmt.Println(conv_a)
	fmt.Println(conv_b)
	fmt.Println(conv_c)
	fmt.Println(conv_d)
}
