package advent_2015_day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMining(t *testing.T) {
	t.Run("test demo", func(t *testing.T) {
		assert.Equal(t, 609043, mine("abcdef"))
		assert.Equal(t, 1048970, mine("pqrstuv"))
	})
}