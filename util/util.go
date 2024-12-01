package util

import (
	"bufio"
	"log"
	"os"
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