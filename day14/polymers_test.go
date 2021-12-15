package day14

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPolymers(t *testing.T) {
	demoInputString := "NNCB\n\nCH -> B\nHH -> N\nCB -> H\nNH -> C\nHB -> C\nHC -> B\nHN -> C\nNN -> C\nBH -> H\nNC -> B\nNB -> B\nBN -> B\nBB -> N\nBC -> B\nCC -> N\nCN -> C"
	demoInput := strings.Split(demoInputString, "\n")

	t.Run("Test parse input", func(t *testing.T) {
		polymer, insertions := parseInput(demoInput)
		assert.Equal(t, "NNCB", polymer)
		assert.Equal(t, 16, len(insertions))
		assert.Equal(t, "B", insertions["HC"])
	})

	t.Run("Test demo compact", func(t *testing.T) {
		polymer, insertions := parseInput(demoInput)
		expanded := expandNTimes(10, countPairs(polymer), insertions)

		counts := countLetters(expanded, polymer[len(polymer)-1:])
		assert.Equal(t, 1749, counts["B"])
		assert.Equal(t, 161, counts["H"])
	})

	t.Run("Test demo alternative", func(t *testing.T) {
		pairCount := countPairs("NNCB")
		assert.Equal(t, 3, len(pairCount))
		assert.Equal(t, 1, pairCount["NN"])
		pairCount = countPairs("NNNCB")
		assert.Equal(t, 3, len(pairCount))
		assert.Equal(t, 2, pairCount["NN"])
	})

	t.Run("Test part1 demo", func(t *testing.T) {
		assert.Equal(t, "most frequent B (1749) minus least frequent H (161) = 1588", Part1(demoInput))
	})

}
