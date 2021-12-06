package advent

import (
	"bufio"
	"os"
	"strconv"
)

func MustAtoi(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return value
}

func MustReadLines(path string) []string {
	lines, err := ReadLines(path)
	if err != nil {
		panic(err)
	}

	return lines
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}