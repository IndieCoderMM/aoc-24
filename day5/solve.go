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
	rules, pages, err := readFile(filePath)
	if err != nil {
		panic(err)
	}

	// Sort all pages in rules
	for _, rule := range rules {
		utils.QuickSort(rule, 0, len(rule)-1, false)
	}

	fmt.Printf("Rules: %v\n", rules)

	ans := scanIncorrectPages(rules, pages)

	return ans
}

func scanIncorrectPages(rules map[int][]int, pages [][]int) int {
	ans := 0

	for _, line := range pages {
		valid := true
		// scan the line
		for i, page := range line {
			beforePages := rules[page]
			if len(beforePages) == 0 {
				continue
			}
			// the previous pages must not be in this list
			for j := i - 1; j >= 0; j-- {
				if _, found := utils.BinarySearch(beforePages, line[j]); found {
					// Invalid
					valid = false
					break
				}
			}
		}
		// Skip valid line
		if valid {
			continue
		}

		// Swap order in place
		for i := len(line) - 1; i >= 0; i-- {
			beforePages := rules[line[i]]
			if len(beforePages) == 0 {
				continue
			}

			swapped := false
			for j := i - 1; j >= 0; j-- {
				if _, found := utils.BinarySearch(beforePages, line[j]); found {
					line[i], line[j] = line[j], line[i]
					swapped = true
				}
			}

			if swapped {
				i += 1
			}
		}

		mid := len(line) / 2
		ans += line[mid]
	}

	return ans
}

func scanCorrectPages(rules map[int][]int, pages [][]int) int {
	ans := 0
	for _, line := range pages {
		valid := true
		for i, page := range line {
			beforePages := rules[page]
			if len(beforePages) > 0 {
				// the previous pages must not be in this list
				for j := i - 1; j >= 0; j-- {
					if _, found := utils.BinarySearch(beforePages, line[j]); found {
						// Invalid
						valid = false
						break
					}
				}
			}
		}
		if valid {
			fmt.Printf("Valid line %v\n", line)
			mid := len(line) / 2
			ans += line[mid]
		}
	}

	return ans
}

func readFile(filePath string) (map[int][]int, [][]int, error) {
	rules := make(map[int][]int)
	pages := [][]int{}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to open file %v\n", err)
	}
	defer file.Close()

	ruleEnds := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if !ruleEnds {
			rule := strings.Split(text, "|")
			if len(rule) != 2 {
				ruleEnds = true
				continue
			}

			left, err := strconv.Atoi(rule[0])
			if err != nil {
				return nil, nil, fmt.Errorf("Failed to parse %v, %v \n", rule, err)
			}

			right, err := strconv.Atoi(rule[1])
			if err != nil {
				panic(err)
			}

			curr, ok := rules[left]
			if ok {
				rules[left] = append(curr, right)
			} else {
				rules[left] = []int{right}
			}

			continue
		}

		// Reading pages...
		nums := []int{}
		for _, s := range strings.Split(text, ",") {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)

		}
		pages = append(pages, nums)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("Failed to scan file %v\n", err)
	}

	return rules, pages, nil
}
