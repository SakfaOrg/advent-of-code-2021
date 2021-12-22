package advent_2015_day17

import (
	advent "advent2021/utils"
	"fmt"
	"math"
)

func collectCombinations(total int, containers []int) [][]int {
	return collectCombinationsInner(total, []int{}, containers)
}

func collectCombinationsInner(total int, usedSoFar []int, containers []int) [][]int {
	var result [][]int
	for idx, container := range containers {
		if total >= container {
			used := make([]int, len(usedSoFar))
			copy(used, usedSoFar)
			used = append(used, container)

			left := total - container
			if left == 0 {
				result = append(result, used)
			} else {
				result = append(result, collectCombinationsInner(total-container, used, containers[idx+1:])...)
			}
		}
	}

	return result
}

func filterSmallest(combinations [][]int) [][]int {
	shortestLength := math.MaxInt
	for i := 0; i < len(combinations); i++ {
		if len(combinations[i]) < shortestLength {
			shortestLength = len(combinations[i])
		}
	}

	var result [][]int
	for i := 0; i < len(combinations); i++ {
		if len(combinations[i]) == shortestLength {
			result = append(result, combinations[i])
		}
	}

	return result
}

func countCombinations(total int, containers []int) int {
	result := 0
	for idx, container := range containers {
		if container == total {
			result += 1
		} else if total > container {
			result += countCombinations(total-container, containers[idx+1:])
		}
	}

	return result
}

func parseContainers(lines []string) (containers []int) {
	containers = make([]int, len(lines))
	for idx, line := range lines {
		containers[idx] = advent.MustAtoi(line)
	}
	return
}

func Part1(lines []string) string {
	containers := parseContainers(lines)

	return fmt.Sprintf("Number of ways to fit 150 of eggnog: %d", countCombinations(150, containers))
}

func Part2(lines []string) string {
	containers := parseContainers(lines)

	smallest := filterSmallest(collectCombinations(150, containers))
	return fmt.Sprintf("Smallest number of containers is %d, there are %d ways to arrange them.",
		len(smallest[0]), len(smallest),
	)
}
