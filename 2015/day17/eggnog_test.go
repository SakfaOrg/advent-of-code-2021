package advent_2015_day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEggnog(t *testing.T) {
	t.Run("Test none", func(t *testing.T) {
		assert.Equal(t, 0, countCombinations(5, []int {}))
		assert.Equal(t, 0, countCombinations(5, []int {3}))

		// requirement: filling all containers ""entirely**
		assert.Equal(t, 0, countCombinations(5, []int {6}))
		assert.Equal(t, 0, countCombinations(0, []int {5, 5, 5}))
	})

	t.Run("Test one", func(t *testing.T) {
		assert.Equal(t, 1, countCombinations(5, []int {5}))
	})

	t.Run("Test collect", func(t *testing.T) {
		assert.Equal(t, [][]int{{5}}, collectCombinations(5, []int {5}))
	})

	t.Run("Test demo 1", func(t *testing.T) {
		assert.Equal(t, 1, countCombinations(25, []int{20, 5}))
	})

	t.Run("Test demo 2", func(t *testing.T) {
		assert.Equal(t, 2, countCombinations(25, []int{20, 5, 15, 10}))
	})

	t.Run("Test collect demo 2", func(t *testing.T) {
		assert.Equal(t, [][]int{{20, 5}, {15,10}}, collectCombinations(25, []int{20, 5, 15, 10}))
	})

	t.Run("Test demo full", func(t *testing.T) {
		result := countCombinations(25, []int {20, 15, 10, 5, 5})
		assert.Equal(t, 4, result)
	})

	t.Run("Test collect demo full", func(t *testing.T) {
		result := collectCombinations(25, []int {20, 15, 10, 5, 5})

		assert.Equal(t, [][]int{{20,5}, {20,5}, {15,10}, {15,5,5} }, result)
	})

	t.Run("Filter smallest", func(t *testing.T) {
		assert.Equal(t, [][]int{{20,5}, {20,5}, {15,10}}, filterSmallest([][]int{{20,5}, {20,5}, {15,10}, {15,5,5}}))
	})

	t.Run("Puzzle", func(t *testing.T) {
		result := countCombinations(150, []int {43, 3, 4, 10, 21, 44, 4, 6, 47, 41, 34, 17, 17, 44, 36, 31, 46, 9, 27, 38})
		assert.Equal(t, 1638, result)
	})
}