package main

import "fmt"

func main() {
	filePath := "data/day-6.txt"

	ans := Solve(filePath)

	fmt.Printf("Distinct positions: %d\n", ans)
}

// part1 -> 5564
// too low -> 1934
// 1976
