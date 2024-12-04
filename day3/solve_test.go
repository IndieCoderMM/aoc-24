package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/IndieCoderMM/aoc-24/utils"
)

func TestCalByCondition(t *testing.T) {
	testCases := utils.GetTestData("day-3-test.txt")
	strResults := utils.GetTestData("day-3-result.txt")

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			ans, _ := CalByCondition(tc, false)
			n, err := strconv.Atoi(strResults[i])
			if err != nil {
				panic("Failed to convert result to int")
			}
			if ans != n {
				t.Errorf("Expected %s, got %d", strResults[i], ans)
			}
		})
	}
}

// func TestCalMulipleSum(t *testing.T) {
// 	testCases := utils.GetTestData("day-3-test.txt")
// 	strResults := utils.GetTestData("day-3-result.txt")
//
// 	for i, tc := range testCases {
// 		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
// 			ans := CalMultipleSum(strings.Split(tc, ""))
// 			n, err := strconv.Atoi(strResults[i])
// 			if err != nil {
// 				panic("Failed to convert result to int")
// 			}
// 			if ans != n {
// 				t.Errorf("Expected %s, got %d", strResults[i], ans)
// 			}
// 		})
// 	}
// }

func TestSolve(t *testing.T) {
	file := utils.GetFilePath(3)
	expected := utils.GetAnswer(3)

	t.Run("Day-3 Solution", func(t *testing.T) {
		ans := Solve(file)
		if ans != expected {
			t.Errorf("Expected %d, got %d", expected, ans)
		}
	})
}
