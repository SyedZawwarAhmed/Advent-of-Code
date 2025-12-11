package main

import "testing"

const testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestPart1(t *testing.T) {
	result := part1(testInput)
	expected := 3
	if result != expected {
		t.Errorf("Part 1: expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	expected := 14
	if result != expected {
		t.Errorf("Part 2: expected %d, got %d", expected, result)
	}
}
