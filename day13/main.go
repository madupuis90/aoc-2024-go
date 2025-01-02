package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 37901 in 187.5µs
Part 2: 77407675412647 in 162.9µs
*/
func main() {
	start := time.Now()
	part1 := part1("input.txt")
	fmt.Printf("Part 1: %v in %s\n", part1, time.Since(start))

	start = time.Now()
	part2 := part2("input.txt")
	fmt.Printf("Part 2: %v in %s\n", part2, time.Since(start))
}

func part1(f string) int64 {
	scanner := util.CreateScannerFromFile(f)
	var total int64 = 0

	var ax, ay, bx, by, X, Y int64

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Button A") {
			xStr := line[strings.IndexRune(line, 'X')+2 : strings.IndexRune(line, ',')]
			yStr := line[strings.IndexRune(line, 'Y')+2:]
			ax, _ = strconv.ParseInt(xStr, 10, 64)
			ay, _ = strconv.ParseInt(yStr, 10, 64)
		}

		if strings.Contains(line, "Button B") {
			xStr := line[strings.IndexRune(line, 'X')+2 : strings.IndexRune(line, ',')]
			yStr := line[strings.IndexRune(line, 'Y')+2:]
			bx, _ = strconv.ParseInt(xStr, 10, 64)
			by, _ = strconv.ParseInt(yStr, 10, 64)
		}

		if strings.Contains(line, "Prize") {
			xStr := line[strings.IndexRune(line, 'X')+2 : strings.IndexRune(line, ',')]
			yStr := line[strings.IndexRune(line, 'Y')+2:]
			X, _ = strconv.ParseInt(xStr, 10, 64)
			Y, _ = strconv.ParseInt(yStr, 10, 64)
		}

		if line == "" {
			total += compute(ax, ay, bx, by, X, Y)
		}
	}
	total += compute(ax, ay, bx, by, X, Y)

	return total
}

func part2(f string) int64 {
	scanner := util.CreateScannerFromFile(f)
	var total int64 = 0

	var ax, ay, bx, by, X, Y int64

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Button A") {
			xStr := line[strings.IndexRune(line, 'X')+2 : strings.IndexRune(line, ',')]
			yStr := line[strings.IndexRune(line, 'Y')+2:]
			ax, _ = strconv.ParseInt(xStr, 10, 64)
			ay, _ = strconv.ParseInt(yStr, 10, 64)
		}

		if strings.Contains(line, "Button B") {
			xStr := line[strings.IndexRune(line, 'X')+2 : strings.IndexRune(line, ',')]
			yStr := line[strings.IndexRune(line, 'Y')+2:]
			bx, _ = strconv.ParseInt(xStr, 10, 64)
			by, _ = strconv.ParseInt(yStr, 10, 64)
		}

		if strings.Contains(line, "Prize") {
			xStr := line[strings.IndexRune(line, 'X')+2 : strings.IndexRune(line, ',')]
			yStr := line[strings.IndexRune(line, 'Y')+2:]
			X, _ = strconv.ParseInt(xStr, 10, 64)
			Y, _ = strconv.ParseInt(yStr, 10, 64)
			X += 10000000000000
			Y += 10000000000000
		}

		if line == "" {
			total += compute(ax, ay, bx, by, X, Y)
		}
	}
	total += compute(ax, ay, bx, by, X, Y)
	return total
}

func compute(ax, ay, bx, by, X, Y int64) int64 {
	var a, b int64

	Bnumerator := X*ay - Y*ax
	Bdenomiator := bx*ay - ax*by
	if Bdenomiator == 0 || Bnumerator%Bdenomiator != 0 {
		return 0
	}
	b = Bnumerator / Bdenomiator
	Anumerator := Y - b*by
	if Anumerator%ay != 0 {
		return 0
	}
	a = Anumerator / ay

	return 3*a + b
}
