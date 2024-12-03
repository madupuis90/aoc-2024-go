package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 171183089 in 371.2µs
Part 2: 63866497 in 797.9µs
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

	for scanner.Scan() {
		line := scanner.Text()
		total = total + scanUncorruptedInstructions(line)
	}

	return total
}

/*
For this one I had to parse the whole file in one chunk instead of line by line; turns out the regex need to span multiple lines
The idea here is to remove everything between don't() and do() and then the problem is the same as part1

Notes:
1) a little gotcha is that the input ends with a don't() with no following do() so we need to also check for end of string "$"
2) we also need to set the flag for Dot to match newline characters "(?s)""
*/
func part2(f string) int {
	total := 0
	data, _ := os.ReadFile(f)
	r := regexp.MustCompile(`(?s)don't\(\).*?(do\(\)|$)`)
	parsed := r.ReplaceAllLiteralString(string(data), "")
	total = total + scanUncorruptedInstructions(parsed)

	return total
}

func scanUncorruptedInstructions(line string) int {
	sum := 0
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(line, -1)
	for _, m := range matches {
		n1, _ := strconv.Atoi(m[1]) // m[0] is the entire match
		n2, _ := strconv.Atoi(m[2])
		sum = sum + n1*n2
	}
	return sum
}
