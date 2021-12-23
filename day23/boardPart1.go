package day23

import (
	advent "advent2021/utils"
	"fmt"
)

// BoardPart1
// board consists of 19 fields, mapped to map as follow:
// each field is either EMPTY or has an amphipod on it
//	#############
//	#0123456789a#
//	###b#d#f#h###
//	  #c#e#g#i#
//    #########
//
type BoardPart1 []Field

func (b BoardPart1) validMoves() []Move {
	return validMoves(b, 19, 2)
}

func NewBoardPart1() *BoardPart1 {
	board := make(BoardPart1, 19)
	return &board
}

func (b BoardPart1) String() string {
	corridor := ""
	for i := 0; i <= 10; i++ {
		corridor += b[i].String()
	}
	return fmt.Sprintf("#############\n#%s#\n  #%s#%s#%s#%s#  \n  #%s#%s#%s#%s#  \n  #########  ",
		corridor,
		b[AMBER.room2()].String(), b[BRONZE.room2()].String(), b[COPPER.room2()].String(), b[DESERT.room2()].String(),
		b[AMBER.room2()+1].String(), b[BRONZE.room2()+1].String(), b[COPPER.room2()+1].String(), b[DESERT.room2()+1].String(),
	)
}

// return state of the board. We can pack this in quite a small number by leveraging the fact there are exactly 2
// amphipods of any given type and each amphipods position can be encoded on 5 bits (1 of 19 fields) for a total
// of 40 bits needed to encode a state
func (b BoardPart1) signature() Signature {
	var as, bs, cs, ds []int
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case AMBER:
			as = append(as, i)
		case BRONZE:
			bs = append(bs, i)
		case COPPER:
			cs = append(cs, i)
		case DESERT:
			ds = append(ds, i)
		}
	}

	if SANITY_CHECKS {
		if len(as) != 2 {
			panic("Expected exactly 2 as!")
		}
		if len(bs) != 2 {
			panic("Expected exactly 2 bs!")
		}
		if len(cs) != 2 {
			panic("Expected exactly 2 cs!")
		}
		if len(ds) != 2 {
			panic("Expected exactly 2 ds!")
		}
	}

	return Signature{
		a: int64(advent.Min(as...)<<0+advent.Max(as...)<<5+
			advent.Min(bs...)<<10+advent.Max(bs...)<<15+
			advent.Min(cs...)<<20+advent.Max(cs...)<<25) +
			int64(advent.Min(ds...))<<30 + int64(advent.Max(ds...))<<35,
		b: 0,
	}
}

// BoardPart1.isArranged
// returns true if all amphipods are on valid positions
func (b BoardPart1) isArranged() bool {
	amphipods := []Field{AMBER, BRONZE, COPPER, DESERT}
	allInRooms := true
	for _, amphipod := range amphipods {
		if b[amphipod.room2()] != amphipod || b[amphipod.room2()+1] != amphipod {
			allInRooms = false
			break
		}
	}

	if SANITY_CHECKS && allInRooms {
		for i := 0; i <= 10; i++ {
			if b[i] != EMPTY {
				panic("Arranged board but hallway is not empty???")
			}
		}
	}

	return allInRooms
}

func (b BoardPart1) apply(move Move) BoardInterface {
	newBoard := *NewBoardPart1()
	for i := 0; i < 19; i++ {
		newBoard[i] = b[i]
	}
	newBoard[move.to] = newBoard[move.from]
	newBoard[move.from] = EMPTY
	return newBoard
}
