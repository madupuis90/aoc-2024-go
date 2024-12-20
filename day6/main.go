package main

import (
	"errors"
	"fmt"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 4973 in 491.5µs
Part 2: 1482 in 1.4188431s
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
	var pos Vector
	var dir rune

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	for i, row := range grid {
		for j, rune := range row {
			if rune == '^' {
				pos = Vector{x: i, y: j}
				dir = grid[i][j]
				break
			}
		}
	}

	visited, _ := walk(grid, pos, dir)

	return len(visited)
}

func part2(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)
	var grid [][]rune
	var pos Vector
	var dir rune

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	for i, row := range grid {
		for j, rune := range row {
			if rune == '^' {
				pos = Vector{x: i, y: j}
				dir = grid[i][j]
				break
			}
		}
	}

	visited, _ := walk(grid, pos, dir)

	// I spent a bit too much time trying to come up with something clever, ended up brute forcing by putting an obstacle on every location
	// that the guard previously walked on
	for v := range visited {
		grid[v.x][v.y] = '#'
		_, err := walk(grid, pos, dir)
		if err != nil {
			total++
		}
		grid[v.x][v.y] = '.'
	}

	return total
}

func walk(grid [][]rune, pos Vector, dir rune) (map[Vector]Empty, error) {
	visited := make(map[Vector]Empty)
	vecs := make(map[PreviousPath]Empty)
	v := Vector{x: 0, y: 0}

	for {
		visited[pos] = Empty{}
		var nextPos Vector
		var nextDir rune

		switch dir {
		case '^':
			nextPos = Vector{x: pos.x - 1, y: pos.y}
			nextDir = '>'
			v.x--
		case '>':
			nextPos = Vector{x: pos.x, y: pos.y + 1}
			nextDir = 'v'
			v.y++
		case 'v':
			nextPos = Vector{x: pos.x + 1, y: pos.y}
			nextDir = '<'
			v.x++
		case '<':
			nextPos = Vector{x: pos.x, y: pos.y - 1}
			nextDir = '^'
			v.y--
		}

		// Check if out of bound
		if nextPos.x < 0 || nextPos.x >= len(grid) || nextPos.y < 0 || nextPos.y >= len(grid[0]) {
			break
		}

		if grid[nextPos.x][nextPos.y] == '#' {
			PreviousPath := PreviousPath{dir: dir, vec: v, pos: pos}
			if _, exists := vecs[PreviousPath]; exists {
				return nil, errors.New("looping")
			} else {
				vecs[PreviousPath] = Empty{}
			}
			v = Vector{}
			dir = nextDir
		} else {
			pos = nextPos
		}
	}
	return visited, nil
}

type Empty = struct{}

type Vector struct {
	x, y int
}

type PreviousPath struct {
	vec Vector
	dir rune
	pos Vector
}
