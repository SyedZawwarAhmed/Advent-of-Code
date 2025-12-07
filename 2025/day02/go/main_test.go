package main

import "testing"

const testInput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`

func TestPart1(t *testing.T) {
	result := part1(testInput)
	expected := 1227775554
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
