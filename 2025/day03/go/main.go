package main

import (
	"fmt"
	"math"
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
	arr := strings.Split(input, "\n")
	for i := range arr {
		bankStr := strings.Split(arr[i], "")
		firstBattery := 0
		secondBattery := 0
		firstBatteryIndex := 0
		for k := range len(bankStr) - 1 {
			battery, _ := strconv.Atoi(bankStr[k])
			if battery > firstBattery {
				firstBattery = battery
				firstBatteryIndex = k
			}
		}
		for k := firstBatteryIndex + 1; k < len(bankStr); k++ {
			battery, _ := strconv.Atoi(bankStr[k])
			if battery > secondBattery {
				secondBattery = battery
			}
		}
		result += firstBattery*10 + secondBattery
	}
	return result
}

func part2(input string) int {
	result := 0
	arr := strings.Split(input, "\n")
	for i := range arr {
		bankStr := strings.Split(arr[i], "")
		batteries := []int{}
		previousBatteryIndex := -1
		for j := range 12 {
			tempBattery := 0
			for k := previousBatteryIndex + 1; k < len(bankStr)+j-11; k++ {
				battery, _ := strconv.Atoi(bankStr[k])
				if battery > tempBattery {
					tempBattery = battery
					previousBatteryIndex = k
				}
			}
			batteries = append(batteries, tempBattery)
		}
		for index, battery := range batteries {
			power := len(batteries) - index - 1
			result += battery * int(math.Pow10(power))
		}
	}
	return result
}

func main() {
	input := readInput("../input.txt")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
