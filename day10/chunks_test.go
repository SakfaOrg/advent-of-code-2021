package day10

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const demoInput = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestChunks(t *testing.T) {
	t.Run("Test part1 demo lines", func(t *testing.T) {
		err, score, _ := processLine("foo")
		assert.Equal(t, "Syntax error: Unrecognized symbol 'f'.", err)

		err, score, _ = processLine(")")
		assert.Equal(t, "Syntax error: found closing bracket ')' but nothing was opened.", err)

		err, score, _ = processLine("{([(<{}[<>[]}>{[]{[(<()>")
		assert.Equal(t, "Syntax error: Expected ], but found } instead.", err)
		assert.Equal(t, 1197, score)

		err, score, _ = processLine("[[<[([]))<([[{}[[()]]]")
		assert.Equal(t, "Syntax error: Expected ], but found ) instead.", err)
		assert.Equal(t, 3, score)

		err, score, _ = processLine("[{[{({}]{}}([{[{{{}}([]")
		assert.Equal(t, "Syntax error: Expected ), but found ] instead.", err)
		assert.Equal(t, 57, score)

		err, score, _ = processLine("[<(<(<(<{}))><([]([]()")
		assert.Equal(t, "Syntax error: Expected >, but found ) instead.", err)
		assert.Equal(t, 3, score)

		err, score, _ = processLine("<{([([[(<>()){}]>(<<{{")
		assert.Equal(t, "Syntax error: Expected ], but found > instead.", err)
		assert.Equal(t, 25137, score)
	})

	t.Run("Test part2 demo lines", func(t *testing.T) {
		message, _, score := processLine("[({(<(())[]>[[{[]{<()<>>")
		assert.Equal(t, "OK", message)
		assert.Equal(t, 288957, score)

		message, _, score = processLine("[(()[<>])]({[<{<<[]>>(")
		assert.Equal(t, "OK", message)
		assert.Equal(t, 5566, score)

		message, _, score = processLine("(((({<>}<{<{<>}{[]{[]{}")
		assert.Equal(t, "OK", message)
		assert.Equal(t, 1480781, score)

		message, _, score = processLine("{<[[]]>}<{[{[{[]{()[[[]")
		assert.Equal(t, "OK", message)
		assert.Equal(t, 995444, score)

		message, _, score = processLine("<{([{{}}[<[[[<>{}]]]>[]]")
		assert.Equal(t, "OK", message)
		assert.Equal(t, 294, score)
	})

	t.Run("Test part1", func(t *testing.T) {
		assert.Equal(t, "Syntax error score: 26397", Part1(strings.Split(demoInput, "\n")))
	});

	t.Run("Test part2", func(t *testing.T) {
		assert.Equal(t, "Middle score: 288957", Part2(strings.Split(demoInput, "\n")))
	});
}

