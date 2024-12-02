package main

import "fmt"

func main() {
	filePath := "data/day-2.txt"
	ans := Solve(filePath)
	fmt.Printf("Safe reports: %d\n", ans)
}
