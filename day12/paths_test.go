package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaths(t *testing.T) {
	demoInput := []string{"start-A", "start-b", "A-c", "A-b", "b-d", "A-end", "b-end"}
	demoInput2 := []string{"dc-end", "HN-start", "start-kj", "dc-start", "dc-HN",
		"LN-dc", "HN-end", "kj-sa", "kj-HN", "kj-dc"}
	demoInput3 := []string{"fs-end", "he-DX", "fs-he", "start-DX", "pj-DX", "end-zg",
		"zg-sl", "zg-pj", "pj-he", "RW-he", "fs-DX", "pj-RW", "zg-RW", "start-pj",
		"he-WI", "zg-he", "pj-fs", "start-RW"}

	t.Run("Parse demo input", func(t *testing.T) {
		start := parseInput(demoInput)
		assert.Equal(t, "start", start.name)
		assert.Equal(t, "end", start.connections["A"].connections["b"].connections["end"].name)
	})

	t.Run("Count paths demo input", func(t *testing.T) {
		start := parseInput(demoInput)
		paths := countPaths(start, make(map[string]bool), false)
		assert.Equal(t, 10, paths)
	})

	t.Run("Count paths demo input with revisit", func(t *testing.T) {
		start := parseInput(demoInput)
		paths := countPaths(start, make(map[string]bool), true)
		assert.Equal(t, 36, paths)
	})

	t.Run("Count paths demo input 2", func(t *testing.T) {
		start := parseInput(demoInput2)
		assert.Equal(t, 19, countPaths(start, make(map[string]bool), false))
	})

	t.Run("Count paths demo input 2 with revisit", func(t *testing.T) {
		start := parseInput(demoInput2)
		assert.Equal(t, 103, countPaths(start, make(map[string]bool), true))
	})

	t.Run("Count paths demo input 3", func(t *testing.T) {
		start := parseInput(demoInput3)
		assert.Equal(t, 226, countPaths(start, make(map[string]bool), false))
	})

	t.Run("Count paths demo input 3 with revisit", func(t *testing.T) {
		start := parseInput(demoInput3)
		assert.Equal(t, 3509, countPaths(start, make(map[string]bool), true))
	})
}
