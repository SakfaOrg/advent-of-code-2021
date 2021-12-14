package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransform(t *testing.T) {
	demoInput := []string{"6,10", "0,14", "9,10", "0,3", "10,4", "4,11", "6,0", "6,12", "4,1", "0,13", "10,12",
		"3,4", "3,0", "8,4", "1,10", "2,14", "8,10", "9,0", "", "fold along y=7", "fold along x=5"}

	t.Run("Test vertical transform", func(t *testing.T) {
		fold := HorizontalFold{beforeHeight: 7, beforeWidth: 3, foldY: 4}
		assert.Equal(t, 4, fold.height())
		assert.Equal(t, 3, fold.width())

		assert.Equal(t, Point{0, 0}, fold.transform(Point{0, 0}))
		assert.Equal(t, Point{0, 1}, fold.transform(Point{0, 1}))
		assert.Equal(t, Point{0, 2}, fold.transform(Point{0, 2}))
		assert.Equal(t, Point{0, 3}, fold.transform(Point{0, 3}))
		assert.Equal(t, Point{0, 3}, fold.transform(Point{0, 5}))
		assert.Equal(t, Point{0, 2}, fold.transform(Point{0, 6}))
	})

	t.Run("Test horizontal transform", func(t *testing.T) {
		fold := VerticalFold{beforeHeight: 2, beforeWidth: 7, foldX: 3}

		assert.Equal(t, 2, fold.height())
		assert.Equal(t, 3, fold.width())
		assert.Equal(t, Point{0, 0}, fold.transform(Point{0, 0}))
		assert.Equal(t, Point{1, 0}, fold.transform(Point{1, 0}))
		assert.Equal(t, Point{2, 0}, fold.transform(Point{2, 0}))
		assert.Equal(t, Point{2, 0}, fold.transform(Point{4, 0}))
		assert.Equal(t, Point{1, 0}, fold.transform(Point{5, 0}))
		assert.Equal(t, Point{0, 0}, fold.transform(Point{6, 0}))
	})

	t.Run("Test parse demo input", func(t *testing.T) {
		points, folds := parseInput(demoInput)

		assert.Equal(t, 18, len(points))
		assert.Equal(t, Point{0, 3}, points[3])
		assert.Equal(t, 2, len(folds))
		assert.Equal(t, HorizontalFold{11, 15, 7}, folds[0])
		assert.Equal(t, VerticalFold{11, 7, 5}, folds[1])
	})

	t.Run("Test demo input", func(t *testing.T) {
		assert.Equal(t, "Points in result: 17", Part1(demoInput))
	})
}
