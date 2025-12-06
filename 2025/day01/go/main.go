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
	arr := strings.Split(input, "\n")
	dialAt := 50
	password := 0

	for i := range arr {
		var rotationStr string
		var left, right bool
		rotationStr, left = strings.CutPrefix(arr[i], "L")
		if !left {
			rotationStr, right = strings.CutPrefix(arr[i], "R")
		}

		rotation, err := strconv.Atoi(rotationStr)
		if err == nil {
			rotation %= 100
			if left {
				dialAt -= rotation
				if dialAt < 0 {
					dialAt = 100 + dialAt
				}
			}
			if right {
				dialAt += rotation
				if dialAt > 99 {
					dialAt = dialAt - 100
				}
			}
		}

		if dialAt == 0 {
			password++
		}
	}

	return password
}

func part2(input string) int {
	arr := strings.Split(input, "\n")
	dialAt := 50
	password := 0

	for i := range arr {
		var rotationStr string
		var left, right bool
		rotationStr, left = strings.CutPrefix(arr[i], "L")
		if !left {
			rotationStr, right = strings.CutPrefix(arr[i], "R")
			left = !right
		}
		rotation, err := strconv.Atoi(rotationStr)

		if err == nil {
			password += rotation / 100
			rotation %= 100
			if rotation > 0 {
				if left {
					dialAt -= rotation
					if dialAt == 0 {
						password++
					}
					if dialAt < 0 {
						if (rotation + dialAt) != 0 {
							password++
						}
						dialAt = 100 + dialAt
					}
				}
				if right {
					dialAt += rotation
					if dialAt > 99 {
						password++
						dialAt = dialAt - 100
					}
				}
			}
		}
	}
	return password
}

func main() {
	input := readInput("../input.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
