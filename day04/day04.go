package main

import (
	"fmt"

	"github.com/jacob187/advent-of-code-2024/utils"
)

func partOne(filePath string) int {
	wordSearch := utils.ParseFileToStringArray(filePath)

	lineLength, wordSearchLength, count := getWordMapInformationAndCount(wordSearch)

	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < lineLength; j++ {
			// bounds := j > 2 && j < lineLength-3
			// Horizontal: Left -> Right
			if j < lineLength-3 && string(wordSearch[i][j]) == "X" && string(wordSearch[i][j+1]) == "M" && string(wordSearch[i][j+2]) == "A" && string(wordSearch[i][j+3]) == "S" {
				count++
			}

			// Horizontal: Left <- Right
			if j > 2 && string(wordSearch[i][j]) == "X" && string(wordSearch[i][j-1]) == "M" && string(wordSearch[i][j-2]) == "A" && string(wordSearch[i][j-3]) == "S" {
				count++
			}

			// Vertical: Up -> Down
			if i < wordSearchLength-3 && string(wordSearch[i][j]) == "X" && string(wordSearch[i+1][j]) == "M" && string(wordSearch[i+2][j]) == "A" && string(wordSearch[i+3][j]) == "S" {
				count++
			}

			// Vertical: Up <- Down
			if i > 2 && string(wordSearch[i][j]) == "X" && string(wordSearch[i-1][j]) == "M" && string(wordSearch[i-2][j]) == "A" && string(wordSearch[i-3][j]) == "S" {
				count++
			}

			// ...S......
			// ..A.......
			// .M........
			// X.........
			if j < lineLength-3 && i > 2 && string(wordSearch[i][j]) == "X" && string(wordSearch[i-1][j+1]) == "M" && string(wordSearch[i-2][j+2]) == "A" && string(wordSearch[i-3][j+3]) == "S" {
				count++
			}

			// ...X......
			// ..M.......
			// .A........
			// S.........
			if j < lineLength-3 && i > 2 && string(wordSearch[i][j]) == "S" && string(wordSearch[i-1][j+1]) == "A" && string(wordSearch[i-2][j+2]) == "M" && string(wordSearch[i-3][j+3]) == "X" {
				count++
			}

			// X.........
			// .M........
			// ..A.......
			// ...S......
			if j < lineLength-3 && i < wordSearchLength-3 && string(wordSearch[i][j]) == "X" && string(wordSearch[i+1][j+1]) == "M" && string(wordSearch[i+2][j+2]) == "A" && string(wordSearch[i+3][j+3]) == "S" {
				count++
			}

			// S.........
			// .A........
			// ..M.......
			// ...X......

			if j < lineLength-3 && i < wordSearchLength-3 && string(wordSearch[i][j]) == "S" && string(wordSearch[i+1][j+1]) == "A" && string(wordSearch[i+2][j+2]) == "M" && string(wordSearch[i+3][j+3]) == "X" {
				count++
			}

		}

	}
	return count
}

func partTwo(filePath string) int {

	wordSearch := utils.ParseFileToStringArray(filePath)
	lineLength, wordSearchLength, count := getWordMapInformationAndCount(wordSearch)

	// M.S
	// .A.
	// M.S
	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < lineLength; j++ {
			if i < wordSearchLength-2 && j < lineLength-2 && string(wordSearch[i][j]) == "M" && string(wordSearch[i+1][j+1]) == "A" && string(wordSearch[i+2][j+2]) == "S" && string(wordSearch[i][j+2]) == "S" && string(wordSearch[i+2][j]) == "M" {
				count++
			}
			if i < wordSearchLength-2 && j < lineLength-2 && string(wordSearch[i][j]) == "M" && string(wordSearch[i+1][j+1]) == "A" && string(wordSearch[i+2][j+2]) == "S" && string(wordSearch[i][j+2]) == "M" && string(wordSearch[i+2][j]) == "S" {
				count++
			}

			if i < wordSearchLength-2 && j < lineLength-2 && string(wordSearch[i][j]) == "S" && string(wordSearch[i+1][j+1]) == "A" && string(wordSearch[i+2][j+2]) == "M" && string(wordSearch[i][j+2]) == "S" && string(wordSearch[i+2][j]) == "M" {
				count++
			}

			if i < wordSearchLength-2 && j < lineLength-2 && string(wordSearch[i][j]) == "S" && string(wordSearch[i+1][j+1]) == "A" && string(wordSearch[i+2][j+2]) == "M" && string(wordSearch[i][j+2]) == "M" && string(wordSearch[i+2][j]) == "S" {
				count++
			}
		}

	}
	return count
}

func getWordMapInformationAndCount(wordSearch []string) (lineLength int, wordSearchLength int, count int) {
	return len(wordSearch[0]), len(wordSearch), 0
}

func main() {
	fmt.Println(partOne("input.txt"))
	fmt.Println(partTwo("input.txt"))
}
