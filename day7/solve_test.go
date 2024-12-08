package main

import (
	"testing"

	"github.com/IndieCoderMM/aoc-24/utils"
)

func TestSolve(t *testing.T) {
	file := utils.GetFilePath(7)
	expected := utils.GetAnswer(7)

	t.Run("Day-7 Solution", func(t *testing.T) {
		ans := Solve(file)
		if ans != uint64(expected) {
			t.Errorf("Expected %d, got %d", expected, ans)
		}
	})
}

// func TestCheckTest(t *testing.T) {
// 	testCases := []struct {
// 		testVal  uint64
// 		line     []uint64
// 		expected bool
// 	}{
// 		{190, []uint64{10, 19}, true},
// 		{3267, []uint64{81, 40, 27}, true},
// 		{161011, []uint64{16, 10, 13}, false},
// 		{192, []uint64{17, 8, 14}, false},
// 		{292, []uint64{11, 6, 16, 20}, true},
// 	}
//
// 	for _, tc := range testCases {
// 		t.Run("Case", func(t *testing.T) {
// 			ans := CheckTest(tc.testVal, tc.line)
// 			if ans != tc.expected {
// 				t.Errorf("Expected %v, got %v\n", tc.expected, ans)
// 			}
// 		})
// 	}
// }
