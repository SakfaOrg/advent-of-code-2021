package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange(t *testing.T) {
	t.Run("Test demo", func(t *testing.T) {
		result := Part1([]string{"target area: x=20..30, y=-10..-5"})
		assert.Equal(t, "Highest shot x=6,y=9 reaches target after 20 steps with y high=45", result)
	})

	t.Run("Test Part1", func(t *testing.T) {
		result := Part1([]string{"target area: x=34..67, y=-215..-186"})
		assert.Equal(t, "Highest shot x=8,y=214 reaches target after 430 steps with y high=23005", result)
	})

	t.Run("Test findAll", func(t *testing.T) {
		result := findAll(Range{20, 30}, Range{-10, -5})
		assert.Equal(t, 112, result)
	})
}
