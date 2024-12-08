package main

import "fmt"

func main() {
	// filePath := "data/test.txt"
	filePath := "data/day-7.txt"
	ans := Solve(filePath)
	fmt.Printf("Total calibration %d\n", ans)
}

// Part1
// Too low -> 49083398265
// Too low -> 49103603864
// Too low -> 52146785024
// Wrong   -> 303878202962
// Wrong   -> 303850277708
// Wrong   -> 303848560401
// Correct -> 303876485655
