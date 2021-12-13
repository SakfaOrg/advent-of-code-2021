package day13

import (
	advent "advent2021/utils"
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

type Transform interface {
	height() (after int)

	width() (after int)

	transform(before Point) (after Point)
}

type VerticalFold struct {
	beforeWidth  int
	beforeHeight int
	foldX        int
}

type HorizontalFold struct {
	beforeWidth  int
	beforeHeight int
	foldY        int
}

func (v VerticalFold) String() string {
	return fmt.Sprintf("vertical: h=%d,w=%d, x=%d", v.height(), v.width(), v.foldX)
}

func (v VerticalFold) height() (after int) {
	return v.beforeHeight
}

func (v VerticalFold) width() (after int) {
	after = v.beforeWidth - v.foldX - 1
	if !(v.foldX <= (v.beforeWidth / 2)) {
		panic(fmt.Sprintf("Unsupported fold of sheet width=%d at x=%d that is higher than middle.", v.beforeHeight, v.foldX))
	}
	return
}

func (v VerticalFold) transform(before Point) (after Point) {
	after.y = before.y
	if before.x < v.foldX {
		after.x = (before.x - v.foldX + 1) * -1
	} else if before.x == v.foldX {
		panic(fmt.Sprintf("Illegal transform of point %d,%d on fold line %d", before.x, before.y, v.foldX))
	} else {
		after.x = before.x - v.foldX - 1
	}
	return
}

func (h HorizontalFold) String() string {
	return fmt.Sprintf("horizontal: h=%d,w=%d, y=%d", h.height(), h.width(), h.foldY)
}

func (h HorizontalFold) height() (after int) {
	after = h.foldY
	if !(after >= (h.beforeHeight / 2)) {
		panic(fmt.Sprintf("Unsupported fold of sheet height=%d at y=%d that is higher than middle.", h.beforeHeight, h.foldY))
	}
	return
}

func (h HorizontalFold) width() (after int) {
	return h.beforeWidth
}

func (h HorizontalFold) transform(before Point) (after Point) {
	after.x = before.x
	if before.y < h.foldY {
		after.y = before.y
	} else if before.y == h.foldY {
		panic(fmt.Sprintf("Illegal transform of point %d,%d on fold line %d", before.x, before.y, h.foldY))
	} else {
		after.y = h.foldY - (before.y - h.foldY)
	}
	return
}

func parseInput(lines []string) (points []Point, transforms []Transform) {
	maxX := 0
	maxY := 0
	parsingPoint := true
	for _, line := range lines {
		if parsingPoint {
			if line == "" {
				parsingPoint = false
			} else {
				xy := strings.Split(line, ",")
				point := Point{
					x: advent.MustAtoi(xy[0]),
					y: advent.MustAtoi(xy[1]),
				}
				points = append(points, point)
				if point.x > maxX {
					maxX = point.x
				}
				if point.y > maxY {
					maxY = point.y
				}
			}
		} else {
			split := strings.Split(line, "=")
			foldAlong := split[0]
			axis := foldAlong[len(foldAlong)-1]
			coord := advent.MustAtoi(split[1])

			var width int
			var height int
			if len(transforms) == 0 {
				width = maxX + 1
				height = maxY + 1
			} else {
				lastTransform := transforms[len(transforms)-1]
				width = lastTransform.width()
				height = lastTransform.height()
			}

			if axis == 'y' {
				transforms = append(transforms, HorizontalFold{beforeWidth: width, beforeHeight: height, foldY: coord})
			} else if axis == 'x' {
				transforms = append(transforms, VerticalFold{beforeWidth: width, beforeHeight: height, foldX: coord})
			}
		}
	}
	return
}

func Part1(lines []string) string {
	points, transforms := parseInput(lines)
	isInResult := make(map[Point]bool)

	firstTransform := transforms[0]

	for _, point := range points {
		isInResult[firstTransform.transform(point)] = true
	}

	return fmt.Sprintf("Points in result: %d", len(isInResult))
}

func Part2(lines []string) string {
	points, transforms := parseInput(lines)
	isInResult := make(map[Point]bool)

	for _, point := range points {
		current := point
		for _, transform := range transforms {
			current = transform.transform(current)
		}
		isInResult[current] = true
	}

	var resultLines []string
	finalTransform := transforms[len(transforms)-1]
	for i := 0; i < finalTransform.height(); i++ {
		resultLine := ""
		// seems my input is a mirror transformation of actual digits, print it backwards. Not sure if that's a bug
		// or an extra twist in the puzzle
		for j := finalTransform.width() - 1; j >= 0; j-- {
			if (isInResult[Point{j, i}]) {
				resultLine += "*"
			} else {
				resultLine += " "
			}
		}
		resultLines = append(resultLines, resultLine)
	}
	return fmt.Sprintf("Code: \n%s\n", strings.Join(resultLines, "\n"))
}
