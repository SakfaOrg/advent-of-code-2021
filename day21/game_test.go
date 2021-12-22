package day21

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame(t *testing.T) {
	t.Run("test deterministic die", func(t *testing.T) {
		die := NewDeterministicDie(3)
		assert.Equal(t, 1, die.roll())
		assert.Equal(t, 2, die.roll())
		assert.Equal(t, 3, die.roll())
		assert.Equal(t, 1, die.roll())
	})

	t.Run("parse player", func(t *testing.T) {
		player := parsePlayer("Player 1 starting position: 4")
		assert.Equal(t, 1, player.number)
		assert.Equal(t, 4, player.position)
		assert.Equal(t, 0, player.score)
	})

	t.Run("test part1 demo", func(t *testing.T) {
		result := Part1([]string{
			"Player 1 starting position: 4",
			"Player 2 starting position: 8",
		})
		assert.Equal(t, "Loser score 745, rolls 993, result = 739785", result)
	})

	t.Run("test part2 demo", func(t *testing.T) {
		result := Part2([]string{
			"Player 1 starting position: 4",
			"Player 2 starting position: 8",
		})
		assert.Equal(t, "done after 19 loops, player 1 victories 444356092776315, player 2 victories 341960390180808", result)
	})
}
