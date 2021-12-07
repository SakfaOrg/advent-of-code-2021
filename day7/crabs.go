package day7

import (
	advent "advent2021/utils"
	"fmt"
	"math"
	"strings"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type FuelCalculator func(int, int) int

func solve(input string, fuelCalculator FuelCalculator) (position int, fuel int) {
	var maxPosition int
	var positionsInts = []int{}
	for _, p := range strings.Split(input, ",") {
		position := advent.MustAtoi(p)
		positionsInts = append(positionsInts, position)
		if position > maxPosition {
			maxPosition = position
		}
	}

	fuel = math.MaxInt
	for i := 0; i <= maxPosition; i++ {
		fuelCost := 0
		for _, p := range positionsInts {
			fuelCost += fuelCalculator(p, i)
		}

		if fuelCost < fuel {
			fuel = fuelCost
			position = i
		}
	}
	return
}

func Part1(lines []string) string {
	position, fuel := solve(lines[0], func(from, to int) int {
		return abs(from - to)
	})

	return fmt.Sprintf("position=%d (fuel=%d)", position, fuel)
}

func Part2(lines []string) string {
	position, fuel := solve(lines[0], func(from, to int) int {
		distance := abs(from - to)
		return ((1 + distance) * distance) / 2
	})

	return fmt.Sprintf("position=%d (fuel=%d)", position, fuel)
}