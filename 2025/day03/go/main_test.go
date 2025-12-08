package main

import "testing"

const testInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPart1(t *testing.T) {
	result := part1(testInput)
	expected := 357
	if result != expected {
		t.Errorf("Part 1: expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	expected := 3121910778619
	if result != expected {
		t.Errorf("Part 2: expected %d, got %d", expected, result)
	}
}
