package main

import "fmt"

func main() {
	filePath := "data/day-5.txt"
	ans := Solve(filePath)

	fmt.Printf("Sum of middle pages: %d\n", ans)
}

// Too high -> 4941
// Too high -> 4938
// 4713
