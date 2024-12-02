package main

import "fmt"

func main() {
	filePath := "data/day-1.txt"
	ans := Solve(filePath)

	fmt.Printf("Similarity score: %d\n", ans)
}
