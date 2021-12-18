package advent_2015_day6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLights(t *testing.T) {
	t.Run("Test parse line", func(t *testing.T) {
		assert.Equal(t, Instruction{Rect{Point{301,3}, Point{808,453}}, TurnOff}, parseLine("turn off 301,3 through 808,453"))
		assert.Equal(t, Instruction{Rect{Point{351,678}, Point{951,908}}, TurnOn}, parseLine("turn on 351,678 through 951,908"))
		assert.Equal(t, Instruction{Rect{Point{720,196}, Point{897,994}}, Toggle}, parseLine("toggle 720,196 through 897,994"))
	})

	t.Run("Solve", func(t *testing.T) {
		instructions := []Instruction {
			{Rect{Point{0,0}, Point{2,2}}, TurnOn}, // 9
			{Rect{Point{1,1}, Point{2,2}}, TurnOff}, // 5
			{Rect{Point{0,0}, Point{2,2}}, Toggle}, // 4
		}

		result := solve(Rect{Point{0, 0}, Point{3, 3}}, instructions)
		assert.Equal(t, 4, result)
	})
}