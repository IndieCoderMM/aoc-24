package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve(filePath string) int {
	grid, pos := readFile(filePath)

	ans := CalculatePath(grid, pos)

	return ans
}

type Pos = struct {
	x   int
	y   int
	dir string
}

func CalculatePath(grid [][]string, pos Pos) int {
	ans := 0
	visited := initVisited(grid)
	directions := map[string]Pos{
		">": {1, 0, "v"},
		"v": {0, 1, "<"},
		"<": {-1, 0, "^"},
		"^": {0, -1, ">"},
	}

	dir, ok := directions[pos.dir]
	if !ok {
		err := fmt.Sprintf("Invalid direction: %s\n", pos.dir)
		panic(err)
	}

	// for _, row := range grid {
	// 	for _, i := range row {
	// 		fmt.Printf("| %s |", i)
	// 	}
	// 	fmt.Println()
	// }

	var dfs func(x int, y int, dir Pos)
	dfs = func(x int, y int, dir Pos) {
		if x <= 0 || x >= len(grid[y])-1 || y <= 0 || y >= len(grid)-1 {
			// outside
			return
		}

		curr := grid[y][x]
		next := grid[y+dir.y][x+dir.x]

		if !visited[y][x] && curr != "#" {
			visited[y][x] = true
			ans += 1
		}

		if next == "#" {
			// found obstacle
			rotate90, ok := directions[dir.dir]
			if !ok {
				panic("Invalid direction")
			}
			dfs(x+rotate90.x, y+rotate90.y, rotate90)
		} else {
			// continue in same dir
			dfs(x+dir.x, y+dir.y, dir)
		}

	}

	dfs(pos.x, pos.y, dir)

	return ans + 1
}

func initVisited(grid [][]string) [][]bool {
	visited := [][]bool{}
	for i := range grid {
		temp := []bool{}
		for j := 0; j < len(grid[i]); j++ {
			temp = append(temp, false)
		}
		visited = append(visited, temp)
	}

	return visited
}

func readFile(filePath string) ([][]string, Pos) {
	lines := [][]string{}
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pos := Pos{-1, -1, ""}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Split(text, "")

		// Scan current guard position
		for i := range line {
			if line[i] != "#" && line[i] != "." {
				pos.x = i
				pos.y = len(lines) + 1
				pos.dir = line[i]
			}
		}

		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines, pos
}
