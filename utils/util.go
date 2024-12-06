package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func scanData(filePath string) (*bufio.Scanner, *os.File) {
	data, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(data), data
}
func ParseFileToStringArray(filePath string) []string {

	var out []string
	scanner, file := scanData(filePath)
	defer file.Close()

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return out
}

func ParseFileToString(filePath string) string {

	var out string
	scanner, file := scanData(filePath)
	defer file.Close()

	for scanner.Scan() {
		out += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return out
}

func ConvertToInteger(s string) int {
	myInt, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return myInt
}

func GetSignInteger(n int) int8 {
	switch {
	case n > 0:
		return 1
	case n < 0:
		return -1
	default:
		return 0
	}
}
