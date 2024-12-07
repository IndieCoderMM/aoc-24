package main

import (
	"testing"

	"github.com/IndieCoderMM/aoc-24/utils"
)

func TestSolve(t *testing.T) {
	file := utils.GetFilePath(6)
	expected := utils.GetAnswer(6)

	t.Run("Day-6 Solution", func(t *testing.T) {
		ans := Solve(file)
		if ans != expected {
			t.Errorf("Expected %d, got %d", expected, ans)
		}
	})
}
