package day23

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestAmphipods(t *testing.T) {
	t.Run("Test board paths", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3}, path(1, 3, 2))
		assert.Equal(t, []int{3, 2, 1}, path(3, 1, 2))
		assert.Equal(t, []int{11, 12}, path(11, 12, 2))
		assert.Equal(t, []int{16, 15}, path(16, 15, 2))
		assert.Equal(t, []int{0, 1, 2, 11}, path(0, 11, 2))

		assert.Equal(t, []int{12, 11, 2, 1, 0}, path(12, 0, 2))
		assert.Equal(t, []int{12, 11, 2, 3, 4, 5, 6, 15, 16}, path(12, 16, 2))
		assert.Equal(t, []int{13, 12, 11, 2, 3, 4, 15, 16}, path(13, 16, 4))
		assert.Equal(t, []int{11, 12, 13, 14}, path(11, 14, 4))
		assert.Equal(t, []int{14, 13, 12, 11}, path(14, 11, 4))
	})

	t.Run("Test signature", func(t *testing.T) {
		board := solvedBoard()
		expected, _ := strconv.ParseInt("1001010001100000111101110011010110001011", 2, 0)
		assert.Equal(t, Signature{a: expected, b: 0}, board.signature())
	})

	t.Run("Test valid moves", func(t *testing.T) {
		moves := solvedBoard().validMoves()
		assert.Equal(t, 0, len(moves))
		fmt.Printf("%#v\n", moves)

		board := solvedBoard().apply(Move{from: 11, to: 0}).apply(Move{from: 14, to: 3})
		moves = board.validMoves()
		fmt.Printf("%#v\n", moves)
		assert.Equal(t, 2, len(moves))
		assert.Equal(t, Move{from: 0, to: 11, cost: 3}, moves[0])
		assert.Equal(t, Move{from: 13, to: 14, cost: 10}, moves[1])

		test := demoBoard().apply(Move{from: 13, to: 3})
		fmt.Printf("%s\n", test)
		moves = test.validMoves()
		fmt.Printf("%s\n", test)
		for _, move := range moves {
			fmt.Printf("move %s\n", move)
		}
		assert.Equal(t, 14, len(test.validMoves()))
	})

	t.Run("Test demo moves", func(t *testing.T) {
		board := demoBoard()
		expectedMoves := []Move{
			{from: 15, to: 3, cost: 40},
			{from: 13, to: 5, cost: 200},
			{from: 5, to: 15, cost: 200},
			{from: 14, to: 5, cost: 3000},
			{from: 3, to: 14, cost: 30},
			{from: 11, to: 3, cost: 20},
			{from: 3, to: 13, cost: 20},
			{from: 17, to: 7, cost: 2000},
			{from: 18, to: 9, cost: 3},
			{from: 7, to: 18, cost: 3000},
			{from: 5, to: 17, cost: 4000},
			{from: 9, to: 11, cost: 8},
		}
		for _, expectedMove := range expectedMoves {
			moves := board.validMoves()
			had := false
			for _, move := range moves {
				if move == expectedMove {
					had = true
				}
			}
			if !had {
				fmt.Printf("Did not find move %d -> %d (%d cost) from board\n%s",
					expectedMove.from, expectedMove.to, expectedMove.cost, board)
				t.Fail()
			}
			board = board.apply(expectedMove)
		}
		assert.Equal(t, true, board.isArranged())
	})

	t.Run("Test arrange", func(t *testing.T) {
		board := solvedBoard().apply(Move{from: 11, to: 0}).apply(Move{from: 14, to: 3})

		_, solution := arrange(board)
		assert.NotNil(t, solution)
	})

	t.Run("Test demo", func(t *testing.T) {
		_, solution := arrange(demoBoard())
		if solution != nil {
			assert.Equal(t, 12521, solution.cost)
		} else {
			t.Fail()
		}
	})

	t.Run("Task board", func(t *testing.T) {
		_, solution := arrange(taskBoard())
		assert.Equal(t, 11417, solution.cost)
	})

	t.Run("Task board Part2", func(t *testing.T) {
		boardPart1 := taskBoard().(BoardPart1)
		board := BoardPart2FromPart1(&boardPart1)
		_, solution := arrange(board)
		assert.Equal(t, 49529, solution.cost)
	})

	t.Run("Test extend board", func(t *testing.T) {
		boardPart1 := demoBoard().(BoardPart1)
		board := BoardPart2FromPart1(&boardPart1)
		assert.Equal(t, "#############\n#           #\n  #B#C#B#D#  \n  #D#C#B#A#  \n  #D#B#A#C#  \n  #A#D#C#A#  \n  #########  ", board.String())

		_, solution := arrange(board)
		if solution != nil {
			assert.Equal(t, 44169, solution.cost)
		} else {
			t.Fail()
		}
	})

	t.Run("Test demo moves Part 2", func(t *testing.T) {
		boardPart1 := demoBoard().(BoardPart1)
		board := *BoardPart2FromPart1(&boardPart1)
		expectedMoves := []Move{
			{from: 23, to: 10, cost: 3000},
			{from: 24, to: 0, cost: 10},
			{from: 19, to: 9, cost: 40},
			{from: 20, to: 7, cost: 30},
			{from: 21, to: 1, cost: 8},
			{from: 15, to: 5, cost: 200},
			{from: 5, to: 21, cost: 400},
			{from: 16, to: 5, cost: 300},
			{from: 5, to: 20, cost: 300},
			{from: 17, to: 5, cost: 40},
			{from: 18, to: 3, cost: 5000},
			{from: 5, to: 18, cost: 50},
			{from: 7, to: 17, cost: 60},
			{from: 9, to: 16, cost: 70},
			{from: 25, to: 7, cost: 400},
			{from: 7, to: 19, cost: 200},
			{from: 26, to: 9, cost: 5},
			{from: 3, to: 26, cost: 9000},
			{from: 11, to: 3, cost: 20},
			{from: 3, to: 15, cost: 20},
			{from: 12, to: 3, cost: 3000},
			{from: 3, to: 25, cost: 8000},
			{from: 13, to: 3, cost: 4000},
			{from: 1, to: 13, cost: 4},
			{from: 0, to: 12, cost: 4},
			{from: 3, to: 24, cost: 7000},
			{from: 9, to: 11, cost: 8},
			{from: 10, to: 23, cost: 3000},
		}
		for _, expectedMove := range expectedMoves {
			moves := board.validMoves()
			had := false
			for _, move := range moves {
				if move == expectedMove {
					had = true
				}
			}
			if !had {
				fmt.Printf("Did not find move %d -> %d (%d cost) from board\n%s",
					expectedMove.from, expectedMove.to, expectedMove.cost, board)
				t.Fail()
			}
			board = board.apply(expectedMove).(BoardPart2)
		}
		assert.Equal(t, true, board.isArranged())
	})
}

func demoBoard() BoardInterface {
	board := *NewBoardPart1()
	board[11] = BRONZE
	board[12] = AMBER
	board[13] = COPPER
	board[14] = DESERT
	board[15] = BRONZE
	board[16] = COPPER
	board[17] = DESERT
	board[18] = AMBER
	return board
}

func solvedBoard() BoardInterface {
	board := *NewBoardPart1()
	board[11] = AMBER
	board[12] = AMBER
	board[13] = BRONZE
	board[14] = BRONZE
	board[15] = COPPER
	board[16] = COPPER
	board[17] = DESERT
	board[18] = DESERT
	return board
}
