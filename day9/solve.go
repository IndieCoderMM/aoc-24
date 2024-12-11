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
	fmt.Printf("Extracted -> %d\n", len(extracted))

	ans = insert(extracted)

	return ans
}

func insert(nums []int) uint64 {
	var ans uint64 = 0
	left := 0
	right := len(nums) - 1

	for j := right; j > left; j-- {
		for right > left && nums[right] == -1 {
			right -= 1
		}

		rightSize := 0
		for nums[right] != -1 {
			rightSize += 1
			if right-1 < 0 {
				break
			}
			if nums[right] == nums[right-1] {
				right -= 1
			} else {
				break
			}
		}

		for nums[left] != -1 {
			left += 1
			if left >= len(nums) {
				break
			}
		}
		// left at the first space

		// counting spaces
		spaces := 0
		for spaces <= rightSize && left < right {
			// fmt.Printf("Space calculating...%d\n", left)
			for left < right && nums[left] == -1 {
				spaces += 1
				left += 1
			}
			// left is at non-space index
			if spaces >= rightSize {
				break
			} else {
				// if space is not enough
				// Find next space start
				spaces = 0
				// Move left to next first space
				for left < right && nums[left] != -1 {
					left += 1
				}
				if nums[left] != -1 {
					break
				}
			}
		}

		// Swap [leftmost..(spaces)..left....right..(rightSize)..]
		// fmt.Printf(">%d..(%d)..%d..%d..(%d)..[%d]\n", leftmost, spaces, left, right, rightSize, nums[right])
		if spaces >= rightSize {
			j := rightSize - 1
			for i := 0; i < rightSize; i++ {
				nums[left-spaces+i], nums[right+j] = nums[right+j], nums[left-spaces+i]
				j -= 1
			}
			// fmt.Printf("Swapped: \n %v\n", nums)
		} else {
			// No swap
			// fmt.Printf("No Swapped: \n %v\n", nums)
		}

		// Reset left
		for i := 0; i < right; i++ {
			if nums[i] == -1 {
				left = i
				break
			}
		}

		right -= 1
		for right > left && nums[right] == -1 {
			right -= 1
		}

		if right < left {
			break
		}
	}

	// fmt.Printf("RESULT\n %v\n", nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			continue
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
