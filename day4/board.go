package day4

import (
	"advent2021/utils"
	"fmt"
	"regexp"
	"strings"
)

type Field struct {
	number uint8
	striked bool
}

func NewField(number uint8) *Field {
	return &Field{
		number: number,
		striked: false,
	}
}

type Board struct {
	number int
	fields [][]*Field
}

func NewBoard(number int, lines []string) *Board {
	splitter := regexp.MustCompile("\\s+")
	fields := make([][]*Field, len(lines))
	for row, line := range lines {
		numbers := splitter.Split(strings.TrimLeft(line, " "), -1)
		fields[row] = make([]*Field, len(numbers))
		for col, number := range numbers {
			fields[row][col] = NewField(uint8(advent.MustAtoi(number)))
		}
	}
	return &Board{
		number: number,
		fields: fields,
	}
}

func (b Board) Strike(number uint8) {
	for _, line := range b.fields {
		for _, field := range line {
			if field.number == number {
				field.striked = true
			}
		}
	}
}

func (b Board) Wins() bool {
	for _, line := range b.fields {
		wins := true
		for _, field := range line {
			if !field.striked {
				wins = false
				break
			}
		}
		if wins {
			return true
		}
	}

	for col := 0; col < len(b.fields[0]); col++ {
		wins := true
		for row := 0; row < len(b.fields); row++ {
			if !b.fields[row][col].striked {
				wins = false
				break
			}
		}
		if wins {
			return true
		}
	}

	return false;
}

func (b Board) SumOfFieldsLeft() int {
	left := 0
	for _, lines := range b.fields {
		for _, field := range lines {
			if !field.striked {
				left += int(field.number)
			}
		}
	}
	return left
}

func (b Board) String() string {
	lines := []string{}
	for _, fields := range b.fields {
		pieces := []string{}
		for _, field := range fields {
			striked := " "
			if field.striked {
				striked = "*"
			}
			pieces = append(pieces, fmt.Sprintf("%s%02d", striked, field.number))
		}
		lines = append(lines, strings.Join(pieces, " "))
	}
	return strings.Join(lines, "\n")
}