package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/IndieCoderMM/aoc-24/utils"
)

func Solve(filePath string) int {
	left, right, err := readFile(filePath)
	if err != nil {
		panic(err)
	}

	ans := FindTotalDistance(left, right)
	fmt.Printf("Total distance: %d\n", ans)

	score := FindSimilarityScore(left, right)

	return score
}

func FindSimilarityScore(left []int, right []int) int {
	counter := make(map[int]int)
	for _, n := range right {
		counter[n] += 1
	}

	var score int = 0
	for _, n := range left {
		score += n * counter[n]
	}

	return score
}

func FindTotalDistance(left []int, right []int) int {
	utils.QuickSort(left, 0, len(left)-1, false)
	utils.QuickSort(right, 0, len(right)-1, false)

	if len(left) != len(right) {
		panic("Arrays of different sizes")
	}

	var totalDistance int = 0
	for i := range left {
		dist := left[i] - right[i]
		if dist < 0 {
			dist *= -1
		}

		totalDistance += dist
	}

	return totalDistance
}

func readFile(filePath string) ([]int, []int, error) {
	var left []int
	var right []int

	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ") // Split by 3 spaces
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("Invalid line format: %s", line)
		}

		leftNum, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, nil, fmt.Errorf("Failed to parse left number: %v", err)
		}

		rightNum, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, nil, fmt.Errorf("Failed to parse right number: %v", err)
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("Error reading file: %v", err)
	}

	return left, right, nil
}
