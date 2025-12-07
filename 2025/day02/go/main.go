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

func isValidIdPartOne(input string) bool {
	inputLength := len(input)
	if inputLength%2 != 0 {
		return true
	} else {
		firstHalf := input[0 : inputLength/2]
		secondHalf := input[inputLength/2 : inputLength]
		if strings.EqualFold(firstHalf, secondHalf) {
			return false
		}
	}
	return true
}

func isValidIdPartTwo(input string) bool {
	inputLength := len(input)
	if inputLength == 1 {
		return true
	}
	isInvalid := true
	for numberOfSplits := 2; numberOfSplits <= inputLength; numberOfSplits++ {
		if inputLength%numberOfSplits != 0 {
			continue
		}
		sequenceLength := inputLength / numberOfSplits
		sequence := input[0:sequenceLength]

		isInvalid = true
		for i := 0; i < numberOfSplits; i++ {
			subString := input[i*sequenceLength : (i*sequenceLength)+sequenceLength]
			if !strings.EqualFold(sequence, subString) {
				isInvalid = false
				break
			}
		}
		if isInvalid {
			break
		}
	}
	return !isInvalid
}

func part1(input string) int {
	result := 0
	arr := strings.Split(input, ",")
	for i := range arr {
		rangeArr := strings.Split(arr[i], "-")
		rangeStart, errStart := strconv.Atoi(rangeArr[0])
		rangeEnd, errEnd := strconv.Atoi(rangeArr[1])

		if errStart == nil && errEnd == nil {
			for id := rangeStart; id <= rangeEnd; id++ {
				if !isValidIdPartOne(strconv.Itoa(id)) {
					result += id
				}
			}
		}
	}
	return result
}

func part2(input string) int {
	result := 0
	arr := strings.Split(input, ",")
	for i := range arr {
		rangeArr := strings.Split(arr[i], "-")
		rangeStart, errStart := strconv.Atoi(rangeArr[0])
		rangeEnd, errEnd := strconv.Atoi(rangeArr[1])

		if errStart == nil && errEnd == nil {
			for id := rangeStart; id <= rangeEnd; id++ {
				if !isValidIdPartTwo(strconv.Itoa(id)) {
					result += id
				}
			}
		}
	}
	return result
}

func main() {
	input := readInput("../input.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
