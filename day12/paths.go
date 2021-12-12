package day12

import (
	"fmt"
	"strings"
)

type Cave struct {
	name string
	connections map[string]*Cave
}

func NewCave(name string) *Cave {
	return &Cave{
		name: name,
		connections: make(map[string]*Cave),
	}
}

func (c Cave) isStart() bool {
	return c.name == "start"
}

func (c Cave) isEnd() bool {
	return c.name == "end"
}

func (c Cave) isSmall() bool {
	return 'a' <= c.name[0] && c.name[0] <= 'z'
}

func parseInput(lines []string) (start *Cave) {
	caves := make(map[string]*Cave)
	getOrCreateCave := func(name string) *Cave {
		if _, ok := caves[name]; !ok {
			caves[name] = NewCave(name)
		}
		return caves[name]
	}

	for _, line := range lines {
		split := strings.Split(line, "-")
		from := getOrCreateCave(split[0])
		if from.isStart() {
			start = from
		}

		to := getOrCreateCave(split[1])
		from.connections[to.name] = to
		to.connections[from.name] = from
	}

	return
}

func clone(m map[string]bool) map[string]bool {
	cloned := make(map[string]bool)
	for k, v := range m {
		cloned[k] = v
	}
	return cloned
}

func countPaths(cave *Cave, visitedSmallCaves map[string]bool, canRevisitSmall bool) int {
	result := 0
	if cave.isSmall() {
		visitedSmallCaves[cave.name] = true
	}

	for _, connection := range cave.connections {
		if connection.isEnd() {
			result += 1
			continue
		}
		if connection.isStart() || (connection.isSmall() && visitedSmallCaves[connection.name] && !canRevisitSmall) {
			continue
		}

		result += countPaths(connection, clone(visitedSmallCaves), canRevisitSmall && !(connection.isSmall() && visitedSmallCaves[connection.name]))
	}

	return result
}

func Part1(lines []string) string {
	return fmt.Sprintf("Number of paths: %d", countPaths(parseInput(lines), make(map[string]bool), false))
}

func Part2(lines []string) string {
	return fmt.Sprintf("Number of paths: %d", countPaths(parseInput(lines), make(map[string]bool), true))
}