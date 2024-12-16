package main

import "fmt"

func main() {
	filePath := "data/day-12.txt"
	ans := Solve(filePath)

	fmt.Printf("Total fence price: %d\n", ans)
}
