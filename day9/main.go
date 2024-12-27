package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 6382875730645 in 2.8964ms
Part 2: 6420913943576 in 180.3321ms
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

	var diskmap []int
	if scanner.Scan() {
		line := scanner.Text()
		blockIdx := 0
		isFileBlock := true

		for _, r := range line {
			var blocks []int
			num := int(r - '0')
			if isFileBlock {
				blocks = slices.Repeat([]int{blockIdx}, num)
				blockIdx = blockIdx + 1
			} else {
				blocks = slices.Repeat([]int{-1}, num)
			}
			diskmap = append(diskmap, blocks...)
			isFileBlock = !isFileBlock
		}
	}

	start := 0
	end := len(diskmap) - 1
	for {
		for diskmap[start] != -1 {
			start++
		}
		if start >= end {
			break
		}
		diskmap[start] = diskmap[end]
		diskmap[end] = -1
		end--
	}

	for i, num := range diskmap {
		if num == -1 {
			break
		}
		total = total + i*num
	}

	return total
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	total := 0

	var diskmap []Block
	if scanner.Scan() {
		line := scanner.Text()
		blockIdx := 0
		isFileBlock := true

		for _, r := range line {
			num := int(r - '0')
			var block Block
			if isFileBlock {
				block = Block{Id: blockIdx, Size: num, IsFree: !isFileBlock}
				blockIdx = blockIdx + 1
			} else {
				block = Block{Id: -1, Size: num, IsFree: !isFileBlock}
			}
			diskmap = append(diskmap, block)
			isFileBlock = !isFileBlock
		}
	}

	for b := len(diskmap) - 1; b >= 0; b-- {
		if diskmap[b].IsFree {
			continue
		}
		for i := 0; i < b; i++ {
			if diskmap[i].IsFree && diskmap[i].Size >= diskmap[b].Size {
				// create new block
				newBlock := diskmap[b]
				// mark end block as free
				diskmap[b].IsFree = true
				diskmap[b].Id = -1
				// shrink current
				diskmap[i].Size = diskmap[i].Size - diskmap[b].Size
				// insert
				diskmap = slices.Insert(diskmap, i, newBlock)
				break
			}
		}
	}

	offset := 0
	for _, b := range diskmap {
		if !b.IsFree {
			for j := 0; j < b.Size; j++ {
				total = total + b.Id*(j+offset)
			}
		}
		offset = offset + b.Size
	}

	return total
}

// used to print the diskmap
func BlockToString(diskmap []Block) string {
	var sb strings.Builder
	for _, b := range diskmap {
		for j := 0; j < b.Size; j++ {
			if b.IsFree {
				sb.WriteString(".")
			} else {
				sb.WriteString(fmt.Sprintf("%v", b.Id))
			}
		}
	}
	return sb.String()
}

type Block struct {
	Id     int
	Size   int
	IsFree bool
}
