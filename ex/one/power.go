package main

import (
	"fmt"
	"math"
)

// Create a program which calculate 2 to the power of 11
func power(power int) {
    res := math.Pow(2,float64(power));
    fmt.Println(res);
}
