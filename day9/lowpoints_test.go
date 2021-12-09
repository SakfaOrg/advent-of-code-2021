package day9

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLowpoints(t *testing.T) {
	demoInput :=
`2199943210
3987894921
9856789892
8767896789
9899965678`
	demoInputLines := strings.Split(demoInput, "\n")

	t.Run("Test demo input", func(t *testing.T) {
		assert.Equal(t, "sum of risk points=15", Part1(demoInputLines))
	})

	t.Run("Test demo input part2", func(t *testing.T) {
		assert.Equal(t, "product of top three basins=1134", Part2(demoInputLines))
	})
}