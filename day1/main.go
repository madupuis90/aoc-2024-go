package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 3714264 in 207.8µs
Part 2: 18805872 in 773.2µs
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
	loc1 := []int{}
	loc2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])
		loc1 = append(loc1, num1)
		loc2 = append(loc2, num2)
	}

	slices.Sort(loc1)
	slices.Sort(loc2)

	for i := 0; i < len(loc1); i++ {
		diff := loc1[i] - loc2[i]
		if diff < 0 { // math.Abs only works with floats in Go?
			diff = -diff
		}
		total = total + diff
	}
	return total
}

func part2(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)
	cache := map[int]int{}
	loc1 := []int{}
	loc2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])
		loc1 = append(loc1, num1)
		loc2 = append(loc2, num2)
	}

	count := 0
	for i := 0; i < len(loc1); i++ {
		value, exist := cache[loc1[i]]
		if exist {
			total = total + loc1[i]*value
			continue
		}
		for j := 0; j < len(loc2); j++ {
			if loc2[j] == loc1[i] {
				count++
			}
		}
		cache[loc1[i]] = count
		total = total + loc1[i]*count
		count = 0
	}
	return total
}
