package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}

type Point struct {
	x int
	y int
}

func getAdjacentRollsPosition(x int, y int, rows int, columns int) []Point {
	minimum := Point{x: x - 1, y: y - 1}
	maximum := Point{x: x + 1, y: y + 1}

	var result []Point

	for i := minimum.x; i <= maximum.x; i++ {
		for j := minimum.y; j <= maximum.y; j++ {
			if i < 0 || i >= columns || j < 0 || j >= rows || (i == x && j == y) {
				continue
			}
			result = append(result, Point{x: i, y: j})
		}
	}

	return result
}

func part1(input string) int {
	result := 0
	arr := strings.Split(input, "\n")
	for i, row := range arr {
		for j, roll := range row {
			if roll == '@' {
				adjacentRollCount := 0
				adjacentRolls := getAdjacentRollsPosition(i, j, len(arr), len(row))
				for _, adjacentRoll := range adjacentRolls {
					if arr[adjacentRoll.x][adjacentRoll.y] == '@' {
						adjacentRollCount++
					}
				}
				if adjacentRollCount < 4 {

					result++
				}
			}
		}
	}
	return result
}

func part2(input string) int {
	// TODO: Implement part 2
	return 0
}

func main() {
	input := readInput("../input.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
