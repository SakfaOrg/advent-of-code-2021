package day19

import (
	advent "advent2021/utils"
	"advent2021/utils/geometry"
	"fmt"
	"regexp"
	"strings"
)

type Scanner struct {
	id         int
	position   geometry.Point3D
	seenProbes []geometry.Point3D
	// maintain a set of all distances between points. Why? This allows us to optimize a lot: if
	// scanner sees some set of points which will be in some distance from each other, then if other scanner sees the
	// same set of points it will see overlapping 1+2+3+...+n same distances. This happens regardless of orientation!
	// so by quickly checking how many same distances we see, if there's less than 66 we know that scanners definitely
	// don't see set of same 12 points! (66 = 1+2+3+...+9+10+11)
	distances *map[int]int
}

func (s Scanner) scannerRotations() []Scanner {
	rotatedScanners := make([]Scanner, 24)
	for i := 0; i < 24; i++ {
		rotatedScanners[i] = Scanner{
			id:         s.id,
			position:   s.position,
			seenProbes: make([]geometry.Point3D, len(s.seenProbes)),
			distances:  s.distances,
		}
	}

	for idx, probe := range s.seenProbes {
		rotations := geometry.Rotations3D(probe)
		for i := 0; i < 24; i++ {
			rotatedScanners[i].seenProbes[idx] = rotations[i]
		}
	}

	return rotatedScanners
}

func (s Scanner) moveBy(vector geometry.Point3D) Scanner {
	moved := Scanner{
		id:         s.id,
		position:   s.position.Plus(vector),
		seenProbes: make([]geometry.Point3D, len(s.seenProbes)),
		distances:  s.distances,
	}

	for idx, probe := range s.seenProbes {
		moved.seenProbes[idx] = probe.Minus(vector)
	}
	return moved
}

func (s Scanner) String() string {
	lines := []string{
		fmt.Sprintf("--- scanner %02d (at %s)---", s.id, s.position),
	}
	for _, probe := range s.seenProbes {
		lines = append(lines, probe.String())
	}
	return strings.Join(lines, "\n")
}

func readScanner(lines []string) Scanner {
	scanner, _ := readScannerFrom(lines, 0)
	return scanner
}

func readScannerFrom(lines []string, startPos int) (scanner Scanner, pos int) {
	regex := *regexp.MustCompile("--- scanner (\\d+) ---")
	submatches := regex.FindStringSubmatch(lines[startPos])
	if submatches == nil {
		panic("first line doesn't start with 'scanner'")
	}

	scanner = Scanner{
		id:       advent.MustAtoi(submatches[1]),
		position: geometry.Point3D{0, 0, 0},
	}

	for pos = startPos + 1; pos < len(lines) && lines[pos] != ""; pos++ {
		probe := strings.Split(lines[pos], ",")
		scanner.seenProbes = append(scanner.seenProbes, geometry.Point3D{
			X: advent.MustAtoi(probe[0]),
			Y: advent.MustAtoi(probe[1]),
			Z: advent.MustAtoi(probe[2]),
		})
	}

	distancesMap := make(map[int]int)
	for i := 0; i < len(scanner.seenProbes); i++ {
		for j := i + 1; j < len(scanner.seenProbes); j++ {
			distance := scanner.seenProbes[i].SquaredDistanceTo(scanner.seenProbes[j])
			distancesMap[distance] += 1
		}
	}
	scanner.distances = &distancesMap

	if pos < len(lines) {
		pos += 1
	}

	return
}

type Pair struct{ left, right int }

func (p Pair) normalize() Pair {
	if p.left > p.right {
		return Pair{p.right, p.left}
	}
	return p
}

