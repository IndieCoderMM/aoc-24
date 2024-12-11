package main

import "fmt"

func main() {
	filePath := "data/day-9.txt"

	ans := Solve(filePath)

	fmt.Printf("Checksum: %d\ns", ans)
}

// Too low -> 6366824700663
// Part1 -> 6367087064415
// Part2 -> 6390781891880
