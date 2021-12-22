package advent_2015_day2

import (
	advent "advent2021/utils"
	"fmt"
	"sort"
	"strings"
)

type Box struct{ l, w, h int }

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	} else {
		return c
	}
}

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	} else {
		return c
	}
}

func paperNeeded(box Box) int {
	a := box.l * box.w
	b := box.w * box.h
	c := box.h * box.l
	s := min(a, b, c)
	return 2*(a+b+c) + s
}

func ribbonNeeded(box Box) int {
	dims := []int{box.w, box.h, box.l}
	sort.Ints(dims)
	return 2*(dims[0]+dims[1]) + box.w*box.h*box.l
}

type calcFunc func(Box) int

func solve(lines []string, calc calcFunc) int {
	needed := 0
	for _, line := range lines {
		split := strings.Split(line, "x")
		box := Box{advent.MustAtoi(split[0]), advent.MustAtoi(split[1]), advent.MustAtoi(split[2])}
		needed += calc(box)
	}
	return needed
}

func Part1(lines []string) string {
	return fmt.Sprintf("Paper needed: %d", solve(lines, paperNeeded))
}

func Part2(lines []string) string {
	return fmt.Sprintf("Ribbon needed: %d", solve(lines, ribbonNeeded))
}
