package advent_2015_day3

import "fmt"

type Point struct{ x, y int }

func move(point *Point, step rune) {
	switch step {
	case '>':
		point.x++
	case '<':
		point.x--
	case '^':
		point.y++
	case 'v':
		point.y--
	default:
		panic("Invalid step " + string(step))
	}
}

func solve(instructions string) int {
	visited := make(map[Point]int)
	santa := Point{0, 0}
	visited[santa] = 1

	for _, step := range instructions {
		move(&santa, step)
		visited[santa]++
	}

	return len(visited)
}

func solveWithRoboSanta(instructions string) int {
	visited := make(map[Point]int)
	santa := Point{0, 0}
	roboSanta := Point{0, 0}
	visited[santa] = 1

	for idx, step := range instructions {
		if idx%2 == 0 {
			move(&santa, step)
			visited[santa]++
		} else {
			move(&roboSanta, step)
			visited[roboSanta]++
		}
	}

	return len(visited)
}

func Part1(lines []string) string {
	return fmt.Sprintf("Houses with presents: %d", solve(lines[0]))
}

func Part2(lines []string) string {
	return fmt.Sprintf("Houses with presents santa+robosanta: %d", solveWithRoboSanta(lines[0]))
}
