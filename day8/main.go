package main

import "fmt"

func main() {
	filePath := "data/day-8.txt"
	ans, _, _ := Solve(filePath)

	fmt.Printf("Unique antidode locations %d\n", ans)

	// DrawMap(grid, highlights)
}
