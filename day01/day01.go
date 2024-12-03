package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/jacob187/advent-of-code-2024/utils"
)

func partOne(filepath string) int {

	colOne, colTwo := parseLines(filepath)
	sort.Ints(colOne)
	sort.Ints(colTwo)

	var sum int

	for i := 0; i < len(colOne); i++ {
		sum += int(math.Abs(float64(colOne[i]) - float64(colTwo[i])))
	}

	return sum
}

func partTwo(filepath string) int {

	colOne, colTwo := parseLines(filepath)

	var similarityScore int
	for i := 0; i < len(colOne); i++ {
		var numberOfOccurrences int

		for j := 0; j < len(colTwo); j++ {
			if colOne[i] == colTwo[j] {
				numberOfOccurrences += 1
			}
		}
		similarityScoreIncreaseAmount := numberOfOccurrences * colOne[i]
		similarityScore += similarityScoreIncreaseAmount
	}

	return similarityScore

}

func parseLines(filepath string) ([]int, []int) {
	lines := utils.ParseFileToStringArray(filepath)

	var colOne []int
	var colTwo []int
	for _, line := range lines {
		line := strings.Split(line, "   ")
		colOne = append(colOne, utils.ConvertToInteger(line[0]))
		colTwo = append(colTwo, utils.ConvertToInteger(line[1]))
	}

	return colOne, colTwo
}

func main() {
	partOneAnswer := partOne("input.txt")
	fmt.Printf("%d \n", partOneAnswer)
	partTwoAnswer := partTwo("input.txt")
	fmt.Printf("%d", partTwoAnswer)
}
