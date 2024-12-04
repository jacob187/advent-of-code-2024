package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ParseFileToStringArray(filePath string) []string {
	var out []string

	data, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		out = append(out, scanner.Text())
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
