package main

import (
	"fmt"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 2534 in 354.4µs
Part 2: 1866 in 167.3µs
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

	var grid [][]rune

	for scanner.Scan() {
		runes := []rune(scanner.Text())
		grid = append(grid, runes)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'X' {
				for k := -1; k <= 1; k++ {
					for l := -1; l <= 1; l++ {
						if i+k*3 < 0 || i+k*3 >= len(grid) || j+l*3 < 0 || j+l*3 >= len(grid[i]) {
							continue // check bounds
						}
						if grid[i+k*1][j+l*1] == 'M' && grid[i+k*2][j+l*2] == 'A' && grid[i+k*3][j+l*3] == 'S' {
							total++
						}
					}
				}
			}
		}
	}

	return total
}

func part2(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)

	var grid [][]rune

	for scanner.Scan() {
		runes := []rune(scanner.Text())
		grid = append(grid, runes)
	}
	for i := range grid {
		if i == 0 || i == len(grid)-1 {
			continue // remove outer rows
		}
		for j := range grid[i] {
			if j == 0 || j == len(grid[i])-1 {
				continue // remove outer columns
			}
			if grid[i][j] == 'A' {
				if true &&
					(grid[i-1][j-1] == 'S' && grid[i+1][j-1] == 'S' && grid[i-1][j+1] == 'M' && grid[i+1][j+1] == 'M') ||
					(grid[i-1][j-1] == 'M' && grid[i+1][j-1] == 'M' && grid[i-1][j+1] == 'S' && grid[i+1][j+1] == 'S') ||
					(grid[i-1][j-1] == 'M' && grid[i+1][j-1] == 'S' && grid[i-1][j+1] == 'M' && grid[i+1][j+1] == 'S') ||
					(grid[i-1][j-1] == 'S' && grid[i+1][j-1] == 'M' && grid[i-1][j+1] == 'S' && grid[i+1][j+1] == 'M') {
					total++
				}
			}
		}
	}

	return total
}
