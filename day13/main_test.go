package main

import (
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample.txt")
	var want int64 = 480

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestSamplePart2(t *testing.T) {
	result := part2("sample.txt")
	var want int64 = 875318608908

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}
