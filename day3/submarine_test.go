package day3

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
	demoLines := strings.Split("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010", "\n")

	t.Run("puzzle demo part1", func(t *testing.T) {
		result := Part1(demoLines)

		assert.Equal(t, result, "gamma*epsilon=result: 22*9=198")
	})

	t.Run("puzzle demo part 2", func(t *testing.T) {
		result := Part2(demoLines)

		assert.Equal(t, "oxygen*co2=result: 23*10=230", result)
	})
}
