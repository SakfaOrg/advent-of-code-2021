package day9

import (
	"fmt"
	"sort"
)

type Field struct {
	value   uint8
	visited bool
}

func parseInput(lines []string) [][]Field {
	points := make([][]Field, len(lines))

	for row, line := range lines {
		points[row] = make([]Field, len(line))
		for col := 0; col < len(line); col++ {
			points[row][col] = Field{
				value:   line[col] - '0',
				visited: false,
			}
		}
	}

	return points
}

func floodFill(points [][]Field, row, col int) int {
	height := len(points)
	width := len(points[0])

	points[row][col].visited = true
	visited := 1

	neighbours := [][]int{
		{row - 1, col},
		{row + 1, col},
		{row, col - 1},
		{row, col + 1},
	}

	for ni := 0; ni < len(neighbours); ni++ {
		nrow := neighbours[ni][0]
		ncol := neighbours[ni][1]

		if inBounds(nrow, ncol, height, width) {
			if !points[nrow][ncol].visited && points[nrow][ncol].value != 9 {
				visited += floodFill(points, nrow, ncol)
			}
		}
	}

	return visited
}

func Part2(lines []string) string {
	points := parseInput(lines)
	height := len(points)
	width := len(points[0])
	var basinSizes []int
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if inBounds(row, col, height, width) {
				if !points[row][col].visited && points[row][col].value != 9 {
					basinSize := floodFill(points, row, col)
					basinSizes = append(basinSizes, basinSize)
				}
			}
		}
	}

	sort.Ints(basinSizes)
	topThree := basinSizes[len(basinSizes)-3 : len(basinSizes)]
	result := 1
	for _, basin := range topThree {
		result *= basin
	}

	return fmt.Sprintf("product of top three basins=%d", result)
}

func inBounds(row, col, height, width int) bool {
	return row >= 0 && row < height && col >= 0 && col < width
}

func Part1(lines []string) string {
	points := parseInput(lines)
	height := len(points)
	width := len(points[0])
	var result int
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			isLocalMinimum := true
			neighbours := [][]int{
				{row - 1, col},
				{row + 1, col},
				{row, col - 1},
				{row, col + 1},
			}

			for ni := 0; ni < len(neighbours); ni++ {
				nrow := neighbours[ni][0]
				ncol := neighbours[ni][1]

				if !inBounds(nrow, ncol, height, width) {
					continue
				}

				if points[nrow][ncol].value <= points[row][col].value {
					isLocalMinimum = false
					break
				}
			}

			if isLocalMinimum {
				result += int(points[row][col].value) + 1
			}
		}
	}

	return fmt.Sprintf("sum of risk points=%d", result)

}
