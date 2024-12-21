package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 538191549061 in 17.9063ms
Part 2: 34612812972206 in 1.6353173s
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
	for scanner.Scan() {
		line := scanner.Text()
		colonIdx := strings.IndexRune(line, ':')
		rStr := line[:colonIdx]
		rNums := line[colonIdx+1:]
		r, _ := strconv.Atoi(rStr)
		nums := util.SliceAtoi(strings.Fields(rNums))

		root := buildTree(nums)

		if isValidOperation(r, root) {
			total += r
		}
	}

	return total
}

func part2(f string) int {
	scanner := util.CreateScannerFromFile(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		colonIdx := strings.IndexRune(line, ':')
		rStr := line[:colonIdx]
		rNums := line[colonIdx+1:]
		r, _ := strconv.Atoi(rStr)
		nums := util.SliceAtoi(strings.Fields(rNums))

		root := buildTree2(nums)

		if isValidOperation2(r, root) {
			total += r
		}
	}

	return total
}

func isValidOperation(result int, node *Node) bool {
	if node.Left == nil {
		return node.Value == result
	}

	if isValidOperation(result, node.Left) || isValidOperation(result, node.Right) {
		return true
	}

	return false
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func buildTree(numbers []int) *Node {
	if len(numbers) == 0 {
		return nil
	}
	root := &Node{Value: numbers[0]}

	buildSubTree(root, numbers[1:])

	return root
}

func buildSubTree(node *Node, numbers []int) {
	if len(numbers) == 0 {
		return
	}
	node.Left = &Node{Value: node.Value + numbers[0]}
	node.Right = &Node{Value: node.Value * numbers[0]}

	buildSubTree(node.Left, numbers[1:])
	buildSubTree(node.Right, numbers[1:])
}

func isValidOperation2(result int, node *Node2) bool {
	if node.First == nil {
		return node.Value == result
	}

	if isValidOperation2(result, node.First) || isValidOperation2(result, node.Second) || isValidOperation2(result, node.Third) {
		return true
	}

	return false
}

type Node2 struct {
	Value  int
	First  *Node2
	Second *Node2
	Third  *Node2
}

func buildTree2(numbers []int) *Node2 {
	if len(numbers) == 0 {
		return nil
	}
	root := &Node2{Value: numbers[0]}

	buildSubTree2(root, numbers[1:])

	return root
}

func buildSubTree2(node *Node2, numbers []int) {
	if len(numbers) == 0 {
		return
	}
	v1 := strconv.Itoa(node.Value)
	v2 := strconv.Itoa(numbers[0])
	n, _ := strconv.Atoi(v1 + v2)
	node.First = &Node2{Value: node.Value + numbers[0]}
	node.Second = &Node2{Value: node.Value * numbers[0]}
	node.Third = &Node2{Value: n}

	buildSubTree2(node.First, numbers[1:])
	buildSubTree2(node.Second, numbers[1:])
	buildSubTree2(node.Third, numbers[1:])
}
