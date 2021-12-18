package main

import (
	"advent2021/day1"
	"advent2021/day10"
	"advent2021/day11"
	"advent2021/day12"
	"advent2021/day13"
	"advent2021/day14"
	"advent2021/day15"
	"advent2021/day16"
	"advent2021/day17"
	"advent2021/day18"
	"advent2021/day2"
	"advent2021/day3"
	"advent2021/day4"
	"advent2021/day5"
	"advent2021/day6"
	"advent2021/day7"
	"advent2021/day8"
	"advent2021/day9"
	advent2 "advent2021/utils"
	"fmt"
	"time"
)

type PartRunner func([]string) string

func run(day int, part1 PartRunner, part2 PartRunner) (timeTaken time.Duration) {
	fmt.Printf("\n%30s DAY %2d %30s\n\n", "", day, "")
	lines := advent2.MustReadLines(fmt.Sprintf("day%d/input", day))
	timeTaken = time.Duration(0)
	timeTaken += advent2.Timed("part1", func() string {
		return part1(lines)
	})
	if part2 != nil {
		timeTaken += advent2.Timed("part2", func() string {
			return part2(lines)
		})
	}

	return
}

func main() {
	timeTaken := time.Duration(0)
	timeTaken += run(1, day1.Part1, day1.Part2)
	timeTaken += run(2, day2.Part1, day2.Part2)
	timeTaken += run(3, day3.Part1, day3.Part2)
	timeTaken += run(4, day4.Part1, day4.Part2)
	timeTaken += run(5, day5.Part1, day5.Part2)
	timeTaken += run(6, day6.Part1, day6.Part2)
	timeTaken += run(7, day7.Part1, day7.Part2)
	timeTaken += run(8, day8.Part1, day8.Part2)
	timeTaken += run(9, day9.Part1, day9.Part2)
	timeTaken += run(10, day10.Part1, day10.Part2)
	timeTaken += run(11, day11.Part1, day11.Part2)
	timeTaken += run(12, day12.Part1, day12.Part2)
	timeTaken += run(13, day13.Part1, day13.Part2)
	timeTaken += run(14, day14.Part1, day14.Part2)
	timeTaken += run(15, day15.Part1, day15.Part2)
	timeTaken += run(16, day16.Part1, day16.Part2)
	timeTaken += run(17, day17.Part1, day17.Part2)
	timeTaken += run(18, day18.Part1, day18.Part2)

	fmt.Printf("\n%30s Summary %30s\n\n", "", "")
	fmt.Printf("All solutions combined took: %s\n", timeTaken)
}
