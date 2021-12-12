package day10

import (
	"fmt"
	"sort"
)

type stack []rune

func (s stack) push(v rune) stack {
	return append(s, v)
}

func (s stack) empty() bool {
	return len(s) == 0
}

func (s stack) peek() rune {
	return s[len(s)-1]
}

func (s stack) pop() (stack, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func isOpeningBracket(v rune) bool {
	return v == '(' || v == '[' || v == '{' || v == '<'
}

func isClosingBracket(v rune) bool {
	return v == ')' || v == ']' || v == '}' || v == '>'
}

func processLine(line string) (message string, syntaxScore, finishScore int) {
	message = "OK"
	syntaxScore, finishScore = 0, 0

	var symbols stack
	bracketsMap := make(map[rune]rune)
	bracketsMap['('] = ')'
	bracketsMap['['] = ']'
	bracketsMap['{'] = '}'
	bracketsMap['<'] = '>'

	pointsMap := make(map[rune]int)
	pointsMap[')'] = 3
	pointsMap[']'] = 57
	pointsMap['}'] = 1197
	pointsMap['>'] = 25137

	closingPointsMap := make(map[rune]int)
	closingPointsMap[')'] = 1
	closingPointsMap[']'] = 2
	closingPointsMap['}'] = 3
	closingPointsMap['>'] = 4

	for _, symbol := range line {
		if isOpeningBracket(symbol) {
			symbols = symbols.push(symbol)
		} else if isClosingBracket(symbol) {
			if symbols.empty() {
				message = fmt.Sprintf("Syntax error: found closing bracket '%s' but nothing was opened.", string(symbol))
				return
			}

			var opened rune
			symbols, opened = symbols.pop()
			expected := bracketsMap[opened]
			if expected != symbol {
				syntaxScore = pointsMap[symbol]
				message = fmt.Sprintf("Syntax error: Expected %s, but found %s instead.", string(expected), string(symbol))
				return
			}
		} else {
			message = fmt.Sprintf("Syntax error: Unrecognized symbol '%s'.", string(symbol))
			return
		}
	}

	for !symbols.empty() {
		var symbol rune
		symbols, symbol = symbols.pop()
		closingSymbol := bracketsMap[symbol]
		finishScore = finishScore*5 + closingPointsMap[closingSymbol]
	}

	return
}

func Part1(lines []string) string {
	points := 0
	for _, line := range lines {
		_, syntaxScore, _ := processLine(line)
		points += syntaxScore
	}
	return fmt.Sprintf("Syntax error score: %d", points)
}

func Part2(lines []string) string {
	var scores []int
	for _, line := range lines {
		_, syntaxScore, finishScore := processLine(line)
		if syntaxScore == 0 {
			scores = append(scores, finishScore)
		}
	}

	sort.Ints(scores)
	return fmt.Sprintf("Middle score: %d", scores[len(scores)/2])
}
