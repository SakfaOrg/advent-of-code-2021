package advent

import (
	"fmt"
	"time"
)

func Timed(name string, code func() string) (timeTaken time.Duration) {
	startedAt := time.Now()
	result := code()
	timeTakenCold := time.Now().Sub(startedAt)

	startedAt = time.Now()
	tries := 3
	for i := 0; i < tries; i++ {
		retry := code()
		if retry != result {
			fmt.Printf("First result:\n%s\n%d th try:\n%s\n", result, retry)
			panic("Code doesn't give stable results! See above")
		}
	}

	timeTakenWarm := time.Now().Sub(startedAt) / 10

	fmt.Printf("%s: %s (took %s for first, then %s on average for next 10)\n", name, result, timeTakenCold, timeTakenWarm)
	return timeTakenCold
}
