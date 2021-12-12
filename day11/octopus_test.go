package day11

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestOctopuses(t *testing.T) {
	demoInput := strings.Split(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`, "\n")

	t.Run("Test Cave", func(t *testing.T) {
		c := NewCave([]string{"123", "789"})
		assert.Equal(t, 2, c.height)
		assert.Equal(t, 3, c.width)
		assert.Equal(t, Octopus{0, 0, 1}, c.octopi[0][0])
		assert.Equal(t, Octopus{1, 2, 9}, c.octopi[1][2])
	})

	t.Run("Test tick", func(t *testing.T) {
		c := NewCave([]string{"193", "559"})
		flashed := tick(c)
		assert.Equal(t, 2, flashed)
	})

	t.Run("Test neighbours of", func(t *testing.T) {
		c := NewCave([]string{"987", "456"})
		neighbours := c.neighboursOf(c.octopi[0][1])
		assert.Equal(t, 5, len(neighbours))
	})

	t.Run("Test tick easy multiple generations", func(t *testing.T) {
		c := NewCave([]string{"988", "657"})
		flashed := tick(c)
		assert.Equal(t, 6, flashed)
	})

	t.Run("Test demo input small", func(t *testing.T) {
		c := NewCave([]string{"11111", "19991", "19191", "19991", "11111"})
		flashed := tick(c)
		assert.Equal(t, 9, flashed)
	})

	t.Run("Test demo input big", func(t *testing.T) {
		c := NewCave(demoInput)
		flashed := simulate(c, 10)
		assert.Equal(t, 204, flashed)
	})

	t.Run("Test demo input part1", func(t *testing.T) {
		assert.Equal(t, "Flashed after 100 days: 1656", Part1(demoInput))
	})

	t.Run("Test demo input part2", func(t *testing.T) {
		assert.Equal(t, "All octopi flashed on day 195", Part2(demoInput))
	})
}
