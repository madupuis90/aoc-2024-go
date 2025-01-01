package main

import (
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample.txt")
	want := 140

	result2 := part1("sample2.txt")
	want2 := 772

	result3 := part1("sample3.txt")
	want3 := 1930

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}

	if result2 != want2 {
		t.Fatalf(`Wanted %v, but got %v`, want2, result2)
	}

	if result3 != want3 {
		t.Fatalf(`Wanted %v, but got %v`, want3, result3)
	}
}

func TestSamplePart2(t *testing.T) {
	result := part2("sample.txt")
	want := 80

	result2 := part2("sample2.txt")
	want2 := 436

	result3 := part2("sample3.txt")
	want3 := 1206

	result4 := part2("sample4.txt")
	want4 := 236

	result5 := part2("sample5.txt")
	want5 := 368

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}

	if result2 != want2 {
		t.Fatalf(`Wanted %v, but got %v`, want2, result2)
	}

	if result3 != want3 {
		t.Fatalf(`Wanted %v, but got %v`, want3, result3)
	}

	if result4 != want4 {
		t.Fatalf(`Wanted %v, but got %v`, want4, result4)
	}

	if result5 != want5 {
		t.Fatalf(`Wanted %v, but got %v`, want5, result5)
	}
}
