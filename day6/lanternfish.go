package day6

import (
	advent "advent2021/utils"
	"fmt"
	"strings"
)

const days = 80
const cycle = 7
const cooldown = 2

/**
 * builds a lookup table telling "how many fish will a fish born on day i spawn until days"?
 *
 * for instance, if days = 10 a fish born on day 5 will spawn none, but a fish spawned on day 1 will spawn 1 fish (on day 9) so a table will look like thos
 * [1,1,0,0,0,0,0,0,0]
 * but if days = 18 things become interesting as fish born on day 1 will spawn a fish on day 9 and 15 + whatever was spawned by fish born on day 9 which
 * can be answered quickly if we build this array in reverse.
 */
func buildSpawnTable(cycle, cooldown, days int) []int {
	result := make([]int, days+1)
	for day := days - cycle - cooldown; day >= 0; day-- {
		spawned := 0
		firstChild := day + cycle + cooldown
		if firstChild <= days {
			spawned = 1 /* this child */ + result[firstChild] /* and add whatever that child will spawn */
		}

		for nextChild := firstChild + cycle; nextChild <= days; nextChild += cycle {
			spawned += 1 + result[nextChild]
		}

		result[day] = spawned
	}
	return result
}

func fishOnDay(initialTimes []int, day int) int {
	result := len(initialTimes)
	spawnTable := buildSpawnTable(cycle, cooldown, day)

	for _, fishCounter := range initialTimes {
		for nextChild := fishCounter + 1; nextChild <= day; nextChild += cycle {
			result = result + 1 + spawnTable[nextChild]
		}
	}

	return result
}

func parseInput(lines []string) []int {
	initialTimersStrings := strings.Split(lines[0], ",")
	var initialTimes []int
	for _, v := range initialTimersStrings {
		initialTimes = append(initialTimes, advent.MustAtoi(v))
	}
	return initialTimes
}

func Part1(lines []string) string {
	initialTimes := parseInput(lines)

	return fmt.Sprintf("fishes after day 80=%d", fishOnDay(initialTimes, 80))
}

func Part2(lines []string) string {
	initialTimes := parseInput(lines)

	return fmt.Sprintf("fishes after day 256=%d", fishOnDay(initialTimes, 256))
}
