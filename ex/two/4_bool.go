package main

import "fmt"

// Problem 04 - Are you a hooman?
// Identify several species around you. Include yourself if you so desire.
// For each of them, name them if they aren’t named already & identify if they are human or not.
//
//
// Examples:
// Is Gary a hooman?
// false
// Gary is a cat!
//
// Is Steve a hooman?
// true

func typeBoolean() {
	q1 := "Is Gary a hooman?"
	q2 := "Is Steve a hooman?"
	var yep bool = true
	var nah bool = false

	fmt.Println(q1)
	fmt.Println(nah)
	fmt.Println("Gary is a cat!")
	fmt.Println(q2)
	fmt.Println(yep)
}
