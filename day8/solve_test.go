package main

import (
	"testing"

	"github.com/IndieCoderMM/aoc-24/utils"
)

func TestSolve(t *testing.T) {
	file := utils.GetFilePath(8)
	expected := utils.GetAnswer(8)

	t.Run("Day-8 Solution", func(t *testing.T) {
		ans, _, _ := Solve(file)
		if ans != expected {
			t.Errorf("Expected %d, got %d", expected, ans)
		}
	})
}