func match(scanners []Scanner) []Scanner {
	// we will do a lot of matching of scanners, we can pre-calculate rotations to make this slightly faster
	scannerToRotations := make(map[int][]Scanner)
	for i := 1; i < len(scanners); i++ {
		scannerToRotations[scanners[i].id] = scanners[i].scannerRotations()
	}

	// because of how we loop over known,left pairs all the time it may happen that we will try to match scanner in
	// next iteration that we already checked in previous. Remember which pairs we checked already
	pairsThatDontMatch := make(map[Pair]bool)

	knownIds := make(map[int]bool)
	knownList := []Scanner{scanners[0]}
	knownIds[scanners[0].id] = true

	leftList := scanners[1:]
	for len(knownList) != len(scanners) {
	KNOWN:
		for _, knownCandidate := range knownList {
			for _, leftCandidate := range leftList {
				if _, ok := knownIds[leftCandidate.id]; !ok {
					checkedPair := Pair{knownCandidate.id, leftCandidate.id}.normalize()
					if _, dontMatch := pairsThatDontMatch[checkedPair]; !dontMatch {
						matched := matchScanners(knownCandidate, leftCandidate, scannerToRotations[leftCandidate.id])
						if matched != nil {
							knownList = append(knownList, *matched)
							knownIds[matched.id] = true
							break KNOWN
						} else {
							pairsThatDontMatch[checkedPair] = true
						}
					}
				}
			}
		}
	}
	return knownList
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countSameDistances(s1, s2 Scanner) int {
	result := 0
	for distance, count := range *s1.distances {
		if otherCount, ok := (*s2.distances)[distance]; ok {
			result += min(count, otherCount)
		}
	}
	return result
}

func matchScanners(fixed Scanner, other Scanner, otherRotations []Scanner) *Scanner {
	// if these scanner saw the same set of 12 points they would see set of same 66 (1+2+...+10+11) pairs of points that
	// are some distance apart from each other. No 66 same distances? Then clearly these don't match.
	if countSameDistances(fixed, other) < 66 {
		return nil
	}

	// matching algorithm works as follow:
	//
	// find some pair of points seen by fixed scanner, we will check ALL possible pairs of points it sees
	for i := 0; i < len(fixed.seenProbes); i++ {
		for j := i + 1; j < len(fixed.seenProbes); j++ {
			fixedA := fixed.seenProbes[i]
			fixedB := fixed.seenProbes[j]
			// calculate distance between these 2 points. Why? Because if these 2 points are seen by other scanner
			// they will be **the same** distance apart. The idea is, moving or rotating scanner doesn't change
			// relative distances between points
			distance := fixedA.SquaredDistanceTo(fixedB)

			// now check all point pairs in other scanner, we're looking for 2 points that are distance apart of each other
			for x := 0; x < len(other.seenProbes); x++ {
				for y := x + 1; y < len(other.seenProbes); y++ {
					matcheeDistance := other.seenProbes[x].SquaredDistanceTo(other.seenProbes[y])
					if distance == matcheeDistance { // gotcha!
						// so now we will align that other scanner: since we don't know it's exact orientation we will
						// check all 24 of possible rations
						for _, rotation := range otherRotations {
							// how to check scanners match? Grab 2 points from the pair. Move scanner in such a
							// way that 1 of that points will align with one of the points from fixed scanner.
							// did other point align as well? Then it seems we have a potential match!
							moveVector := rotation.seenProbes[x].Minus(fixedA)
							movedB := rotation.seenProbes[y].Minus(moveVector)

							if movedB == fixedB { // gotcha! We now found a pair of points that in this specific position
								// and rotation of scanner is seen in the same coordinates relative to
								// scanner 0
								// the final step is to move ALL points and see how many of them are the same.
								moved := rotation.moveBy(moveVector)
								commonPoints := geometry.IntersectPoints(fixed.seenProbes, moved.seenProbes)
								if len(commonPoints) >= 12 { // if 12 we've got a match! (according to puzzle authors at least 12 points will overlap)
									return &moved
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func solve(lines []string) []Scanner {
	var scanners []Scanner
	for pos := 0; pos < len(lines); {
		var scanner Scanner
		scanner, pos = readScannerFrom(lines, pos)
		scanners = append(scanners, scanner)
	}

	return match(scanners)
}

func Part1(lines []string) string {
	matchedScanners := solve(lines)

	points := matchedScanners[0].seenProbes
	for i := 1; i < len(matchedScanners); i++ {
		points = geometry.SumPoints(points, matchedScanners[i].seenProbes)
	}

	return fmt.Sprintf("All scanners see %d points in total\n", len(points))
}

func Part2(lines []string) string {
	matchedScanners := solve(lines)

	maxDistance := 0
	for i := 0; i < len(matchedScanners); i++ {
		for j := i + 1; j < len(matchedScanners); j++ {
			distance := matchedScanners[i].position.ManhattanDistanceTo(matchedScanners[j].position)
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return fmt.Sprintf("Furthest scanners are %d apart\n", maxDistance)
}
