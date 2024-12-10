package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solve(filePath string) int {
	grid := readFile(filePath)

	ans := Search(grid)

	return ans
}

func Search(grid [][]int) int {
	ans := 0

	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
			return 0
		}

		if grid[i][j] == 9 {
			return 1
		}

		neighbors := [4][2]int{
			{1, 0}, {0, 1}, {-1, 0}, {0, -1},
		}

		res := 0
		for _, n := range neighbors {
			nextI := i + n[0]
			nextJ := j + n[1]
			if nextI < 0 || nextI >= len(grid) || nextJ < 0 || nextJ >= len(grid[nextI]) {
				continue
			}
			if grid[nextI][nextJ] == grid[i][j]+1 {
				res += dfs(nextI, nextJ)
			}
		}
		return res
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				ans += dfs(i, j)
			}
		}
	}

	return ans
}

func readFile(filePath string) [][]int {
	grid := [][]int{}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), "")
		line := []int{}
		for _, s := range str {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			line = append(line, n)
		}

		grid = append(grid, line)
	}

	return grid
}
