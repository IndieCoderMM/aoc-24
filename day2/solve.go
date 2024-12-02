package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(filePath string) int {
	lines, err := readFile(filePath)
	if err != nil {
		panic(err)
	}

	ans := GetSafeLines(lines)

	return ans
}

func GetSafeLines(lines [][]int) int {
	var ans int = 0

	for i := range lines {
		line := lines[i]
		mod := getLevelOrder(line)
		if mod == 0 && len(line) > 1 {
			// Unsafe level
			fmt.Printf("Skipping line %v\n", line)
			continue
		}
		safe := checkLine(line, mod)
		if safe {
			ans += 1
		}
		fmt.Printf("Checking line %v, result (%v)\n", line, safe)
	}

	return ans
}

func getLevelOrder(line []int) int {
	inCount := 0
	deCount := 0
	for i := 0; i < len(line)-1; i++ {
		diff := line[i+1] - line[i]
		if diff > 0 {
			inCount += 1
		} else {
			deCount += 1
		}
	}

	if inCount > deCount {
		return 1
	} else if deCount > inCount {
		return -1
	}

	return 0
}

// Line is safe if:
// - All increasing/decreasing
// - Adjacent levels differ by 1, 2, 3
// - Able to remove one unsafe level (dampable)
func checkLine(line []int, mod int) bool {
	if len(line) == 1 {
		return true
	}

	return helper(line, 0, false, mod)
}

func helper(line []int, index int, damped bool, mod int) bool {
	if index >= len(line)-1 {
		return true
	}

	isSafe := isSafeLevels(line[index], line[index+1], mod)
	if isSafe {
		return helper(line, index+1, damped, mod)
	}

	if !damped {
		// Able to damp one level
		// Try dumping next
		// check if curr and nNext are safe
		if index+1 == len(line)-1 || isSafeLevels(line[index], line[index+2], mod) {
			// fmt.Printf("Removed %d @%d from line %v\n", line[index+1], index+1, line)
			return helper(line, index+2, true, mod)
		}

		// Try dumping current
		// check if prev and next are safe
		if index == 0 || isSafeLevels(line[index-1], line[index+1], mod) {
			// fmt.Printf("Removed %d @%d from line %v\n", line[index], index, line)
			return helper(line, index+1, true, mod)
		}
	}

	// if damped {
	// 	fmt.Printf("Already damped, %v\n", line)
	// } else {
	// 	fmt.Printf("Cannot damp any, %v\n", line)
	// }

	return false
}

func isSafeLevels(curr int, next int, mod int) bool {
	diff := next - curr

	if diff == 0 {
		return false
	}

	return diff*mod > 0 && diff*mod <= 3
}

func readFile(filePath string) ([][]int, error) {
	var lines [][]int

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strLine := strings.Split(scanner.Text(), " ") // Split by 1 space
		line := []int{}

		for _, n := range strLine {
			intN, err := strconv.Atoi(n)
			if err != nil {
				return nil, fmt.Errorf("Failed to parse int: %v", err)
			}
			line = append(line, intN)
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	return lines, nil
}
