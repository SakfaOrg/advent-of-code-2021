package main

import (
	advent_2015_day1 "advent2021/2015/day1"
	advent_2015_day17 "advent2021/2015/day17"
	advent_2015_day2 "advent2021/2015/day2"
	advent_2015_day3 "advent2021/2015/day3"
	advent_2015_day5 "advent2021/2015/day5"
	advent_2015_day6 "advent2021/2015/day6"
	advent "advent2021/utils"
	"fmt"
	"time"
)

type PartRunner func([]string) string

func run(day int, part1 PartRunner, part2 PartRunner) (timeTaken time.Duration) {
	fmt.Printf("\n%30s DAY %2d %30s\n\n", "", day, "")
	lines := advent.MustReadLines(fmt.Sprintf("2015/day%d/input", day))
	timeTaken = time.Duration(0)
	timeTaken += advent.Timed("part1", 0, func() string {
		return part1(lines)
	})
	if part2 != nil {
		timeTaken += advent.Timed("part2", 0, func() string {
			return part2(lines)
		})
	}

	return
}

func main() {
	timeTaken := time.Duration(0)
	timeTaken += run(1, advent_2015_day1.Part1, advent_2015_day1.Part2)
	timeTaken += run(2, advent_2015_day2.Part1, advent_2015_day2.Part2)
	timeTaken += run(3, advent_2015_day3.Part1, advent_2015_day3.Part2)
	//timeTaken += run(4, advent_2015_day4.Part1, advent_2015_day4.Part2)
	timeTaken += run(5, advent_2015_day5.Part1, advent_2015_day5.Part2)
	timeTaken += run(6, advent_2015_day6.Part1, advent_2015_day6.Part2)
	timeTaken += run(17, advent_2015_day17.Part1, advent_2015_day17.Part2)

	fmt.Printf("\n%30s Summary %30s\n\n", "", "")
	fmt.Printf("All solutions combined took: %s\n", timeTaken)
}
