package day17

import (
	advent "advent2021/utils"
	"fmt"
	"regexp"
)

type Range struct {
	from int
	to   int
}

func findAll(xr, yr Range) int {
	validPairs := 0
	for x := 1; x <= xr.to; x++ {
		for y := yr.from; y <= -1*yr.from; y++ {
			reachedTarget := false
			for step := 1; ; step++ {
				targetX, targetY := calcTarget(x, y, step)
				if targetX >= xr.from && targetX <= xr.to && targetY >= yr.from && targetY <= yr.to {
					reachedTarget = true
					break
				}
				if targetY < yr.from {
					break
				}
			}
			if reachedTarget {
				validPairs += 1
			}
		}
	}
	return validPairs
}

func calcTarget(x, y, step int) (targetX, targetY int) {
	targetX = (((1 + x) * x) / 2)
	left := x - step
	if left > 0 {
		targetX -= (1 + left) * left / 2
	}
	targetY = (y + y - step + 1) * step / 2
	return
}

func solve(xr, yr Range) string {
	// the solution will most likely involve a vertical fall at the end. Vertical fall will always happen at x = 3+2+1+0+0+0 ...
	// so the final position x, assuming solution finishes with vertical fall, is some number that is a result of (1+n)*n/2
	// equation for natural ns. We also know that number must be somewhere in target bound
	var xs []int
	for n := 1; ; n++ {
		verticalFallAt := ((1 + n) * n) / 2
		if verticalFallAt > xr.to {
			break
		} else if verticalFallAt >= xr.from {
			xs = append(xs, n)
		}
	}

	type shot struct{ x, y, step int }
	highestShot := shot{0, 0, 0}
	for _, x := range xs {
		// now note that x is both starting velocity and also minimum amount of steps after which we achieve vertical fall
		// for each x let's check bunch of y. Just as we assumed that the valid solution will end with vertical fall, I'm
		// assuming that it will end up with y > 0
		for y := 1; y <= -1*yr.from; y++ { // check at most until yr.from, any y higher than that and we will overshoot target after first step since 0
			// y always starts to go up, then evetually it starts to fall and at some point will get back to position 0
			// (why? becuse it starts with i.e., 1,0,-1,...  or 2,1,0,-1,-2,..., you get the drill first few steps cancel
			// each other). We will have to do _at least_ that amount of steps, assuming the target is negative (it is
			// in the puzzle input)
			minimumSteps := y*2 + 1

			// now that we know minimum amount of steps we can see if we will land in
			reachedTarget := false
			reachedTargetStep := 0
			for step := minimumSteps + 1; ; step++ {
				targetX, targetY := calcTarget(x, y, step)
				if targetX >= xr.from && targetX <= xr.to && targetY >= yr.from && targetY <= yr.to {
					reachedTarget = true
					reachedTargetStep = step
					break
				}
				if targetY < yr.from {
					break
				}
			}

			if reachedTarget {
				if y > highestShot.y {
					highestShot = shot{x, y, reachedTargetStep}
				}
			}
		}
	}

	return fmt.Sprintf("Highest shot x=%d,y=%d reaches target after %d steps with y high=%d",
		highestShot.x, highestShot.y, highestShot.step, ((1+highestShot.y)*highestShot.y)/2)
}

func parseInput(lines []string) (xRange, yRange Range) {
	regex := *regexp.MustCompile("target area: x=(\\d+)..(\\d+), y=-(\\d+)..-(\\d+)")
	submatches := regex.FindStringSubmatch(lines[0])
	xRange = Range{advent.MustAtoi(submatches[1]), advent.MustAtoi(submatches[2])}
	yRange = Range{-1 * advent.MustAtoi(submatches[3]), -1 * advent.MustAtoi(submatches[4])}
	return
}

func Part1(lines []string) string {
	xRange, yRange := parseInput(lines)
	return solve(xRange, yRange)
}

func Part2(lines []string) string {
	xRange, yRange := parseInput(lines)
	combinations := findAll(xRange, yRange)
	return fmt.Sprintf("There are %d distinct velocity values to reach target", combinations)
}
