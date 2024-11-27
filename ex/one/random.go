package main

import (
	"fmt"
	"math/rand"
)

// generate random number from 0 to 10
func random() {
    random := rand.Intn(10); 
    fmt.Println(random);
}
