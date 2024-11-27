package main

import (
	"fmt"
	"time"
)

// print the current date
func currentTime() {
	t := time.Now();
    fmt.Println(t.Format(time.RFC850));
}
