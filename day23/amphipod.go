package day23

import (
	"container/heap"
	"fmt"
	"sort"
)

const SANITY_CHECKS = true

type Signature struct {
	a, b int64
}

type BoardInterface interface {
	isArranged() bool
	signature() Signature
	validMoves() []Move
	apply(move Move) BoardInterface
}

// Move
// represents a move: from to which field amphipod moved and at what cost
type Move struct {
	from, to, cost int
}

func (m Move) String() string {
	return fmt.Sprintf("%d -> %d (%d cost)", m.from, m.to, m.cost)
}

type Solution struct {
	board BoardInterface
	moves []Move
	cost  int
}

type SolutionsQueue struct {
	nodes []*Solution
}

func (sq SolutionsQueue) Len() int {
	return len(sq.nodes)
}

func (sq SolutionsQueue) Less(i, j int) bool {
	return sq.nodes[i].cost < sq.nodes[j].cost
}

func (sq SolutionsQueue) Swap(i, j int) {
	temp := sq.nodes[i]
	sq.nodes[i] = sq.nodes[j]
	sq.nodes[j] = temp
}

func (sq *SolutionsQueue) Push(x interface{}) {
	sq.nodes = append(sq.nodes, x.(*Solution))
}

func (sq *SolutionsQueue) Pop() interface{} {
	result := sq.nodes[len(sq.nodes)-1]
	sq.nodes = sq.nodes[:len(sq.nodes)-1]
	return result
}

/**
 * find all possible ways to arrange board - with brute force, just check all possible moves each time.
 * This should work because number of legal moves is quite limited
 */
func arrange(initialState BoardInterface) *Solution {
	seenBoards := make(map[Signature]bool) // remember states we already saw so we don't move back and forth
	var toExplore SolutionsQueue
	heap.Init(&toExplore)
	heap.Push(&toExplore, &Solution{initialState, []Move{}, 0})

	explored := 0
	for toExplore.Len() > 0 {
		explored++
		solution := heap.Pop(&toExplore).(*Solution)
		if solution.board.isArranged() {
			return solution
		}

		thisSignature := solution.board.signature()
		if _, ok := seenBoards[thisSignature]; ok {
			continue
		}
		seenBoards[thisSignature] = true
		moves := solution.board.validMoves()
		sort.Slice(moves, func(i, j int) bool {
			return moves[i].cost < moves[j].cost
		})

		for _, move := range moves {
			newMoves := make([]Move, len(solution.moves))
			copy(newMoves, solution.moves)
			nextSolution := Solution{
				solution.board.apply(move),
				append(newMoves, move),
				solution.cost + move.cost,
			}
			heap.Push(&toExplore, &nextSolution)
		}
	}
	return nil
}

func taskBoard() BoardInterface {
	board := *NewBoardPart1()
	board[11] = BRONZE
	board[12] = BRONZE
	board[13] = AMBER
	board[14] = COPPER
	board[15] = AMBER
	board[16] = DESERT
	board[17] = DESERT
	board[18] = COPPER
	return board
}

func Part1(_ []string) string {
	solution := arrange(taskBoard())
	return fmt.Sprintf("Least energy permutation with 2 deep rooms: %d", solution.cost)
}

func Part2(_ []string) string {
	part1Board := taskBoard().(BoardPart1)
	solution := arrange(BoardPart2FromPart1(&part1Board))
	return fmt.Sprintf("Least energy permutation with 4 deep rooms: %d", solution.cost)
}
