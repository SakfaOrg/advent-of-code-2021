package day24

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestALU(t *testing.T) {
	t.Run("Test arithmetic", func(t *testing.T) {
		demoLines := []string {
			"add z w",
			"mod z 2",
			"div w 2",
			"add y w",
			"mod y 2",
			"div w 2",
			"add x w",
			"mod x 2",
			"div w 2",
			"mod w 2",
		}
		var demoInstructions []Instruction
		for _, line := range demoLines {
			demoInstructions = append(demoInstructions, parseInstruction(line))
		}

		state := State{15, 0, 0, 0}
		state.applyUntilInpOrEnd(demoInstructions)
		assert.Equal(t, State{1, 1, 1, 1}, state)

		state = State{6, 0, 0, 0}
		state.applyUntilInpOrEnd(demoInstructions)
		assert.Equal(t, State{0, 1, 1, 0}, state)
	})

	t.Run("Solve generic", func(t *testing.T) {
		generic := []string {
			"inp w",
			"mul w -1",
			"add z w",
			"inp w",
			"mod w 5",
			"add z w",
		}

		assert.Equal(t, "Max input is 49, min is 11", solve(parseProgram(generic)))
	})
}
