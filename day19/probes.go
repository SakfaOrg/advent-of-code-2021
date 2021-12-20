package day19

import (
	"advent2021/day19/geometry"
	advent "advent2021/utils"
	"fmt"
	"regexp"
	"strings"
)

type Scanner struct {
	id int
	position geometry.Point3D
	seenProbes []geometry.Point3D
}

func (s Scanner) scannerRotations() []Scanner {
	rotatedScanners := make([]Scanner, 24)
	for i := 0; i < 24; i++ {
		rotatedScanners[i] = Scanner{
			id: s.id,
			position: s.position,
			seenProbes: make([]geometry.Point3D, len(s.seenProbes)),
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
	moved := Scanner {
		id: s.id,
		position: s.position.Plus(vector),
		seenProbes: make([]geometry.Point3D, len(s.seenProbes)),
	}

	for idx, probe := range s.seenProbes {
		moved.seenProbes[idx] = probe.Minus(vector)
	}
	return moved
}

func sumPoints(listA, listB []geometry.Point3D) (result []geometry.Point3D) {
	points := make(map[geometry.Point3D]bool)
	for _, point := range listA {
		points[point] = true
	}
	for _, point := range listB {
		points[point] = true
	}
	for point, _ := range points {
		result = append(result, point)
	}
	return result
}

func intersectPoints (listA, listB []geometry.Point3D) []geometry.Point3D {
	pointsA := make(map[geometry.Point3D]bool)
	for _, point := range listA {
		pointsA[point] = true
	}
	var result []geometry.Point3D
	for _, point := range listB {
		if pointsA[point] {
			result = append(result, point)
		}
	}
	return result
}

func (s Scanner) String() string {
	lines := []string {
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
		id: advent.MustAtoi(submatches[1]),
		position: geometry.Point3D{0, 0, 0},
	}

	for pos = startPos + 1; pos < len(lines) && lines[pos] != ""; pos++ {
		probe := strings.Split(lines[pos], ",")
		scanner.seenProbes = append(scanner.seenProbes, geometry.Point3D{
			advent.MustAtoi(probe[0]),
			advent.MustAtoi(probe[1]),
			advent.MustAtoi(probe[2]),
		})
	}

	if pos < len(lines) {
		pos += 1
	}

	return
}

func match(scanners []Scanner) []Scanner {
	knownIds := make(map[int]bool)
	knownList := []Scanner { scanners[0] }
	knownIds[scanners[0].id] = true

	leftList := scanners[1:]
	for len(knownList) != len(scanners) {
		KNOWN: for _, known := range knownList {
			for _, left := range leftList {
				if _, ok := knownIds[left.id]; !ok {
					matched := matchScanners(known, left)
					if matched != nil {
						knownList = append(knownList, *matched)
						knownIds[matched.id] = true
						fmt.Printf("Scanner %d matched to %d!\n", matched.id, known.id)
						break KNOWN
					}
				}
			}
		}
	}
	fmt.Printf("All scanners matched.\n")
	return knownList
}

func matchScanners(scanner Scanner, matchee Scanner) *Scanner {
	for i := 0; i < len(scanner.seenProbes); i++ {
		for j := i + 1; j < len(scanner.seenProbes); j++ {
			probeA := scanner.seenProbes[i]
			probeB := scanner.seenProbes[j]
			distance := probeA.SquaredDistanceTo(probeB)

			for x := 0; x < len(matchee.seenProbes); x++ {
				for y := x + 1; y < len(matchee.seenProbes); y++ {
					matcheeA := matchee.seenProbes[x]
					matcheeB := matchee.seenProbes[y]
					matcheeDistance := matcheeA.SquaredDistanceTo(matcheeB)
					if distance == matcheeDistance {
						for _, rotation := range matchee.scannerRotations() {
							moved := rotation.moveBy(rotation.seenProbes[x].Minus(probeA))
							movedA := moved.seenProbes[x]
							movedB := moved.seenProbes[y]

							if movedA != probeA {
								panic("Uh-oh, we moved points to equal but they don't...")
							}
							if movedB == probeB {
								commonPoints := intersectPoints(scanner.seenProbes, moved.seenProbes)
								if len(commonPoints) >= 12 {
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

var part1Solution *[]Scanner

func Part1(lines []string) string {
	var scanners []Scanner
	for pos := 0; pos < len(lines); {
		var scanner Scanner
		scanner, pos = readScannerFrom(lines, pos)
		scanners = append(scanners, scanner)
	}

	matchedScanners := match(scanners)
	part1Solution = &matchedScanners
	points := matchedScanners[0].seenProbes
	for i := 1; i < len(matchedScanners); i++ {
		points = sumPoints(points, matchedScanners[i].seenProbes)
	}

	return fmt.Sprintf("All scanners see %d points in total\n", len(points))
}

func Part2(lines []string) string {
	var matchedScanners []Scanner
	var message string

	if part1Solution != nil {
		message = "(reused part1 solution)"
		matchedScanners = *part1Solution
	} else {
		message = "(computed from scratch)"
		var scanners []Scanner
		for pos := 0; pos < len(lines); {
			var scanner Scanner
			scanner, pos = readScannerFrom(lines, pos)
			scanners = append(scanners, scanner)
		}

		matchedScanners = match(scanners)
	}

	points := matchedScanners[0].seenProbes
	for i := 1; i < len(matchedScanners); i++ {
		points = sumPoints(points, matchedScanners[i].seenProbes)
	}

	maxDistance := 0
	for i := 0; i < len(matchedScanners); i++ {
		for j := i + 1; j < len(matchedScanners); j++ {
			distance := matchedScanners[i].position.ManhattanDistanceTo(matchedScanners[j].position)
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return fmt.Sprintf("Furthest scanners are %d apart %s\n", maxDistance, message)
}