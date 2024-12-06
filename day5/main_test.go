package main

import (
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample.txt")
	want := 143

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestSamplePart2(t *testing.T) {
	result := part2("sample.txt")
	want := 123

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}
