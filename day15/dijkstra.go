package day15

import (
	advent "advent2021/utils"
	"container/heap"
	"fmt"
)

type Point struct {
	x, y int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Node struct {
	coords         Point
	cumulativeRisk int
	from           *Node
}

type NodesQueue struct {
	nodes []*Node
	destination Point
}

func NewNode(coords Point, risk int, from *Node) *Node {
	return &Node {
		coords:         coords,
		cumulativeRisk: risk,
		from:           from,
	}
}

func (nq NodesQueue) Len() int {
	return len(nq.nodes)
}

func (nq NodesQueue) Less(i, j int) bool {
	return nq.nodes[i].cumulativeRisk < nq.nodes[j].cumulativeRisk
}

func (nq NodesQueue) Swap(i, j int) {
	tmp := nq.nodes[i]
	nq.nodes[i] = nq.nodes[j]
	nq.nodes[j] = tmp
}

func (nq *NodesQueue) Push(x interface{}) {
	nq.nodes = append(nq.nodes, x.(*Node))
}

func (nq *NodesQueue) Pop() interface{} {
	nl := len(nq.nodes)
	node := nq.nodes[nl- 1]
	nq.nodes[nl- 1] = nil
	nq.nodes = nq.nodes[0 : nl- 1]
	return node
}

func NewNodesQueue(target Point) *NodesQueue {
	queue := NodesQueue {
		nodes: make([]*Node, 0),
		destination: target,
	}
	heap.Init(&queue)
	return &queue
}

type RiskMap interface {
	getWidth() int
	getHeight() int
	getRisk(coord Point) int
}

type BasicRiskMap struct {
	height, width int
	risks map[Point]int
}

func (b BasicRiskMap) getWidth() int {
	return b.width
}

func (b BasicRiskMap) getHeight() int {
	return b.height
}

func (b BasicRiskMap) getRisk(coord Point) int {
	return b.risks[coord]
}

type ExtendedRiskMap struct {
	underlying RiskMap
	ratio int
}

func (e ExtendedRiskMap) getWidth() int {
	return e.underlying.getWidth() * e.ratio
}

func (e ExtendedRiskMap) getHeight() int {
	return e.underlying.getHeight() * e.ratio
}

func (e ExtendedRiskMap) getRisk(coord Point) int {
	basicPoint := Point{
		x: coord.x % e.underlying.getHeight(),
		y: coord.y % e.underlying.getWidth(),
	}
	cellX := coord.x / e.underlying.getHeight()
	cellY := coord.y / e.underlying.getWidth()
	risk := e.underlying.getRisk(basicPoint) + cellX + cellY
	
	return ((risk - 1) % 9) + 1
}

func extendRiskMap(underlying RiskMap, ratio int) RiskMap {
	return &ExtendedRiskMap{
		underlying: underlying,
		ratio: ratio,
	}
} 

func parseInput(lines []string) RiskMap {
	risks := make(map[Point]int)
	for y, line := range lines {
		for x := 0; x < len(lines[0]); x++ {
			risks[Point{x, y}] = advent.MustAtoi(line[x:x+1])
		}
	}
	return BasicRiskMap{
		len(lines),
		len(lines[0]),
		risks,
	}
}

func findShortestPath(from Point, to Point, riskMap RiskMap) (found *Node, visited int) {
	nodesQueue := NewNodesQueue(to)
	visitedNodes := make(map[Point]bool)
	heap.Push(nodesQueue, NewNode(from, 0, nil))
	visited = 0

	for nodesQueue.Len() > 0 {
		node := heap.Pop(nodesQueue).(*Node)
		coords := node.coords
		visited += 1

		neighboursCoords := [][]int{
			{coords.x - 1, coords.y},
			{coords.x, coords.y - 1},
			{coords.x + 1, coords.y},
			{coords.x, coords.y + 1},
		}

		for i := 0; i < len(neighboursCoords); i++ {
			nx := neighboursCoords[i][0]
			ny := neighboursCoords[i][1]
			if nx >= 0 && ny >= 0 && nx < riskMap.getWidth() && ny < riskMap.getHeight() { // in bounds
				ncoords := Point{nx, ny}
				if _, ok := visitedNodes[ncoords]; !ok { // not visited yet, let's go there
					nextNode := NewNode(ncoords, node.cumulativeRisk + riskMap.getRisk(ncoords), node)
					if nextNode.coords.x == to.x && nextNode.coords.y == to.y { // btw, this is our destination, return it!
						return nextNode, visited
					}
					heap.Push(nodesQueue, nextNode)
					visitedNodes[ncoords] = true
				}
			}
		}
	}

	panic("Path not found.")
}

func Part1(lines []string) string {
	riskMap := parseInput(lines)
	from := Point{0, 0}
	to := Point{riskMap.getWidth() - 1, riskMap.getHeight() - 1}
	shortest, visited := findShortestPath(from, to, riskMap)

	return fmt.Sprintf("Least risky path %d (checked %d nodes)", shortest.cumulativeRisk, visited)
}

func Part2(lines []string) string {
	riskMap := extendRiskMap(parseInput(lines), 5)
	from := Point{0, 0}
	to := Point{riskMap.getWidth() - 1, riskMap.getHeight() - 1}
	shortest, visited := findShortestPath(from, to, riskMap)

	return fmt.Sprintf("Least risky path %d (checked %d nodes)", shortest.cumulativeRisk, visited)
}