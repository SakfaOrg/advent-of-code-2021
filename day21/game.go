package day21

import (
	advent "advent2021/utils"
	"fmt"
	"strings"
)

type Die interface {
	roll() int
}

type DeterministicDie struct {
	sides int
	value int
	rolls int
}

type Player struct {
	number, position, score int
}

// CompactPlayer
// in Part2 position and score is limited to 10 an 21 respectively. We are remembering a lot of game states, we can
// compact this structure a lot by using uint8.
// Separate structure cause part1 goes all the way to score = 1000 so it needs larger structs
type CompactPlayer struct {
	number, position, score uint8
}

func (p Player) compact() CompactPlayer {
	return CompactPlayer{
		number:   uint8(p.number),
		position: uint8(p.position),
		score:    uint8(p.score),
	}
}

func NewDeterministicDie(sides int) Die {
	return &DeterministicDie{
		sides: sides,
		value: 1,
	}
}

func (d *DeterministicDie) roll() int {
	result := d.value
	d.rolls += 1
	if d.value == d.sides {
		d.value = 1
	} else {
		d.value = d.value + 1
	}
	return result
}

func parsePlayer(line string) Player {
	segments := strings.Split(line, " ")
	return Player{
		number:   advent.MustAtoi(segments[1]),
		position: advent.MustAtoi(segments[len(segments)-1]),
		score:    0,
	}
}

func makeMove(die Die, player *Player, winningScore int) (won bool) {
	return applyRoll(die.roll()+die.roll()+die.roll(), player, winningScore)
}

func applyRoll(rolled int, player *Player, winningScore int) (won bool) {
	nextPosition := (player.position-1+rolled)%10 + 1
	player.position = nextPosition
	player.score += player.position
	return player.score >= winningScore
}

func applyRollCompact(rolled uint8, player *CompactPlayer, winningScore uint8) (won bool) {
	nextPosition := (player.position-1+rolled)%10 + 1
	player.position = nextPosition
	player.score += player.position
	return player.score >= winningScore
}

type GameState struct {
	next, other CompactPlayer
}

func (gs *GameState) isFinished() bool {
	return gs.next.score >= 21 || gs.other.score >= 21
}

func (gs *GameState) applyRoll(roll uint8) GameState {
	nextCopy := gs.next
	otherCopy := gs.other
	applyRollCompact(roll, &nextCopy, 21)
	return GameState{next: otherCopy, other: nextCopy}
}

func Part2(lines []string) string {
	universeCounter := make(map[GameState]int)
	player1 := parsePlayer(lines[0]).compact()
	player2 := parsePlayer(lines[1]).compact()
	universeCounter[GameState{next: player1, other: player2}] = 1

	// each player will spawn 27 universes (3 rolls each spawning 3 universes - 3^3) but
	// there are only 7 unique sums of rolls (3 to 9). Map below maps roll outcome -> number of universes with that outcome
	diracDieCounts := make(map[uint8]int)
	for roll1 := uint8(1); roll1 <= 3; roll1++ {
		for roll2 := uint8(1); roll2 <= 3; roll2++ {
			for roll3 := uint8(1); roll3 <= 3; roll3++ {
				diracDieCounts[roll1+roll2+roll3] += 1
			}
		}
	}

	var loop int
	for loop = 0; ; loop++ {
		pendingGames := make(map[GameState]int)
		for gameState, counter := range universeCounter {
			if !gameState.isFinished() {
				pendingGames[gameState] = counter
				delete(universeCounter, gameState)
			}
		}

		if len(pendingGames) == 0 {
			break
		}

		for gameState, universesInThisState := range pendingGames {
			for roll, rolls := range diracDieCounts {
				nextState := gameState.applyRoll(roll)
				universeCounter[nextState] += rolls * universesInThisState
			}
		}
	}

	playerVictories := make(map[uint8]int)
	for gs, counter := range universeCounter {
		if gs.next.score >= 21 {
			playerVictories[gs.next.number] += counter
		} else {
			playerVictories[gs.other.number] += counter
		}
	}

	return fmt.Sprintf("done after %d loops, player 1 victories %d, player 2 victories %d", loop, playerVictories[1], playerVictories[2])
}

func Part1(lines []string) string {
	player1 := parsePlayer(lines[0])
	player2 := parsePlayer(lines[1])
	die := NewDeterministicDie(100)

	var current *Player
	var other *Player
	var loser *Player

	current = &player1
	other = &player2
	for {
		if makeMove(die, current, 1000) {
			loser = other
			break
		}

		tmp := current
		current = other
		other = tmp
	}

	return fmt.Sprintf("Loser score %d, rolls %d, result = %d", loser.score, die.(*DeterministicDie).rolls,
		loser.score*die.(*DeterministicDie).rolls)
}
