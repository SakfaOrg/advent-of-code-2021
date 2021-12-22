package geometry

import "fmt"

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
}

func (p Point3D) SquaredDistanceTo(other Point3D) int {
	return (p.X-other.X)*(p.X-other.X) + (p.Y-other.Y)*(p.Y-other.Y) + (p.Z-other.Z)*(p.Z-other.Z)
}

func (p Point3D) Minus(other Point3D) Point3D {
	return Point3D{
		X: p.X - other.X,
		Y: p.Y - other.Y,
		Z: p.Z - other.Z,
	}
}

func (p Point3D) Plus(other Point3D) Point3D {
	return Point3D{
		X: other.X + p.X,
		Y: other.Y + p.Y,
		Z: other.Z + p.Z,
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (p Point3D) ManhattanDistanceTo(position Point3D) int {
	return abs(p.X-position.X) + abs(p.Y-position.Y) + abs(p.Z-position.Z)
}

func Rotations3D(p Point3D) []Point3D {
	return []Point3D{
		{p.X, p.Y, p.Z},
		{p.X, -p.Y, -p.Z},
		{-p.X, p.Y, -p.Z},
		{-p.X, -p.Y, p.Z},

		{p.X, p.Z, -p.Y},
		{p.X, -p.Z, p.Y},
		{-p.X, p.Z, p.Y},
		{-p.X, -p.Z, -p.Y},

		{p.Y, p.Z, p.X},
		{p.Y, -p.Z, -p.X},
		{-p.Y, p.Z, -p.X},
		{-p.Y, -p.Z, p.X},

		{p.Y, p.X, -p.Z},
		{p.Y, -p.X, p.Z},
		{-p.Y, p.X, p.Z},
		{-p.Y, -p.X, -p.Z},

		{p.Z, p.X, p.Y},
		{p.Z, -p.X, -p.Y},
		{-p.Z, p.X, -p.Y},
		{-p.Z, -p.X, p.Y},

		{p.Z, p.Y, -p.X},
		{p.Z, -p.Y, p.X},
		{-p.Z, p.Y, p.X},
		{-p.Z, -p.Y, -p.X},
	}
}

func SumPoints(listA, listB []Point3D) (result []Point3D) {
	points := make(map[Point3D]bool)
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

func IntersectPoints(listA, listB []Point3D) []Point3D {
	pointsA := make(map[Point3D]bool)
	for _, point := range listA {
		pointsA[point] = true
	}
	var result []Point3D
	for _, point := range listB {
		if pointsA[point] {
			result = append(result, point)
		}
	}
	return result
}
