package day1

import (
	"advent2021/utils"
	"fmt"
	"math"
)

/*
 *	https://adventofcode.com/2021/day/1
 */

func Part1(lines []string) string {
	prev := math.MaxInt
	var counter int
	for _, line := range lines {
		cur := advent.MustAtoi(line)
		if cur > prev {
			counter++
		}
		prev = cur
	}

	return fmt.Sprintf("depth increases: %d", counter)
}

func windowSum(window [3]int) int {
	return window[0] + window[1] + window[2]
}

func Part2(lines []string) string {
	prevWindow := [3]int {
		advent.MustAtoi(lines[0]),
		advent.MustAtoi(lines[1]),
		advent.MustAtoi(lines[2]),
	}

	var counter int
	for current := 3; current < len(lines); current++ {
		line := lines[current]
		cur := advent.MustAtoi(line)

		newWindow := [3]int{
			prevWindow[1],
			prevWindow[2],
			cur,
		}

		if windowSum(newWindow) > windowSum(prevWindow) {
			counter++
		}

		prevWindow = newWindow
	}

	return fmt.Sprintf("depth increases in moving window width=3: %d", counter)
}

