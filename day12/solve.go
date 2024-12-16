package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve(filePath string) int {
	grid := readFile(filePath)

	// drawGrid(grid)
	ans := calculate(grid)

	return ans
}

type Cell = struct {
	row int
	col int
}

var directions = [4]Cell{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func calculate(grid [][]string) int {
	visited := initVisited(grid)
	ans := 0

	for r, row := range grid {
		for c, p := range row {
			if visited[r][c] {
				continue
			}

			stack := []Cell{{r, c}}
			visited[r][c] = true
			currArea := 1
			currPerimeter := 0
			for len(stack) > 0 {
				cell := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				perimeter := 0

				for _, dir := range directions {
					newRow, newCol := cell.row+dir.row, cell.col+dir.col
					if !isSamePlant(grid, newRow, newCol, p) {
						perimeter++ // increase the perimeter if touching the outside or different plant
						continue
					}
					if visited[newRow][newCol] {
						continue
					}

					visited[newRow][newCol] = true
					stack = append(stack, Cell{newRow, newCol})
					currArea++
				}

				currPerimeter += perimeter
			}

			// fmt.Printf("Total Area %d, Perimeter (%d)\n", currArea, currPerimeter)
			ans += currArea * currPerimeter
		}
	}

	return ans
}

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
