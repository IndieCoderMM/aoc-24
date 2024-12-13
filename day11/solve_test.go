package main

import (
	"testing"

	"github.com/IndieCoderMM/aoc-24/utils"
)

func TestSolve(t *testing.T) {
	file := utils.GetFilePath(11)
	expected := utils.GetAnswer(11)

	t.Run("Day-11 Solution", func(t *testing.T) {
		ans := Solve(file)
		if ans != uint64(expected) {
			t.Errorf("Expected %d, got %d", expected, ans)
		}
	})
}

func BenchmarkSolve(b *testing.B) {
	file := utils.GetFilePath(11)
	// expected := utils.GetAnswer(11)

	for i := 0; i < b.N; i++ {
		Solve(file)
	}
}
