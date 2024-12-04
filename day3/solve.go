package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solve(filePath string) int {
	ans, err := readFile(filePath)
	if err != nil {
		panic(err)
	}

	return ans
}

// Find sum of all valid mul(x,y)
func CalMultipleSum(chars []string) int {
	var mulFormat []string = []string{"m", "u", "l", "("}
	ans := 0
	index := 0
	firstNum := ""
	secNum := ""
	curr := "first"

	resetVars := func() {
		index = 0
		firstNum = ""
		secNum = ""
		curr = "first"
	}

	for _, c := range chars {
		if index == 4 {
			// get first num
			if isNumeric(c) {
				if curr == "first" {
					firstNum += c
				} else if curr == "second" {
					secNum += c
				}
			} else if c == "," && curr == "first" {
				curr = "second"
			} else if c == ")" && curr == "second" {
				if firstNum != "" && secNum != "" {
					// fmt.Printf("Multiplying %s * %s\n", firstNum, secNum)
					// Numbers are complete
					first, err := strconv.Atoi(firstNum)
					if err != nil {
						panic(err)
					}
					second, err := strconv.Atoi(secNum)
					if err != nil {
						panic(err)
					}

					// multiply and add it to answer
					ans += first * second
				}

				// Reset index
				resetVars()
			} else {
				// fmt.Printf("Not numeric -> %s\n", c)
				// Reset index, not matched
				resetVars()
			}
			continue
		}

		if c == mulFormat[index] {
			// fmt.Printf("Format matched %s == %s\n", c, mulFormat[index])
			index += 1
		} else {
			index = 0
		}
	}

	return ans
}

// Conditionally decide which part of a string to calculate sum
func CalByCondition(text string, skipFirst bool) (int, bool) {
	ans := 0
	disableKey := "don't()"
	enableKey := "do()"
	endsWithDont := false

	// split by don't -> {g1}don't(){g2}
	groups := strings.Split(text, disableKey)

	for j := 0; j < len(groups); j++ {
		if j == 0 && !skipFirst {
			// Safely calculate the first part if the previous line doesn't end with *dont*
			ans += CalMultipleSum(strings.Split(groups[0], ""))
			continue
		}

		enabledParts := strings.Split(groups[j], enableKey)

		// Whether the line endsWithDont; use for the next line
		if j == len(groups)-1 && len(enabledParts) == 1 {
			endsWithDont = true
		}

		// Skip the first part
		for k := 1; k < len(enabledParts); k++ {
			ans += CalMultipleSum(strings.Split(enabledParts[k], ""))
		}
	}

	return ans, endsWithDont
}

func readFile(filePath string) (int, error) {
	ans := 0
	skipFirst := false

	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		sum, endsWithDont := CalByCondition(text, skipFirst)
		ans += sum
		skipFirst = endsWithDont
		fmt.Printf("Skipping next: %v\n", skipFirst)
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("Error reading file: %v", err)
	}

	return ans, nil
}

func isNumeric(str string) bool {
	return regexp.MustCompile(`\d`).MatchString(str)
}
