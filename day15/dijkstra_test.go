package day15

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	t.Run("Test queue", func(t *testing.T) {
		queue := NewNodesQueue(Point{10, 10})
		heap.Push(queue, NewNode(Point{0, 0}, 0, nil)) // 20
		heap.Push(queue, NewNode(Point{1, 0}, 15, nil)) // 34
		heap.Push(queue, NewNode(Point{2, 0}, 10, nil)) // 28
		heap.Push(queue, NewNode(Point{3, 0}, 5, nil)) // 22
		heap.Push(queue, NewNode(Point{3, 3}, 3, nil)) // 17
		heap.Push(queue, NewNode(Point{1, 1}, 3, nil)) // 21
		heap.Push(queue, NewNode(Point{5, 5}, 3, nil)) // 13

		node := heap.Pop(queue).(*Node)
		assert.Equal(t, Point{5, 5}, node.coords)
		assert.Equal(t, 3, node.cumulativeRisk)

		node = heap.Pop(queue).(*Node)
		assert.Equal(t, Point{3, 3}, node.coords)
		assert.Equal(t, 3, node.cumulativeRisk)

		node = heap.Pop(queue).(*Node)
		assert.Equal(t, Point{0, 0}, node.coords)
		assert.Equal(t, 0, node.cumulativeRisk)
	})

	demoInput := strings.Split("1163751742\n1381373672\n2136511328\n3694931569\n7463417111\n1319128137\n1359912421\n3125421639\n1293138521\n2311944581", "\n")

	t.Run("Test demo input", func(t *testing.T) {
		riskMap := parseInput(demoInput)
		assert.Equal(t, 10, riskMap.getWidth())
		assert.Equal(t, 10, riskMap.getHeight())
		assert.Equal(t, 1, riskMap.getRisk(Point{0, 0}))
		assert.Equal(t, 6, riskMap.getRisk(Point{2, 0}))
		assert.Equal(t, 5, riskMap.getRisk(Point{2, 6}))
	})

	t.Run("Test extended map", func(t *testing.T) {
		extended := extendRiskMap(parseInput(demoInput), 5)
		assert.Equal(t, 50, extended.getWidth())
		assert.Equal(t, 50, extended.getHeight())
		assert.Equal(t, 1, extended.getRisk(Point{0, 0}))
		assert.Equal(t, 6, extended.getRisk(Point{2, 0}))
		assert.Equal(t, 5, extended.getRisk(Point{2, 6}))

		assert.Equal(t, 7, extended.getRisk(Point{5, 48}))
		assert.Equal(t, 3, extended.getRisk(Point{10, 10}))
		assert.Equal(t, 1, extended.getRisk(Point{19, 7}))
		assert.Equal(t, 9, extended.getRisk(Point{49, 49}))
	})

	t.Run("Test find shortest path", func(t *testing.T) {
		riskMap := parseInput([]string {"19999", "19111", "11191"})
		from := Point{0, 0}
		to := Point{riskMap.getWidth() - 1, riskMap.getHeight() - 1}
		shortest, _ := findShortestPath(from, to, riskMap)
		assert.Equal(t, to, shortest.coords)
		assert.Equal(t, 8, shortest.cumulativeRisk)
	})

	t.Run("Test find shortest path demo input", func(t *testing.T) {
		riskMap := parseInput(demoInput)
		from := Point{0, 0}
		to := Point{riskMap.getWidth() - 1, riskMap.getHeight() - 1}
		shortest, _ := findShortestPath(from, to, riskMap)
		assert.Equal(t, to, shortest.coords)
		assert.Equal(t, 40, shortest.cumulativeRisk)
	})

	t.Run("Test find shortest path demo input extended", func(t *testing.T) {
		riskMap := extendRiskMap(parseInput(demoInput), 5)
		from := Point{0, 0}
		to := Point{riskMap.getWidth() - 1, riskMap.getHeight() - 1}
		shortest, _ := findShortestPath(from, to, riskMap)
		assert.Equal(t, to, shortest.coords)
		assert.Equal(t, 315, shortest.cumulativeRisk)
	})
}
