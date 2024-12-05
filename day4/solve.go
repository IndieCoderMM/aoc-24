package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve() int {
	lines, err := readFile("data/day-4.txt")
	if err != nil {
		panic(err)
	}

	return FindX_MAS(lines)
}

// Find X-MAS
func FindX_MAS(lines [][]string) int {
	ans := 0

	isX_MAS := func(row int, col int) bool {
		// Skip first and last rows/cols
		if row <= 0 || col <= 0 || row >= len(lines)-1 || col >= len(lines[row])-1 {
			return false
		}

		if lines[row][col] != "A" {
			return false
		}

		topLeft := lines[row-1][col-1]
		topRight := lines[row-1][col+1]
		botLeft := lines[row+1][col-1]
		botRight := lines[row+1][col+1]

		// S . S
		// . A .
		// M . M

		// M . M
		// . A .
		// S . S

		// TL -> BR
		if (topLeft == "M" && botRight == "S") || (topLeft == "S" && botRight == "M") {
			// TR -> BL
			if (topRight == "M" && botLeft == "S") || (topRight == "S" && botLeft == "M") {
				// fmt.Printf("%s.%s\n", topLeft, topRight)
				// fmt.Printf(".%s.\n", lines[row][col])
				// fmt.Printf("%s.%s\n", botLeft, botRight)

				return true
			}
		}

		return false
	}

	for i := range lines {
		for j := range lines[i] {
			if isX_MAS(i, j) {
				ans += 1
			}
		}
	}

	return ans
}

// Part1: Finding XMAS
func FindXMAS(lines [][]string) int {
	ans := 0
	xmas := []string{"X", "M", "A", "S"}

	var directions = [8][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	var dfs func(row int, col int, current int, dir [2]int)
	dfs = func(row int, col int, current int, dir [2]int) {
		if row < 0 || col < 0 {
			return
		}

		if row >= len(lines) || col >= len(lines[row]) {
			return
		}

		// Found current match
		if lines[row][col] == xmas[current] {
			if current == len(xmas)-1 {
				// Complete word
				ans += 1
				return
			} else {
				// Keep searching in same direction
				dfs(row+dir[0], col+dir[1], current+1, dir)
			}
		}
	}

	for i := range lines {
		for j := range lines[i] {
			for _, dir := range directions {
				dfs(i, j, 0, dir)
			}
		}
	}

	return ans
}

func readFile(filePath string) ([][]string, error) {
	lines := [][]string{}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Split(text, "")
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	return lines, nil
}
