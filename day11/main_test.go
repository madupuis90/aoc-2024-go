package main

import (
	"testing"
)

func TestSamplePart1(t *testing.T) {
	result := part1("sample.txt")
	want := 55312

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}

func TestPow10(t *testing.T) {
	if result, want := pow10(0), 1; result != want {
		t.Fatalf(`Test 0: Wanted %v, but got %v`, want, result)
	}
	if result, want := pow10(1), 10; result != want {
		t.Fatalf(`Test 0: Wanted %v, but got %v`, want, result)
	}
	if result, want := pow10(2), 100; result != want {
		t.Fatalf(`Test 0: Wanted %v, but got %v`, want, result)
	}
	if result, want := pow10(5), 100000; result != want {
		t.Fatalf(`Test 0: Wanted %v, but got %v`, want, result)
	}
}
func TestCountDigits(t *testing.T) {
	if result, want := countDigits(0), 1; result != want {
		t.Fatalf(`Test 0: Wanted %v, but got %v`, want, result)
	}
	if result, want := countDigits(1), 1; result != want {
		t.Fatalf(`Test 1: Wanted %v, but got %v`, want, result)
	}
	if result, want := countDigits(22), 2; result != want {
		t.Fatalf(`Test 2: Wanted %v, but got %v`, want, result)
	}
	if result, want := countDigits(333), 3; result != want {
		t.Fatalf(`Test 3: Wanted %v, but got %v`, want, result)
	}
	if result, want := countDigits(1234567), 7; result != want {
		t.Fatalf(`Test 4: Wanted %v, but got %v`, want, result)
	}
}

func TestSamplePart2(t *testing.T) {
	result := part2("sample.txt")
	want := 0

	if result != want {
		t.Fatalf(`Wanted %v, but got %v`, want, result)
	}
}
