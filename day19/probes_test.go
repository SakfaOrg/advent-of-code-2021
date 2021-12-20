package day19

import (
	"advent2021/day19/geometry"
	advent2 "advent2021/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestProbes(t *testing.T) {
	t.Run("Test probe read and rotates", func(t *testing.T) {
		scanner := readScanner(strings.Split("--- scanner 0 ---\n-1,-1,1\n-2,-2,2\n-3,-3,3\n-2,-3,1\n5,6,-4\n8,0,7", "\n"))

		for _, s := range scanner.scannerRotations() {
			fmt.Println(s)
		}
	})

	t.Run("Test manhattan distance", func(t *testing.T) {
		pointsA := geometry.Point3D{1105, -1205, 1229}
		pointsB := geometry.Point3D{-92, -2380, -20}
		assert.Equal(t, 3621, pointsA.ManhattanDistanceTo(pointsB))
	})

	t.Run("Test demo report", func(t *testing.T) {
		lines := advent2.MustReadLines("demoReport")
		var scanners []Scanner
		for pos := 0; pos < len(lines); {
			var scanner Scanner
			scanner, pos = readScannerFrom(lines, pos)
			scanners = append(scanners, scanner)
		}

		assert.Equal(t, 5, len(scanners))
		matchedScanners := match(scanners)
		points := matchedScanners[0].seenProbes
		fmt.Printf("Probe %d sees %d points\n", matchedScanners[0].id, len(points))
		for i := 1; i < len(matchedScanners); i++ {
			points = geometry.SumPoints(points, matchedScanners[i].seenProbes)
			fmt.Printf("+ with probe %d we have %d points in total\n", matchedScanners[i].id, len(points))
		}

		assert.Equal(t, 79, len(points))
	})
}