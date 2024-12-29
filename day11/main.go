package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 211306 in 497.1µs
Part 2: 250783680217283 in 23.747ms
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
	var stones = make(map[int]int)

	if scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		for _, f := range fields {
			s, _ := strconv.Atoi(f)
			stones[s]++
		}
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	return countStones(stones)
}

func countStones(stones map[int]int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	return sum
}

func blink(stones map[int]int) map[int]int {
	newStones := make(map[int]int)
	for k, v := range stones {
		if k == 0 { // rule 1
			newStones[1] += v
		} else if digits := countDigits(k); digits%2 == 0 { // rule 2
			m := pow10(digits / 2)
			left := k / m
			right := k - (left * m)
			newStones[left] += v
			newStones[right] += v
		} else { // rule 3
			newStones[k*2024] += v
		}
	}
	return newStones
}

func pow10(power int) int {
	if power == 0 {
		return 1
	}
	result := 10
	for i := 1; i < power; i++ {
		result = result * 10
	}
	return result
}

func countDigits(num int) int {
	count := 1
	if num == 0 {
		return count
	}
	if num < 0 {
		num = -num
	}
	for n := num; n >= 10; n = n / 10 {
		count++
	}
	return count
}
func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	var stones = make(map[int]int)

	if scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		for _, f := range fields {
			s, _ := strconv.Atoi(f)
			stones[s]++
		}
	}

	for i := 0; i < 75; i++ {
		stones = blink(stones)
	}

	return countStones(stones)
}
