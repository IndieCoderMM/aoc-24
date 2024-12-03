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

var mulFormat []string = []string{"m", "u", "l", "("}
var dontFormat []string = []string{"d", "o", "n", "'", "t", "(", ")"}
var doFormat []string = []string{"d", "o", "(", ")"}

func CalMultipleSum(chars []string) int {
	ans := 0
	index := 0
	firstNum := ""
	secNum := ""
	curr := "first"

	fmt.Printf("Calculating: %v\n", chars)

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

func CalByCondition(text string) int {
	ans := 0
	disableKey := "don't()"
	enableKey := "do()"

	groups := strings.Split(text, disableKey)
	fmt.Printf("[DONT] Splitted: %v\n", len(groups))    // >> [all-dos, (dont),]
	ans += CalMultipleSum(strings.Split(groups[0], "")) // >> First part all safe

	for j := 1; j < len(groups); j++ {
		enabledParts := strings.Split(groups[j], enableKey)
		fmt.Printf("[DO] Splitted: %v\n", len(enabledParts))

		// Skip the first part
		for k := 1; k < len(enabledParts); k++ {
			ans += CalMultipleSum(strings.Split(enabledParts[k], ""))
		}
	}

	return ans
}

func readFile(filePath string) (int, error) {
	ans := 0

	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		ans += CalByCondition(text)
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("Error reading file: %v", err)
	}

	return ans, nil
}

func isNumeric(str string) bool {
	return regexp.MustCompile(`\d`).MatchString(str)
}
