package advent_2015_day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNice(t *testing.T) {
	t.Run("Is nice", func(t *testing.T) {
		assert.Equal(t, true, isNice("ugknbfddgicrmopn"))
		assert.Equal(t, true, isNice("aaa"))
		assert.Equal(t, false, isNice("jchzalrnumimnmhp"))
		assert.Equal(t, false, isNice("haegwjzuvuyypxyu"))
		assert.Equal(t, false, isNice("dvszwmarrgswjxmb"))
	})

	t.Run("Is nice2", func(t *testing.T) {
		assert.Equal(t, true, isNicePart2("qjhvhtzxzqqjkmpb"))
		assert.Equal(t, true, isNicePart2("xxyxx"))
		assert.Equal(t, false, isNicePart2("uurcxstgmygtbstg"))
		assert.Equal(t, false, isNicePart2("ieodomkazucvgmuy"))
	})
}
