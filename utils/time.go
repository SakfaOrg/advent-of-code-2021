package advent

import (
	"fmt"
	"time"
)

func Timed(name string, warmRetries int, code func() string) (timeTaken time.Duration) {
	startedAt := time.Now()
	result := code()
	timeTakenCold := time.Now().Sub(startedAt)

	startedAt = time.Now()
	if warmRetries > 0 {
		tries := warmRetries
		for i := 0; i < tries; i++ {
			retry := code()
			if retry != result {
				fmt.Printf("First result:\n%s\n%d th try:\n%s\n", result, i, retry)
				panic("Code doesn't give stable results! See above")
			}
		}

		timeTakenWarm := time.Now().Sub(startedAt) / 10

		fmt.Printf("%s: %s (took %s for first, then %s on average for next %d)\n", name, result, timeTakenCold, timeTakenWarm, warmRetries)
	} else {
		fmt.Printf("%s: %s (took %s for first)\n", name, result, timeTakenCold)
	}

	return timeTakenCold
}
