package day22

import (
	advent "advent2021/utils"
	"fmt"
	"sort"
	"strings"
)

type Action uint8

const (
	ActionOff Action = iota
	ActionOn
)

type Range struct {
	from, to int
}

func (r Range) normalize() Range {
	if r.from > r.to {
		return Range{r.to, r.from}
	}
	return r
}

type Cube struct {
	xRange, yRange, zRange Range
}

func (c Cube) normalize() Cube {
	return Cube{
		xRange: c.xRange.normalize(),
		yRange: c.yRange.normalize(),
		zRange: c.zRange.normalize(),
	}
}

func (c Cube) String() string {
	return fmt.Sprintf("%s,%s,%s", c.xRange, c.yRange, c.zRange)
}

func (r Range) String() string {
	return fmt.Sprintf("%d..%d", r.from, r.to)
}

func parseRange(str string) Range {
	split := strings.Split(strings.Split(str, "=")[1], "..")
	return Range{
		from: advent.MustAtoi(split[0]),
		to:   advent.MustAtoi(split[1]),
	}
}

func parseCube(line string) (cube Cube) {
	var rest string
	if line[0:2] == "on" {
		rest = line[3:]
	} else if line[0:3] == "off" {
		rest = line[4:]
	} else {
		panic("Expected cube to start with 'on' or 'off'")
	}

	chunks := strings.Split(rest, ",")
	cube = Cube{
		xRange: parseRange(chunks[0]),
		yRange: parseRange(chunks[1]),
		zRange: parseRange(chunks[2]),
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type CoordSet map[int]bool

func (cs CoordSet) add(r Range) {
	cs[r.from] = true
	cs[r.to+1] = true
}
func (cs CoordSet) uniqueAndSorted() []int {
	var result []int
	for k, _ := range cs {
		result = append(result, k)
	}
	sort.Ints(result)
	return result
}

func idxOf(is []int, i int) int {
	for idx, value := range is {
		if value == i {
			return idx
		}
	}
	panic(fmt.Sprintf("Value %d not found in coords!", i))
}

func Part2(lines []string) string {
	xsSet := make(CoordSet)
	ysSet := make(CoordSet)
	zsSet := make(CoordSet)
	for _, line := range lines {
		cube := parseCube(line)
		xsSet.add(cube.xRange)
		ysSet.add(cube.yRange)
		zsSet.add(cube.zRange)
	}

	xs := xsSet.uniqueAndSorted()
	ys := ysSet.uniqueAndSorted()
	zs := zsSet.uniqueAndSorted()

	space := make([][][]bool, len(xs))
	for x := 0; x < len(xs); x++ {
		space[x] = make([][]bool, len(ys))
		for y := 0; y < len(ys); y++ {
			space[x][y] = make([]bool, len(zs))
		}
	}

	for _, line := range lines {
		cube := parseCube(line).normalize()
		counter := 0
		for xi := idxOf(xs, cube.xRange.from); xs[xi] <= cube.xRange.to; xi++ {
			for yi := idxOf(ys, cube.yRange.from); ys[yi] <= cube.yRange.to; yi++ {
				for zi := idxOf(zs, cube.zRange.from); zs[zi] <= cube.zRange.to; zi++ {
					counter++
					space[xi][yi][zi] = line[0:2] == "on"
				}
			}
		}
		//fmt.Printf("Matching cube %s: mutated %d spaces.\n", cube, counter)
	}

	onCounter := 0
	for x := 0; x < len(xs); x++ {
		for y := 0; y < len(ys); y++ {
			for z := 0; z < len(zs); z++ {
				if space[x][y][z] {
					w := xs[x+1] - xs[x]
					h := ys[y+1] - ys[y]
					d := zs[z+1] - zs[z]
					//fmt.Printf("Space %d..%d, %d..%d, %d..%d, volume=%d*%d*%d=%d\n", xs[x], xs[x+1]-1, ys[y], ys[y+1]-1, zs[z], zs[z+1]-1,
					//	w,h,d,w*h*d)
					onCounter += w * h * d
				}
			}
		}
	}

	return fmt.Sprintf("Splitted space size := %d * %d * %d, pixels on: %d", len(xsSet), len(ysSet), len(zsSet), onCounter)
}

func Part1(lines []string) string {
	var pixels [101][101][101]bool
	for _, line := range lines {
		cube := parseCube(line)
		for x := max(cube.xRange.normalize().from, -50); x <= min(cube.xRange.normalize().to, 50); x++ {
			for y := max(cube.yRange.normalize().from, -50); y <= min(cube.yRange.normalize().to, 50); y++ {
				for z := max(cube.zRange.normalize().from, -50); z <= min(cube.zRange.normalize().to, 50); z++ {
					pixels[x+50][y+50][z+50] = line[0:2] == "on"
				}
			}
		}
	}

	counter := 0
	for _, pixels1 := range pixels {
		for _, pixels2 := range pixels1 {
			for _, pixels3 := range pixels2 {
				if pixels3 {
					counter += 1
				}
			}
		}
	}

	return fmt.Sprintf("pixels on: %d", counter)
}
