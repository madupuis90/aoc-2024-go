package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 670 in 353.7µs
Part 2: 700 in 381.9µs
*/
func main() {
	start := time.Now()
	part1 := part1("input.txt")
	fmt.Printf("Part 1: %v in %s\n", part1, time.Since(start))

	start = time.Now()
	part2 := part2("input.txt")
	fmt.Printf("Part 2: %v in %s\n", part2, time.Since(start))
}

func part1(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		fieldsInt := util.StringSliceAtoi(fields)
		total = total + isSafe(fieldsInt)
	}

	return total
}

func isSafe(s []int) int {

	var diff int
	increasing := s[1] > s[0]

	for i := 1; i < len(s); i++ {
		if increasing {
			diff = s[i] - s[i-1]
		} else {
			diff = s[i-1] - s[i]
		}
		if diff < 1 || diff > 3 {
			return 0
		}
	}
	return 1
}

func part2(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		fieldsInt := util.StringSliceAtoi(fields)
		total = total + isSafe2(fieldsInt)
	}

	return total
}

func isSafe2(s []int) int {

	var diff int
	increasing := s[1] > s[0]

	for i := 1; i < len(s); i++ {
		if increasing {
			diff = s[i] - s[i-1]
		} else {
			diff = s[i-1] - s[i]
		}
		if diff < 1 || diff > 3 {
			// In case of error, brute force removing 1 report at the time
			// I had more complex logic before, but was off by 2 reports
			for j := range s {
				newSlice := make([]int, len(s))
				_ = copy(newSlice, s)
				newSlice = slices.Delete(newSlice, j, j+1)
				if isSafe(newSlice) == 1 {
					return 1
				}
			}
			return 0
		}
	}
	return 1
}
