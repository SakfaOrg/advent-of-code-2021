package day23

import (
	"fmt"
	"sort"
)

// BoardPart2
// board consists of 27 fields, mapped to map as follow:
// each field is either EMPTY or has an amphipod on it
//	#############
//	#0123456789a#
//	###b#f#j#n###
//	  #c#g#k#o#
//	  #d#h#l#p#
//	  #e#i#m#q#
//    #########
//
type BoardPart2 []Field

func (b BoardPart2) signature() Signature {
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
		if len(as) != 4 {
			panic("Expected exactly 2 as!")
		}
		if len(bs) != 4 {
			panic("Expected exactly 2 bs!")
		}
		if len(cs) != 4 {
			panic("Expected exactly 2 cs!")
		}
		if len(ds) != 4 {
			panic("Expected exactly 2 ds!")
		}
	}

	sort.Ints(as)
	sort.Ints(bs)
	sort.Ints(cs)
	sort.Ints(ds)

	return Signature{
		a: int64(as[0]<<0+as[1]<<5+
			bs[0]<<10+bs[1]<<15+
			cs[0]<<20+cs[1]<<25) +
			int64(ds[0])<<30 + int64(ds[1])<<35,
		b: int64(as[2]<<0+as[3]<<5+
			bs[2]<<10+bs[3]<<15+
			cs[2]<<20+cs[3]<<25) +
			int64(ds[2])<<30 + int64(ds[3])<<35,
	}
}

func (b BoardPart2) isArranged() bool {
	amphipods := []Field{AMBER, BRONZE, COPPER, DESERT}
	allInRooms := true
	for _, amphipod := range amphipods {
		if b[amphipod.room4()] != amphipod || b[amphipod.room4()+1] != amphipod || b[amphipod.room4()+2] != amphipod || b[amphipod.room4()+3] != amphipod {
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

func NewBoardPart2() *BoardPart2 {
	board := make(BoardPart2, 27)
	return &board
}

func BoardPart2FromPart1(board *BoardPart1) *BoardPart2 {
	newBoard := NewBoardPart2()
	for i := 0; i <= 10; i++ {
		(*newBoard)[i] = (*board)[i]
	}
	(*newBoard)[11] = (*board)[11]
	(*newBoard)[12] = DESERT
	(*newBoard)[13] = DESERT
	(*newBoard)[14] = (*board)[12]

	(*newBoard)[15] = (*board)[13]
	(*newBoard)[16] = COPPER
	(*newBoard)[17] = BRONZE
	(*newBoard)[18] = (*board)[14]

	(*newBoard)[19] = (*board)[15]
	(*newBoard)[20] = BRONZE
	(*newBoard)[21] = AMBER
	(*newBoard)[22] = (*board)[16]

	(*newBoard)[23] = (*board)[17]
	(*newBoard)[24] = AMBER
	(*newBoard)[25] = COPPER
	(*newBoard)[26] = (*board)[18]

	return newBoard
}

func (b BoardPart2) String() string {
	corridor := ""
	for i := 0; i <= 10; i++ {
		corridor += b[i].String()
	}
	return fmt.Sprintf("#############\n#%s#\n  #%s#%s#%s#%s#  \n  #%s#%s#%s#%s#  \n  #%s#%s#%s#%s#  \n  #%s#%s#%s#%s#  \n  #########  ",
		corridor,
		b[AMBER.room4()].String(), b[BRONZE.room4()].String(), b[COPPER.room4()].String(), b[DESERT.room4()].String(),
		b[AMBER.room4()+1].String(), b[BRONZE.room4()+1].String(), b[COPPER.room4()+1].String(), b[DESERT.room4()+1].String(),
		b[AMBER.room4()+2].String(), b[BRONZE.room4()+2].String(), b[COPPER.room4()+2].String(), b[DESERT.room4()+2].String(),
		b[AMBER.room4()+3].String(), b[BRONZE.room4()+3].String(), b[COPPER.room4()+3].String(), b[DESERT.room4()+3].String(),
	)
}

func (b BoardPart2) validMoves() []Move {
	return validMoves(b, 27, 4)
}

func (b BoardPart2) apply(move Move) BoardInterface {
	newBoard := *NewBoardPart2()
	for i := 0; i < 27; i++ {
		newBoard[i] = b[i]
	}
	newBoard[move.to] = newBoard[move.from]
	newBoard[move.from] = EMPTY
	return newBoard
}
