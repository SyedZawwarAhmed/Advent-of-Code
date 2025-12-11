package main

import "testing"

const testInput = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +`

func TestPart1(t *testing.T) {
	result := part1(testInput)
	expected := 4277556
	if result != expected {
		t.Errorf("Part 1: expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	expected := 0
	if result != expected {
		t.Errorf("Part 2: expected %d, got %d", expected, result)
	}
}
