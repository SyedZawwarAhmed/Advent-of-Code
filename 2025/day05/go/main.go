package main

import (
	"fmt"
	"os"
	"sort"
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
	ingredientIds := strings.SplitSeq(splittedInput[1], "\n")

	for ingredientId := range ingredientIds {
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
	result := 0
	splittedInput := strings.Split(input, "\n\n")
	freshIdRanges := strings.Split(splittedInput[0], "\n")

	sort.Slice(freshIdRanges, func(i, j int) bool {
		firstMin, err := strconv.Atoi(strings.Split(freshIdRanges[i], "-")[0])
		secondMin, err := strconv.Atoi(strings.Split(freshIdRanges[j], "-")[0])
		if err != nil {
			panic(err)
		}
		return firstMin < secondMin
	})

	globalMax := 0
	for i := range freshIdRanges {
		current := strings.Split(freshIdRanges[i], "-")
		currentMin, err := strconv.Atoi(current[0])
		currentMax, err := strconv.Atoi(current[1])
		if err != nil {
			panic(err)
		}

		if i > 0 {
			previous := strings.Split(freshIdRanges[i-1], "-")
			previousMin, err := strconv.Atoi(previous[0])
			previousMax, err := strconv.Atoi(previous[1])
			if previousMax > globalMax {
				globalMax = previousMax
			}
			if err != nil {
				panic(err)
			}
			if currentMin == previousMin {
				if currentMax > globalMax {
					result += currentMax - globalMax
				}
			} else {
				if currentMin < globalMax {
					if currentMax > globalMax {
						result += currentMax - globalMax
					}
				} else if currentMin == globalMax {
					result += currentMax - currentMin
				} else {
					result += currentMax - currentMin + 1
				}
			}
		} else {
			result += currentMax - currentMin + 1
			globalMax = currentMax
		}
	}
	return result
}

func main() {
	input := readInput("../input.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
