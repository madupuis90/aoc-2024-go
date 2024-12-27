package main

import (
	"fmt"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 538191549061 in 17.9063ms
Part 2: 34612812972206 in 1.6353173s
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
	scanner := util.CreateScannerFromFile(f)
	total := 0

	return total
}

func part2(f string) int {
	// scanner := util.CreateScannerFromFile(f)
	total := 0

	return total
}
