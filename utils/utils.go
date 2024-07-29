package aoc_utils

import (
	"bufio"
	"os"
)

func ScanFile() *bufio.Scanner {
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	return scanner
}
