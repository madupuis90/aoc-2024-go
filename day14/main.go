package main

import (
	"fmt"
	"slices"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 209409792 in 839.8µs
Part 2: 8006 in 727.7722ms
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
	bathroom := &Bathroom{}
	if f == "sample.txt" {
		bathroom.lenX = 11
		bathroom.lenY = 7
	} else {
		bathroom.lenX = 101
		bathroom.lenY = 103
	}

	for scanner.Scan() {
		robot := &Robot{}
		fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &robot.position.X, &robot.position.Y, &robot.velocity.X, &robot.velocity.Y)
		robot.move(100, bathroom)
		countQuadrant(robot, bathroom)
	}

	return bathroom.safetyFactor()
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	bathroom := &Bathroom{}
	var robots []*Robot

	if f == "sample.txt" {
		bathroom.lenX = 11
		bathroom.lenY = 7
	} else {
		bathroom.lenX = 101
		bathroom.lenY = 103
	}
	for scanner.Scan() {
		robot := &Robot{}
		fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &robot.position.X, &robot.position.Y, &robot.velocity.X, &robot.velocity.Y)
		robots = append(robots, robot)
	}

	max := 0
	idx := 0
	for i := 0; i < 10000; i++ {
		moveAllRobots(robots, 1, bathroom)
		m := maxXconsecutive(robots)
		if m > max {
			max = m
			idx = i
		}
	}

	return idx + 1
}

func countQuadrant(robot *Robot, bathroom *Bathroom) {
	thresholdX := bathroom.lenX / 2
	thresholdY := bathroom.lenY / 2
	switch {
	case robot.position.X == thresholdX || robot.position.Y == thresholdY:
		return
	case robot.position.X < thresholdX && robot.position.Y < thresholdY:
		bathroom.topleft++
	case robot.position.X > thresholdX && robot.position.Y < thresholdY:
		bathroom.topright++
	case robot.position.X < thresholdX && robot.position.Y > thresholdY:
		bathroom.bottomleft++
	case robot.position.X > thresholdX && robot.position.Y > thresholdY:
		bathroom.bottomright++
	}
}

type Bathroom struct {
	lenX, lenY, topleft, topright, bottomleft, bottomright int
}

func (bathroom *Bathroom) safetyFactor() int {
	return bathroom.topleft * bathroom.topright * bathroom.bottomleft * bathroom.bottomright
}

type Robot struct {
	position, velocity Vector
}

func (robot *Robot) move(times int, bathroom *Bathroom) {
	robot.position.X = mod(robot.position.X+times*robot.velocity.X, bathroom.lenX)
	robot.position.Y = mod(robot.position.Y+times*robot.velocity.Y, bathroom.lenY)
}

func moveAllRobots(robots []*Robot, times int, bathroom *Bathroom) [][]rune {

	grid := make([][]rune, bathroom.lenY)
	for i := 0; i < bathroom.lenY; i++ {
		grid[i] = make([]rune, bathroom.lenX)
	}

	for _, r := range robots {
		r.move(times, bathroom)
		grid[r.position.Y][r.position.X] = 'O'
	}

	return grid
}

func maxXconsecutive(robots []*Robot) int {

	m := make(map[int][]int)
	for _, r := range robots {
		m[r.position.Y] = append(m[r.position.Y], r.position.X)
	}

	max := 0
	for _, v := range m {
		slices.Sort(v)
		n := countMaxSequential(v)
		if n > max {
			max = n
		}
	}

	return max
}

func countMaxSequential(a []int) int {
	sum := 1
	if len(a) <= 1 {
		return sum
	}
	maxSum := 1
	for i := 1; i < len(a); i++ {
		if a[i]-a[i-1] == 1 {
			sum++
			maxSum = sum
		} else {
			sum = 1
		}
	}
	return maxSum
}

type Vector struct {
	X, Y int
}

func mod(a, m int) int {
	return ((a % m) + m) % m
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Printf("|")
		for _, v := range row {
			if v == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%c", v)
			}
		}
		fmt.Printf("|")
		fmt.Printf("\n")
	}
}
