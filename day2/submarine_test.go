package day2

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestPart2(t *testing.T) {
	demoLines := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	t.Run("puzzle demo part1", func(t *testing.T) {
		result := Part1(demoLines)

		assert.Equal(t, result, "horizontal*depth=result: 15*10=150")
	})

	t.Run("puzzle demo part 2", func(t *testing.T) {
		result := Part2(demoLines)

		assert.Equal(t, result, "horizontal*depth=result: 15*60=900")
	})
}
