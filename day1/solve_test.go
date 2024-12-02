package main

import (
	"testing"

	"github.com/IndieCoderMM/aoc-24/utils"
)

func TestSolve(t *testing.T) {
	filePath := utils.GetFilePath(1)
	expected := utils.GetAnswer(1)

	t.Run("Day-1 Solution", func(t *testing.T) {
		ans := Solve(filePath)
		if ans != expected {
			t.Errorf("Expected %d, got %d", expected, ans)
		}
	})
}
