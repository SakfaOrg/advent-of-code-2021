package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrabs(t *testing.T) {
	t.Run("Test demo input", func(t *testing.T) {
		input := []string { "16,1,2,0,4,2,7,1,2,14" }

		assert.Equal(t, "position=2 (fuel=37)", Part1(input))
	})

	t.Run("Test demo input 2", func(t *testing.T) {
		input := []string { "16,1,2,0,4,2,7,1,2,14" }

		assert.Equal(t, "position=5 (fuel=168)", Part2(input))
	})
}
