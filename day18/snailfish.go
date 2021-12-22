package day18

import (
	"fmt"
)

type SnailfishNumber struct {
	value                     int
	leftChild, rightChild     *SnailfishNumber
	leftSibling, rightSibling *SnailfishNumber
	parent                    *SnailfishNumber
}

func (s SnailfishNumber) isPair() bool {
	return s.leftChild != nil && s.rightChild != nil
}

func (s SnailfishNumber) String() string {
	if s.isPair() {
		return fmt.Sprintf("[%s,%s]", s.leftChild, s.rightChild)
	} else {
		return fmt.Sprintf("%d", s.value)
	}
}

func (s SnailfishNumber) magnitude() int {
	if s.isPair() {
		return 3*s.leftChild.magnitude() + 2*s.rightChild.magnitude()
	} else {
		return s.value
	}
}

func add(a, b *SnailfishNumber) *SnailfishNumber {
	sum := &SnailfishNumber{
		leftChild:  a,
		rightChild: b,
	}
	a.parent = sum
	b.parent = sum

	aRight := findRightmostLeaf(a)
	bLeft := findLeftmostLeaf(b)
	aRight.rightSibling = bLeft
	bLeft.leftSibling = aRight

	reduce(sum)
	return sum
}

func findLeftmostLeaf(n *SnailfishNumber) *SnailfishNumber {
	if !n.isPair() {
		return n
	} else {
		return findLeftmostLeaf(n.leftChild)
	}
}

func findRightmostLeaf(n *SnailfishNumber) *SnailfishNumber {
	if !n.isPair() {
		return n
	} else {
		return findRightmostLeaf(n.rightChild)
	}
}

func parse(number string) *SnailfishNumber {
	var stack []*SnailfishNumber
	var lastLiteral *SnailfishNumber
	for i := 0; i < len(number); i++ {
		chr := number[i]
		switch chr {
		case '[':
			newPair := &SnailfishNumber{}
			if len(stack) > 0 {
				newPair.parent = stack[len(stack)-1]
			}
			stack = append(stack, newPair)
		case ']':
			rightOperand := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			stack[len(stack)-1].rightChild = rightOperand
		case ',':
			leftOperand := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			stack[len(stack)-1].leftChild = leftOperand
		default:
			if !(chr >= '0' && chr <= '9') {
				panic("Unexpected symbol, expected one of [], or digit but got " + string(chr))
			}

			digit := int(chr - '0')
			top := stack[len(stack)-1]
			if top.value > 0 {
				top.value = top.value*10 + digit
			} else {
				newLiteral := &SnailfishNumber{
					value:  digit,
					parent: stack[len(stack)-1],
				}

				if lastLiteral != nil {
					lastLiteral.rightSibling = newLiteral
					newLiteral.leftSibling = lastLiteral
				}
				lastLiteral = newLiteral
				stack = append(stack, newLiteral)
			}
		}
	}

	if len(stack) != 1 {
		panic(fmt.Sprintf("Expected exactly 1 item leftChild on stack but got %d", len(stack)))
	} else {
		return stack[0]
	}
}

func reduce(number *SnailfishNumber) {
	for didSomething := true; didSomething; {
		didSomething = explodeWalk(0, number) || splitWalk(number)
	}
}

func explodeNumber(number *SnailfishNumber) {
	// prepare number that will replace us: value 0, same parent
	replacement := &SnailfishNumber{
		value:  0,
		parent: number.parent,
	}

	// now check if we had siblings: if so, add our value to them and fix the siblink links
	if number.leftChild.leftSibling != nil {
		number.leftChild.leftSibling.value += number.leftChild.value
		number.leftChild.leftSibling.rightSibling = replacement
		replacement.leftSibling = number.leftChild.leftSibling
	}
	if number.rightChild.rightSibling != nil {
		number.rightChild.rightSibling.value += number.rightChild.value
		number.rightChild.rightSibling.leftSibling = replacement
		replacement.rightSibling = number.rightChild.rightSibling
	}

	// finally replace node in the tree
	replaceChild(number, replacement)
}

func explodeWalk(level int, number *SnailfishNumber) (hadExploded bool) {
	if !number.isPair() {
		return false
	} else {
		if level == 4 {
			explodeNumber(number)
			return true
		} else {
			return explodeWalk(level+1, number.leftChild) || explodeWalk(level+1, number.rightChild)
		}
	}
}

func splitNumber(number *SnailfishNumber) {
	leftChild := &SnailfishNumber{ // left child is number divided by 2 rounded down
		value: number.value / 2,
	}
	rightChild := &SnailfishNumber{ // right is number divided by 2...
		value: number.value / 2,
	}
	if number.value%2 == 1 { // ...rounded up
		rightChild.value += 1
	}

	// fix siblings
	leftChild.rightSibling = rightChild
	rightChild.leftSibling = leftChild
	if number.leftSibling != nil {
		leftChild.leftSibling = number.leftSibling
		number.leftSibling.rightSibling = leftChild
	}
	if number.rightSibling != nil {
		rightChild.rightSibling = number.rightSibling
		number.rightSibling.leftSibling = rightChild
	}

	replacement := &SnailfishNumber{
		leftChild:  leftChild,
		rightChild: rightChild,
		parent:     number.parent,
	}

	// and set correct parent on children
	leftChild.parent = replacement
	rightChild.parent = replacement

	replaceChild(number, replacement)
}

func splitWalk(number *SnailfishNumber) bool {
	if !number.isPair() {
		if number.value >= 10 {
			splitNumber(number)
			return true
		}
		return false
	} else {
		return splitWalk(number.leftChild) || splitWalk(number.rightChild)
	}
}

func replaceChild(number *SnailfishNumber, replacement *SnailfishNumber) {
	if number.parent.leftChild == number {
		number.parent.leftChild = replacement
	} else if number.parent.rightChild == number {
		number.parent.rightChild = replacement
	} else {
		panic("Child is not a child of its parent?!")
	}
}

func addLines(lines []string) *SnailfishNumber {
	current := parse(lines[0])
	for i := 1; i < len(lines); i++ {
		next := parse(lines[i])
		current = add(current, next)
	}
	return current
}

func Part1(lines []string) string {
	number := addLines(lines)
	return fmt.Sprintf("Result: %s, magnitude: %d", number, number.magnitude())
}

func Part2(lines []string) string {
	var maxSum int
	var maxA, maxB string
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i != j {
				ni := parse(lines[i])
				nj := parse(lines[j])
				sum := add(ni, nj).magnitude()
				if sum > maxSum {
					maxSum = sum
					maxA = lines[i]
					maxB = lines[j]
				}
			}
		}
	}

	return fmt.Sprintf("Maximum magnitude %d when adding %s + %s", maxSum, maxA, maxB)
}
