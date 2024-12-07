package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jacob187/advent-of-code-2024/utils"
)

const mulRegex = `mul\(\d*,\d*\)`

func partOne(filePath string) int {
	puzzleMemory := utils.ParseFileToString(filePath)
	re := regexp.MustCompile(mulRegex)

	mulInstructions := re.FindAllString(puzzleMemory, -1)
	return sumMulInstructions(mulInstructions)
}

func partTwo(filePath string) int {
	puzzleMemory := utils.ParseFileToString(filePath)

	invalidToValidMulInstructionsRegex := regexp.MustCompile(`don't\(\).*?do\(\)`)
	removeInvalidMulInstructions := invalidToValidMulInstructionsRegex.ReplaceAllString(puzzleMemory, "")

	invalidCommandsRegex := regexp.MustCompile(`don't\(\).*`)
	removeRemainingInvalidMulValues := invalidCommandsRegex.ReplaceAllString(removeInvalidMulInstructions, "")

	re := regexp.MustCompile(mulRegex)

	mulInstructions := re.FindAllString(removeRemainingInvalidMulValues, -1)

	return sumMulInstructions(mulInstructions)
}

func sumMulInstructions(mulInstructions []string) int {
	var total int = 0
	for _, i := range mulInstructions {
		sum := sumMulValues(i)
		total += sum

	}

	return total
}

func sumMulValues(mulInstructions string) int {
	prefixCut := strings.ReplaceAll(mulInstructions, "mul", "")
	mulNumbers := strings.Split(strings.Trim(prefixCut, "()"), ",")
	return utils.ConvertToInteger(mulNumbers[0]) * utils.ConvertToInteger(mulNumbers[1])
}

func main() {
	fmt.Println(partOne("input.txt"))
	fmt.Println(partTwo("input.txt"))
}
