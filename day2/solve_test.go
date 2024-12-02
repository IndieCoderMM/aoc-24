package main

import (
	"fmt"
	"testing"

	"github.com/IndieCoderMM/aoc-24/utils"
)

func TestSolve(t *testing.T) {
	file := utils.GetFilePath(2)
	expected := utils.GetAnswer(2)

	t.Run("Day-2 Solution", func(t *testing.T) {
		ans := Solve(file)
		if ans != expected {
			t.Errorf("Expected %d, got %d", expected, ans)
		}
	})
}

// Custom test cases

type TestCase = struct {
	name     string
	lines    [][]int
	expected int
}

var testCases = []TestCase{
	{
		name: "All safe lines",
		lines: [][]int{
			{1, 2, 3, 4, 5, 6},
			{10, 12, 14, 16},
			{20, 17, 14, 11, 8},
			{1, 4, 7, 10, 13},
			{987, 988, 989, 990},
		},
		expected: 5,
	},
	{
		name: "Mixed safe lines (increasing and decreasing)",
		lines: [][]int{
			{1, 4, 7, 8, 11, 14, 16, 17},
			{5, 3, 1, -1, -3},
			{8, 11, 12, 13, 15, 17, 18},
			{78, 75, 72, 69, 66, 63, 60},
			{-64, -67, -70, -73, -76, -79},
		},
		expected: 5,
	},
	{
		name: "Dampable lines",
		lines: [][]int{
			{10, 7, 4, 1},
			{1, 3, 5, 7},
			{1, 3, 5, 8, 7},   // 8 is dampable
			{1, 3, 5, 7, 19},  // 19 is dampable
			{33, 2, 5, 8, 11}, // 33 is dampable
		},
		expected: 5,
	},
	{
		name: "Mixed dampable lines",
		lines: [][]int{
			{1, 3, 5, 8, 7},
			{1, 3, 5, 7, 19},
			{33, 2, 5, 8, 11},
			{47, 45, 43, 41, 89, 38, 35, 33},
		},
		expected: 4,
	},
	{
		name: "Unfixable lines",
		lines: [][]int{
			{1, 5, 2, 8, 4},
			{10, 7, 1, 5, 0},
			{2, 5, 8, 12, 13},
			{2, 3, 3, 4, 6, 7, 100},
			{45, 43, 59, 42, 40, 35},
		},
		expected: 0,
	},
	{
		name:     "Empty lines",
		lines:    [][]int{},
		expected: 0,
	},
	{
		name: "Single-number lines",
		lines: [][]int{
			{5},
			{10},
		},
		expected: 2,
	},
	{
		name: "Two-number lines",
		lines: [][]int{
			{1, 3},
			{10, 7},
			{5, 5},
		},
		expected: 3,
	},
}

func TestGetSafeLines(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Printf("Case: %v\n", tc.name)
			ans := GetSafeLines(tc.lines)
			if ans != tc.expected {
				t.Errorf("Expected %d for %v, got %d", tc.expected, tc.lines, ans)
			}
		})
	}
}
