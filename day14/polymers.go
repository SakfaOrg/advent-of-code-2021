package day14

import (
	"fmt"
	"math"
	"strings"
)

func parseInput(lines []string) (polymer string, insertions map[string]string) {
	polymer = lines[0]
	insertions = make(map[string]string)
	for i := 2; i < len(lines); i++ {
		split := strings.Split(lines[i], " -> ")
		insertions[split[0]] = split[1]
	}
	return
}

func countPairs(polymer string) map[string]int {
	result := make(map[string]int)
	for i := 0; i < len(polymer) - 1; i++ {
		result[polymer[i:i+2]] += 1
	}
	return result
}

func expand(pairCounts map[string]int, insertions map[string]string) map[string]int {
	result := make(map[string]int)
	for pair, count := range pairCounts {
		if insertion, ok := insertions[pair]; ok {
			result[pair[0:1] + insertion] += count
			result[insertion + pair[1:2]] += count
		} else {
			result[pair] += count
		}
	}
	return result
}

func expandNTimes(times int, polymer map[string]int, insertions map[string]string) map[string]int {
	current := polymer
	for i := 0; i < times; i++ {
		current = expand(current, insertions)
	}
	return current
}

func countLetters(polymer map[string]int, lastLetter string) map[string]int {
	result := make(map[string]int)
	for pair, count := range polymer {
		result[pair[0:1]] += count
	}
	result[lastLetter] += 1
	return result
}

func findMostAndLeastCommon(counts map[string]int) (mostCommon, leastCommon string, mostCommonCount, leastCommonCount int) {
	leastCommonCount = math.MaxInt
	for letter, frequency := range counts {
		if frequency > mostCommonCount {
			mostCommonCount = frequency
			mostCommon = letter
		}
		if frequency < leastCommonCount {
			leastCommonCount = frequency
			leastCommon = letter
		}
	}
	return
}

func solve(lines[] string, steps int) string {
	polymer, insertions := parseInput(lines)
	expanded := expandNTimes(steps, countPairs(polymer), insertions)
	counts := countLetters(expanded, polymer[len(polymer)-1:])
	mostCommon, leastCommon, mostCommonCount, leastCommonCount := findMostAndLeastCommon(counts)
	return fmt.Sprintf("most frequent %s (%d) minus least frequent %s (%d) = %d", mostCommon, mostCommonCount,
		leastCommon, leastCommonCount, mostCommonCount - leastCommonCount)
}

func Part1(lines []string) string {
	return solve(lines, 10)
}

func Part2(lines []string) string {
	return solve(lines, 40)
}