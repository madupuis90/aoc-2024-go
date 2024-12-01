package main

import (
	"fmt"
	"time"

	"example.com/aoc-2024-go/util"
)

func main() {
	start := time.Now()
	part1 := part1("input.txt")
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part1 took %s\n", time.Since(start))

	start = time.Now()
	part2 := part2("input.txt")
	fmt.Printf("Part 2: %v\n", part2)
	fmt.Printf("Part2 took %s\n", time.Since(start))
}

func part1(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)

	return total
}

func part2(f string) int {
	total := 0

	return total
}
