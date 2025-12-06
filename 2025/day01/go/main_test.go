package main

import "testing"

const testInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestPart1(t *testing.T) {
	result := part1(testInput)
	expected := 3
	if result != expected {
		t.Errorf("Part 1: expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	expected := 6
	if result != expected {
		t.Errorf("Part 2: expected %d, got %d", expected, result)
	}
}
