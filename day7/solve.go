package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(filePath string) uint64 {
	testValues, lines := readFile(filePath)

	ans := GetValidTests(testValues, lines)

	return ans
}

func GetValidTests(testValues []uint64, lines [][]uint64) uint64 {
	var ans uint64 = 0
	validTests := []uint64{}

	for i, line := range lines {
		isValid := CheckTest(testValues[i], line)
		if isValid {
			ans += testValues[i]
			validTests = append(validTests, testValues[i])
		}
	}

	fmt.Printf("Valid tests %v\n", validTests)

	return ans
}

func CheckTest(expected uint64, line []uint64) bool {
	stack := []uint64{}
	operations := []string{"+", "*", "~"}

	stack = append(stack, line[0])

	for i := 1; i < len(line); i++ {
		newStack := []uint64{}
		for _, v := range stack {
			res := CalOperations(v, line[i], operations)
			for _, r := range res {
				if r > expected {
					continue
				}
				newStack = append(newStack, r)
			}
		}
		stack = newStack
	}

	for _, v := range stack {
		if v == expected {
			return true
		}
	}

	return false
}

func CalOperations(a uint64, b uint64, ops []string) []uint64 {
	ans := []uint64{}

	for _, op := range ops {
		ans = append(ans, Operate(a, b, op))
	}

	return ans
}

func Operate(a uint64, b uint64, op string) uint64 {
	var ans uint64 = 0
	switch op {
	case "+":
		ans = a + b
	case "*":
		ans = a * b
	case "~":
		n, err := strconv.ParseUint(fmt.Sprintf(`%d%d`, a, b), 10, 64)
		if err != nil {
			panic(err)
		}
		ans = n
	}

	return ans
}

func readFile(filePath string) ([]uint64, [][]uint64) {
	testValues := []uint64{}
	lines := [][]uint64{}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ": ")
		testVal, err := strconv.ParseUint(text[0], 10, 64)
		if err != nil {
			panic(err)
		}

		line := []uint64{}
		for i := 1; i < len(text); i++ {
			strNums := strings.Split(text[i], " ")
			for _, str := range strNums {
				num, err := strconv.ParseUint(str, 10, 64)
				if err != nil {
					panic(err)
				}
				line = append(line, num)
			}
			lines = append(lines, line)
		}

		testValues = append(testValues, testVal)
	}

	return testValues, lines
}
