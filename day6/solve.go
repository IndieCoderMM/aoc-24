package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve(filePath string) int {
	grid, pos, dir := readFile(filePath)

	ans := CalculatePath(grid, pos, dir)

	return ans
}

type Dir = struct {
	x    int
	y    int
	curr string
	next string
}

type Pos = struct {
	x int
	y int
}

type Path = struct {
	pos Pos
	dir Dir
}

// loops happens when curr postion has been visited + same direction
// store Vector -> visited
// 1. run dfs once to find the normal path
// 2. for each position in normal path, scan for potential looping

func CalculatePath(grid [][]string, startPos Pos, startDir string) int {
	ans := 0
	uniquePos := initVisited(grid)
	path := []Path{}

	directions := map[string]Dir{
		">": {1, 0, ">", "v"},
		"v": {0, 1, "v", "<"},
		"<": {-1, 0, "<", "^"},
		"^": {0, -1, "^", ">"},
	}

	dir, ok := directions[startDir]
	if !ok {
		err := fmt.Sprintf("Invalid direction: %s\n", startDir)
		panic(err)
	}

	var dfs func(x int, y int, dir Dir)
	dfs = func(x int, y int, dir Dir) {
		if x <= 0 || x >= len(grid[y])-1 || y <= 0 || y >= len(grid)-1 {
			// outside
			uniquePos[y][x] = true
			path = append(path, Path{Pos{x, y}, dir})
			return
		}

		nextX, nextY := x+dir.x, y+dir.y
		next := grid[nextY][nextX]

		if !uniquePos[y][x] {
			path = append(path, Path{Pos{x, y}, dir})
			uniquePos[y][x] = true
		}

		if next == "#" {
			// found obstacle
			nextDir := directions[dir.next]

			dfs(x+nextDir.x, y+nextDir.y, nextDir)
		} else {
			// continue in same dir
			dfs(x+dir.x, y+dir.y, dir)
		}
	}

	dfs(startPos.x, startPos.y, dir)

	totalUniques := 0
	for i := range uniquePos {
		for j := range uniquePos[i] {
			if uniquePos[i][j] {
				totalUniques += 1
			}
		}
	}
	// Part1
	fmt.Printf("Unique locations: %d\n", totalUniques)

	walk := func(x int, y int, dir Dir) Path {
		if x < 0 || x > len(grid[y])-1 || y < 0 || y > len(grid)-1 {
			// outside
			return Path{Pos{-1, -1}, dir}
		}

		if grid[y][x] == "#" {
			x -= dir.x
			y -= dir.y
			dir = directions[dir.next]
		}

		nextX, nextY := x+dir.x, y+dir.y
		if nextX < 0 || nextX > len(grid[y])-1 || nextY < 0 || nextY > len(grid)-1 {
			return Path{Pos{-1, -1}, dir}
		}
		next := grid[nextY][nextX]

		if next == "#" {
			// found obstacle
			nextDir := directions[dir.next]

			return Path{Pos{x + nextDir.x, y + nextDir.y}, nextDir}
		}

		return Path{Pos{nextX, nextY}, dir}
	}

	placements := make(map[Pos]bool)

	for _, curr := range path {
		if _, ok := placements[curr.pos]; ok || curr.pos == startPos {
			continue
		}
		// Imagine obstacle
		temp := grid[curr.pos.y][curr.pos.x]
		grid[curr.pos.y][curr.pos.x] = "#"

		p1 := Path{startPos, dir}
		p2 := Path{startPos, dir}

		for {
			if p1.pos.x == -1 || p1.pos.y == -1 || p2.pos.x == -1 || p2.pos.y == -1 {
				// out of bounds
				break
			}

			p1 = walk(p1.pos.x, p1.pos.y, p1.dir)
			p1 = walk(p1.pos.x, p1.pos.y, p1.dir)

			p2 = walk(p2.pos.x, p2.pos.y, p2.dir)

			if p1 == p2 {
				// Is a loop
				ans += 1
				placements[curr.pos] = true
				break
			}
		}

		grid[curr.pos.y][curr.pos.x] = temp
	}

	// for p := range placements {
	// 	showMap(grid, p.x, p.y)
	// }

	return ans
}

func showMap(grid [][]string, x int, y int) {
	for i, row := range grid {
		for j, c := range row {
			if i == y && j == x {
				if c != "." {
					fmt.Printf("%s|", "X")
				} else {
					fmt.Printf("%s|", "o")
				}
			} else {
				fmt.Printf("%s|", c)
			}
		}
		fmt.Println()
	}
	fmt.Println("-------------------------------")
}

func getDistBetweenLastTwo(obstacles []Pos) int {
	if len(obstacles) < 2 {
		return 0
	}

	dist := 0
	curr := obstacles[len(obstacles)-1]
	prev := obstacles[len(obstacles)-2]
	if curr.x == prev.x {
		dist = curr.y - prev.y
	} else if curr.y == prev.y {
		dist = curr.x - prev.x
	}

	if dist < 0 {
		dist *= -1
	}

	return dist
}

func isWalkable(grid [][]string, x int, y int, dist int, dir Dir) bool {
	for i := 1; i <= dist; i++ {
		nextX, nextY := x+(i*dir.x), y+(i*dir.y)

		if !isEmpty(grid, nextX, nextY) {
			return false
		}
	}

	return true
}

func isEmpty(grid [][]string, x int, y int) bool {
	return isInBounds(x, y, len(grid[y]), len(grid)) && grid[y][x] != "#"
}

func isInBounds(x int, y int, maxX int, maxY int) bool {
	return x >= 0 && y >= 0 && x <= maxX-1 && y <= maxY-1
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

func readFile(filePath string) ([][]string, Pos, string) {
	lines := [][]string{}
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pos := Pos{-1, -1}
	dir := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Split(text, "")

		// Scan current guard position
		for i := range line {
			if line[i] != "#" && line[i] != "." {
				pos.x = i
				pos.y = len(lines)
				dir = line[i]
			}
		}

		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines, pos, dir
}

// INFO: Attempt to find loop by drawing rectangles
// Doesnt work with larger loop
//
// curr := Pos{x, y}
// obstacles = append(obstacles, curr)
//
// dist := getDistBetweenLastTwo(obstacles)
//
// if dist > 0 && prevDist > 0 {
// 	fmt.Printf("Two dists: %d - %d\n", prevDist, dist)
//
// 	// Skip if obstacle exists
// 	if isWalkable(grid, x, y, prevDist, nextDir) {
// 		// Rotate 90 and walk for dist[1]
//
// 		rotate := directions[nextDir.dir]
// 		if isWalkable(grid, x, y, dist, rotate) {
// 			showMap(grid, x+((prevDist+1)*nextDir.x), y+((prevDist+1)*nextDir.y))
// 			fmt.Printf("Placement successful--\n")
// 			ans += 1
// 		}
// 	}
// }
// prevDist = dist
