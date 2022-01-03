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

func Min64(numbers ...int64) int64 {
	min := int64(math.MaxInt64)
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

func Max64(numbers ...int64) int64 {
	max := int64(math.MinInt64)
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
