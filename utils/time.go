package advent

import (
	"fmt"
	"time"
)

func Timed(name string, code func() string) {
	startedAt := time.Now()
	result := code()

	timeTaken := time.Now().Sub(startedAt)
	fmt.Printf("%s: %s (took %s)\n", name, result, timeTaken)
}