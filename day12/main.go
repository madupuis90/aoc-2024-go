package main

import (
	"fmt"
	"slices"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 1424006 in 6.6118ms
Part 2: 858684 in 11.4734ms
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

	visited := make(map[Vector]Empty)

	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	for x, row := range grid {
		for y := range row {
			v := Vector{X: x, Y: y}
			if _, exists := visited[v]; exists {
				continue
			} else {
				r := exploreRegion(grid, v, visited)
				total += r.Area * r.Edges
			}
		}
	}
	return total
}

func exploreRegion(grid [][]rune, pos Vector, visited map[Vector]Empty) Region {

	region := Region{Edges: 0, Area: 1}
	plant := grid[pos.X][pos.Y]
	visited[pos] = Empty{}

	// up
	if nextX := pos.X - 1; nextX >= 0 && grid[nextX][pos.Y] == plant {
		if _, exists := visited[Vector{X: nextX, Y: pos.Y}]; !exists {
			nr := exploreRegion(grid, Vector{X: nextX, Y: pos.Y}, visited)
			region.Area += nr.Area
			region.Edges += nr.Edges
		}
	} else {
		region.Edges++
	}

	// down
	if nextX := pos.X + 1; nextX < len(grid) && grid[nextX][pos.Y] == plant {
		if _, exists := visited[Vector{X: nextX, Y: pos.Y}]; !exists {
			nr := exploreRegion(grid, Vector{X: nextX, Y: pos.Y}, visited)
			region.Area += nr.Area
			region.Edges += nr.Edges
		}
	} else {
		region.Edges++
	}

	// left
	if nextY := pos.Y - 1; nextY >= 0 && grid[pos.X][nextY] == plant {
		if _, exists := visited[Vector{X: pos.X, Y: nextY}]; !exists {
			nr := exploreRegion(grid, Vector{X: pos.X, Y: nextY}, visited)
			region.Area += nr.Area
			region.Edges += nr.Edges
		}
	} else {
		region.Edges++
	}

	// right
	if nextY := pos.Y + 1; nextY < len(grid[0]) && grid[pos.X][nextY] == plant {
		if _, exists := visited[Vector{X: pos.X, Y: nextY}]; !exists {
			nr := exploreRegion(grid, Vector{X: pos.X, Y: nextY}, visited)
			region.Area += nr.Area
			region.Edges += nr.Edges
		}
	} else {
		region.Edges++
	}
	return region
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	total := 0
	visited := make(map[Vector]Empty)

	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	for x, row := range grid {
		for y := range row {
			v := Vector{X: x, Y: y}
			if _, exists := visited[v]; exists {
				continue
			} else {
				edges := Edges{
					up:    make(map[int][]int),
					down:  make(map[int][]int),
					left:  make(map[int][]int),
					right: make(map[int][]int),
				}
				r := exploreRegion2(grid, v, visited, &edges)
				total += r.Area * countEdges(&edges)
			}
		}
	}
	return total
}

func exploreRegion2(grid [][]rune, pos Vector, visited map[Vector]Empty, edges *Edges) Region2 {

	region := Region2{
		Area: 1,
	}
	plant := grid[pos.X][pos.Y]
	visited[pos] = Empty{}

	// up
	if nextX := pos.X - 1; nextX >= 0 && grid[nextX][pos.Y] == plant {
		if _, exists := visited[Vector{X: nextX, Y: pos.Y}]; !exists {
			nr := exploreRegion2(grid, Vector{X: nextX, Y: pos.Y}, visited, edges)
			region.Area += nr.Area
		}
	} else {
		edges.up[pos.X] = append(edges.up[pos.X], pos.Y)
	}

	// down
	if nextX := pos.X + 1; nextX < len(grid) && grid[nextX][pos.Y] == plant {
		if _, exists := visited[Vector{X: nextX, Y: pos.Y}]; !exists {
			nr := exploreRegion2(grid, Vector{X: nextX, Y: pos.Y}, visited, edges)
			region.Area += nr.Area
		}
	} else {
		edges.down[pos.X] = append(edges.down[pos.X], pos.Y)
	}

	// left
	if nextY := pos.Y - 1; nextY >= 0 && grid[pos.X][nextY] == plant {
		if _, exists := visited[Vector{X: pos.X, Y: nextY}]; !exists {
			nr := exploreRegion2(grid, Vector{X: pos.X, Y: nextY}, visited, edges)
			region.Area += nr.Area
		}
	} else {
		edges.left[pos.Y] = append(edges.left[pos.Y], pos.X)
	}

	// right
	if nextY := pos.Y + 1; nextY < len(grid[0]) && grid[pos.X][nextY] == plant {
		if _, exists := visited[Vector{X: pos.X, Y: nextY}]; !exists {
			nr := exploreRegion2(grid, Vector{X: pos.X, Y: nextY}, visited, edges)
			region.Area += nr.Area
		}
	} else {
		edges.right[pos.Y] = append(edges.right[pos.Y], pos.X)
	}
	return region
}

func countEdges(e *Edges) int {
	sum := 0

	for _, v := range e.up {
		slices.Sort(v)
		sum += countSequential(v)
	}
	for _, v := range e.down {
		slices.Sort(v)
		sum += countSequential(v)
	}
	for _, v := range e.left {
		slices.Sort(v)
		sum += countSequential(v)
	}
	for _, v := range e.right {
		slices.Sort(v)
		sum += countSequential(v)
	}
	return sum
}

func countSequential(a []int) int {
	sum := 1
	if len(a) == 1 {
		return sum
	}
	for i := 1; i < len(a); i++ {
		if a[i]-a[i-1] != 1 {
			sum++
		}
	}
	return sum
}

type Vector struct {
	X, Y int
}

type Empty = struct{}

type Region struct {
	Edges int
	Area  int
}

type Edges struct {
	up, down, left, right map[int][]int
}
type Region2 struct {
	Area int
}
