package main

import (
	"testing"

	"github.com/IndieCoderMM/aoc-24/utils"
)

func TestSolve(t *testing.T) {
	file := utils.GetFilePath(9)
	expected := utils.GetAnswer(9)

	t.Run("Day-9 Solution", func(t *testing.T) {
		ans := Solve(file)
		if ans != uint64(expected) {
			t.Errorf("Expected %d, got %d", expected, ans)
		}
	})
}
