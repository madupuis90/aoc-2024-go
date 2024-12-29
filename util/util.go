package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// DISCLAIMER: You should not do this in a normal program; You should close the file when done scanning
func CreateScannerFromFile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Expected %v to be in the current folder", filename)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}

func StringSliceAtoi(s []string) []int {
	var result = make([]int, len(s))
	for i, str := range s {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't convert []string to []int")
		}
		result[i] = num
	}
	return result
}

func RuneSliceAtoi(s []rune) []int {
	var result = make([]int, len(s))
	for i, r := range s {
		result[i] = int(r - '0')
	}
	return result
}

func PrintGrid[T any](grid [][]T) {
	for _, r := range grid {
		fmt.Println(r)
		// for _, c := range r {
		// 	fmt.Printf("%v ", c)
		// }
		// fmt.Printf("\n")
	}
}
