package main

import "fmt"

func main() {
	filePath := "data/day-8.txt"
	ans := Solve(filePath)

	fmt.Printf("Unique antidode locations %d\n", ans)
}
