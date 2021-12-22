package advent_2015_day5

import (
	"fmt"
	"strings"
)

func isVowel(r uint8) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}

func isNaughtySubstring(letter uint8, next uint8) bool {
	if (letter == 'a' && next == 'b') ||
		(letter == 'c' && next == 'd') ||
		(letter == 'p' && next == 'q') ||
		(letter == 'x' && next == 'y') {
		return true
	}
	return false
}

func isNicePart2(s string) bool {
	ok := false
	for i := 0; i < len(s)-3; i++ {
		pair := s[i : i+2]
		if strings.Contains(s[i+2:], pair) {
			ok = true
			break
		}
	}
	if !ok {
		return false
	}

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}

	return false
}

func isNice(s string) bool {
	vowels := 0
	hadDuplicatedLetter := false
	for i := 0; i < len(s)-1; i++ {
		letter := s[i]
		if isVowel(letter) {
			vowels++
		}
		if s[i] == s[i+1] {
			hadDuplicatedLetter = true
		}
		if isNaughtySubstring(s[i], s[i+1]) {
			return false
		}
	}
	if isVowel(s[len(s)-1]) {
		vowels++
	}

	return vowels >= 3 && hadDuplicatedLetter
}

type NiceChecker func(string) bool

func solve(lines []string, checker NiceChecker) string {
	nice := 0
	for _, line := range lines {
		if checker(line) {
			nice++
		}
	}
	return fmt.Sprintf("Nice strings: %d", nice)
}

func Part1(lines []string) string {
	return solve(lines, isNice)
}

func Part2(lines []string) string {
	return solve(lines, isNicePart2)
}
