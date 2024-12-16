package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve(filePath string) int {
	grid := readFile(filePath)
	visited := initVisited(grid)
	ans := 0
	calculate(grid, visited, false, &ans)

	// Run(grid, visited, func() (int, bool) { return calculate(grid, visited, false, &ans) })

	return ans
}

type Cell = struct {
	row int
	col int
}

var directions = [4]Cell{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func calculate(grid [][]string, visited [][]bool, step bool, ans *int) (int, bool) {
	completed := false
	processed := false

	for r, row := range grid {
		for c, p := range row {
			if visited[r][c] {
				continue
			}

			stack := []Cell{{r, c}}
			visited[r][c] = true
			currArea := 1
			currSides := 0
			regions := make(map[Cell][4]bool)

			// Getting area
			for len(stack) > 0 {
				cell := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				regions[cell] = [4]bool{}
				for _, dir := range directions {
					newRow, newCol := cell.row+dir.row, cell.col+dir.col

					if !isSamePlant(grid, newRow, newCol, p) {
						// Part1: perimeter++ // increase the perimeter if touching the outside or different plant
						continue
					}

					if visited[newRow][newCol] {
						continue
					}

					visited[newRow][newCol] = true
					stack = append(stack, Cell{newRow, newCol})
					currArea++
				}
			}

			// For cell in each region set index to true if it touching outside
			// direction -> top, right, bottom, left
			for cell, neighbors := range regions {
				for i, dir := range directions {
					newRow, newCol := cell.row+dir.row, cell.col+dir.col
					if !isSamePlant(grid, newRow, newCol, p) {
						neighbors[i] = true
						regions[cell] = neighbors
					}
				}
			}

			// fmt.Printf("Current region %d, %v\n", len(regions), regions)
			// Count sides +1 if cell[dir] is true and above cell[dir] is false
			for cell, neighbors := range regions {
				// fmt.Printf("Current cell %v, [%v]\n", cell, neighbors)
				for i := range directions {
					// fmt.Printf("checking direction %s @i%d\n", dir.dir, i)
					if neighbors[i] && !regions[Cell{cell.row, cell.col - 1}][i] && !regions[Cell{cell.row - 1, cell.col}][i] {
						currSides++
					}
				}
			}

			// drawEdges(grid, allEdges)
			// fmt.Printf("[%s]: Area %d, Sides (%d)\n", p, currArea, currSides)

			// Update ans
			*ans += currArea * currSides

			if step {
				// Only increment once
				processed = true
				break
			}
		}
		if step && processed {
			// Only increment once
			break
		}
	}

	// Check if all cells are visited
	completed = true
	for r := range visited {
		for _, v := range visited[r] {
			if !v {
				completed = false
				break
			}
		}
	}

	return *ans, completed
}

//
// Draw cell at x,y with the given plant, edges
// +-+-+-+
// |p|p|p|
// +-+-+-+
// each cell is 3x3 grid

func drawGrid[T any](grid [][]T) {
	for r := range grid {
		for _, v := range grid[r] {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}
}

func isSamePlant(grid [][]string, row int, col int, plant string) bool {
	return !isOutside(grid, row, col) && grid[row][col] == plant
}

func isOutside(grid [][]string, row int, col int) bool {
	return row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row])
}

func initVisited(grid [][]string) [][]bool {
	visited := [][]bool{}
	for r := range grid {
		row := []bool{}
		for j := 0; j < len(grid[r]); j++ {
			row = append(row, false)
		}
		visited = append(visited, row)
	}

	return visited
}

func readFile(filePath string) [][]string {
	file, e := os.Open(filePath)
	if e != nil {
		panic(e)
	}

	grid := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		row := []string{}
		for _, s := range strings.Split(text, "") {
			row = append(row, s)
		}
		grid = append(grid, row)
	}

	return grid
}
