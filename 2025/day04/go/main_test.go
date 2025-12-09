package main

import "testing"

const testInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestPart1(t *testing.T) {
	result := part1(testInput)
	expected := 13
	if result != expected {
		t.Errorf("Part 1: expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	expected := 43
	if result != expected {
		t.Errorf("Part 2: expected %d, got %d", expected, result)
	}
}
