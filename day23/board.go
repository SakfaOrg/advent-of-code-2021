package day23

import (
	advent "advent2021/utils"
	"strconv"
)

type Field uint8

const (
	EMPTY Field = iota
	AMBER
	BRONZE
	COPPER
	DESERT
)

func (f Field) cost() int {
	switch f {
	case AMBER:
		return 1
	case BRONZE:
		return 10
	case COPPER:
		return 100
	case DESERT:
		return 1000
	case EMPTY:
		panic("checking cost of empty field")
	default:
		panic("unknown field " + strconv.Itoa(int(f)))
	}
}

func (f Field) room2() int {
	return f.room(2)
}

func (f Field) room4() int {
	return f.room(4)
}

func (f Field) room(roomSize int) int {
	return 11 + roomSize*int(f-1)
}

func (f Field) String() string {
	switch f {
	case AMBER:
		return "A"
	case BRONZE:
		return "B"
	case COPPER:
		return "C"
	case DESERT:
		return "D"
	case EMPTY:
		return " "
	default:
		panic("unknown field " + strconv.Itoa(int(f)))
	}
}

func validMoves(b []Field, boardSize int, roomSize int) []Move {
	var result []Move
NEXT:
	for from := 0; from < boardSize; from++ {
		if b[from] != EMPTY {
			amphipod := b[from]
			roomTop := amphipod.room(roomSize)
			roomBottom := roomTop + roomSize
			// case 1: amphipod is already in a correct room
			if from >= roomTop && from < roomBottom {
				// case 1.1: we're in our room, but there are only empty spots below us
				var i int
				for i = from + 1; i < roomBottom && b[i] == EMPTY; i++ {
				}
				if i != from+1 {
					result = append(result, Move{from, i - 1, amphipod.cost()*(i - from - 1)})
					continue NEXT
				}

				// case 1.2: we're in our room, and there are OTHER amphibos we're blocking below us. We never leave our room
				// unless we're blocking someone, hence this check
				blocking := false
				for i = from + 1; i < roomBottom; i++ {
					if b[i] != EMPTY && b[i] != amphipod {
						blocking = true
						break
					}
				}
				if !blocking {
					continue NEXT // case 1.3 we're in correct room, no room to go down, we're not blocking anyone. Stay where you are.
				}
			}

			// case 2: amphipod is either in the incorrect room, or in correct room but it blocks someone in the incorrect room.
			// from that position we can move to any reachable spot in hallway
			if from > 10 {
				entrance := roomEntrance(from, roomSize)
				pathToEntrance := path(from, entrance, roomSize)
				if isPathFree(b, pathToEntrance) {
					for i := entrance; i >= 0; i-- {
						if b[i] != EMPTY {
							break
						}
						if !isRoomEntrance(i) {
							moveLength := len(pathToEntrance) + advent.Abs(entrance-i)
							result = append(result, Move{from, i, amphipod.cost() * (moveLength - 1)})
						}
					}
					for i := entrance; i <= 10; i++ {
						if b[i] != EMPTY {
							break
						}
						if !isRoomEntrance(i) {
							moveLength := len(pathToEntrance) + advent.Abs(entrance-i)
							result = append(result, Move{from, i, amphipod.cost() * (moveLength - 1)})
						}
					}
				}
			}

			// case 3; amphipod is in the corridor and wants to travel to the bottom of its room if it's free
			if from <= 10 {
				for i := roomTop; i < roomBottom; i++ {
					if b[i] != EMPTY && b[i] != amphipod {
						continue NEXT // we're not going to our room, there's someone else there!
					}
				}

				if b[roomTop] == EMPTY {
					var target int
					for target = roomTop; target < roomBottom && b[target] == EMPTY; target++ {
					}
					path := path(from, target-1, roomSize)
					if isPathFree(b, path) {
						result = append(result, Move{from, target - 1, amphipod.cost() * (len(path) - 1)})
					}
				}
			}
		}
	}
	return result
}

// returns a shortest path from one field to another
func path(from, to, roomSize int) []int {
	if from == to {
		return []int{from}
	}

	if from <= 10 && to <= 10 { // hallway to hallway, easy peasy (straight line)
		if from < to {
			var result []int
			for i := from; i <= to; i++ {
				result = append(result, i)
			}
			return result
		} else {
			return reverse(path(to, from, roomSize))
		}
	}
	if from > 10 { // room to somewhere else: easy: get out of the room and then go to destination
		// which is either hallway<->hallway path we already covered or hallway<->room path
		// which is same as already covered room<->hallway just in reverse
		top := roomTop(from, roomSize)
		entrance := roomEntrance(from, roomSize)
		if to >= top && to < top+roomSize {
			var result []int
			if from > to {
				for i := from; i >= to; i-- {
					result = append(result, i)
				}
				return result
			} else {
				return reverse(path(to, from, roomSize))
			}
		} else {
			var toEntrance []int
			for i := from; i >= top; i-- {
				toEntrance = append(toEntrance, i)
			}
			return append(toEntrance, path(entrance, to, roomSize)...)
		}
	} else if to > 10 && from <= 10 { // hallway to room, just reverse it so we have room to hallway which we already covered
		return reverse(path(to, from, roomSize))
	} else {
		panic("Illegal room to room move")
	}
}

func isRoomEntrance(idx int) bool {
	return idx == 2 || idx == 4 || idx == 6 || idx == 8
}

func isPathFree(fields []Field, path []int) bool {
	for i := 1; i < len(path); i++ {
		if fields[path[i]] != EMPTY {
			return false
		}
	}
	return true
}

func roomEntrance(idx int, roomSize int) int {
	roomNumber := (idx-11)/roomSize + 1
	return 2 * roomNumber
}

func roomTop(idx int, roomSize int) int {
	roomNumber := (idx - 11) / roomSize
	return 11 + roomNumber*roomSize
}

func reverse(ints []int) (result []int) {
	result = make([]int, len(ints))
	for i := len(ints) - 1; i >= 0; i-- {
		result[len(ints)-1-i] = ints[i]
	}
	return
}