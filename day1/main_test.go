package main

import (
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample.txt")
	want := 0

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestSamplePart2(t *testing.T) {
	result := part2("sample.txt")
	want := 0

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}
