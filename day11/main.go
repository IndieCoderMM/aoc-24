package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	data := "data/day-11.txt"

	ans := Solve(data)

	fmt.Printf("Stones: %d\n", ans)
	end := time.Now()
	fmt.Printf("Duration: %v\n", end.Sub(start))
}

// Too low -> 791625185
