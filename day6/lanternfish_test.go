package day6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay6(t *testing.T) {
	t.Run("Test build table - zeroes only", func(t *testing.T) {
		table := buildSpawnTable(7, 2, 5)
		assert.Equal(t, []int{0, 0, 0, 0, 0, 0}, table)
	})

	t.Run("Test build table", func(t *testing.T) {
		assert.Equal(t, []int{1, 0, 0, 0}, buildSpawnTable(2, 1, 3))
		assert.Equal(t, []int{2, 1, 1, 0, 0, 0}, buildSpawnTable(2, 1, 5))
		assert.Equal(t, []int{3, 2, 1, 1, 0, 0, 0}, buildSpawnTable(2, 1, 6))
		assert.Equal(t, []int{6, 4, 3, 2, 1, 1, 0, 0, 0}, buildSpawnTable(2, 1, 8))
	})

	t.Run("Test demo input", func(t *testing.T) {
		times := []int{3, 4, 3, 1, 2}
		assert.Equal(t, 26, fishOnDay(times, 18))
		assert.Equal(t, 5934, fishOnDay(times, 80))
		assert.Equal(t, 26984457539, fishOnDay(times, 256))
	})
}
