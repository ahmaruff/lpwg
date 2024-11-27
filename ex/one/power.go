package main

import (
	"fmt"
	"math"
)

func power(power int) {
    res := math.Pow(2,float64(power));
    fmt.Println(res);
}
