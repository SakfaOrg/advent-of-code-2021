package advent_2015_day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaperNeeded(t *testing.T) {
	t.Run("test paper needed", func(t *testing.T) {
		assert.Equal(t, 43, paperNeeded(Box{1, 1, 10}))
		assert.Equal(t, 58, paperNeeded(Box{2, 3, 4}))
	})

	t.Run("test ribbon needed", func(t *testing.T) {
		assert.Equal(t, 14, ribbonNeeded(Box{1, 1, 10}))
		assert.Equal(t, 34, ribbonNeeded(Box{2, 3, 4}))
	})
}
