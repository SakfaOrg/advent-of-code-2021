package day25

import (
	"fmt"
	"strings"
)

type Field uint8
const (
	EMPTY Field = '.'
	EAST  Field = '>'
	SOUTH Field = 'v'
)

type Region [][]Field

func parseLine(line string) []Field {
	result := make([]Field, len(line))
	for idx, chr := range line {
		result[idx] = Field(chr)
	}
	return result
}

type Point struct {
	i,j int
}
type Move struct {
	from, to Point
}

func (r Region) isValid(move Move) bool {
	return r[move.to.i][move.to.j] == EMPTY
}

func (r Region) apply(moves []Move) {
	for _, move := range moves {
		r[move.to.i][move.to.j] = r[move.from.i][move.from.j]
		r[move.from.i][move.from.j] = EMPTY
	}
}

func (r Region) step() int {
	movesApplied := 0
	var moves []Move
	for i, line := range r {
		for j, field := range line {
			if field == EAST {
				move := Move{Point{i,j}, Point{i, (j + 1) % len(line)}}
				if r.isValid(move) {
					moves = append(moves, move)
				}
			}
		}
	}

	movesApplied += len(moves)
	r.apply(moves)
	moves = []Move{}

	for i, line := range r {
		for j, field := range line {
			if field == SOUTH {
				move := Move{Point{i,j}, Point{(i + 1) % len(r), j}}
				if r.isValid(move) {
					moves = append(moves, move)
				}
			}
		}
	}

	movesApplied += len(moves)
	r.apply(moves)

	return movesApplied
}

func (r Region) String() string {
	var lines []string
	for _, row := range r {
		line := ""
		for _, field := range row {
			line += string(field)
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func parseRegion(lines []string) Region {
	result := make(Region, len(lines))
	for idx, line := range lines {
		result[idx] = parseLine(line)
	}
	return result
}

func Part1(lines []string) string {
	var stepsCounter int
	region := parseRegion(lines)

	for stepsCounter = 1; ; stepsCounter++ {
		if region.step() == 0 {
			break;
		}
	}

	return fmt.Sprintf("Cucumbers stopped moving after %d steps", stepsCounter)
}