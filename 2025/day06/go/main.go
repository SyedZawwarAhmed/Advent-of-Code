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
	return string(data)
}

func part1(input string) int {
	var numbers [][]int
	var operators []string
	result := 0
	temp := ""
	row := []int{}

	for _, character := range input {
		switch character {
		case '\n':
			if temp != "" {
				number, err := strconv.Atoi(temp)
				if err != nil {
					operators = append(operators, temp)
				}
				row = append(row, number)
				temp = ""
			}
			numbers = append(numbers, row)
			row = []int{}
		case ' ':
			if temp != "" {
				number, err := strconv.Atoi(temp)
				if err != nil {
					operators = append(operators, temp)
				}
				row = append(row, number)
				temp = ""
			}
		default:
			temp += string(character)
		}
	}
	operators = append(operators, temp)

	for i, operator := range operators {
		var solution int
		switch operator {
		case "+":
			solution = 0
		case "*":
			solution = 1
		}
		for j := range len(numbers) {
			switch operator {
			case "+":
				solution += numbers[j][i]
			case "*":
				solution *= numbers[j][i]
			}
		}
		result += solution
	}

	return result
}

func isOperator(character rune) bool {
	return character == '+' || character == '*'
}

func part2(input string) int {
	result := 0
	arr := strings.Split(strings.TrimRight(input, "\n"), "\n")

	column := []string{}
	for i := (len(arr[0]) - 1); i >= 0; i-- {
		temp := ""
		for j := range arr {
			character := arr[j][i]
			if isOperator(rune(character)) {
				column = append(column, temp)
				temp = ""
				sum := 0
				product := 1
				for _, numberStr := range column {
					number, err := strconv.Atoi(numberStr)
					if err != nil {
						panic(err)
					}
					if character == '+' {
						sum += number
					}
					if character == '*' {
						product *= number
					}
				}
				if character == '+' {
					result += sum
				}
				if character == '*' {
					result += product
				}
				column = []string{}
			} else if character != ' ' {
				temp += string(character)
			}
		}
		if temp != "" {
			column = append(column, temp)
		}
	}
	return result
}

func main() {
	input := readInput("../input.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
