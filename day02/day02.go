package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jacob187/advent-of-code-2024/utils"
)

func partOne(filepath string) int {
	reports := utils.ParseFileToStringArray(filepath)
	totalValidLines := 0

	for _, line := range reports {
		numbers := strings.Split(line, " ")

		if isValidSequence(numbers) {
			totalValidLines++
		}

	}
	return totalValidLines
}

func partTwo(filepath string) int {
	reports := utils.ParseFileToStringArray(filepath)
	totalValidLines := 0

	for _, line := range reports {
		numbers := strings.Split(line, " ")

		if isValidSequence(numbers) {
			totalValidLines++
			continue
		}

		for i := 0; i < len(numbers); i++ {
			newNumbers := make([]string, 0)
			newNumbers = append(newNumbers, numbers[:i]...)
			newNumbers = append(newNumbers, numbers[i+1:]...)

			if isValidSequence(newNumbers) {
				totalValidLines++
				break
			}
		}
	}
	return totalValidLines
}

func isValidSequence(numbers []string) bool {
	if len(numbers) <= 1 {
		return true
	}

	direction := utils.GetSignInteger(utils.ConvertToInteger(numbers[1]) - utils.ConvertToInteger(numbers[0]))

	for i := 1; i < len(numbers); i++ {
		diff := utils.ConvertToInteger(numbers[i]) - utils.ConvertToInteger(numbers[i-1])
		if math.Abs(float64(diff)) > 3 || diff == 0 ||
			utils.GetSignInteger(diff) != direction {
			return false
		}
	}
	return true
}

func main() {
	partOneAnswer := partOne("input.txt")
	fmt.Println(partOneAnswer)
	partTwoAnswer := partTwo("input.txt")
	fmt.Println(partTwoAnswer)
}
