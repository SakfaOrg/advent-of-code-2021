package day5

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const demoInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestDay5(t *testing.T) {
	t.Run("Test parse line", func(t *testing.T) {
		actual := parseLine("0,9 -> 5,9")
		expected := line{
			from: point{0, 9},
			to:   point{5, 9},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("Test fill vertical line", func(t *testing.T) {
		points := make(map[point]int16)
		line := parseLine("0,9 -> 2,9")

		fillPoints(points, line)

		assert.Equal(t, 3, len(points))
		assert.Equal(t, int16(1), points[point{0, 9}])
		assert.Equal(t, int16(1), points[point{1, 9}])
		assert.Equal(t, int16(1), points[point{2, 9}])
	})

	t.Run("Test fill horizontal line", func(t *testing.T) {
		points := make(map[point]int16)
		line := parseLine("0,9 -> 0,7")

		fillPoints(points, line)

		assert.Equal(t, 3, len(points))
		assert.Equal(t, int16(1), points[point{0, 9}])
		assert.Equal(t, int16(1), points[point{0, 8}])
		assert.Equal(t, int16(1), points[point{0, 7}])
	})

	t.Run("Test fill 2 lines", func(t *testing.T) {
		points := make(map[point]int16)

		fillPoints(points, parseLine("0,9 -> 0,7"))
		fillPoints(points, parseLine("0,9 -> 2,9"))

		assert.Equal(t, 5, len(points))
		assert.Equal(t, int16(2), points[point{0, 9}])
		assert.Equal(t, int16(1), points[point{0, 8}])
		assert.Equal(t, int16(1), points[point{0, 7}])
		assert.Equal(t, int16(1), points[point{1, 9}])
		assert.Equal(t, int16(1), points[point{2, 9}])
	})

	t.Run("Test part 1", func(t *testing.T) {
		result := Part1(strings.Split(demoInput, "\n"))

		assert.Equal(t, "overlapping points = 5", result)
	})

	t.Run("Test fill diagonal 45", func(t *testing.T) {
		points := make(map[point]int16)
		fillPoints(points, parseLine("0,1 -> 2,3"))
		fillPoints(points, parseLine("4,5 -> 2,3"))

		assert.Equal(t, 5, len(points))
		assert.Equal(t, int16(1), points[point{0, 1}])
		assert.Equal(t, int16(1), points[point{1, 2}])
		assert.Equal(t, int16(2), points[point{2, 3}])
		assert.Equal(t, int16(1), points[point{3, 4}])
		assert.Equal(t, int16(1), points[point{4, 5}])
	})

	t.Run("Test fill diagonal 135", func(t *testing.T) {
		points := make(map[point]int16)
		fillPoints(points, parseLine("0,4 -> 2,2"))
		fillPoints(points, parseLine("2,2 -> 4,0"))

		assert.Equal(t, 5, len(points))
		assert.Equal(t, int16(1), points[point{0, 4}])
		assert.Equal(t, int16(1), points[point{1, 3}])
		assert.Equal(t, int16(2), points[point{2, 2}])
		assert.Equal(t, int16(1), points[point{3, 1}])
		assert.Equal(t, int16(1), points[point{4, 0}])
	})

	t.Run("Test part 2", func(t *testing.T) {
		result := Part2(strings.Split(demoInput, "\n"))

		assert.Equal(t, "overlapping points = 12", result)
	})
}
