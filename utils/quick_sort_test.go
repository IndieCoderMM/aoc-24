package utils

import (
	"testing"
)

var testCases = []struct {
	input    []int // Input array
	reverse  bool  // Sort in descending order?
	expected []int // Expected output
}{
	// Basic cases
	{[]int{3, 5, 1, 4, 2}, false, []int{1, 2, 3, 4, 5}}, // Ascending sort
	{[]int{3, 5, 1, 4, 2}, true, []int{5, 4, 3, 2, 1}},  // Descending sort

	// Cases with duplicates
	{[]int{3, 5, 1, 4, 2, 3}, false, []int{1, 2, 3, 3, 4, 5}}, // Ascending with duplicates
	{[]int{3, 5, 1, 4, 2, 3}, true, []int{5, 4, 3, 3, 2, 1}},  // Descending with duplicates

	// Already sorted arrays
	{[]int{1, 2, 3, 4, 5}, false, []int{1, 2, 3, 4, 5}}, // Ascending, no change needed
	{[]int{1, 2, 3, 4, 5}, true, []int{5, 4, 3, 2, 1}},  // Descending

	// Reverse sorted arrays
	{[]int{5, 4, 3, 2, 1}, false, []int{1, 2, 3, 4, 5}}, // Ascending
	{[]int{5, 4, 3, 2, 1}, true, []int{5, 4, 3, 2, 1}},  // Descending, no change needed

	// Edge cases
	{[]int{}, false, []int{}},                     // Empty array
	{[]int{}, true, []int{}},                      // Empty array descending
	{[]int{1}, false, []int{1}},                   // Single element, ascending
	{[]int{1}, true, []int{1}},                    // Single element, descending
	{[]int{1, 1, 1, 1}, false, []int{1, 1, 1, 1}}, // All elements identical
	{[]int{1, 1, 1, 1}, true, []int{1, 1, 1, 1}},  // All elements identical, descending

	// Large input
	{[]int{100, 99, 98, 97, 96, 95}, false, []int{95, 96, 97, 98, 99, 100}},       // Large, ascending
	{[]int{100, 99, 98, 97, 96, 95}, true, []int{100, 99, 98, 97, 96, 95}},        // Large, descending
	{[]int{10, 20, 30, 40, 50, 20, 10}, false, []int{10, 10, 20, 20, 30, 40, 50}}, // Large with duplicates, ascending
	{[]int{10, 20, 30, 40, 50, 20, 10}, true, []int{50, 40, 30, 20, 20, 10, 10}},  // Large with duplicates, descending

	// Negative numbers
	{[]int{-3, -1, -4, -2, 0}, false, []int{-4, -3, -2, -1, 0}}, // Ascending with negatives
	{[]int{-3, -1, -4, -2, 0}, true, []int{0, -1, -2, -3, -4}},  // Descending with negatives

	// Mixed positive and negative numbers
	{[]int{-3, 5, -1, 4, 2, -4}, false, []int{-4, -3, -1, 2, 4, 5}}, // Mixed, ascending
	{[]int{-3, 5, -1, 4, 2, -4}, true, []int{5, 4, 2, -1, -3, -4}},  // Mixed, descending

	// Large numbers
	{[]int{1_000_000, 500_000, -1_000_000, 0}, false, []int{-1_000_000, 0, 500_000, 1_000_000}}, // Large numbers, ascending
	{[]int{1_000_000, 500_000, -1_000_000, 0}, true, []int{1_000_000, 500_000, 0, -1_000_000}},  // Large numbers, descending
}

func TestQuickSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run("Quick sort testing", func(t *testing.T) {
			QuickSort(tc.input, 0, len(tc.input)-1, tc.reverse)
			for i := 0; i < len(tc.input); i++ {
				if tc.input[i] != tc.expected[i] {
					t.Errorf("Expected %v, got %v", tc.expected, tc.input)
				}
			}
		})
	}
}
