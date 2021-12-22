package advent_2015_day1

import "fmt"

func Part1(lines []string) string {
	opens := 0
	closes := 0
	for _, chr := range lines[0] {
		if chr == '(' {
			opens++
		} else if chr == ')' {
			closes++
		}
	}

	return fmt.Sprintf("Floor = %d", opens-closes)
}

func Part2(lines []string) string {
	floor := 0
	for idx, chr := range lines[0] {
		if chr == '(' {
			floor++
		} else if chr == ')' {
			floor--
		}
		if floor < 0 {
			return fmt.Sprintf("Basement entered at %d", idx+1)
		}
	}

	panic("Santa never entered basement!")
}
