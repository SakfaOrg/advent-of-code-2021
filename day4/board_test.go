package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard(t *testing.T) {
	t.Run("test board creation, striking and printing", func(t *testing.T) {
		board := NewBoard(0, []string{" 1  0", "22 34"})

		assert.Equal(t, " 01  00\n 22  34", board.String())
		board.Strike(22)
		assert.Equal(t, " 01  00\n*22  34", board.String())
	})

	t.Run("test line win", func(t *testing.T) {
		board := NewBoard(0, []string{" 1  0", "22 34"})
		assert.Equal(t, false, board.Wins())

		board.Strike(22)
		assert.Equal(t, false, board.Wins())

		board.Strike(34)
		assert.Equal(t, true, board.Wins())

		assert.Equal(t, 1, board.SumOfFieldsLeft())
	})

	t.Run("test column win", func(t *testing.T) {
		board := NewBoard(0, []string{" 1  0", "22 34"})
		assert.Equal(t, false, board.Wins())

		board.Strike(0)
		assert.Equal(t, false, board.Wins())

		board.Strike(34)
		assert.Equal(t, true, board.Wins())

		assert.Equal(t, 23, board.SumOfFieldsLeft())
	})
}
