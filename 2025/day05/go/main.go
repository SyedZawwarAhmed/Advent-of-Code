package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}

func part1(input string) int {
	result := 0
	splittedInput := strings.Split(input, "\n\n")
	freshIdRanges := strings.Split(splittedInput[0], "\n")
	ingredientIds := strings.Split(splittedInput[1], "\n")

	for _, ingredientId := range ingredientIds {
		for _, freshIdRange := range freshIdRanges {
			freshIdRangeArr := strings.Split(freshIdRange, "-")
			ingredient, err := strconv.Atoi(ingredientId)
			freshMin, err := strconv.Atoi(freshIdRangeArr[0])
			freshMax, err := strconv.Atoi(freshIdRangeArr[1])

			if err != nil {
				panic(err)
			}
			if ingredient >= freshMin && ingredient <= freshMax {
				result++
				break
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
