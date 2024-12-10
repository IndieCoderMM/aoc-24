package main

import "fmt"

func main() {
	filePath := "data/day-10.txt"

	ans := Solve(filePath)

	fmt.Printf("Trailhead score: %d\n", ans)
}
