package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(filePath string) uint64 {
	var ans uint64 = 0
	line := readFile(filePath)

	extracted := extract(line)

	fmt.Printf("Line -> %v\n", len(line))
	fmt.Printf("Extracted -> \n%v\n", len(extracted))

	ans = insert(extracted)

	return ans
}

func insert(nums []int) uint64 {
	var ans uint64 = 0
	left := 0
	right := len(nums) - 1

	for left < right {
		for nums[left] != -1 {
			left += 1
		}
		for nums[right] == -1 {
			right -= 1
		}

		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left += 1
			right -= 1
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			break
		}
		ans += uint64(nums[i] * i)
	}

	return ans
}

func extract(line []int) []int {
	extracted := []int{}
	index := 0

	for i := 0; i < len(line); i += 2 {
		totalBlock := line[i]
		file := []int{}
		for j := 0; j < totalBlock; j++ {
			file = append(file, index)
		}
		extracted = append(extracted, file...)

		j := i + 1
		if j >= len(line) {
			continue
		}
		totalSpace := line[j]
		spaces := []int{}
		for j := 0; j < totalSpace; j++ {
			spaces = append(spaces, -1)
		}
		extracted = append(extracted, spaces...)
		index += 1
	}

	fmt.Printf("Last index %d\n", index)

	return extracted
}

func readFile(filePath string) []int {
	line := []int{}
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		splits := strings.Split(text, "")

		for _, n := range splits {
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			line = append(line, num)
		}
	}

	return line
}
