package day8

import (
	"fmt"
	"sort"
	"strings"
)

func filterSignalsByLength(signals []string, length int) []string {
	var result []string
	for _, signal := range signals {
		if len(signal) == length {
			result = append(result, signal)
		}
	}
	return result
}

/**
 * finds a signal with given length, panic if ambiguous
 */
func findUniqueSignalByLength(signals []string, length int) string {
	result := filterSignalsByLength(signals, length)
	if len(result) != 1 {
		panic(fmt.Sprintf("Expected only 1 signal with length %d but input %#v has %d!", length, signals, len(result)))
	}
	return result[0]
}

/**
 * finds a signal of length length that has all the wires from mustHaveWires
 */
func findUniqueSignalByLengthWithWires(signals []string, length int, mustHaveWires string) string {
	var result []string
Signal:
	for _, maybeResult := range signals {
		if len(maybeResult) == length {
			for i := 0; i < len(mustHaveWires); i++ {
				wire := mustHaveWires[i : i+1]
				if !strings.Contains(maybeResult, wire) {
					continue Signal
				}
			}
			result = append(result, maybeResult)
		}
	}

	if len(result) != 1 {
		panic(fmt.Sprintf("Expected only 1 signal with length %d and wires %s in input %#v but found %d!", length, mustHaveWires, signals, len(result)))
	}
	return result[0]
}

/**
 * returns the only wire in signal that is not present in knownWires, panics if ambigous
 */
func getMissingWire(signal string, knownWires string) string {
	var result []string
	for i := 0; i < len(signal); i++ {
		wire := signal[i : i+1]
		if !strings.Contains(knownWires, wire) {
			result = append(result, wire)
		}
	}
	if len(result) != 1 {
		panic(fmt.Sprintf("Expected only 1 missing signal in %s minus %s but got %d", signal, knownWires, len(result)))
	}
	return result[0]
}

func inferWires(signals []string) map[string]string {
	// first, find codes for those numbers that are unambiguous:
	oneSignal := findUniqueSignalByLength(signals, 2)
	fourSignal := findUniqueSignalByLength(signals, 4)
	sevenSignal := findUniqueSignalByLength(signals, 3)

	var abcdefgWire = "abcdefg"
	var aWire string
	var bWire string
	var cWire string
	var cfWires string
	var dWire string
	var eWire string
	var fWire string
	var bdWires string
	var gWire string
	for i := 0; i < 7; i++ {
		wire := abcdefgWire[i : i+1]
		if strings.Contains(sevenSignal, wire) && !strings.Contains(oneSignal, wire) {
			aWire = wire
		}
		if strings.Contains(sevenSignal, wire) && strings.Contains(oneSignal, wire) {
			cfWires += wire
		}
		if strings.Contains(fourSignal, wire) && !strings.Contains(oneSignal, wire) {
			bdWires += wire
		}
	}

	// now we can figure out which signal codes "g": it's a signal for 9 minus a, cf and bd
	// signal for 9 has length of 6 and is the only signal of length 6 that contains a,c,f,b and d:
	acfbdWire := aWire + cfWires + bdWires
	nineSignal := findUniqueSignalByLengthWithWires(signals, 6, acfbdWire)
	gWire = getMissingWire(nineSignal, acfbdWire)

	// then we can disambiguate bd wires: d is in '3' signal and 3 signal is the only signal of length 5 that contains acfg
	acfgWire := aWire + cfWires + gWire
	threeSignal := findUniqueSignalByLengthWithWires(signals, 5, acfgWire)
	dWire = getMissingWire(threeSignal, acfgWire)
	bWire = getMissingWire(bdWires, dWire)

	// now use similar logic to disambiguate cf wires using 5 signal:
	abdgWire := aWire + bWire + dWire + gWire
	fiveSignal := findUniqueSignalByLengthWithWires(signals, 5, abdgWire)
	fWire = getMissingWire(fiveSignal, abdgWire)
	cWire = getMissingWire(cfWires, fWire)

	// and the final wire: e
	eWire = getMissingWire(abcdefgWire, aWire+bWire+cWire+dWire+fWire+gWire)

	result := make(map[string]string)
	result[aWire] = "a"
	result[bWire] = "b"
	result[cWire] = "c"
	result[dWire] = "d"
	result[eWire] = "e"
	result[fWire] = "f"
	result[gWire] = "g"
	return result
}

func decodeDigit(seenWires string, wiresMap map[string]string) int {
	actualWiresStr := ""
	for i := 0; i < len(seenWires); i++ {
		actualWiresStr += wiresMap[seenWires[i:i+1]]
	}

	actualWiresSlice := strings.Split(actualWiresStr, "")
	sort.Strings(actualWiresSlice)
	actualWiresSorted := strings.Join(actualWiresSlice, "")

	sevenSegmentCodes := make(map[string]int)
	sevenSegmentCodes["abcefg"] = 0
	sevenSegmentCodes["cf"] = 1
	sevenSegmentCodes["acdeg"] = 2
	sevenSegmentCodes["acdfg"] = 3
	sevenSegmentCodes["bcdf"] = 4
	sevenSegmentCodes["abdfg"] = 5
	sevenSegmentCodes["abdefg"] = 6
	sevenSegmentCodes["acf"] = 7
	sevenSegmentCodes["abcdefg"] = 8
	sevenSegmentCodes["abcdfg"] = 9

	digit, has := sevenSegmentCodes[actualWiresSorted]
	if !has {
		panic("Unrecognized sequence " + actualWiresSorted)
	}
	return digit
}

func solveLine(line string) int {
	signalsAndDigits := strings.Split(line, " | ")
	signals := strings.Split(signalsAndDigits[0], " ")
	wiresMap := inferWires(signals)

	digits := strings.Split(signalsAndDigits[1], " ")
	result := 0
	position := 1000
	for _, digitCode := range digits {
		digit := decodeDigit(digitCode, wiresMap)
		result += position * digit
		position = position / 10
	}

	return result
}

func Part2(lines []string) string {
	result := 0
	for _, line := range lines {
		result += solveLine(line)
	}
	return fmt.Sprintf("sum of values: %d", result)
}

func Part1(lines []string) string {
	counter := 0
	for _, line := range lines {
		splitted := strings.Split(line, " | ")
		for _, digits := range strings.Split(splitted[1], " ") {
			if (len(digits) >= 2 && len(digits) <= 4) || len(digits) == 7 {
				counter++
			}
		}
	}

	return fmt.Sprintf("unique digits=%d", counter)
}
