package day2

import (
	"advent2021/utils"
	"fmt"
	"strings"
)

/*
 *	https://adventofcode.com/2021/day/2
 */

func Part1(lines []string) string {
	horizontal := 0
	depth := 0

	for _, line := range lines {
		words := strings.Fields(line)
		command := words[0]
		arg := advent.MustAtoi(words[1])

		switch (command) {
		case "forward":
			horizontal += arg
		case "down":
			depth += arg
		case "up":
			depth -= arg
		default:
			panic("Unknown command " + command)
		}
	}

	return fmt.Sprintf("horizontal*depth=result: %d*%d=%d", horizontal, depth, horizontal * depth)
}

func Part2(lines []string) string {
	aim := 0
	horizontal := 0
	depth := 0

	for _, line := range lines {
		words := strings.Fields(line)
		command := words[0]
		arg := advent.MustAtoi(words[1])

		switch (command) {
		case "forward":
			horizontal += arg
			depth += aim * arg
		case "down":
			aim += arg
		case "up":
			aim -= arg
		default:
			panic("Unknown command " + command)
		}
	}

	return fmt.Sprintf("horizontal*depth=result: %d*%d=%d", horizontal, depth, horizontal * depth)
}