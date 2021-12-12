package day4

import (
	"advent2021/utils"
	"fmt"
	"strings"
)

/*
 *	https://adventofcode.com/2021/day/4
 */

func loadInput(lines []string) (numbers []uint8, boards []*Board) {
	numbers = []uint8{}
	for _, stringyNumber := range strings.Split(lines[0], ",") {
		numbers = append(numbers, uint8(advent.MustAtoi(stringyNumber)))
	}
	boards = []*Board{}

	idx := 0
	for i := 2; i < len(lines); i += 6 {
		boards = append(boards, NewBoard(idx, lines[i:i+5]))
		idx++
	}
	return
}

func Part1(lines []string) string {
	numbers, boards := loadInput(lines)

	for _, number := range numbers {
		for _, board := range boards {
			board.Strike(number)
			if board.Wins() {
				left := board.SumOfFieldsLeft()
				return fmt.Sprintf("winner: board=%d, called*left=result; %d*%d=%d", board.number, number, left, int(number)*left)
			}
		}
	}

	return "No one won!"
}

func Part2(lines []string) string {
	numbers, boards := loadInput(lines)

	for _, number := range numbers {
		var pendingBoards []*Board
		for _, board := range boards {
			board.Strike(number)
			if len(boards) == 1 && board.Wins() {
				left := board.SumOfFieldsLeft()
				return fmt.Sprintf("last winner: board=%d, called*left=result; %d*%d=%d", board.number, number, left, int(number)*left)
			}

			if !board.Wins() {
				pendingBoards = append(pendingBoards, board)
			}
		}

		boards = pendingBoards
	}

	return fmt.Sprintf("%d boards left after all numbers!", len(boards))
}
