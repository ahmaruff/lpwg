package main

import (
	"fmt"
	"time"
)

func currentTime() {
	t := time.Now();
    fmt.Println(t.Format(time.RFC850));
}
