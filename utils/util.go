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
