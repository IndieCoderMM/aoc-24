package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	BLINK       = 75
	MAX_WORKERS = 8
)

func Solve(filePath string) uint64 {
	line := readFile(filePath)

	ans := run(line, len(line)/MAX_WORKERS, MAX_WORKERS)
	// ans := blinkForTimes(line)

	return ans
}

func run(line []uint64, chunkSize int, maxWorkers int) uint64 {
	var ans uint64
	var wg sync.WaitGroup

	results := make(chan uint64, len(line))
	workers := make(chan struct{}, maxWorkers)

	for i := 0; i < len(line); i += chunkSize {
		end := i + chunkSize
		if end > len(line) {
			end = len(line)
		}

		chunk := line[i:end]
		workers <- struct{}{}

		wg.Add(1)
		go func(chunk []uint64) {
			defer wg.Done()
			processChunk(chunk, results)
			<-workers
		}(chunk)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		ans += res
	}

	return ans
}

func processChunk(chunk []uint64, result chan uint64) {
	total := blinkForTimes(chunk)

	result <- total
}

func blinkForTimes(stones []uint64) uint64 {
	var ans uint64
	stoneBlinkMap := make(map[uint64]uint64)

	// Blink one stone for times
	var blink func(num uint64, times int) uint64
	blink = func(num uint64, times int) uint64 {
		if times == 0 {
			return 1
		}
		key := (num << 32) | uint64(times)
		if cached, ok := stoneBlinkMap[key]; ok {
			return cached
		}

		blinked := Blink(num)

		var res uint64
		for _, s := range blinked {
			res += blink(s, times-1)
		}

		stoneBlinkMap[key] = res
		return res
	}

	for _, s := range stones {
		ans += blink(s, BLINK)
	}

	return ans
}

func Blink(num uint64) []uint64 {
	if num == 0 {
		return []uint64{1}
	}

	digits := lenDigits(num)
	if digits%2 == 0 {
		first, sec := split(num, digits/2)
		return []uint64{first, sec}
	}

	return []uint64{num * 2024}
}

func split(n uint64, from int) (uint64, uint64) {
	leftDivisor := uint64(math.Pow(10, float64(from)))

	leftPart := n / leftDivisor
	rightPart := n % leftDivisor

	return leftPart, rightPart
}

func lenDigits(n uint64) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(float64(n))) + 1
}

func readFile(filePath string) []uint64 {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	line := []uint64{}

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")
		for _, s := range strs {
			n, e := strconv.ParseUint(s, 10, 64)
			if e != nil {
				panic(e)
			}
			line = append(line, n)
		}
	}

	return line
}
