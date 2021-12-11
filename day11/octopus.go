package day11

import (
	"fmt"
	"strings"
)

type Octopus struct {
	row int
	col int
	energy uint8
}

type OctopiStack []Octopus

func (s OctopiStack) push(v Octopus) OctopiStack {
	return append(s, v)
}

func (s OctopiStack) empty() bool {
	return len(s) == 0
}

func (s OctopiStack) peek() Octopus {
	return s[len(s)-1]
}

func (s OctopiStack) pop() (OctopiStack, Octopus) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (o Octopus) isFlashing() bool {
	return o.energy == 10
}

func (o Octopus) String() string {
	return fmt.Sprintf("%d,%d (%d)", o.row, o.col, o.energy)
}

func (o *Octopus) excite() (hasFlashed bool) {
	if o.energy < 10 {
		o.energy += 1
		if o.energy == 10 {
			return true
		}
	}
	return false
}

func (o *Octopus) reset() {
	o.energy = 0
}

type Cave struct {
	width int
	height int
	octopi [][]Octopus
}

func NewCave(lines []string) *Cave {
	width := len(lines[0])
	height := len(lines)
	octopi := make([][]Octopus, height)
	for i := 0; i < height; i++ {
		octopi[i] = make([]Octopus, width)
		for j := 0; j < width; j++ {
			octopi[i][j] = Octopus{
				row: i,
				col: j,
				energy: lines[i][j] - '0',
			}
		}
	}
	return &Cave{
		width: width,
		height: height,
		octopi: octopi,
	}
}

func (c Cave) String() string {
	var lines []string
	for i := 0; i < c.height; i++ {
		var line string
		for j := 0; j < c.width; j++ {
			energy := c.octopi[i][j].energy
			if energy == 10 {
				line += "*"
			} else {
				line += fmt.Sprintf("%d", energy)
			}
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func (c Cave) inBounds(row, col int) bool {
	return row >= 0 && row < c.height && col >= 0 && col < c.width
}

func (c Cave) neighboursOf(o Octopus) (neighbours []*Octopus) {
	neighboursCoords := [][]int {
		{o.row - 1, o.col - 1},
		{o.row - 1, o.col},
 		{o.row - 1, o.col + 1},
		{o.row, o.col - 1},
		{o.row, o.col + 1},
		{o.row + 1, o.col - 1},
		{o.row + 1, o.col},
		{o.row + 1, o.col + 1},
	}

	for i := 0; i < len(neighboursCoords); i++ {
		row := neighboursCoords[i][0]
		col := neighboursCoords[i][1]
		if c.inBounds(row, col) {
			neighbours = append(neighbours, &c.octopi[row][col])
		}
	}
	return
}

func tick(c *Cave) (flashedCounter int) {
	//fmt.Printf("cave:\n%s\n\n", c)
	var flashedOctopi OctopiStack
	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			if c.octopi[i][j].excite() {
				flashedOctopi = flashedOctopi.push(c.octopi[i][j])
			}
		}
	}

	for !flashedOctopi.empty() {
		var o Octopus
		flashedOctopi, o = flashedOctopi.pop()
		for _, neighbour := range c.neighboursOf(o) {
			if neighbour.excite() {
				flashedOctopi = flashedOctopi.push(*neighbour)
			}
		}
	}

	for i := 0; i < c.height; i++ {
		for j := 0; j < c.width; j++ {
			if c.octopi[i][j].isFlashing() {
				flashedCounter++
				c.octopi[i][j].reset()
			}
		}
	}

	return
}

func simulate(c *Cave, days int) (flashed int) {
	for day := 0; day < days; day ++ {
		flashed += tick(c)
	}
	return flashed
}

func Part1(lines []string) string {
	flashed := simulate(NewCave(lines), 100)
	return fmt.Sprintf("Flashed after %d days: %d", 100, flashed)
}

func Part2(lines []string) string {
	c := NewCave(lines)

	var flashed int
	var days int
	for ; flashed != 100; days++ {
		flashed = tick(c)
	}

	return fmt.Sprintf("All octopi flashed on day %d", days)
}