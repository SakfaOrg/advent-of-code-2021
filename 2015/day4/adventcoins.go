package advent_2015_day4

import (
	"crypto/md5"
	"fmt"
)

func mine(secret string) int {
	for i := 1; ; i++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", secret, i)))
		if hash[0] == 0 && hash[1] == 0 && hash[2] < 16 {
			return i
		}
	}
}

func mine6(secret string) int {
	for i := 1; ; i++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", secret, i)))
		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			return i
		}
	}
}

func Part1(lines []string) string {
	return fmt.Sprintf("Number that produces hash with 5 zeroes: %d", mine(lines[0]))
}

func Part2(lines []string) string {
	return fmt.Sprintf("Number that produces hash with 6 zeroes: %d", mine6(lines[0]))
}
