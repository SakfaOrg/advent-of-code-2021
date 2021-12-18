package advent_2015_day6

import (
	advent "advent2021/utils"
	"fmt"
	"regexp"
)

type Action uint8

func (a Action) apply(i *int) {
	switch a {
	case TurnOn:
		*i = 1
	case TurnOff:
		*i = 0
	case Toggle:
		if *i == 1 {
			*i = 0
		}  else {
			*i = 1
		}
	}
}

func (a Action) applyPart2(i *int) {
	switch a {
	case TurnOn:
		*i++
	case TurnOff:
		*i--
		if *i < 0 {
			*i = 0
		}
	case Toggle:
		*i += 2
	}
}

const (
	TurnOn Action = iota
	TurnOff
	Toggle
)

type Point struct {x,y int}
type Rect struct {topLeft, bottomRight Point}
type Instruction struct {
	rect Rect
	action Action
}

func makeLightMap(rect Rect) (lightMap [][]int) {
	lightMap = make([][]int, rect.bottomRight.y)
	for y := 0; y < rect.bottomRight.y; y++ {
		lightMap[y] = make([]int, rect.bottomRight.x)
	}
	return
}

func solve(rect Rect, instructions []Instruction) int {
	lightMap := makeLightMap(rect)

	for _, instruction := range instructions {
		for x := instruction.rect.topLeft.x; x <= instruction.rect.bottomRight.x; x++ {
			for y := instruction.rect.topLeft.y; y <= instruction.rect.bottomRight.y; y++ {
				instruction.action.apply(&lightMap[x][y])
			}
		}
	}

	lit := 0
	for _, row := range lightMap {
		for _, i := range row {
			if i > 0 {
				lit += 1
			}
		}
	}
	return lit
}

func solvePart2(rect Rect, instructions []Instruction) int {
	lightMap := makeLightMap(rect)

	for _, instruction := range instructions {
		for x := instruction.rect.topLeft.x; x <= instruction.rect.bottomRight.x; x++ {
			for y := instruction.rect.topLeft.y; y <= instruction.rect.bottomRight.y; y++ {
				instruction.action.applyPart2(&lightMap[x][y])
			}
		}
	}

	lit := 0
	for _, row := range lightMap {
		for _, i := range row {
			lit += i
		}
	}
	return lit
}

func parseLine(line string) Instruction {
	regex := *regexp.MustCompile("(turn on|turn off|toggle) (\\d+),(\\d+) through (\\d+),(\\d+)")
	submatches := regex.FindStringSubmatch(line)

	var action Action
	switch submatches[1] {
	case "turn on":
		action = TurnOn
	case "turn off":
		action = TurnOff
	case "toggle":
		action = Toggle
	default:
		panic("Unrecognized action " + submatches[1])
	}

	return Instruction{
		rect: Rect{
			Point{advent.MustAtoi(submatches[2]), advent.MustAtoi(submatches[3])},
			Point{advent.MustAtoi(submatches[4]), advent.MustAtoi(submatches[5])},
		},
		action: action,
	}
}

func parseInput(lines []string) (result []Instruction) {
	for _, line := range lines {
		result = append(result, parseLine(line))
	}
	return result
}

func Part1(lines[] string) string {
	instructions := parseInput(lines)
	return fmt.Sprintf("Lights lit = %d", solve(Rect{Point{0,0},Point{1000,1000}}, instructions))
}

func Part2(lines[] string) string {
	instructions := parseInput(lines)
	return fmt.Sprintf("Lights brightness = %d", solvePart2(Rect{Point{0,0},Point{1000,1000}}, instructions))
}