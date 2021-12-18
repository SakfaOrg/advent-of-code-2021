package advent_2015_day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPresents(t *testing.T) {
	t.Run("test demo", func(t *testing.T) {
		assert.Equal(t, 2, solve(">"))
		assert.Equal(t, 4, solve("^>v<"))
		assert.Equal(t, 2, solve("^v^v^v^v^v"))
	})

	t.Run("test robosanta demo", func(t *testing.T) {
		assert.Equal(t, 3, solveWithRoboSanta("^v"))
		assert.Equal(t, 3, solveWithRoboSanta("^>v<"))
		assert.Equal(t, 11, solveWithRoboSanta("^v^v^v^v^v"))
	})
}
