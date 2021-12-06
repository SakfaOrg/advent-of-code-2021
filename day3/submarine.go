package day3

import (
	"fmt"
	"strconv"
)

/*
 *	https://adventofcode.com/2021/day/3
 */

func Part1(lines []string) string {
	numberOfBits := len(lines[0])
	idxToOnes := countOnes(lines)

	power := 1
	gamma := 0
	epsilon := 0
	for idx := numberOfBits - 1; idx >= 0; idx-- {
		numberOfOnes := idxToOnes[idx]
		if numberOfOnes > len(lines) / 2 {
			gamma += power
		} else {
			epsilon += power
		}
		power *= 2
	}

	return fmt.Sprintf("gamma*epsilon=result: %d*%d=%d", gamma, epsilon, gamma * epsilon)
}

func countOnes(lines []string) map[int]int {
	result := make(map[int]int)
	for _, line := range lines {
		for bitIdx, bit := range line {
			if bit == '1' {
				result[bitIdx] += 1
			}
		}
	}
	return result
}

func filterLines(lines []string, bitPosition int, bitValue uint8) []string {
	var result []string
	for _, line := range lines {
		if line[bitPosition] == bitValue {
			result = append(result, line)
		}
	}
	return result
}

func getResult(name string, lines []string) int {
	if len(lines) != 1 {
		panic(fmt.Sprintf("Expected exactly 1 line left for %s but got %d", name, len(lines)))
	}

	line := lines[0]
	result, err := strconv.ParseInt(line, 2, 0)
	if err != nil {
		panic(fmt.Sprintf("Final line left '%s' is not valid binary: %s", line, err.Error()))
	}
	return int(result)
}

func processBit(lines []string, bitPosition int, oneSelector selector) []string {
	if len(lines) < 2 {
		return lines
	}

	numberOfOnes := countOnes(lines)[bitPosition]
	numberOfLines := len(lines)
	filter := oneSelector(numberOfOnes, numberOfLines)
	linesLeft := filterLines(lines, bitPosition, filter)
	return linesLeft
}

type selector = func(numberOfOnes, numberOfLines int) uint8;
var MostCommonSelector = func(numberOfOnes, numberOfLines int) uint8 {
	if numberOfLines % 2 == 0 && numberOfOnes == numberOfLines / 2 {
		return '1'
	} else if numberOfOnes > numberOfLines / 2 {
		return '1'
	} else {
		return '0'
	}
}
var LeastCommonSelector = func(numberOfOnes, numberOfLines int) uint8 {
	if numberOfLines % 2 == 0 && numberOfOnes == numberOfLines / 2 {
		return '0'
	} else if numberOfOnes <= numberOfLines / 2 {
		return '1'
	} else {
		return '0'
	}
}

func Part2(lines []string) string {
	numberOfBits := len(lines[0])
	oxygenLines := lines
	co2Lines := lines

	for idx := 0; idx < numberOfBits; idx++ {
		oxygenLines = processBit(oxygenLines, idx, MostCommonSelector)
		co2Lines = processBit(co2Lines, idx, LeastCommonSelector)
	}

	oxygen := getResult("oxygen", oxygenLines)
	co2 := getResult("co2", co2Lines)

	return fmt.Sprintf("oxygen*co2=result: %d*%d=%d", oxygen, co2, oxygen * co2)
}