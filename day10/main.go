package main

import (
	"fmt"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 659 in 337.2µs
Part 2: 1463 in 178.8µs
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

	var grid [][]int
	visited := make(map[UniqueReachable]struct{})

	for scanner.Scan() {
		grid = append(grid, util.RuneSliceAtoi([]rune(scanner.Text())))
	}

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 0 {
				total = total + climb(grid, &Vector{x, y}, 0, visited, &Vector{x, y}, true)
			}
		}
	}

	return total
}

func climb(grid [][]int, pos *Vector, height int, visited map[UniqueReachable]struct{}, originalPos *Vector, removeDuplicate bool) int {
	sum := 0
	nextHeight := height + 1

	// I was actually doing the part 2 logic in part 1...I added this janky map to remove duplicates
	if removeDuplicate {
		if height == 9 {
			u := UniqueReachable{peak: *pos, origin: *originalPos}
			if _, exists := visited[u]; exists {
				return sum
			} else {
				visited[u] = struct{}{}
				return sum + 1
			}
		}
	} else {
		if height == 9 {
			return sum + 1
		}
	}

	// up
	if nextX := pos.X - 1; nextX >= 0 && grid[nextX][pos.Y] == nextHeight {
		sum += climb(grid, &Vector{X: nextX, Y: pos.Y}, nextHeight, visited, originalPos, removeDuplicate)
	}
	// down
	if nextX := pos.X + 1; nextX < len(grid) && grid[nextX][pos.Y] == nextHeight {
		sum += climb(grid, &Vector{X: nextX, Y: pos.Y}, nextHeight, visited, originalPos, removeDuplicate)
	}
	// left
	if nextY := pos.Y - 1; nextY >= 0 && grid[pos.X][nextY] == nextHeight {
		sum += climb(grid, &Vector{X: pos.X, Y: nextY}, nextHeight, visited, originalPos, removeDuplicate)
	}
	// right
	if nextY := pos.Y + 1; nextY < len(grid[0]) && grid[pos.X][nextY] == nextHeight {
		sum += climb(grid, &Vector{X: pos.X, Y: nextY}, nextHeight, visited, originalPos, removeDuplicate)
	}

	return sum
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	total := 0

	var grid [][]int
	visited := make(map[UniqueReachable]struct{})

	for scanner.Scan() {
		grid = append(grid, util.RuneSliceAtoi([]rune(scanner.Text())))
	}

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 0 {
				total = total + climb(grid, &Vector{x, y}, 0, visited, &Vector{x, y}, false)
			}
		}
	}

	return total
}

type Vector struct {
	X, Y int
}

type UniqueReachable struct {
	origin, peak Vector
}
