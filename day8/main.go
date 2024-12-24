package main

import (
	"fmt"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 311 in 110.8µs
Part 2: 1115 in 250.4µs
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

	var grid [][]rune
	antennaMap := make(map[rune][]Vector)
	antinodeMap := make(map[Vector]struct{})

	for scanner.Scan() {
		runes := []rune(scanner.Text())
		grid = append(grid, runes)
	}

	for y := range grid {
		for x, r := range grid[y] {
			if r != '.' {
				antennaMap[r] = append(antennaMap[r], Vector{x: x, y: y})
			}
		}
	}

	for _, antennas := range antennaMap {
		for i, a1 := range antennas {
			for _, a2 := range antennas[i+1:] {
				diff := Vector{x: a2.x - a1.x, y: a2.y - a1.y}
				antinode1 := Vector{x: a1.x - diff.x, y: a1.y - diff.y}
				antinode2 := Vector{x: a2.x + diff.x, y: a2.y + diff.y}
				if antinode1.x >= 0 && antinode1.x < len(grid[0]) && antinode1.y >= 0 && antinode1.y < len(grid) {
					antinodeMap[antinode1] = Empty{}
				}
				if antinode2.x >= 0 && antinode2.x < len(grid[0]) && antinode2.y >= 0 && antinode2.y < len(grid) {
					antinodeMap[antinode2] = Empty{}
				}
			}
		}
	}
	return len(antinodeMap)
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)

	var grid [][]rune
	antennaMap := make(map[rune][]Vector)
	antinodeMap := make(map[Vector]struct{})

	for scanner.Scan() {
		runes := []rune(scanner.Text())
		grid = append(grid, runes)
	}

	for y := range grid {
		for x, r := range grid[y] {
			if r != '.' {
				antennaMap[r] = append(antennaMap[r], Vector{x: x, y: y})
			}
		}
	}

	for _, antennas := range antennaMap {
		for i, a1 := range antennas {
			for _, a2 := range antennas[i+1:] {
				antinodeMap[a1] = Empty{}
				antinodeMap[a2] = Empty{}
				diff := Vector{x: a2.x - a1.x, y: a2.y - a1.y}
				count := 1
				for {
					antinode1 := Vector{x: a1.x - diff.x*count, y: a1.y - diff.y*count}
					if antinode1.x >= 0 && antinode1.x < len(grid[0]) && antinode1.y >= 0 && antinode1.y < len(grid) {
						antinodeMap[antinode1] = Empty{}
						count++
					} else {
						break
					}
				}
				count = 1
				for {
					antinode2 := Vector{x: a2.x + diff.x*count, y: a2.y + diff.y*count}
					if antinode2.x >= 0 && antinode2.x < len(grid[0]) && antinode2.y >= 0 && antinode2.y < len(grid) {
						antinodeMap[antinode2] = Empty{}
						count++
					} else {
						break
					}
				}
			}
		}
	}
	return len(antinodeMap)
}

type Empty = struct{}
type Vector struct {
	x, y int
}
