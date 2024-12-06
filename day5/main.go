package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"example.com/aoc-2024-go/util"
)

/*
Part 1: 5087 in 1.0286ms
Part 2: 4971 in 2.5012ms
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

	dag := make(map[int][]int)
	var updates [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") { // rule line
			s := strings.Split(line, "|")
			pair := util.SliceAtoi(s)
			dag[pair[0]] = append(dag[pair[0]], pair[1])
		} else if strings.Contains(line, ",") { // update line
			s := strings.Split(line, ",")
			pages := util.SliceAtoi(s)
			updates = append(updates, pages)
		}
	}

	for _, update := range updates {
		// fmt.Printf("%v %v\n", update, isValidTopologicalSort(dag, update))
		if isValidTopologicalSort(dag, update) {
			total = total + update[len(update)/2]
		}
	}
	return total
}

func isValidTopologicalSort(dag map[int][]int, update []int) bool {
	position := make(map[int]int)
	for i, node := range update {
		position[node] = i
	}

	for u, neighbors := range dag {
		if pos1, exists := position[u]; exists {
			for _, v := range neighbors {
				// fmt.Printf("u: %-5v neighbors: %v v: %-5v\n", u, neighbors, v)
				if pos2, exists := position[v]; exists {
					if pos1 > pos2 {
						return false
					}
				}
			}
		}

	}
	return true
}

func topologicalSort(dag map[int][]int) ([]int, error) {
	inDegree := make(map[int]int)
	queue := []int{}
	topoSort := []int{}

	// in-degrees
	for node, neighbors := range dag {
		if _, exists := inDegree[node]; !exists {
			inDegree[node] = 0
		}
		for _, n := range neighbors {
			if _, exists := inDegree[n]; !exists {
				inDegree[n] = 0
			}
			inDegree[n]++
		}
	}

	// init the queue with 0 in-degree nodes
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	// go through queue, remove in-degree from neighbors
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		topoSort = append(topoSort, node)
		for _, neighbor := range dag[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// check that we processed all nodes in dag
	if len(topoSort) == len(inDegree) {
		return topoSort, nil
	} else {
		return []int{}, errors.New("DAG is cyclic, can't topo sort")
	}
}

func part2(f string) int {
	total := 0
	scanner := util.CreateScannerFromFile(f)

	dag := make(map[int][]int)
	var updates [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") { // rule line
			s := strings.Split(line, "|")
			pair := util.SliceAtoi(s)
			dag[pair[0]] = append(dag[pair[0]], pair[1])
		} else if strings.Contains(line, ",") { // update line
			s := strings.Split(line, ",")
			pages := util.SliceAtoi(s)
			updates = append(updates, pages)
		}
	}

	// The sample DAG was acyclic, but not the input one
	// Initital idea was to return a possible sorted order and adjust each update base one that
	// Instead I will create a mini-dag for each update...not ideal but spent too much time here
	for _, update := range updates {
		if !isValidTopologicalSort(dag, update) {
			miniDag := createMiniDag(update, dag)
			sortedUpdate, err := topologicalSort(miniDag)
			if err != nil {
				panic(err)
			}
			total = total + sortedUpdate[len(sortedUpdate)/2]
		}
	}
	return total
}

func createMiniDag(update []int, dag map[int][]int) map[int][]int {
	miniDag := make(map[int][]int)
	for _, node := range update {
		newNeighbors := []int{}
		for _, n := range dag[node] {
			if slices.Contains(update, n) {
				newNeighbors = append(newNeighbors, n)
			}
		}
		miniDag[node] = newNeighbors
	}
	return miniDag
}
