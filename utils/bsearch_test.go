package utils

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var testCases = []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 2, 1},
	}
	for _, tc := range testCases {
		t.Run("Binary search testing", func(t *testing.T) {
			result, _ := BinarySearch(tc.input, tc.target)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
