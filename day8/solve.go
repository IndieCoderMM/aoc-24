package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func Solve(filePath string) int {
	ans := 0
	grid := readFile(filePath)
	atnMap := getAntennaMap(grid)

	antinodes := make(map[Pos]bool)
	for k, v := range atnMap {
		fmt.Printf("Antenna freq: %s...\n", k)
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				nodes := getAntinodePos(grid, v[i], v[j])
				for _, pos := range nodes {
					if _, ok := antinodes[pos]; !ok {
						antinodes[pos] = true
					}
				}
			}
		}
	}

	markers := []Pos{}
	for k := range antinodes {
		markers = append(markers, k)
	}

	ans = len(markers)

	drawMap(grid, markers...)

	return ans
}

type Pos = struct {
	x int
	y int
}

func getAntinodePos(grid [][]string, a Pos, b Pos) []Pos {
	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1

	antinodes := []Pos{}
	dirAB := Pos{b.x - a.x, b.y - a.y}
	dirBA := Pos{a.x - b.x, a.y - b.y}

	// fmt.Printf("%v -> %v : %v\n", a, b, dirAB)
	// fmt.Printf("%v -> %v : %v\n", b, a, dirBA)
	// fmt.Println()

	vectors := [][2]Pos{}
	vectors = append(vectors, [2]Pos{b, dirAB})
	vectors = append(vectors, [2]Pos{a, dirBA})

	for _, v := range vectors {
		poss := placeAntinodes(v[0], v[1], maxRow, maxCol)
		antinodes = append(antinodes, poss...)
	}

	return antinodes
}

func placeAntinodes(start Pos, dir Pos, maxRow int, maxCol int) []Pos {
	antinodes := []Pos{}
	pos := Pos{start.x, start.y}

	for {
		pos.x += dir.x
		pos.y += dir.y

		if pos.x < 0 || pos.x > maxCol || pos.y < 0 || pos.y > maxRow {
			break
		}
		antinodes = append(antinodes, pos)
		break
	}

	return antinodes
}

func getAntennaMap(grid [][]string) map[string][]Pos {
	antennas := make(map[string][]Pos)
	for y, row := range grid {
		for x, c := range row {
			if c == "." {
				continue
			}

			if _, ok := antennas[c]; ok {
				antennas[c] = append(antennas[c], Pos{x, y})
			} else {
				antennas[c] = []Pos{{x, y}}
			}
		}
	}

	return antennas
}

func drawMap(grid [][]string, markers ...Pos) {
	escape := "\033[1;30;45m"

	fmt.Printf("  ")
	for i := 0; i < len(grid[0]); i++ {
		fmt.Printf("%d", i%10)
	}
	fmt.Printf("\n  ")
	for i := 0; i < len(grid[0]); i++ {
		fmt.Print("-")
	}
	fmt.Println()
	for i := 0; i < len(grid); i++ {
		fmt.Printf("%d|", i%10)
		for j := 0; j < len(grid[i]); j++ {
			if slices.Index(markers, Pos{j, i}) != -1 {
				fmt.Printf(escape)
				// fmt.Printf(escape+str+"\033[0m")
			}
			fmt.Printf("%s\033[0m", grid[i][j])
		}
		fmt.Println()
	}

	fmt.Printf(" ")
	for i := 0; i < len(grid[0]); i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func readFile(filePath string) [][]string {
	var lines [][]string
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		chars := strings.Split(text, "")
		lines = append(lines, chars)
	}

	return lines
}
