package day5

import (
	"advent2021/utils"
	"fmt"
	"strings"
)

/*
 *	https://adventofcode.com/2021/day/5
 */

type point struct {
	x,y int16
}

type line struct {
	from, to point
}

func parsePoint(stringified string) point {
	splitted := strings.Split(stringified, ",")
	return point{
		int16(advent.MustAtoi(splitted[0])),
		int16(advent.MustAtoi(splitted[1])),
	}
}

func parseLine(input string) line {
	splitted := strings.Split(input, " -> ")
	return line{
		parsePoint(splitted[0]),
		parsePoint(splitted[1]),
	}
}

func (l line) isVertical() bool {
	return l.from.y == l.to.y
}

func (l line) isHorizontal() bool {
	return l.from.x == l.to.x
}

func smaller(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}

func larger(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func sign(x int16) int16 {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func fillPoints(points map[point]int16, l line) {
	if l.isVertical() {
		for x := smaller(l.to.x, l.from.x); x <= larger(l.to.x, l.from.x); x++ {
			points[point{x, l.from.y}] = points[point{x, l.from.y}] + 1
		}
	} else if l.isHorizontal() {
		for y := smaller(l.to.y, l.from.y); y <= larger(l.to.y, l.from.y); y++ {
			points[point{l.from.x, y}] = points[point{l.from.x, y}] + 1
		}
	} else {
		dx := sign(l.to.x - l.from.x)
		dy := sign(l.to.y - l.from.y);
		for x, y := l.from.x, l.from.y; x != l.to.x + dx; x, y = x + dx, y + dy {
			points[point{x, y}] = points[point{x, y}] + 1
		}
	}
}

func countOverlapping(points map[point]int16) int {
	count := 0
	for _, v := range points {
		if v > 1 {
			count++
		}
	}
	return count
}

func Part1(lines []string) string {
	points := make(map[point]int16)

	for _, ls := range lines {
		l := parseLine(ls)
		if l.isVertical() || l.isHorizontal() {
			fillPoints(points, l)
		}
	}

	overlapping := countOverlapping(points)
	return fmt.Sprintf("overlapping points = %d", overlapping)
}

func Part2(lines []string) string {
	points := make(map[point]int16)

	for _, ls := range lines {
		l := parseLine(ls)
		fillPoints(points, l)
	}

	overlapping := countOverlapping(points)
	return fmt.Sprintf("overlapping points = %d", overlapping)
}