package day24

import (
	advent "advent2021/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

type State struct {
	w,x,y,z int
}

func (s State) String() string {
	return fmt.Sprintf("%d,%d,%d,%d", s.w,s.x,s.y,s.z)
}

type Command uint8
const (
	INP Command = iota
	ADD
	MUL
	DIV
	MOD
	EQL
)

type OperandType uint8
const (
	W = iota
	X
	Y
	Z
	LITERAL
	EMPTY
)

type Operand struct {
	register OperandType
	value    int
}

type Instruction struct {
	cmd      Command
	operandA Operand
	operandB Operand
}

func (o Operand) read(s *State) int {
	switch o.register {
	case LITERAL: return o.value
	case EMPTY: panic("Can't read EMPTY")
	case W: return s.w
	case X: return s.x
	case Y: return s.y
	case Z: return s.z
	default: panic("Invalid register " + string(o.register))
	}
}

func (o Operand) store(s *State, value int) {
	switch o.register {
	case LITERAL:
		panic("Can't store to LITERAL")
	case EMPTY:
		panic("Can't store to EMPTY")
	case W:
		s.w = value
	case X:
		s.x = value
	case Y:
		s.y = value
	case Z:
		s.z = value
	default:
		panic("Invalid register " + string(o.register))
	}
}

func (s *State) applyInp(instruction Instruction, digit int) {
	if instruction.cmd != INP {
		panic("Only INP can be applied this way.")
	}

	instruction.operandA.store(s, digit)
}

func (s *State) apply(instruction Instruction) {
	if instruction.cmd == INP {
		panic("INP can't be applied this way")
	}

	result := 0
	a := instruction.operandA.read(s)
	b := instruction.operandB.read(s)
	switch instruction.cmd {
	case ADD:
		result = a + b
	case MUL:
		result = a * b
	case DIV:
		result = a / b
	case MOD:
		result = a % b
	case EQL:
		if a == b {
			result = 1
		}
	}

	instruction.operandA.store(s, result)
}

func (s *State) applyUntilInpOrEnd(instructions []Instruction) (remaining []Instruction) {
	for idx, instruction := range instructions {
		if instruction.cmd != INP {
			s.apply(instruction)
		} else {
			return instructions[idx:]
		}
	}
	return []Instruction{}
}

func parseInstruction(instruction string) Instruction {
	splitted := strings.Split(instruction, " ")
	var operation Command
	switch splitted[0] {
	case "inp":
		operation = INP
	case "add":
		operation = ADD
	case "mul":
		operation = MUL
	case "div":
		operation = DIV
	case "mod":
		operation = MOD
	case "eql":
		operation = EQL
	default:
		panic("Unknown operation '" + splitted[0] + "'")
	}
	operandA := parseOperand(splitted[1])
	var operandB Operand
	if operation == INP {
		operandB = Operand{EMPTY, 0}
	} else {
		operandB = parseOperand(splitted[2])
	}
	return Instruction{operation, operandA, operandB}
}

func parseOperand(operand string) Operand {
	switch operand {
	case "w": return Operand{W, 0}
	case "x": return Operand{X, 0}
	case "y": return Operand{Y, 0}
	case "z": return Operand{Z, 0}
	default: return Operand{LITERAL, advent.MustAtoi(operand)}
	}
}

type StateAndMinMax struct {
	s State
	min, max int64
}

func deduplicate(states []StateAndMinMax) []StateAndMinMax {
	sort.Slice(states, func(i int, j int) bool {
		a := states[i]
		b := states[j]
		if a.s.w < b.s.w {
			return true
		} else if a.s.w == b.s.w {
			if a.s.x < b.s.x {
				return true
			} else if a.s.x == b.s.x {
				if a.s.y < b.s.y {
					return true
				} else if a.s.y == b.s.y {
					return a.s.z < b.s.z
				}
			}
		}
		return false
	})

	var deduped []StateAndMinMax
	current := states[0]
	for i := 1; i < len(states); i++ {
		if states[i].s == current.s {
			current.min = advent.Min64(states[i].min, current.min)
			current.max = advent.Max64(states[i].max, current.max)
		} else {
			deduped = append(deduped, current)
			current = states[i]
		}
	}
	deduped = append(deduped, current)
	return deduped
}

func solveNoHashMap(program Program) string {
	states := []StateAndMinMax{{State{0, 0, 0, 0}, 0, 0}}
	for digitIdx := 1; digitIdx <= program.digitsCount(); digitIdx++ {
		var nextStates []StateAndMinMax
		for _, stateAndMinMax := range states {
			state := stateAndMinMax.s
			for digit := 1; digit <= 9; digit++ {
				// apply current INP for current digit
				newState := State{state.w, state.x, state.y, state.z}
				newState.w = digit // note this works because all inp are `inp w`
				program.run(digitIdx, &newState)
				// an ugly shortcut but it works for my input: ALL inp instructions store to 'w' register, we can move ahead
				// and nullify this register already to save quite a lot on number of states we will explore in next step
				// samy thing with x and y registers, they seem to be multipled by 0 before they are used. In other words:
				// after each segment, only value remaining in z matters. Curiously, it doesn't help a lot.
				newState.w = 0
				newState.x = 0
				newState.y = 0

				newMax := int64(10)*stateAndMinMax.max + int64(digit)
				newMin := int64(10)*stateAndMinMax.min + int64(digit)
				nextStates = append(nextStates, StateAndMinMax{newState, newMin, newMax})
			}
		}
		states = deduplicate(nextStates)
		fmt.Printf("Done inp %d (have %d states)\n", digitIdx, len(states))
	}

	min := int64(math.MaxInt64)
	max := int64(0)
	for _, state := range states {
		if state.s.z == 0 {
			if state.min < min {
				min = state.min
			}
			if state.max > max {
				max = state.max
			}
		}
	}

	return fmt.Sprintf("Max input is %d, min is %d", max, min)
}

type MinAndMax struct {
	min, max int64
}

func merge(states map[State]MinAndMax, newState State, newMin, newMax int64) {
	if seenValue, seen := states[newState]; seen {
		if seenValue.min < newMin {
			newMin = seenValue.min
		}
		if seenValue.max > newMax {
			newMax = seenValue.max
		}
	}
	states[newState] = MinAndMax{newMin, newMax}
}

func processDigit(program Program, digit int, digitIdx int, states map[State]MinAndMax, result chan map[State]MinAndMax) {
	go func() {
		nextStates := make(map[State]MinAndMax)
		for state, minMax := range states {
			newState := State{state.w, state.x, state.y, state.z}
			newState.w = digit // note this works because all inp are `inp w`
			program.run(digitIdx, &newState)
			// an ugly shortcut but it works for my input: ALL inp instructions store to 'w' register, we can move ahead
			// and nullify this register already to save quite a lot on number of states we will explore in next step
			// samy thing with x and y registers, they seem to be multipled by 0 before they are used. In other words:
			// after each segment, only value remaining in z matters. Curiously, it doesn't help a lot.
			newState.w = 0
			newState.x = 0
			newState.y = 0

			newMin := int64(10)*minMax.min + int64(digit)
			newMax := int64(10)*minMax.max + int64(digit)
			merge(nextStates, newState, newMin, newMax)
		}
		result <- nextStates
	}()
}

func solve(program Program) string {
	states := make(map[State]MinAndMax) // a map of state to value of input that led to that state
	states[State{0,0,0,0}] = MinAndMax{0, 0}

	for digitIdx := 1; digitIdx <= program.digitsCount(); digitIdx++ {
		nextStates := make(map[State]MinAndMax)
		digitResult := make(chan map[State]MinAndMax)

		for digit := 1; digit <= 9; digit++ {
			processDigit(program, digit, digitIdx, states, digitResult)
		}
		for digit := 1; digit <= 9; digit++ {
			for newState, minMax := range <-digitResult {
				merge(nextStates, newState, minMax.min, minMax.max)
			}
		}

		states = nextStates
		fmt.Printf("Done inp %d (have %d states)\n", digitIdx, len(states))
	}

	min := int64(math.MaxInt64)
	max := int64(0)
	for state, minMax := range states {
		if state.z == 0 {
			if minMax.min < min {
				min = minMax.min
			}
			if minMax.max > max {
				max = minMax.max
			}
		}
	}

	return fmt.Sprintf("Max input is %d, min is %d", max, min)
}

type Program interface {
	digitsCount() int
	run(digitIdx int, s *State)
}

type CompiledProgram struct {}

func (c CompiledProgram) digitsCount() int {
	return 14
}

func (c CompiledProgram) run(digitIdx int, s *State) {
	s.runBlock(digitIdx)
}

type InterpretedProgram struct {
	blocks map[int][]Instruction
}

func (i InterpretedProgram) digitsCount() int {
	return len(i.blocks)
}

func (i InterpretedProgram) run(digitIdx int, s *State) {
	s.applyUntilInpOrEnd(i.blocks[digitIdx])
}

func parseProgram(lines []string) InterpretedProgram {
	var block []Instruction
	inpIdx := 0
	blocks := make(map[int][]Instruction)

	for _, line := range lines {
		instruction := parseInstruction(line)
		if instruction.cmd == INP {
			if instruction.operandA.register != W {
				panic("Only `inp w` is supported currently.")
			}

			if inpIdx > 0 {
				blocks[inpIdx] = block
				block = []Instruction{}
			}
			inpIdx++
		} else {
			block = append(block, instruction)
		}
	}
	blocks[inpIdx] = block
	return InterpretedProgram{blocks}
}

func greater(a, b int64) bool {
	return a > b
}

func smaller(a, b int64) bool {
	return a < b
}

func Part1And2Fast(_ []string) string {
	return solveNoHashMap(CompiledProgram{})
}

func Part1And2(_ []string) string {
	return solve(CompiledProgram{})
}
