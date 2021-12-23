package advent

import "math"

func Min(numbers ...int) int {
	min := math.MaxInt
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}

func Max(numbers ...int) int {
	max := math.MinInt
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
