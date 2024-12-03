package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func GetFilePath(day int) string {
	filename := fmt.Sprintf("day-%d.txt", day)
	cwd, err := os.Getwd()
	if err != nil {
		panic("Failed to get current working directory")
	}

	filePath := filepath.Join(cwd, "../data", filename)

	return filePath
}

func GetTestData(fileName string) []string {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Failed to get current working directory")
	}

	filePath := filepath.Join(cwd, "../data", "tests", fileName)

	data := readLines(filePath)

	return data
}

func readLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic("Failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic("Failed to read file")
	}

	return data
}

func GetAnswer(day int) int {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Failed to get current working directory")
	}

	filePath := filepath.Join(cwd, "../data", "answers.txt")

	file, err := os.Open(filePath)
	if err != nil {
		panic("Failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var d int
		var a int
		fmt.Sscanf(scanner.Text(), "%d %d", &d, &a)
		if d == day {
			return a
		}
	}
	if err := scanner.Err(); err != nil {
		panic("Failed to read file")
	}

	return -1
}
